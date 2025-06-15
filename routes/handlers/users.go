package handlers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/itemslabs/clubz-api/config"
	"github.com/itemslabs/clubz-api/database"
	"github.com/itemslabs/clubz-api/database/schema"
	"github.com/itemslabs/clubz-api/routes/apiconv"
	"github.com/itemslabs/clubz-api/routes/model"
	"github.com/itemslabs/clubz-api/util"
	connections "github.com/itemslabs/clubz-api/web3"
	"github.com/go-openapi/strfmt"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"

	"github.com/volatiletech/null/v8"
)

var JwtSecretKey string

type Notification struct {
	Tokens   []string               `json:"tokens"`
	Platform int                    `json:"platform"` // 1 for iOS, 2 for Android
	Message  string                 `json:"message"`
	Title    string                 `json:"title,omitempty"`
	Data     map[string]interface{} `json:"data,omitempty"`
}

// NotificationRequest wraps a slice of notifications.
type NotificationRequest struct {
	Notifications []Notification `json:"notifications"`
}

type UpgradeRequest struct {
	SequenceSessionID *string `json:"sequenceSessionID"`
	WalletAddress     *string `json:"walletAddress"`
	Username          *string `json:"username"`
	Email             *string `json:"email"`
}

type LoginRequest struct {
	WalletAddress     *string `json:"walletAddress"`
	SequenceSessionID *string `json:"sequenceSessionID"`
	Email             *string `json:"email"`
}

func init() {
	JwtSecretKey = os.Getenv("JWT_SECRET")
}

var userNameRe, _ = regexp.Compile(`^[a-zA-Z0-9]*$`)

// GetUserByID godoc
// @Summary Get user by ID
// @Description Get user by ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} model.User
// @Router /users/{id} [get]
func (e *Env) GetUserByID(c echo.Context) error {
	user, err := e.Store.GetUserByID(c.Param("id"))
	if err != nil {
		return err
	}

	isFollowing, err := e.Store.IsFollowing(userID(c), user.ID)
	if err != nil {
		return err
	}

	return e.RespondSuccess(c, apiconv.ToUser(user, isFollowing))
}

// CurrentProfile godoc
// @Summary Current profile
// @Description Current profile
// @Tags users
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} model.CurrentUser
// @Router /users/me [get]
func (e *Env) CurrentProfile(c echo.Context) error {
	user, err := e.Store.GetUserByID(userID(c))
	if err != nil {
		return err
	}

	return e.RespondSuccess(c, apiconv.ToCurrentUser(user))
}

// SearchUsers godoc
// @Summary Search users
// @Description Search users
// @Tags users
// @Accept json
// @Produce json
// @Param query query string true "Search query"
// @Success 200 {object}  model.ShortUserArray
// @Router /users/search [get]
func (e *Env) SearchUsers(c echo.Context) error {
	users, err := e.Store.SearchUsers(c.QueryParam("query"))
	if err != nil {
		return err
	}

	return e.RespondSuccess(c, apiconv.ToShortUserArray(users))
}

// UpdateUserName godoc
// @Summary Update user name
// @Description Update user name
// @Tags users
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param body body model.UpdateUserNameRequest true "Update user name request"
// @Success 200 {object} model.CurrentUser
// @Router /users/me/name [put]
func (e *Env) UpdateUserName(c echo.Context) error {
	var in model.UpdateUserNameRequest

	if err := e.ParseBody(c, &in); err != nil {
		return err
	}

	user, err := e.Store.GetUserByID(userID(c))
	if err != nil {
		return err
	}

	// Bypass the name change interval check if the name starts with "Player#"
	if !strings.HasPrefix(user.Name, "Player#") {
		if user.LastNameChange.Valid && time.Since(user.LastNameChange.Time) < config.NameChangeInterval() {
			return ErrNameAlreadyChangedRecently
		}
	}

	if in.Name != nil {
		user.Name = *in.Name
	}

	// Check username using regex
	if !userNameRe.Match([]byte(user.Name)) {
		return ErrInvalidName
	}
	user.Name = strings.Replace(strings.TrimSpace(user.Name), " ", "_", -1)

	if user, err := e.Store.UpdateUserName(user, time.Now()); err != nil {
		if err == database.UserWithNameAlreadyExists {
			err = ErrNameIsAlreadyTaken
		}
		return err
	} else {
		return e.RespondSuccess(c, apiconv.ToCurrentUser(user))
	}
}

// UpdateUserAvatar godoc
// @Summary Update user avatar
// @Description Update user avatar
// @Tags users
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param body body model.UpdatePayPalEmailRequest true "Update user avatar request"
// @Success 200 {object} model.CurrentUser
// @Router /users/me/avatar [put]
func (e *Env) UpdatePayPalEmail(c echo.Context) error {
	var in model.UpdatePayPalEmailRequest

	if err := e.ParseBody(c, &in); err != nil {
		return err
	}

	user, err := e.Store.GetUserByID(userID(c))
	if err != nil {
		return err
	}

	user.PaypalEmail = null.StringFromPtr(in.PaypalEmail)
	if user, err := e.Store.UpdateUserPayPalEmail(user); err != nil {
		return err
	} else {
		return e.RespondSuccess(c, apiconv.ToCurrentUser(user))
	}
}

// UpdateUserAvatar godoc
// @Summary Update user avatar
// @Description Update user avatar
// @Tags users
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param body body model.UseReferralRequest true "Update user avatar request"
// @Success 200 {object} model.CurrentUser
// @Router /users/me/avatar [put]
func (e *Env) UseReferralCode(c echo.Context) error {
	var in model.UseReferralRequest

	if err := e.ParseBody(c, &in); err != nil {
		return err
	}

	// trim and transform to uppercase
	refCode := *in.RefCode
	refCode = strings.ToUpper(refCode)
	refCode = strings.TrimSpace(refCode)

	err := e.Store.Transaction(
		func(s database.Store) error {
			// lock user
			currentUser, err := e.Store.LockUser(userID(c))
			if err != nil {
				return err
			}

			if currentUser.ReferrerID.Valid {
				return ErrReferralCodeAlreadyApplied
			}

			// get referrer
			referrer, err := e.Store.GetUserByReferralCode(refCode)
			if err != nil {
				if err == sql.ErrNoRows {
					return ErrInvalidReferralCode
				}
				return err
			}

			// update referrer link for current user
			currentUser.ReferrerID = null.StringFrom(referrer.ID)
			if _, err = e.Store.UpdateUserReferrer(currentUser); err != nil {
				return err
			}

			return nil
		},
	)

	if err != nil {
		return err
	}

	// retrieve updated current user and return it back
	user, err := e.Store.GetUserByID(userID(c))
	if err != nil {
		return err
	}

	return e.RespondSuccess(c, apiconv.ToCurrentUser(user))
}

// FollowUser godoc
// @Summary Follow user
// @Description Follow user
// @Tags users
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "User ID"
// @Success 204
// @Router /users/{id}/follow [post]
func (e *Env) FollowUser(c echo.Context) error {
	currentUserID := userID(c)
	givenUserID := c.Param("id")

	if givenUserID == currentUserID {
		return ErrCannotFollowYourself
	}

	// restrict follow unless user change a name
	currentUser, err := e.Store.GetUserByID(currentUserID)
	if err != nil {
		return err
	}

	if !currentUser.NameChanged {
		return ErrCannotFollowUntilYouSetYourName
	}

	// restrict follow of guest users
	givenUser, err := e.Store.GetUserByID(givenUserID)
	if err != nil {
		return err
	}

	if !givenUser.NameChanged {
		return ErrCannotFollowGuestUser
	}

	err = e.Store.Transaction(
		func(s database.Store) error {
			if err := s.CreateFollower(currentUserID, givenUserID); err != nil {
				if err == database.FollowerAlreadyExists {
					return ErrYouAlreadyFollowThisUser
				}
				return err
			}

			if err := s.ChangeFollowerCount(givenUserID, 1); err != nil {
				return err
			}

			if err := s.ChangeFollowingCount(currentUserID, 1); err != nil {
				return err
			}

			return util.SendPush(
				e.Store, givenUser.ID, "", "New Follower 👀",
				fmt.Sprintf("%s has just started following you!", currentUser.Name), map[string]string{
					"user_id": currentUser.ID,
					"type":    "new_follower",
				},
			)
		},
	)
	if err != nil {
		return err
	}

	return e.RespondNoContent(c)
}

// UnFollowUser godoc
// @Summary Unfollow user
// @Description Unfollow user
// @Tags users
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "User ID"
// @Success 204
// @Router /users/{id}/unfollow [post]
func (e *Env) UnFollowUser(c echo.Context) error {
	givenUserID := c.Param("id")
	currentUserID := userID(c)

	if givenUserID == currentUserID {
		return ErrCannotUnFollowYourself
	}

	err := e.Store.Transaction(
		func(s database.Store) error {
			if err := s.DeleteFollower(currentUserID, givenUserID); err != nil {
				if err == database.FollowerAlreadyExists {
					return ErrYouAlreadyFollowThisUser
				}
				return err
			}

			if err := s.ChangeFollowerCount(givenUserID, -1); err != nil {
				return err
			}

			if err := s.ChangeFollowingCount(currentUserID, -1); err != nil {
				return err
			}

			return nil
		},
	)
	if err != nil {
		return err
	}

	return e.RespondNoContent(c)
}

// GetFollowers godoc
// @Summary Get followers
// @Description Get followers
// @Tags users
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param search query string false "Search query"
// @Success 200 {object} model.ShortUserArray
// @Router /users/me/followers [get]
func (e *Env) GetFollowers(c echo.Context) error {
	followers, err := e.Store.GetFollowers(userID(c), c.QueryParam("search"))
	if err != nil {
		return err
	}

	return e.RespondSuccess(c, apiconv.ToShortUserArray(followers))
}

// GetFollowings godoc
// @Summary Get followings
// @Description Get followings
// @Tags users
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param search query string false "Search query"
// @Success 200 {object} model.ShortUserArray
// @Router /users/me/followings [get]
func (e *Env) GetFollowings(c echo.Context) error {
	followings, err := e.Store.GetFollowings(userID(c), c.QueryParam("search"))
	if err != nil {
		return err
	}

	return e.RespondSuccess(c, apiconv.ToShortUserArray(followings))
}

// GetUserWalletBalance godoc
// @Summary Get user wallet balance
// @Description Get user wallet balance
// @Tags users
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param contractType path string true "Contract type"
// @Success 200 {object} map[string]string
// @Router /users/me/wallet/balance/{contractType} [get]
func (e *Env) GetUserWalletBalance(c echo.Context) error {
	userID := userID(c) // Assume userID(c) properly extracts the user ID set in the JWT middleware

	user, err := e.Store.GetUserByID(userID)
	if err != nil {
		log.Printf("Error retrieving user: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Error retrieving user.")
	}
	if user == nil || !user.WalletAddress.Valid {
		return echo.NewHTTPError(http.StatusBadRequest, "User does not have a valid wallet address.")
	}

	walletAddress := user.WalletAddress.String
	contractType := c.Param("contractType")

	// Assuming connections.GetBalanceOf is your function to get the wallet balance
	balance, err := connections.GetBalanceOf(walletAddress, contractType)
	if err != nil {
		log.Printf("Error getting wallet balance: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Error getting wallet balance.")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"balance": balance})
}

// GenerateJWTToken creates a token with a specific claim based on user type
func GenerateJWTToken(userID string, isGuest bool) (string, error) {
	// Set expiration to 10 years
	expiration := time.Now().AddDate(10, 0, 0)

	// For guest users, you can choose to keep a shorter duration or the same. Assuming the same here.
	if isGuest {
		expiration = time.Now().AddDate(10, 0, 0) // 10 years for guests as well
	}

	claims := jwt.MapClaims{
		"user_id":  userID,
		"exp":      expiration.Unix(),
		"is_guest": isGuest,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(JwtSecretKey))
}

// Register godoc
// @Summary Register a new user
// @Description Register a new user with email, password, and account type.
// @Tags auth
// @Accept json
// @Produce json
// @Param user body model.RegisterRequest true "Register Information"
// @Success 200 {object} model.RegisterResponse
// @Failure 400 {object} model.ErrorResponse
// @Router /auth/register [post]
func (e *Env) Register(c echo.Context) error {
	var req model.RegisterRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request body")
	}

	safeNullString := func(s *string) null.String {
		if s != nil {
			return null.StringFrom(*s)
		}
		return null.NewString("", false)
	}

	safeNullStringForEmail := func(email *strfmt.Email) null.String {
		if email != nil && email.String() != "" {
			return null.StringFrom(email.String())
		}
		return null.NewString("", false)
	}
	newUser := &schema.User{
		ID:               uuid.New().String(),
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
		Name:             "Player#" + uuid.NewString(),
		Email:            safeNullStringForEmail(req.Email),
		WalletAddress:    safeNullString(req.WalletAddress),
		FirebaseID:       safeNullString(req.FirebaseID),
		AvatarURL:        safeNullString(req.AvatarURL),
		Balance:          1000,
		PaypalEmail:      safeNullString(req.PaypalEmail),
		Verified:         false,
		BonusPowerups:    0,
		ReferralCode:     safeNullString(req.RefCode),
		ReferrerID:       safeNullString(req.RefCodeID),
		AvgPoints:        null.IntFrom(0),
		AvgRank:          null.IntFrom(0),
		FollowerCount:    null.IntFrom(0),
		FollowingCount:   null.IntFrom(0),
		GamesPlayed:      null.IntFrom(0),
		NameChanged:      false,
		Moderator:        false,
		Premium:          req.AccountType != nil && *req.AccountType == "premium",
		Influencer:       false,
		SubscriptionTier: 0,
		AvgRankPercent:   null.IntFrom(0),
		LastNameChange:   null.TimeFromPtr(nil),
	}

	// Call your data store function to create the user
	createdUser, err := e.Store.CreateUserWithPassword(newUser, safeNullString(req.Password).String)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to create user: "+errors.Unwrap(err).Error())
	}

	token, err := GenerateJWTToken(createdUser.ID, false) // false indicates it's a full account
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to generate token")
	}

	return c.JSON(
		http.StatusOK, echo.Map{
			"message": "User registered successfully",
			"token":   token,
		},
	)
}

// Login godoc
// @Summary User login
// @Description Login with email and password
// @Tags auth
// @Accept json
// @Produce json
// @Param login body model.LoginRequest true "Login Information"
// @Success 200 {object} model.LoginResponse
// @Failure 401 {string} string "Unauthorized"
// @Router /auth/login [post]
func (e *Env) Login(c echo.Context) error {
	var req model.LoginRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request body")
	}

	email := ""
	if req.Email != nil {
		email = req.Email.String()
	}

	password := ""
	if req.Password != nil {
		password = *req.Password
	}

	user, err := e.Store.AuthenticateUser(email, password)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
	}

	// if user == nil || !user.WalletAddress.Valid {
	// 	return echo.NewHTTPError(http.StatusBadRequest, "User does not have a valid wallet address.")
	// }

	// walletAddress := user.WalletAddress.String
	// log.Printf("This is the wallet address" + walletAddress)

	// balance, err := connections.GetBalanceOf(walletAddress, "WarChest")
	// // Error getting wallet balance: no contract code at given address
	// if err != nil && err.Error() == "no contract code at given address" {
	// 	return echo.NewHTTPError(http.StatusBadRequest, "User does not have sufficient NFT balance.")
	// }
	// if err != nil {
	// 	log.Printf("Error getting wallet balance: %v", err)
	// 	return echo.NewHTTPError(http.StatusInternalServerError, "Error getting wallet balance.")
	// }

	// zero := big.NewInt(0)

	// // Use the Cmp method to check if balance is 0
	// if balance.Cmp(zero) == 0 {
	// 	return echo.NewHTTPError(http.StatusBadRequest, "User does not have sufficient NFT balance.")
	// }

	// Generate JWT token for the authenticated user
	token, err := GenerateJWTToken(user.ID, false)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to generate token")
	}

	return c.JSON(http.StatusOK, model.LoginResponse{Token: token})
}

// VerifyEmail godoc
// @Summary Verify email
// @Description Verify email using a token sent to the email
// @Tags auth
// @Accept json
// @Produce json
// @Param token query string true "Verification Token"
// @Success 200 {string} string "Email verified successfully"
// @Failure 400 {string} string "Invalid or expired token"
// @Router /auth/verify-email [get]
func (e *Env) VerifyEmail(c echo.Context) error {
	// Retrieve the token from query parameters
	token := c.QueryParam("token")

	// Check if the token is provided
	if token == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Missing token")
	}

	// Proceed with token verification
	err := e.Store.VerifyEmail(token)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid or expired token")
	}

	// Respond back on success
	return c.String(http.StatusOK, "Email verified successfully")
}

// ResendVerificationEmail godoc
// @Summary Resend verification email
// @Description Resend the email verification link
// @Tags auth
// @Accept json
// @Produce json
// @Param user_id query string true "User ID"
// @Success 200 {string} string "Verification email resent successfully"
// @Failure 400 {string} string "Unable to resend verification email"
// @Router /auth/resend-verification [post]
func (e *Env) ResendVerificationEmail(c echo.Context) error {
	email := c.QueryParam("email")
	if email == "" {
		email = c.Param("email")
	}

	if email == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Missing user ID")
	}

	err := e.Store.ResendVerificationEmail(email)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Unable to resend verification email")
	}

	return c.String(http.StatusOK, "Verification email resent successfully")
}

// RegisterUser godoc
// @Summary Register a new user
// @Description Register a new user with email, password, and account type.
// @Tags auth
// @Accept json
// @Produce json
// @Param user body model.RegisterRequest true "Register Information"
// @Success 200 {object} model.RegisterResponse
// @Failure 400 {object} model.ErrorResponse
// @Router /auth/register [post]
func (e *Env) RegisterUser(c echo.Context) error {
	var req model.RegisterRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request body")
	}

	newUser, err := e.Store.CreateUser(
		&schema.User{
			Email:   null.StringFrom(req.Email.String()), // Assume req.Email is of type string
			Balance: 1000,
		},
	)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to create user: "+err.Error())
	}

	// Return the response, possibly including any relevant user info or tokens
	return c.JSON(http.StatusOK, model.RegisterResponse{UserID: newUser.ID})
}

type UpdatePasswordRequest struct {
	OldPassword string `json:"oldPassword"`
	NewPassword string `json:"newPassword"`
}

// UpdatePassword godoc
// @Summary Update user password
// @Description Update user password
// @Tags auth
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param body body model.UpdatePasswordRequest true "Update password request"
// @Success 204
// @Router /auth/update-password [put]
func (e *Env) UpdatePassword(c echo.Context) error {
	var in UpdatePasswordRequest

	// Bind the incoming JSON payload to the UpdatePasswordRequest struct
	if err := c.Bind(&in); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request format")
	}

	// Extract the user ID from the context/session/token
	userID := userID(c) // Ensure this function exists and correctly retrieves the user's ID

	// Attempt to update the password
	err := e.Store.UpdatePassword(userID, in.OldPassword, in.NewPassword)

	// Handle specific errors returned from UpdatePassword
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return echo.NewHTTPError(http.StatusBadRequest, "Old password does not match")
		}
		// Log the error internally for debugging purposes
		log.Printf("Failed to update password for user %s: %v", userID, err)
		// For other errors, return a generic server error message
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to update password due to an internal error")
	}

	// If the password was successfully updated, return no content
	return c.NoContent(http.StatusNoContent)
}

// RegisterGuest creates a guest user and returns a JWT token
func (e *Env) RegisterGuest(c echo.Context) error {
	// Get the next user index for the guest name
	nextIndex, err := e.Store.GetNextUserIndex()
	if err != nil {
		return err
	}

	// Create the guest user
	guestUser := &schema.User{
		ID:      uuid.NewString(),
		Name:    fmt.Sprintf("Guest#%d", nextIndex),
		Balance: 1000,
	}

	createdUser, err := e.Store.CreateUser(guestUser)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to create guest user: "+err.Error())
	}

	// Assign default badges, banners, and frames to the new user
	err = e.assignDefaultItemsToUser(createdUser.ID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to assign default items: "+err.Error())
	}

	// Generate JWT token for the user
	token, err := GenerateJWTToken(createdUser.ID, true) // true indicates it's a guest user
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to generate token")
	}

	// Return the user ID and token in the response
	return c.JSON(
		http.StatusOK, echo.Map{
			"userID": createdUser.ID,
			"token":  token,
		},
	)
}

// assignDefaultItemsToUser assigns default badges, banners, and frames to the user
func (e *Env) assignDefaultItemsToUser(userID string) error {
	// Fetch default badges
	defaultBadges, err := e.Store.GetTypeDefaultBadges()
	if err != nil {
		return err
	}
	// Assign each default badge to the user
	for _, badge := range defaultBadges {
		_, err = e.Store.CreateUserBadge(userID, badge.ID)
		if err != nil {
			return err
		}
	}

	// Fetch default banners
	defaultBanners, err := e.Store.GetTypeDefaultBadges()
	if err != nil {
		return err
	}
	// Assign each default banner to the user
	for _, banner := range defaultBanners {
		_, err = e.Store.CreateUserBanner(userID, banner.ID)
		if err != nil {
			return err
		}
	}

	// Fetch default frames
	defaultFrames, err := e.Store.GetDefaultFrames()
	if err != nil {
		return err
	}
	// Assign each default frame to the user
	for _, frame := range defaultFrames {
		_, err = e.Store.CreateUserFrame(userID, frame.ID)
		if err != nil {
			return err
		}
	}

	return nil
}

// RefreshToken checks the token, renews it if the user is active.
func (e *Env) RefreshToken(c echo.Context) error {
	token := strings.TrimPrefix(c.Request().Header.Get("Authorization"), "Bearer ")
	newToken, err := RefreshToken(token, JwtSecretKey)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "Invalid or expired token: "+err.Error())
	}

	return c.JSON(http.StatusOK, echo.Map{"token": newToken})
}

// RefreshToken function defined to handle the actual refreshing logic
func RefreshToken(currentToken, secret string) (string, error) {
	// Parse the old token
	token, err := jwt.ParseWithClaims(
		currentToken, &jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		},
	)
	if err != nil {
		log.Printf("Error parsing token: %v", err)
		return "", errors.New("invalid token")
	}

	// Check if token is valid and extract claims
	if claims, ok := token.Claims.(*jwt.MapClaims); ok && token.Valid {
		// Update expiration time
		newExp := time.Now().Add(24 * time.Hour).Unix() // adjust the duration as needed
		(*claims)["exp"] = newExp

		// Create a new token with updated expiration
		newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		return newToken.SignedString([]byte(secret))
	} else {
		return "", errors.New("invalid token claims")
	}
}

// UpgradeToPremium allows a guest user to become a premium user
func (e *Env) UpgradeToPremium(c echo.Context) error {
	var req model.UpgradeRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request body")
	}

	userID := userID(c) // Retrieve user ID from JWT
	if userID == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "User ID not available")
	}

	user, err := e.Store.GetUserByID(userID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "User not found: "+err.Error())
	}

	if user.Email.Valid {
		return echo.NewHTTPError(http.StatusBadRequest, "User already upgraded")
	}

	hashedPassword, err := HashPassword(*req.Password)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to hash password")
	}
	if req.Username == nil || *req.Username == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Username must not be empty")
	}

	validPassword, errMsg := validatePassword(*req.Password)
	if !validPassword {
		return echo.NewHTTPError(http.StatusBadRequest, errMsg)
	}

	if req.Email != nil {
		emailStr := string(*req.Email)
		if !isValidEmail(emailStr) {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid email")
		}
		unique, err := e.isEmailUnique(emailStr)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		if !unique {
			return echo.NewHTTPError(http.StatusBadRequest, "Email already exists")
		}
	} else {
		return echo.NewHTTPError(http.StatusBadRequest, "Email must not be empty")
	}
	// will need validation for this
	user.Name = *req.Username
	user.Email = null.StringFrom(req.Email.String())
	user.PasswordHash = null.StringFrom(hashedPassword)
	user.Premium = true
	user.Balance += 1500

	if _, err := e.Store.UpdateUserName(user, time.Now()); err != nil {
		if err == database.UserWithNameAlreadyExists {
			err = ErrNameIsAlreadyTaken
		}
		return err
	}

	if err := e.Store.UpdateUser(user); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	err = e.Store.SendVerificationEmail(user)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	token, err := GenerateJWTToken(user.ID, false)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to generate token")
	}

	return c.JSON(
		http.StatusOK, echo.Map{
			"message": "User upgraded to premium successfully",
			"token":   token,
		},
	)
}

func (e *Env) isEmailUnique(email string) (bool, error) {
	exists, err := e.Store.EmailExists(email)
	if err != nil {
		return false, err
	}
	return !exists, nil
}

func validatePassword(password string) (bool, string) {
	minLength := 4
	var errorMessages []string

	if len(password) < minLength {
		errorMessages = append(errorMessages, fmt.Sprintf("Password must be at least %d characters long", minLength))
	}
	if !strings.ContainsAny(password, "0123456789") {
		errorMessages = append(errorMessages, "Password must include at least one number")
	}
	if !strings.ContainsAny(password, "ABCDEFGHIJKLMNOPQRSTUVWXYZ") {
		errorMessages = append(errorMessages, "Password must include at least one uppercase letter")
	}
	if !strings.ContainsAny(password, "abcdefghijklmnopqrstuvwxyz") {
		errorMessages = append(errorMessages, "Password must include at least one lowercase letter")
	}
	if !strings.ContainsAny(password, "!@#$%^&*()_+-=[]{}|?<>,./") {
		errorMessages = append(errorMessages, "Password must include at least one special character")
	}

	if len(errorMessages) == 0 {
		return true, "Password is valid"
	}

	return false, strings.Join(errorMessages, ", ")
}

func HashPassword(password string) (string, error) {
	hashedNewPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashedNewPassword), err
}

func isValidEmail(email string) bool {
	// This regex allows for a wider range of email addresses
	emailRegex := regexp.MustCompile(`(?i)^[a-z0-9!#$%&'*+/=?^_` + "`" + `{|}~.-]+@[a-z0-9-]+\.[a-z]{2,24}$`)
	return emailRegex.MatchString(email)
}

// RegisterPushToken handles the registration of a push token from a user's device.
func (env *Env) RegisterPushToken(c echo.Context) error {
	userID := userID(c)
	var token struct {
		PushToken string `json:"pushToken"`
	}
	if err := c.Bind(&token); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid data"})
	}

	// Logic to save the token to the database associated with the userID
	err := env.Store.SavePushToken(userID, token.PushToken)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]string{"status": "success"})
}

// // HandlePushNotification sends a push notification using the provided details.
// func HandlePushNotification(c echo.Context) error {
// 	var req NotificationRequest
// 	if err := c.Bind(&req); err != nil {
// 		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid request body"})
// 	}

// 	// Serialize the request body into JSON
// 	payloadBytes, err := json.Marshal(req)
// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "could not encode request body"})
// 	}
// 	body := bytes.NewReader(payloadBytes)

// 	// Define the Gorush service URL
// 	gorushURL := "https://exp.host/--/api/v2/push/send"

// 	// Create a new POST request to the Gorush server
// 	httpRequest, err := http.NewRequest("POST", gorushURL, body)
// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "could not create request"})
// 	}
// 	httpRequest.Header.Set("Content-Type", "application/json")

// 	// Send the request using the http client
// 	client := &http.Client{}
// 	resp, err := client.Do(httpRequest)
// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "failed to send notification"})
// 	}
// 	defer resp.Body.Close()

// 	// Check the response status code
// 	if resp.StatusCode != http.StatusOK {
// 		return c.JSON(
// 			http.StatusInternalServerError,
// 			echo.Map{"error": "failed to send notification, received non-200 status code"},
// 		)
// 	}

// 	// Return success
// 	return c.JSON(http.StatusOK, echo.Map{"message": "Notifications sent successfully"})
// }

// DeleteUser godoc
// @Summary Delete user
// @Description Delete a user and all related data
// @Tags users
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "User ID"
// @Success 204
// @Failure 400 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
// @Router /users/{id} [delete]
func (e *Env) DeleteUser(c echo.Context) error {
	userID := userID(c)
	if userID == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Missing user ID")
	}

	err := e.Store.DeleteUser(userID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to delete user: "+err.Error())
	}

	return c.JSON(
		http.StatusOK, echo.Map{
			"message": "User successfully deleted",
		},
	)
}

// func (e *Env) ForgotPassword(c echo.Context) error {
// 	var req struct {
// 		Email string `json:"email"`
// 	}

// 	if err := c.Bind(&req); err != nil {
// 		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request format")
// 	}

// 	token, err := e.Store.GeneratePasswordResetToken(req.Email)
// 	if err != nil {
// 		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to generate reset token")
// 	}

// 	// Construct reset URL
// 	resetURL := fmt.Sprintf("https://yourapp.com/reset-password?token=%s", token)
// 	emailBody := fmt.Sprintf(
// 		"Subject: Password Reset Request\n\nClick the following link to reset your password: %s",
// 		resetURL,
// 	)

// 	// Convert emailBody to []byte
// 	emailBodyBytes := []byte(emailBody)

// 	// Create a context
// 	ctx := context.Background()

// 	// Send email
// 	err = util.SendEmail(ctx, []string{req.Email}, emailBodyBytes)
// 	if err != nil {
// 		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to send reset email")
// 	}

// 	return c.JSON(http.StatusOK, map[string]string{"message": "Password reset email sent successfully"})
// }

// func (e *Env) ResetPassword(c echo.Context) error {
// 	var req struct {
// 		Token       string `json:"token"`
// 		NewPassword string `json:"newPassword"`
// 	}

// 	if err := c.Bind(&req); err != nil {
// 		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request format")
// 	}

// 	err := e.Store.UpdatePasswordWithToken(req.Token, req.NewPassword)
// 	if err != nil {
// 		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
// 	}

// 	return c.JSON(http.StatusOK, map[string]string{"message": "Password updated successfully"})
// }

// UpgradeToPremium allows a guest user to become a premium user
func (e *Env) UpgradeToPremiumSequence(c echo.Context) error {
	var req UpgradeRequest
	nextIndex, err := e.Store.GetNextUserIndex()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to get next user index")
	}
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request body")
	}

	// Validate session ID
	if req.SequenceSessionID == nil || *req.SequenceSessionID == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Session ID is required")
	}

	// Validate wallet address
	if req.WalletAddress == nil || *req.WalletAddress == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Wallet address is required")
	} else if u, err := e.Store.AuthenticateUserWithSession(*req.WalletAddress); err == nil && u != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Account already exists")
	}

	userID := userID(c) // Retrieve user ID from JWT
	if userID == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "User ID not available")
	}

	user, err := e.Store.GetUserByID(userID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "User not found: "+err.Error())
	}

	user.Name = fmt.Sprintf("Player#%d", nextIndex)
	user.WalletAddress = null.StringFrom(*req.WalletAddress)
	// Update Email if provided
	if req.Email != nil && *req.Email != "" {
		user.Email = null.StringFrom(*req.Email)
	}
	user.EmailVerified = true
	user.Premium = true
	user.Balance += 1500

	if _, err := e.Store.UpdateUserName(user, time.Now()); err != nil {
		if err == database.UserWithNameAlreadyExists {
			err = ErrNameIsAlreadyTaken
		}
		return err
	}

	if err := e.Store.UpdateUser(user); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	token, err := GenerateJWTToken(user.ID, false)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to generate token")
	}

	return c.JSON(
		http.StatusOK, echo.Map{
			"message": "User upgraded to premium successfully",
			"token":   token,
		},
	)
}

// Login godoc
// @Summary User login
// @Description Login with email and session ID
// @Tags auth
// @Accept json
// @Produce json
// @Param login body model.LoginRequest true "Login Information"
// @Success 200 {object} model.LoginResponse
// @Failure 401 {string} string "Unauthorized"
// @Router /auth/login [post]
func (e *Env) LoginSequence(c echo.Context) error {
	var req LoginRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request body")
	}

	// Validate wallet address
	if req.WalletAddress == nil || *req.WalletAddress == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Wallet address is required")
	}

	// Validate sequence session ID
	if req.SequenceSessionID == nil || *req.SequenceSessionID == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Sequence session ID is required")
	}

	walletAddress := *req.WalletAddress
	newSessionID := *req.SequenceSessionID

	user, err := e.Store.AuthenticateUserWithSession(walletAddress)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
	}

	// Update SequenceSessionID
	user.SequenceSessionID = null.StringFrom(newSessionID)

	// Update Email if provided
	if req.Email != nil && *req.Email != "" {
		user.Email = null.StringFrom(*req.Email)
	}

	if err := e.Store.UpdateUser(user); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	// Generate JWT token for the authenticated user
	token, err := GenerateJWTToken(user.ID, false)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to generate token")
	}

	return c.JSON(http.StatusOK, model.LoginResponse{Token: token})
}

// GetBalanceOf godoc
// @Summary Get balance of an address
// @Description Get balance of an address
// @ID get-balance
// @Produce json
// @Param address path string true "Address to get balance of"
// @Param contractType path string true "Contract type"
// @Success 200 {object} string
// @Router /web3/balance/{address}/{contractType} [get]
func (e *Env) PlayerInventory(c echo.Context) error {
	userID := userID(c)
	user, err := e.Store.GetUserByID(userID)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	address := user.WalletAddress.String
	balance, err := connections.GetBalanceOf(address, "Players")
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	// Assuming GetAssignedPlayersByNFTIDs expects a string slice of NFT IDs
	// Convert balance to appropriate format if needed
	balanceStr := balance.String() // If balance is supposed to be a single value
	// balanceSlice := []string{balanceStr} // If balance needs to be a slice of strings

	players, err := e.Store.GetAssignedPlayersByNFTIDs([]string{balanceStr})
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	// Convert players to JSON serializable format
	playersJSON, err := json.Marshal(players)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]string{"balance": string(playersJSON)})
}

// GetUserByWalletAddress godoc
// @Summary Get user by wallet address
// @Description Get user by wallet address
// @Tags users
// @Accept json
// @Produce json
// @Param walletAddress path string true "User's Wallet Address"
// @Success 200 {object} model.User
// @Failure 400 {object} model.ErrorResponse
// @Failure 404 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
// @Router /users/wallet/{walletAddress} [get]
func (e *Env) GetUserByWalletAddress(c echo.Context) error {
	walletAddress := c.Param("walletAddress")
	if walletAddress == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Wallet address is required")
	}

	// Fetch the user by wallet address
	exists, err := e.Store.GetUserByWalletAddress(walletAddress)
	if err != nil {
		// Log the error and return internal server error
		log.Printf("Failed to retrieve user by wallet address: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to retrieve user")
	}

	// Return JSON response with existence result
	return c.JSON(http.StatusOK, echo.Map{"exists": exists})
}

// GetUserByUsernameByWalletAddress godoc
// @Summary Get username by wallet address
// @Description Retrieve the username of a user by their wallet address
// @Tags users
// @Accept json
// @Produce json
// @Param walletAddress path string true "User's Wallet Address"
// @Success 200 {object} map[string]string "Username retrieved successfully"
// @Failure 400 {object} model.ErrorResponse "Invalid or missing wallet address"
// @Failure 404 {object} model.ErrorResponse "User not found"
// @Failure 500 {object} model.ErrorResponse "Internal server error"
// @Router /users/username/wallet/{walletAddress} [get]
func (e *Env) GetUserByUsernameByWalletAddress(c echo.Context) error {
	walletAddress := c.Param("walletAddress")
	if walletAddress == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Wallet address is required")
	}

	username, err := e.Store.GetUsernameByWalletAddress(walletAddress)
	if err != nil {
		if err == sql.ErrNoRows {
			// Instead of creating a new user, return a suggestion to create a user
			return c.JSON(http.StatusNotFound, map[string]string{"username": "User not found"})
		}
		log.Printf("Error retrieving username by wallet address: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to retrieve username")
	}

	return c.JSON(http.StatusOK, map[string]string{"username": username})
}

// UpgradeOrLogin allows a guest user to become a premium user or logs in an existing user
func (e *Env) UpgradeOrLogin(c echo.Context) error {
	var req UpgradeRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request body")
	}

	// Validate session and wallet address
	if req.SequenceSessionID == nil || *req.SequenceSessionID == "" || req.WalletAddress == nil || *req.WalletAddress == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Session ID and Wallet address are required")
	}

	// Try to authenticate the user with the wallet address
	user, err := e.Store.AuthenticateUserWithSession(*req.WalletAddress)
	if err != nil {
		if err.Error() == "no user found for email and wallet address" {
			// No user found, proceed to create a new premium account
			return e.createPremiumUser(c, req)
		}
		return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
	}

	// User exists, proceed with login and session ID update
	// Check if email is provided and set it
	if req.Email != nil && *req.Email != "" {
		user.Email = null.StringFrom(*req.Email)
	} else {
		// If no email is provided, set Email to null
		user.Email = null.String{}
	}
	user.SequenceSessionID = null.StringFrom(*req.SequenceSessionID)
	if err := e.Store.UpdateUser(user); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err, "Failed to update user")
	}

	token, err := GenerateJWTToken(user.ID, false)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to generate token")
	}

	return c.JSON(http.StatusOK, model.LoginResponse{Token: token})
}

// Helper function to create a premium user
func (e *Env) createPremiumUser(c echo.Context, req UpgradeRequest) error {
	nextIndex, err := e.Store.GetNextUserIndex()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to get next user index")
	}

	userID := userID(c) // Retrieve user ID from JWT
	if userID == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "User ID not available")
	}

	user, err := e.Store.GetUserByID(userID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "User not found: "+err.Error())
	}

	user.Name = fmt.Sprintf("Player#%d", nextIndex)
	user.WalletAddress = null.StringFrom(*req.WalletAddress)
	user.Premium = true
	user.Balance += 1500

	if _, err := e.Store.UpdateUserName(user, time.Now()); err != nil {
		if err == database.UserWithNameAlreadyExists {
			err = ErrNameIsAlreadyTaken
		}
		return err
	}

	if err := e.Store.UpdateUser(user); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	token, err := GenerateJWTToken(user.ID, false)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to generate token")
	}

	return c.JSON(
		http.StatusOK, echo.Map{
			"message": "User upgraded to premium successfully",
			"token":   token,
		},
	)
}

// ChangeEmailRequest defines the expected request payload for changing email
type ChangeEmailRequest struct {
	NewEmail string `json:"newEmail"`
}

// ChangeEmail godoc
// @Summary Change user email
// @Description Allows users to change their email address directly without password
// @Tags users
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param body body ChangeEmailRequest true "Change email request"
// @Success 200 {string} string "Email changed successfully"
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal server error"
// @Router /users/me/change-email [put]
func (e *Env) ChangeEmail(c echo.Context) error {
	var req ChangeEmailRequest

	// Parse the request body
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request body")
	}

	// Validate that the new email is not empty
	if req.NewEmail == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "New email is required")
	}

	// Validate the new email format
	if !isValidEmail(req.NewEmail) {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid email format")
	}

	// Check if the email is already in use
	exists, err := e.Store.EmailExists(req.NewEmail)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error checking email uniqueness")
	}
	if exists {
		return echo.NewHTTPError(http.StatusBadRequest, "Email already in use")
	}

	// Get the current user ID from the context
	userID := userID(c)

	// Fetch the user from the database
	user, err := e.Store.GetUserByID(userID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error retrieving user")
	}

	// Update the user's email
	user.Email = null.StringFrom(req.NewEmail)

	// Save the updated user in the database
	if err := e.Store.UpdateUser(user); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error updating user email")
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Email changed successfully"})
}

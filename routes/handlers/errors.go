package handlers

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/labstack/echo/v4"
)

type APIError struct {
	Code       int    `json:"code"`
	Message    string `json:"message"`
	StatusCode int    `json:"-"`
}

func (e *APIError) Error() string {
	return e.Message
}

var (
	ErrInvalidAppToken                 = &APIError{Message: "Invalid app token", StatusCode: http.StatusForbidden}
	ErrUnauthorized                    = &APIError{Message: "Unauthorized", StatusCode: http.StatusUnauthorized}
	ErrObjectNotFound                  = &APIError{Message: "Not found", StatusCode: http.StatusNotFound, Code: 404}
	ErrCannotUploadAvatar              = &APIError{Message: "Upload avatar error"}
	ErrNameIsAlreadyTaken              = &APIError{Message: "Ooops.. this username is taken!"}
	ErrInvalidName                     = &APIError{Message: "Name can only contain english characters and digits"}
	ErrNameAlreadyChangedRecently      = &APIError{Message: "You already changed name in last 30 days"}
	ErrMatchIsNotAvailable             = &APIError{Message: "Match is not available"}
	ErrInvalidNumberOfPicks            = &APIError{Message: "Invalid number of picks"}
	ErrYouBanned                       = &APIError{Message: "You are currently suspended and cannot play this game"}
	ErrPickAlreadyEnded                = &APIError{Message: "Pick already ended"}
	ErrYouAlreadyPickedThisPlayer      = &APIError{Message: "You already picked this player"}
	ErrInvalidPowerUpPosition          = &APIError{Message: "Invalid power-up position"}
	ErrInvalidMatchStatusForPowerUp    = &APIError{Message: "You can apply power-ups when match is active"}
	ErrAlreadyActivePowerUpForPosition = &APIError{Message: "You already have active power-up for this position"}
	ErrNoPowerUpLeft                   = &APIError{Message: "No power-up left"}
	ErrNotEnoughBalanceForWithdraw     = &APIError{Message: "Sorry, you need a minimum balance of $15 to cashout. Win some more games!"}
	ErrInvalidReferralCode             = &APIError{Message: "Invalid referral code"}
	ErrReferralCodeAlreadyApplied      = &APIError{Message: "Referral code already applied"}
	ErrCannotFollowYourself            = &APIError{Message: "Cannot follow yourself"}
	ErrCannotFollowGuestUser           = &APIError{Message: "Cannot follow guest user"}
	ErrCannotFollowUntilYouSetYourName = &APIError{Message: "Cannot follow user until you set your name"}
	ErrYouAlreadyFollowThisUser        = &APIError{Message: "You already follow this user"}
	ErrCannotUnFollowYourself          = &APIError{Message: "Cannot unfollow yourself"}
	ErrTooLongMessage                  = &APIError{Message: "Too Long Message"}
	ErrNotAuthorized                   = &APIError{Message: "Not authorized"}
	ErrPaymentMethod                   = &APIError{
		Message:    "Invalid payments method",
		StatusCode: http.StatusBadRequest,
	}
	ErrPaymentIntent = &APIError{
		Message:    "Failed to create payments intent",
		StatusCode: http.StatusInternalServerError,
	}
	ErrNotEnoughBalanceForPowerup    = &APIError{Message: "Not enough balance to acquire powerups!"}
	ErrSubstitutionMatchInvalidState = &APIError{Message: "You cannot swap players if a match is not in a proper state (waiting to start or playing)"}
)

// ErrorHandler godoc
// @Summary Error handler
// @Description Error handler
// @Accept json
// @Produce json
// @Param err body interface{} true "Error"
// @Success 400 {object} APIError
// @Failure 500 {object} APIError
// @Router /error [post]
func ErrorHandler(err error, c echo.Context) {
	var apiErr *APIError

	switch v := err.(type) {
	case *APIError:
		apiErr = v
	case *echo.HTTPError:
		apiErr = &APIError{Message: fmt.Sprintf("%s", v.Message), Code: v.Code, StatusCode: v.Code}
	case *errors.CompositeError:
		apiErr = &APIError{Message: v.Error(), Code: http.StatusBadRequest, StatusCode: http.StatusBadRequest}
	default:
		fmt.Println(err)
		if err == sql.ErrNoRows {
			apiErr = ErrObjectNotFound
		} else {
			apiErr = &APIError{Message: "Unknown error", Code: 99999, StatusCode: http.StatusInternalServerError}
		}
	}

	if apiErr.Code == 0 {
		apiErr.Code = 99999
	}
	if apiErr.StatusCode == 0 {
		apiErr.StatusCode = http.StatusBadRequest
	}

	_ = c.JSON(
		apiErr.StatusCode, map[string]interface{}{
			"code":    apiErr.Code,
			"message": apiErr.Message,
		},
	)
}

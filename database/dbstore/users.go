package dbstore

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/gameon-app-inc/laliga-matchfantasy-api/database"
	"github.com/gameon-app-inc/laliga-matchfantasy-api/database/schema"
	"github.com/gameon-app-inc/laliga-matchfantasy-api/util"
	"github.com/google/uuid"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"golang.org/x/crypto/bcrypt"
)

func (s *DBStore) GetUserByID(id string) (*schema.User, error) {
	return schema.Users(
		qm.Where("id = ?", id),
		qm.Load("BanPenalties"),
	).One(s.db)
}

func (s *DBStore) GetUserByName(name string) (*schema.User, error) {
	return schema.Users(
		qm.Where("lower(name) = ?", strings.ToLower(name)),
	).One(s.db)
}

func (s *DBStore) GetUserByReferralCode(code string) (*schema.User, error) {
	return schema.Users(
		qm.Where("referral_code = ?", code),
	).One(s.db)
}

func (s *DBStore) LockUser(userID string) (*schema.User, error) {
	return schema.Users(
		qm.Where("id = ?", userID),
		qm.For("update"),
	).One(s.db)
}

func (s *DBStore) GetUserIDByFirebaseID(firebaseID string) (string, error) {
	var userID string
	err := s.db.QueryRow("SELECT id FROM users WHERE firebase_id = ?", "firebaseid1").Scan(&userID)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("No user found for Firebase ID: firebaseid1")
			return "", fmt.Errorf("no user found for Firebase ID: firebaseid1")
		}
		log.Printf("Error retrieving user ID by Firebase ID: %v", err)
		return "", err
	}

	log.Printf("Successfully retrieved User ID: %s for Firebase ID: firebaseid1", userID)
	return userID, nil
}

func (s *DBStore) GetNextUserIndex() (int, error) {
	query := `select nextval('user_guest_index') v`

	var seqResult struct {
		V int
	}
	if err := queries.Raw(query).Bind(nil, s.db, &seqResult); err != nil {
		return 0, err
	}
	return seqResult.V, nil
}

func (s *DBStore) SearchUsers(query string) (schema.UserSlice, error) {
	var users = schema.UserSlice{}
	if err := queries.Raw(
		"select * from users where name_changed = true and name % $1 order by similarity(name, $2) desc limit 15",
		query,
		query,
	).Bind(nil, s.db, &users); err != nil {
		return nil, err
	}
	return users, nil
}

func (s *DBStore) CreateUser(user *schema.User) (*schema.User, error) {
	if user.Name == "" {
		// create user name from guest index
		nextIndex, err := s.GetNextUserIndex()
		if err != nil {
			return nil, err
		}

		user.Name = fmt.Sprintf("Guest#%d", nextIndex)
	}

	for i := 0; i < 10; i++ {
		// generate random referral
		user.ReferralCode = null.StringFrom(util.GenerateReferralCode())

		if user, err := s.insertUser(user); err == nil {
			return user, nil
		} else {
			// try again with different code
			if err == database.ReferralCodeAlreadyExists {
				continue
			}
			return nil, s.transformError(err)
		}
	}

	return nil, database.CannotGenerateUniqueReferralCode
}

func (s *DBStore) insertUser(user *schema.User) (*schema.User, error) {
	err := user.Insert(s.db, boil.Infer())
	if err != nil {
		// check whether this is integrity error
		if strings.Contains(strings.ToLower(err.Error()), "users_firebase_id_key") {
			err = database.UserAlreadyExists
		}
		if strings.Contains(strings.ToLower(err.Error()), "users_email_key") {
			err = database.UserAlreadyExists
		}
		if strings.Contains(strings.ToLower(err.Error()), "users_referral_code_key") {
			err = database.ReferralCodeAlreadyExists
		}
		return nil, err
	}
	return user, nil
}

func (s *DBStore) UpdateUserName(user *schema.User, t time.Time) (*schema.User, error) {
	user.NameChanged = true
	user.LastNameChange = null.TimeFrom(t)
	if _, err := user.Update(s.db, boil.Whitelist("name", "name_changed", "last_name_change")); err != nil {
		return nil, s.transformError(err)
	} else {
		return user, nil
	}
}

func (s *DBStore) UpdateUserPayPalEmail(user *schema.User) (*schema.User, error) {
	if _, err := user.Update(s.db, boil.Whitelist("paypal_email")); err != nil {
		return nil, err
	} else {
		return user, nil
	}
}

func (s *DBStore) UpdateUserReferrer(user *schema.User) (*schema.User, error) {
	if _, err := user.Update(s.db, boil.Whitelist("referrer_id")); err != nil {
		return nil, err
	} else {
		return user, nil
	}
}

func (s *DBStore) UpdateUserReferrerBonusUsed(user *schema.User) error {
	if _, err := user.Update(s.db, boil.Whitelist(schema.UserColumns.ReferralBonusUsed)); err != nil {
		return err
	}
	return nil
}

func (s *DBStore) UpdateBonusPowerUps(user *schema.User) (*schema.User, error) {
	if _, err := user.Update(s.db, boil.Whitelist("bonus_powerups")); err != nil {
		return nil, err
	} else {
		return user, nil
	}
}

func (s *DBStore) UpdateWalletAddress(userID, address string) (err error) {
	if userID == "" {
		return errors.New("userID cannot be empty")
	}
	if address == "" {
		return errors.New("address cannot be empty")
	}

	tx, err := s.db.Begin()
	if err != nil {
		return err
	}

	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		} else if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	_, err = tx.Exec(`UPDATE users SET wallet_address = $1 WHERE id = $2`, address, userID)
	if err != nil {
		return s.transformError(err)
	}

	return nil
}

//func (s *DBStore) GetWithdrawRequests(userID string) (schema.WithdrawRequestSlice, error) {
//	return schema.WithdrawRequests(
//		qm.Where("user_id = ?", userID),
//		qm.OrderBy("created_at desc"),
//	).All(s.db)
//}

//func (s *DBStore) HasActiveWithdrawRequests(userID string) (bool, error) {
//	cnt, err := schema.WithdrawRequests(
//		qm.Where("user_id = ?", userID),
//		qm.WhereIn("status in ?", database.WithdrawRequestStatusRequested, database.WithdrawRequestStatusProcessing),
//	).Count(s.db)
//	if err != nil {
//		return false, err
//	}
//
//	return cnt > 0, nil
//}

func (s *DBStore) CreateFollower(fromUser, toUser string) error {
	record := &schema.Follower{
		FromUserID: fromUser,
		ToUserID:   toUser,
	}
	if err := record.Insert(s.db, boil.Infer()); err != nil {
		return s.transformError(err)
	}
	return nil
}

func (s *DBStore) DeleteFollower(fromUser, toUser string) error {
	_, err := queries.Raw(
		`delete from followers where from_user_id = $1 and to_user_id = $2`,
		fromUser,
		toUser,
	).Exec(s.db)
	if err != nil {
		return s.transformError(err)
	}
	return nil
}

func (s *DBStore) GetFollowers(userID string, search string) (schema.UserSlice, error) {
	var qms = []qm.QueryMod{
		qm.InnerJoin("followers on followers.from_user_id = users.id"),
		qm.Where("followers.to_user_id = ?", userID),
		qm.OrderBy("users.name"),
	}
	if search != "" {
		qms = append(
			qms,
			qm.Where("users.name ilike ?", fmt.Sprintf("%%%s%%", search)),
		)
	}
	return schema.Users(
		qms...,
	).All(s.db)
}

func (s *DBStore) GetFollowings(userID string, search string) (schema.UserSlice, error) {
	var qms = []qm.QueryMod{
		qm.InnerJoin("followers on followers.to_user_id = users.id"),
		qm.Where("followers.from_user_id = ?", userID),
		qm.OrderBy("users.name"),
	}
	if search != "" {
		qms = append(
			qms,
			qm.Where("users.name ilike ?", fmt.Sprintf("%%%s%%", search)),
		)
	}
	return schema.Users(
		qms...,
	).All(s.db)
}

func (s *DBStore) IsFollowing(fromUser, toUser string) (bool, error) {
	_, err := schema.Followers(
		qm.Where("from_user_id = ?", fromUser),
		qm.Where("to_user_id = ?", toUser),
	).One(s.db)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

func (s *DBStore) ChangeFollowerCount(userID string, diff int) error {
	_, err := queries.Raw(
		`update users set follower_count = coalesce(follower_count, 0) + $1 where id = $2`,
		diff,
		userID,
	).Exec(s.db)
	if err != nil {
		return s.transformError(err)
	}
	return nil
}

func (s *DBStore) ChangeFollowingCount(userID string, diff int) error {
	_, err := queries.Raw(
		`update users set following_count = coalesce(following_count, 0) + $1 where id = $2`,
		diff,
		userID,
	).Exec(s.db)
	if err != nil {
		return s.transformError(err)
	}
	return nil
}

func (s *DBStore) transformError(err error) error {
	errStr := strings.ToLower(err.Error())

	if strings.Contains(errStr, "users_name") && strings.Contains(errStr, "uniq") {
		err = database.UserWithNameAlreadyExists
	} else if strings.Contains(errStr, "followers_from_user") && strings.Contains(errStr, "uniq") {
		err = database.FollowerAlreadyExists
	}

	return err
}

func (s *DBStore) CreateUserWithPassword(user *schema.User, password string) (*schema.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("error hashing password: %w", err)
	}

	user.PasswordHash = null.StringFrom(string(hashedPassword))

	newUser, err := s.insertUser(user)
	if err != nil {
		return nil, fmt.Errorf("error creating user: %w", err)
	}

	if err = s.SendVerificationEmail(newUser); err != nil {
		return newUser, fmt.Errorf("error sending verification email: %w", err)
	}

	return newUser, nil
}

func GenerateToken() string {
	return uuid.NewString()
}

func (s *DBStore) SendVerificationEmail(user *schema.User) error {
	token := GenerateToken()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	to := []string{user.Email.String}
	body := util.ConstructEmailMessage(user.Email.String, token)
	done := make(chan error, 1)

	// Send email in a goroutine to run in the background.
	go func() {
		done <- util.SendEmail(ctx, to, body)
	}()

	select {
	case err := <-done:
		if err != nil {
			return errors.New("Something wrong with sending email: " + err.Error())
		} else {
			user.VerificationToken = null.StringFrom(token)
			user.TokenExpiration = null.NewTime(time.Now().Add(time.Hour*24), true)
			if _, err = user.Update(s.db, boil.Whitelist("verification_token", "token_expiration")); err != nil {
				log.Printf("Failed to update user verification token and expiration: %v", err)
				return errors.New("failed to update user verification token and expiration: " + err.Error())
			}
		}
	case <-ctx.Done():
		if ctx.Err() != nil {
			log.Println("Something wrong with sending email: ", ctx.Err())
			return errors.New("Something wrong with sending email: " + ctx.Err().Error())
		}
	}

	return nil
}

func (s *DBStore) VerifyEmail(token string) error {
	user, err := schema.Users(
		qm.Where("verification_token = ?", token),
		qm.Where("token_expiration > ?", time.Now()),
	).One(s.db)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("invalid or expired token")
		}
		return fmt.Errorf("error retrieving user: %w", err)
	}

	user.EmailVerified = true
	user.VerificationToken = null.StringFrom("")
	user.TokenExpiration = null.NewTime(time.Time{}, false)
	if _, err = user.Update(
		s.db,
		boil.Whitelist("email_verified", "verification_token", "token_expiration"),
	); err != nil {
		return fmt.Errorf("error updating user verification status: %w", err)
	}

	return nil
}

func (s *DBStore) AuthenticateUser(email, password string) (*schema.User, error) {
	user, err := schema.Users(qm.Where("email = ?", email)).One(s.db)
	if err != nil {
		return nil, errors.New("invalid login credentials")
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash.String), []byte(password)); err != nil {
		return nil, errors.New("invalid login credentials")
	}

	if !user.EmailVerified {
		return nil, errors.New("account not verified, please check your email")
	}

	return user, nil
}

func (s *DBStore) ResendVerificationEmail(email string) error {
	user, err := schema.Users(qm.Where("email = ?", email)).One(s.db)
	if err != nil {
		return errors.New("user not found")
	}

	if user.EmailVerified {
		return errors.New("email is already verified")
	}

	if err = s.SendVerificationEmail(user); err != nil {
		log.Printf("Error resending verification email: %v", err)
		return errors.New(err.Error())
	}

	return nil
}

func (s *DBStore) UpdatePassword(userID, oldPassword, newPassword string) error {
	// Fetch the user by ID
	user, err := schema.Users(qm.Where("id = ?", userID)).One(s.db)
	if err != nil {
		return fmt.Errorf("user not found: %w", err)
	}

	// Verify the old password
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash.String), []byte(oldPassword)); err != nil {
		return errors.New("old password does not match")
	}

	// Hash the new password
	hashedNewPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("error hashing new password: %w", err)
	}

	// Update the user's password hash
	user.PasswordHash = null.StringFrom(string(hashedNewPassword))
	if _, err = user.Update(s.db, boil.Whitelist("password_hash")); err != nil {
		return fmt.Errorf("failed to update password: %w", err)
	}

	return nil
}

func (s *DBStore) UpdateUser(user *schema.User) error {
	_, err := user.Update(
		s.db, boil.Whitelist(
			"email",
			"password_hash",
			"premium",
			"updated_at",
			"wallet_address",
			"email_verified",
			"sequence_session_id",
			"balance",
			"real_name",
			"paypal_email",
			"bonus_powerups",
		),
	)
	return err
}

func (s *DBStore) SavePushToken(userID, token string) error {
	if userID == "" {
		return errors.New("user ID cannot be empty")
	}
	if token == "" {
		return errors.New("token cannot be empty")
	}

	// Begin a transaction
	tx, err := s.db.Begin()
	if err != nil {
		return fmt.Errorf("error starting transaction: %w", err)
	}
	defer tx.Rollback()

	// Prepare the update statement
	_, err = tx.Exec(`UPDATE users SET firebase_id = $1 WHERE id = $2`, token, userID)
	if err != nil {
		return fmt.Errorf("error updating user push token: %w", err)
	}

	// Commit the transaction
	if err = tx.Commit(); err != nil {
		return fmt.Errorf("error committing transaction: %w", err)
	}

	return nil
}

func (s *DBStore) RemoveFinishedGameByUserID(userID, gameID string) error {
	query := `
        UPDATE users
        SET finished_games = array_remove(finished_games, $2)
        WHERE id = $1;
    `

	_, err := s.db.Exec(query, userID, gameID)
	return err
}

func (s *DBStore) EmailExists(email string) (bool, error) {
	var exists bool
	query := "SELECT EXISTS (SELECT 1 FROM users WHERE email = $1)"
	err := s.db.QueryRow(query, email).Scan(&exists)
	if err != nil {
		log.Printf("Error checking if email exists: %v", err)
		return false, err
	}
	return exists, nil
}

func (s *DBStore) UsernameExists(username string) (bool, error) {
	var exists bool
	query := "SELECT EXISTS (SELECT 1 FROM users WHERE lower(name) = lower($1))"
	err := s.db.QueryRow(query, username).Scan(&exists)
	if err != nil {
		log.Printf("Error checking if username exists: %v", err)
		return false, err
	}
	return exists, nil
}

func (s *DBStore) DeleteUser(userID string) error {
	// Begin a transaction
	tx, err := s.db.Begin()
	if err != nil {
		return fmt.Errorf("error starting transaction: %w", err)
	}
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		} else if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	// Delete related data in the transaction
	if err = s.deleteAllUserRelatedData(tx, userID); err != nil {
		return err
	}

	// Delete the user
	query := "DELETE FROM users WHERE id = $1"
	res, err := tx.Exec(query, userID)
	if err != nil {
		return fmt.Errorf("error deleting user: %w", err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("error fetching rows affected: %w", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("no user found with id: %s", userID)
	}

	return nil
}

func (s *DBStore) deleteAllUserRelatedData(tx *sql.Tx, userID string) error {
	// List of delete queries for all tables that have a foreign key reference to `users`
	queries := map[string]string{
		"game_events":                    "DELETE FROM game_events WHERE game_id IN (SELECT id FROM games WHERE user_id = $1)",
		"game_picks":                     "DELETE FROM game_picks WHERE game_id IN (SELECT id FROM games WHERE user_id = $1)",
		"transactions":                   "DELETE FROM transactions WHERE user_id = $1",
		"match_leaderboard":              "DELETE FROM match_leaderboard WHERE user_id = $1",
		"match_notifications":            "DELETE FROM match_notifications WHERE user_id = $1",
		"game_powerups":                  "DELETE FROM game_powerups WHERE game_id IN (SELECT id FROM games WHERE user_id = $1)",
		"games":                          "DELETE FROM games WHERE user_id = $1",
		"assigned_players":               "DELETE FROM assigned_players WHERE user_id = $1",
		"assigned_card_packs":            "DELETE FROM assigned_card_packs WHERE user_id = $1",
		"store_product_transactions":     "DELETE FROM store_product_transactions WHERE user_id = $1",
		"chat_messages":                  "DELETE FROM chat_messages WHERE sender_id = $1",
		"chat_room_members":              "DELETE FROM chat_room_members WHERE user_id = $1",
		"followers":                      "DELETE FROM followers WHERE from_user_id = $1 OR to_user_id = $1",
		"user_game_week_histories":       "DELETE FROM user_game_week_histories WHERE user_id = $1",
		"user_divisions":                 "DELETE FROM user_divisions WHERE user_id = $1",
		"user_leaderboard_subscriptions": "DELETE FROM user_leaderboard_subscriptions WHERE user_id = $1",
		"subscription_requests":          "DELETE FROM subscription_requests WHERE user_id = $1",
		"manual_subscriptions":           "DELETE FROM manual_subscriptions WHERE user_id = $1",
		"ban_penalties":                  "DELETE FROM ban_penalties WHERE user_id = $1",
		"app_inbox":                      "DELETE FROM app_inbox WHERE user_id = $1",
		"powerup_actions":                "DELETE FROM powerup_actions WHERE powerup_id IN (SELECT id FROM powerups WHERE user_id = $1)",
		"powerups":                       "DELETE FROM powerups WHERE user_id = $1",
	}

	const maxRetries = 5

	// Execute all queries in the transaction context, retrying up to 5 times for each
	for tableName, query := range queries {
		for attempt := 0; attempt < maxRetries; attempt++ {
			if _, err := tx.Exec(query, userID); err != nil {
				if attempt == maxRetries-1 {
					return fmt.Errorf(
						"error deleting related data for user %s in table [%s] after %d attempts: %w",
						userID,
						tableName,
						maxRetries,
						err,
					)
				}
			} else {
				break // Break out of the retry loop if the execution is successful
			}
		}
	}
	return nil
}
func (s *DBStore) GeneratePasswordResetToken(email string) (string, error) {
	_, err := schema.Users(qm.Where("email = ?", email)).One(s.db)
	if err != nil {
		return "", errors.New("user not found")
	}

	token := uuid.NewString()
	expiration := time.Now().Add(1 * time.Hour) // Token valid for 1 hour

	_, err = s.db.Exec(
		"UPDATE users SET reset_password_token = ?, token_expiration = ? WHERE email = ?",
		token,
		expiration,
		email,
	)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *DBStore) VerifyPasswordResetToken(token string) (*schema.User, error) {
	user, err := schema.Users(
		qm.Where("reset_password_token = ?", token),
		qm.Where("token_expiration > ?", time.Now()),
	).One(s.db)
	if err != nil {
		return nil, errors.New("invalid or expired token")
	}

	return user, nil
}

func (s *DBStore) UpdatePasswordWithToken(token, newPassword string) error {
	user, err := s.VerifyPasswordResetToken(token)
	if err != nil {
		return err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("error hashing password: %w", err)
	}

	_, err = s.db.Exec(
		"UPDATE users SET password_hash = ?, reset_password_token = NULL, token_expiration = NULL WHERE id = ?",
		string(hashedPassword),
		user.ID,
	)
	if err != nil {
		return err
	}

	return nil
}

// AuthenticateUserWithSession authenticates a user using email and walletAddress
func (s *DBStore) AuthenticateUserWithSession(walletAddress string) (*schema.User, error) {
	user, err := schema.Users(
		qm.Where("wallet_address = ?", walletAddress),
	).One(s.db)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("no user found for email and wallet address")
		}
		return nil, fmt.Errorf("error retrieving user with email  and wallet address: %s - %w", walletAddress, err)
	}
	return user, nil
}

// GetUserByWalletAddress retrieves a user by their wallet address.
func (s *DBStore) GetUserByWalletAddress(walletAddress string) (bool, error) {
	_, err := schema.Users(
		qm.Where("wallet_address = ?", walletAddress),
	).One(s.db)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil // No user found, return false but no error
		}
		return false, fmt.Errorf("error retrieving user by wallet address: %w", err)
	}
	return true, nil // User found, return true
}

func (s *DBStore) GetUsernameByWalletAddress(walletAddress string) (string, error) {
	user, err := schema.Users(
		qm.Where("wallet_address = ?", walletAddress),
	).One(s.db)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", fmt.Errorf("no user found for wallet address: %s", walletAddress)
		}
		return "", fmt.Errorf("error retrieving user by wallet address: %w", err)
	}

	return user.Name, nil // Assuming the username is stored in the 'Name' field.
}

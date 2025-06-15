package dbstore

import (
	"github.com/itemslabs/clubz-api/database/schema"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

// GetUserBadges retrieves all badges associated with a user.
func (s *DBStore) GetUserBadges(userID string) (schema.UserBadgeSlice, error) {
	return schema.UserBadges(
		qm.Where("user_id = ?", userID),
	).All(s.db)
}

// GetUserBadgeByID retrieves a user badge by its ID.
func (s *DBStore) GetUserBadgeByID(userBadgeID int) (*schema.UserBadge, error) {
	return schema.UserBadges(
		qm.Where("id = ?", userBadgeID),
	).One(s.db)
}

// DeleteUserBadge deletes a user badge by its ID.
func (s *DBStore) DeleteUserBadge(userBadgeID int) error {
	userBadge, err := schema.UserBadges(
		qm.Where("id = ?", userBadgeID),
	).One(s.db)
	if err != nil {
		return err
	}
	_, err = userBadge.Delete(s.db)
	return err
}

// CreateUserBadge creates a user badge for a specific user.
func (s *DBStore) CreateUserBadge(userID string, badgeID int) (*schema.UserBadge, error) {
	userBadge := &schema.UserBadge{
		UserID:  userID,
		BadgeID: badgeID,
	}

	err := userBadge.Insert(s.db, boil.Infer())
	if err != nil {
		return nil, err
	}

	return userBadge, nil
}

func (s *DBStore) ChangeAllUsersBadgesToUnselected() error {
	_, err := s.db.Exec("UPDATE user_badges SET selected = false")
	return err
}

func (s *DBStore) ChangeUserBadgeToSelected(userBadgeID int) error {
	_, err := s.db.Exec("UPDATE user_badges SET selected = true WHERE id = ?", userBadgeID)
	return err
}

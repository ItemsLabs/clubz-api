package dbstore

import (
	"github.com/gameon-app-inc/laliga-matchfantasy-api/database/schema"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

// GetUserBanners retrieves all banners associated with a user.
func (s *DBStore) GetUserBanners(userID string) (schema.UserBannerSlice, error) {
	return schema.UserBanners(
		qm.Where("user_id = ?", userID),
	).All(s.db)
}

// GetUserBannerByID retrieves a user banner by its ID.
func (s *DBStore) GetUserBannerByID(userBannerID int) (*schema.UserBanner, error) {
	return schema.UserBanners(
		qm.Where("id = ?", userBannerID),
	).One(s.db)
}

// DeleteUserBanner deletes a user banner by its ID.
func (s *DBStore) DeleteUserBanner(userBannerID int) error {
	userBanner, err := schema.UserBanners(
		qm.Where("id = ?", userBannerID),
	).One(s.db)
	if err != nil {
		return err
	}
	_, err = userBanner.Delete(s.db)
	return err
}

// CreateUserBanner creates a user banner for a specific user.
func (s *DBStore) CreateUserBanner(userID string, bannerID int) (*schema.UserBanner, error) {
	userBanner := &schema.UserBanner{
		UserID:   userID,
		BannerID: bannerID,
	}

	err := userBanner.Insert(s.db, boil.Infer())
	if err != nil {
		return nil, err
	}

	return userBanner, nil
}

func (s *DBStore) ChangeAllUsersBannersToUnselected() error {
	_, err := s.db.Exec("UPDATE user_banners SET selected = false")
	return err
}

func (s *DBStore) SelectUserBanner(userBannerID int) error {
	_, err := s.db.Exec("UPDATE user_banners SET selected = true WHERE id = ?", userBannerID)
	return err
}

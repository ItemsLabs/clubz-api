package dbstore

import (
	"database/sql"

	"github.com/gameon-app-inc/laliga-matchfantasy-api/database/schema"
	"github.com/gameon-app-inc/laliga-matchfantasy-api/routes/model"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

// GetBanners retrieves all banners, ordered by points in descending order.
func (s *DBStore) GetBanners() (schema.BannerSlice, error) {
	return schema.Banners(
		qm.OrderBy("points desc"),
	).All(s.db)
}

// GetBannerNameByID returns the name of a banner given its ID.
// It returns an empty string and an error if no banner is found for the given ID.
func (s *DBStore) GetBannerNameByID(bannerID int) (string, error) {
	banner, err := schema.FindBanner(s.db, bannerID)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", nil
		}
		return "", err
	}
	return banner.Name, nil
}

// GetBannerByID retrieves a banner by its ID.
func (s *DBStore) GetBannerByID(bannerID string) (*schema.Banner, error) {
	return schema.Banners(
		qm.Where("id = ?", bannerID),
	).One(s.db)
}

// DeleteBanner deletes a banner given its ID.
// It also deletes associated records from the user_banners table if necessary.
func (s *DBStore) DeleteBanner(bannerID string) error {
	_, err := schema.UserBanners(
		qm.Where("banner_id = ?", bannerID),
	).DeleteAll(s.db)
	if err != nil {
		return err
	}

	banner, err := schema.Banners(
		qm.Where("id = ?", bannerID),
	).One(s.db)
	if err != nil {
		return err
	}

	_, err = banner.Delete(s.db)
	return err
}

// UpdateBanner updates a banner's fields given its ID and an options struct.
func (s *DBStore) UpdateBanner(bannerID string, options *model.UpdateBannerRequest) (*schema.Banner, error) {
	banner, err := schema.Banners(qm.Where("id = ?", bannerID)).One(s.db)
	if err != nil {
		return nil, err
	}

	banner.Name = options.Name
	banner.Description = options.Description
	banner.Image = options.Image
	banner.Points = int(options.Points)
	banner.Type = options.Type
	banner.Status = options.Status

	_, err = banner.Update(s.db, boil.Infer())
	if err != nil {
		return nil, err
	}

	return banner, nil
}

// CreateBanners inserts multiple banners into the database and returns the newly inserted banners.
func (s *DBStore) CreateBanners(banners *model.CreateBannersRequest) ([]schema.Banner, error) {
	var insertedBanners []schema.Banner
	var maxID int
	err := s.db.QueryRow("SELECT MAX(id) FROM banners").Scan(&maxID)
	if err != nil {
		return nil, err
	}

	for _, banner := range *banners {
		maxID++

		schemaBanner := schema.Banner{
			ID:          maxID,
			Name:        *banner.Name,
			Description: *banner.Description,
			Image:       *banner.Image,
			Points:      int(*banner.Points),
			Type:        *banner.Type,
			Status:      *banner.Status,
		}

		err := schemaBanner.Insert(s.db, boil.Infer())
		if err != nil {
			return nil, err
		}

		insertedBanners = append(insertedBanners, schemaBanner)
	}

	return insertedBanners, nil
}

func (s *DBStore) GetTypeDefaultBanners() (schema.BannerSlice, error) {
	return schema.Banners(
		qm.Where("type = ?", "default"),
	).All(s.db)
}

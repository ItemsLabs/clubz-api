package dbstore

import (
	"database/sql"

	"github.com/gameon-app-inc/laliga-matchfantasy-api/database/schema"
	"github.com/gameon-app-inc/laliga-matchfantasy-api/routes/model"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

// GetBadges returns a list of badges ordered by points in descending order.
func (s *DBStore) GetBadges() (schema.BadgeSlice, error) {
	return schema.Badges(
		qm.OrderBy("points desc"),
	).All(s.db)
}

// GetBadgeNameByID returns the name of a badge given its ID.
// It returns an empty string and an error if no badge is found for the given ID.
func (s *DBStore) GetBadgeNameByID(badgeID int) (string, error) {
	badge, err := schema.FindBadge(s.db, badgeID)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", nil // or return an error indicating that no badge was found
		}
		return "", err
	}
	return badge.Name, nil
}

// GetBadgeByID retrieves a badge by its ID.
func (s *DBStore) GetBadgeByID(badgeID string) (*schema.Badge, error) {
	return schema.Badges(
		qm.Where("id = ?", badgeID),
	).One(s.db)
}

// DeleteBadge deletes a badge given its ID.
// It also deletes associated records from the user_badges table.
func (s *DBStore) DeleteBadge(badgeID string) error {
	_, err := schema.UserBadges(
		qm.Where("badge_id = ?", badgeID),
	).DeleteAll(s.db)
	if err != nil {
		return err
	}

	badge, err := schema.Badges(
		qm.Where("id = ?", badgeID),
	).One(s.db)
	if err != nil {
		return err
	}

	_, err = badge.Delete(s.db)
	return err
}

// UpdateBadge updates a badge's fields given its ID and an options struct.
func (s *DBStore) UpdateBadge(badgeID string, options *model.UpdateBadgeRequest) (*schema.Badge, error) {
	badge, err := schema.Badges(qm.Where("id = ?", badgeID)).One(s.db)
	if err != nil {
		return nil, err
	}

	badge.Name = options.Name
	badge.Description = options.Description
	badge.Image = options.Image
	badge.Points = int(options.Points)
	badge.Type = options.Type
	badge.Status = options.Status

	_, err = badge.Update(s.db, boil.Infer())
	if err != nil {
		return nil, err
	}

	return badge, nil
}

// CreateBadges inserts multiple badges into the database and returns the newly inserted badges.
func (s *DBStore) CreateBadges(badges *model.CreateBadgesRequest) ([]schema.Badge, error) {
	var insertedBadges []schema.Badge
	var maxID int
	err := s.db.QueryRow("SELECT MAX(id) FROM badges").Scan(&maxID)
	if err != nil {
		return nil, err
	}

	for _, badge := range *badges {
		maxID++

		schemaBadge := schema.Badge{
			ID:          maxID,
			Name:        *badge.Name,
			Description: *badge.Description,
			Image:       *badge.Image,
			Points:      int(*badge.Points),
			Type:        *badge.Type,
			Status:      *badge.Status,
		}

		err := schemaBadge.Insert(s.db, boil.Infer())
		if err != nil {
			return nil, err
		}

		insertedBadges = append(insertedBadges, schemaBadge)
	}

	return insertedBadges, nil
}

func (s *DBStore) GetTypeDefaultBadges() (schema.BadgeSlice, error) {
	return schema.Badges(
		qm.Where("type = 'default'"),
	).All(s.db)
}

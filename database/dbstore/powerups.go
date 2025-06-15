package dbstore

import (
	"github.com/itemslabs/clubz-api/database"
	"github.com/itemslabs/clubz-api/database/schema"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func (s *DBStore) GetPowerUpByID(id int) (*schema.Powerup, error) {
	return schema.Powerups(
		qm.Where("enabled = ?", true),
		qm.Where("id = ?", id),
	).One(s.db)
}

func (s *DBStore) GetPowerUps() (schema.PowerupSlice, error) {
	return schema.Powerups(
		qm.Where("enabled = ?", true),
		qm.OrderBy("case when id = 7 then -1 else id end"),
	).All(s.db)
}

func (s *DBStore) GetPowerUpActions() (schema.PowerupActionSlice, error) {
	return schema.PowerupActions(
		qm.InnerJoin("actions on powerup_actions.action_id = actions.id"),
		qm.OrderBy("actions.score desc, actions.name"),
		qm.Load("Action"),
	).All(s.db)
}

func (s *DBStore) GetSportSubstitutionPowerUp(sportID string) (*schema.Powerup, error) {
	const powerupName = "substitution"
	return schema.Powerups(
		qm.Where("type = ?", database.PowerUpTypeSpecial),
		qm.Where("name ilike ?", powerupName),
		qm.Where("enabled = ?", true),
		qm.Where("sport_id = ?", sportID),
	).One(s.db)
}

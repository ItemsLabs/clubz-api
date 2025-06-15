package dbstore

import (
	"github.com/itemslabs/clubz-api/database/schema"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func (s *DBStore) GetSportByID(id string) (*schema.Sport, error) {
	return schema.Sports(qm.Where("id = ?", id)).One(s.db)
}

func (s *DBStore) GetSports() (schema.SportSlice, error) {
	return schema.Sports().All(s.db)
}

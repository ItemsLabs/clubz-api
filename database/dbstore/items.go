package dbstore

import (
	"github.com/gameon-app-inc/laliga-matchfantasy-api/database/schema"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func (s *DBStore) GetItemByID(id string) (*schema.Item, error) {
	return schema.Items(
		qm.Where("id = ?", id),
	).One(s.db)
}

func (s *DBStore) GetItems() (schema.ItemSlice, error) {
	return schema.Items(
		qm.OrderBy("close_date_at desc"),
	).All(s.db)
}

package dbstore

import (
	"github.com/itemslabs/clubz-api/database/schema"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func (s *DBStore) FindTransaction(matchID, userID string) (*schema.Transaction, error) {
	return schema.Transactions(
		qm.Where("match_id = ?", matchID),
		qm.Where("user_id = ?", userID),
	).One(s.db)
}

func (s *DBStore) GetTransactionByID(id string) (*schema.Transaction, error) {
	return schema.Transactions(
		qm.Where("id = ?", id),
	).One(s.db)
}

func (s *DBStore) CreateTransaction(order *schema.Transaction) (*schema.Transaction, error) {
	return order, order.Insert(s.db, boil.Infer())
}

func (s *DBStore) FindRewardTransaction(matchID, userID string) (*schema.Transaction, error) {
	return schema.Transactions(
		qm.Where("match_id = ?", matchID),
		qm.Where("user_id = ?", userID),
		qm.Where("text ilike ?", `reward for % position in match % vs %`),
	).One(s.db)
}

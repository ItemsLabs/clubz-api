package dbstore

import (
	"github.com/itemslabs/clubz-api/database/schema"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func (s *DBStore) GetOrderByID(id string) (*schema.Order, error) {
	return schema.Orders(
		qm.Where("id = ?", id),
	).One(s.db)
}

func (s *DBStore) CreateOrder(order *schema.Order) (*schema.Order, error) {
	return order, order.Insert(s.db, boil.Infer())
}

func (s *DBStore) UpdateOrder(order *schema.Order) error {
	_, err := order.Update(s.db, boil.Whitelist("payment_platform_uuid", "payment_platform_status", "blockchain_order_status", "blockchain_uuid", "delivered", "purchased_at", "payment_platform_url"))
	return err
}

func (s *DBStore) LockOrder(orderID string) (*schema.Order, error) {
	return schema.Orders(
		qm.Where("id = ?", orderID),
		qm.For("update"),
	).One(s.db)
}

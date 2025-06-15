package dbstore

import (
	"database/sql"
	"errors"
	"time"

	"github.com/itemslabs/clubz-api/database/schema"
	"github.com/itemslabs/clubz-api/routes/model"
	"github.com/google/uuid"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func (s *DBStore) GetStoreProducts() (schema.StoreProductSlice, error) {
	return schema.StoreProducts().All(s.db)
}

func (s *DBStore) GetStoreProductByProductID(productID string) (*schema.StoreProduct, error) {
	return schema.StoreProducts(qm.Where("store_product_id = ?", productID)).One(s.db)
}

func (s *DBStore) GetStoreProductByStoreAndProductID(store, productID string) (*schema.StoreProduct, error) {
	switch store {
	case model.RevenueCatPurchaseRequestEventStoreAPPSTORE:
		return schema.StoreProducts(qm.Where("apple_product_id = ?", productID)).One(s.db)
	case model.RevenueCatPurchaseRequestEventStorePLAYSTORE:
		return schema.StoreProducts(qm.Where("google_product_id = ?", productID)).One(s.db)
	default:
		return nil, errors.New("invalid store")
	}
}

func (s *DBStore) CreateStoreProductTransaction(
	externalTransactionID, originStore, storeProductID string,
	purchaseDate time.Time,
) (string, error) {
	var pt schema.StoreProductTransaction
	pt.ID = uuid.NewString()
	pt.OriginStore = null.StringFrom(originStore)
	pt.ProductID = storeProductID
	pt.ExternalTransactionID = null.StringFrom(externalTransactionID)
	pt.CreatedAt = purchaseDate
	pt.Initiated = true
	pt.InitiatedAt = null.TimeFrom(time.Now())
	if err := pt.Insert(s.db, boil.Infer()); err != nil {
		return "", err
	}
	return pt.ID, nil
}
func (s *DBStore) GetStoreProductTransactionByExternalID(externalTransactionID string) (
	*schema.StoreProductTransaction,
	error,
) {
	if pt, err := schema.StoreProductTransactions(
		qm.Where("external_transaction_id = ?", externalTransactionID),
		qm.Where("initiated = ?", true),
		qm.Where("confirmed = ?", false),
	).One(s.db); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("Transaction does not exist")
		}
		return nil, errors.New("Transaction has an error")
	} else {
		return pt, nil
	}
}

func (s *DBStore) ConfirmStoreProductTransactionByExternalID(userID, transactionID, externalTransactionID string) error {
	if pt, err := schema.StoreProductTransactions(
		qm.Where("external_transaction_id = ?", externalTransactionID),
		qm.Where("initiated = ?", true),
		qm.Where("confirmed = ?", false),
	).One(s.db); err != nil {
		if err == sql.ErrNoRows {
			return errors.New("Transaction does not exist")
		}
		return errors.New("Transaction has an error")
	} else {
		pt.TransactionID = null.StringFrom(transactionID)
		pt.Confirmed = true
		pt.ConfirmedAt = null.TimeFrom(time.Now())
		pt.UpdatedAt = time.Now()
		pt.UserID = null.StringFrom(userID)
		if _, err := pt.Update(s.db, boil.Infer()); err != nil {
			return err
		}
	}
	return nil
}

func (s *DBStore) CancelStoreProductTransaction(externalTransactionID string) error {
	if pt, err := schema.StoreProductTransactions(
		qm.Where("external_transaction_id = ?", externalTransactionID),
	).One(s.db); err != nil {
		if err == sql.ErrNoRows {
			return errors.New("Transaction does not exist")
		}
		return errors.New("Transaction has an error")
	} else {
		if _, err := pt.Delete(s.db); err != nil {
			return err
		}
	}
	return nil
}

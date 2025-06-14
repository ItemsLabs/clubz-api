package dbstore

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/gameon-app-inc/laliga-matchfantasy-api/database/schema"
	"github.com/google/uuid"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

// CreateCardPack creates a new card pack in the database.
func (s *DBStore) CreateCardPack(cp *schema.CardPackType) error {
	if cp == nil {
		return fmt.Errorf("CreateCardPack: the provided CardPackType is nil")
	}

	cp.ID = uuid.New().String()
	cp.CreatedAt = time.Now()
	cp.UpdatedAt = time.Now()

	if err := cp.Insert(s.db, boil.Infer()); err != nil {
		return fmt.Errorf("CreateCardPack: failed to insert card pack: %w", err)
	}

	return nil
}

// GetCardPackByID retrieves a card pack by its ID.
func (s *DBStore) GetCardPackByID(id string) (*schema.CardPackType, error) {
	cp, err := schema.FindCardPackType(s.db, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("GetCardPackByID: no card pack found with ID %s", id)
		}
		return nil, fmt.Errorf("GetCardPackByID: failed to find card pack: %w", err)
	}

	return cp, nil
}

// GetAllCardPacks retrieves all card packs from the database.
func (s *DBStore) GetAllCardPacks() ([]*schema.CardPackType, error) {
	cps, err := schema.CardPackTypes().All(s.db)
	if err != nil {
		return nil, fmt.Errorf("GetAllCardPacks: failed to retrieve all card packs: %w", err)
	}

	return cps, nil
}

// UpdateCardPack updates an existing card pack.
func (s *DBStore) UpdateCardPack(cp *schema.CardPackType) error {
	if cp == nil {
		return fmt.Errorf("UpdateCardPack: the provided CardPackType is nil")
	}

	cp.UpdatedAt = time.Now()

	_, err := cp.Update(s.db, boil.Whitelist("name", "description", "updated_at"))
	if err != nil {
		return fmt.Errorf("UpdateCardPack: failed to update card pack: %w", err)
	}

	return nil
}

// DeleteCardPack deletes a card pack by its ID.
func (s *DBStore) DeleteCardPack(id string) error {
	cp, err := schema.FindCardPackType(s.db, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("DeleteCardPack: no card pack found with ID %s", id)
		}
		return fmt.Errorf("DeleteCardPack: failed to find card pack for deletion: %w", err)
	}

	_, err = cp.Delete(s.db)
	if err != nil {
		return fmt.Errorf("DeleteCardPack: failed to delete card pack: %w", err)
	}

	return nil
}

// GetAssignedCardPacksByType retrieves all assigned card packs for a given card pack type.
func (s *DBStore) GetAssignedCardPacksByType(cardPackTypeID string) ([]*schema.AssignedCardPack, error) {
	cps, err := schema.AssignedCardPacks(
		qm.Where("card_pack_type_id = ?", cardPackTypeID),
	).All(s.db)
	if err != nil {
		return nil, fmt.Errorf("GetAssignedCardPacksByType: failed to retrieve assigned card packs: %w", err)
	}

	return cps, nil
}

// GetCardPackTypeByName retrieves a card pack type by its name.
func (s *DBStore) GetCardPackTypeByName(name string) (*schema.CardPackType, error) {
	return schema.CardPackTypes(qm.Where("name = ?", name)).One(s.db)
}

// GetCardPackTypeByCode retrieves a card pack type by its code.
func (s *DBStore) GetCardPackTypeByCode(code string) (*schema.CardPackType, error) {
	return schema.CardPackTypes(qm.Where("card_pack_code = ?", code)).One(s.db)
}

func (s *DBStore) DeductFromPackLimit(cardPackTypeID string, amount int) error {
	cp, err := schema.FindCardPackType(s.db, cardPackTypeID)
	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("DeductFromPackLimit: no card pack type found with ID %s", cardPackTypeID)
		}
		return fmt.Errorf("DeductFromPackLimit: failed to find card pack type: %w", err)
	}

	newLimit := cp.PackLimits.Int - amount
	if newLimit < 0 {
		return fmt.Errorf("DeductFromPackLimit: deduction results in a negative pack limit")
	}

	cp.PackLimits = null.IntFrom(newLimit)
	cp.UpdatedAt = time.Now()

	_, err = cp.Update(s.db, boil.Whitelist("pack_limits", "updated_at"))
	if err != nil {
		return fmt.Errorf("DeductFromPackLimit: failed to update pack limit: %w", err)
	}

	return nil
}

func (s *DBStore) RestockPackLimit(cardPackTypeID string, amount int) error {
	cp, err := schema.FindCardPackType(s.db, cardPackTypeID)
	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("RestockPackLimit: no card pack type found with ID %s", cardPackTypeID)
		}
		return fmt.Errorf("RestockPackLimit: failed to find card pack type: %w", err)
	}
	newLimit := cp.PackLimits.Int - +amount
	cp.PackLimits = null.IntFrom(newLimit)
	cp.UpdatedAt = time.Now()
	_, err = cp.Update(s.db, boil.Whitelist("pack_limits", "updated_at"))
	if err != nil {
		return fmt.Errorf("RestockPackLimit: failed to update pack limit: %w", err)
	}

	return nil
}

// GetPackLimitsByCardPackName retrieves the pack limits for a given card pack name.
func (s *DBStore) GetPackLimitsByCardPackName(name string) (int, error) {
	cp, err := schema.CardPackTypes(qm.Where("name = ?", name)).One(s.db)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, fmt.Errorf("GetPackLimitsByCardPackName: no card pack type found with name %s", name)
		}
		return 0, fmt.Errorf("GetPackLimitsByCardPackName: failed to find card pack type: %w", err)
	}

	return cp.PackLimits.Int, nil
}

// GetPackLimitsByCardPackCode retrieves the pack limits for a given card pack code.
func (s *DBStore) GetPackLimitsByCardPackCode(code string) (int, error) {
	cp, err := schema.CardPackTypes(qm.Where("card_pack_code = ?", code)).One(s.db)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, fmt.Errorf("GetPackLimitsByCardPackCode: no card pack type found with code %s", code)
		}
		return 0, fmt.Errorf("GetPackLimitsByCardPackCode: failed to find card pack type: %w", err)
	}

	return cp.PackLimits.Int, nil
}

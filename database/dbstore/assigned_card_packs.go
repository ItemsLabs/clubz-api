package dbstore

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/gameon-app-inc/laliga-matchfantasy-api/database/schema"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

// CreateAssignedCardPack creates a new assigned card pack in the database.
func (s *DBStore) CreateAssignedCardPack(acp *schema.AssignedCardPack) (*schema.AssignedCardPack, error) {
	boil.DebugMode = true
	if err := acp.Insert(s.db, boil.Infer()); err != nil {
		log.Printf("Error creating assigned card pack: %v", err)
		log.Printf("Data: %+v", acp) // Logs the entire data structure
		return nil, fmt.Errorf("failed to insert assigned card pack: %w", err)
	}
	boil.DebugMode = false
	return acp, nil
}

// GetAssignedCardPackByID retrieves an assigned card pack by its ID.
func (s *DBStore) GetAssignedCardPackByID(id string) (*schema.AssignedCardPack, error) {
	acp, err := schema.FindAssignedCardPack(s.db, id)
	if err != nil {
		log.Printf("Error finding assigned card pack by ID: %v", err)
		return nil, fmt.Errorf("failed to find assigned card pack with ID %s: %w", id, err)
	}
	return acp, nil
}

// GetAllAssignedCardPacks retrieves all assigned card packs from the database.
func (s *DBStore) GetAllAssignedCardPacks() (schema.AssignedCardPackSlice, error) {
	acps, err := schema.AssignedCardPacks().All(s.db)
	if err != nil {
		log.Printf("Error retrieving all assigned card packs: %v", err)
		return nil, fmt.Errorf("failed to retrieve all assigned card packs: %w", err)
	}
	return acps, nil
}

// UpdateAssignedCardPack updates an existing assigned card pack.
func (s *DBStore) UpdateAssignedCardPack(acp *schema.AssignedCardPack) error {
	_, err := acp.Update(s.db, boil.Whitelist("opened", "card_pack_type_id", "user_id", "updated_at"))
	if err != nil {
		log.Printf("Error updating assigned card pack: %v", err)
		return fmt.Errorf("failed to update assigned card pack: %w", err)
	}
	return nil
}

// DeleteAssignedCardPack deletes an assigned card pack by its ID.
func (s *DBStore) DeleteAssignedCardPack(id string) error {
	acp, err := schema.FindAssignedCardPack(s.db, id)
	if err != nil {
		log.Printf("Error finding assigned card pack for deletion: %v", err)
		return fmt.Errorf("failed to find assigned card pack for deletion with ID %s: %w", id, err)
	}
	_, err = acp.Delete(s.db)
	if err != nil {
		log.Printf("Error deleting assigned card pack: %v", err)
		return fmt.Errorf("failed to delete assigned card pack with ID %s: %w", id, err)
	}
	return nil
}

// GetAssignedCardPacksByUserID retrieves all assigned card packs for a given user ID.
func (s *DBStore) GetAssignedCardPacksByUserID(userID string) (schema.AssignedCardPackSlice, error) {
	acps, err := schema.AssignedCardPacks(
		qm.Where("user_id = ?", userID),
		qm.Load("CardPackType"),
	).All(s.db)
	if err != nil {
		log.Printf("Error retrieving assigned card packs by user ID: %v", err)
		return nil, fmt.Errorf("failed to retrieve assigned card packs for user ID %s: %w", userID, err)
	}
	return acps, nil
}

// GetAssignedCardPacksByCardPackTypeID retrieves all assigned card packs for a given card pack type ID.
func (s *DBStore) GetAssignedCardPacksByCardPackTypeID(cardPackTypeID string) (schema.AssignedCardPackSlice, error) {
	acps, err := schema.AssignedCardPacks(
		qm.Where("card_pack_type_id = ?", cardPackTypeID),
	).All(s.db)
	if err != nil {
		log.Printf("Error retrieving assigned card packs by card pack type ID: %v", err)
		return nil, fmt.Errorf("failed to retrieve assigned card packs for card pack type ID %s: %w", cardPackTypeID, err)
	}
	return acps, nil
}

// OpenAssignedCardPack marks an assigned card pack as opened.
func (s *DBStore) OpenAssignedCardPack(id string) error {
	acp, err := schema.FindAssignedCardPack(s.db, id)
	if err != nil {
		log.Printf("Error finding assigned card pack to open: %v", err)
		return fmt.Errorf("failed to find assigned card pack to open with ID %s: %w", id, err)
	}

	acp.Opened = true
	_, err = acp.Update(s.db, boil.Whitelist("opened", "updated_at"))
	if err != nil {
		log.Printf("Error marking assigned card pack as opened: %v", err)
		return fmt.Errorf("failed to mark assigned card pack as opened with ID %s: %w", id, err)
	}
	return nil
}

// FilterAssignedCardPacksByDateRange retrieves assigned card packs within a specific date range.
func (s *DBStore) FilterAssignedCardPacksByDateRange(start, end time.Time) (schema.AssignedCardPackSlice, error) {
	acps, err := schema.AssignedCardPacks(
		qm.Where("created_at >= ? AND created_at <= ?", start, end),
	).All(s.db)
	if err != nil {
		log.Printf("Error filtering assigned card packs by date range: %v", err)
		return nil, fmt.Errorf("failed to filter assigned card packs by date range: %w", err)
	}
	return acps, nil
}

// CountUnopenedAssignedCardPacks counts all unopened assigned card packs for a given user.
func (s *DBStore) CountUnopenedAssignedCardPacks(userID string) (int64, error) {
	count, err := schema.AssignedCardPacks(
		qm.Where("user_id = ? AND opened = false", userID),
	).Count(s.db)
	if err != nil {
		log.Printf("Error counting unopened assigned card packs: %v", err)
		return 0, fmt.Errorf("failed to count unopened assigned card packs: %w", err)
	}
	return count, nil
}

// GetRecentAssignedCardPacks retrieves the most recently created or updated assigned card packs.
func (s *DBStore) GetRecentAssignedCardPacks(limit int) (schema.AssignedCardPackSlice, error) {
	acps, err := schema.AssignedCardPacks(
		qm.OrderBy("updated_at DESC"),
		qm.Limit(limit),
	).All(s.db)
	if err != nil {
		log.Printf("Error retrieving recent assigned card packs: %v", err)
		return nil, fmt.Errorf("failed to retrieve recent assigned card packs: %w", err)
	}
	return acps, nil
}

func (s *DBStore) BulkUpdateOpenedStatus(ids []string, opened bool) error {
	tx, err := s.db.Begin()
	if err != nil {
		log.Printf("Error starting transaction: %v", err)
		return fmt.Errorf("failed to start transaction: %w", err)
	}

	// Convert []string to []interface{}
	idsInterface := make([]interface{}, len(ids))
	for i, id := range ids {
		idsInterface[i] = id
	}

	_, err = schema.AssignedCardPacks(
		qm.WhereIn("id in ?", idsInterface...),
	).UpdateAll(tx, map[string]interface{}{
		schema.AssignedCardPackColumns.Opened: opened,
	})
	if err != nil {
		tx.Rollback()
		log.Printf("Error in bulk updating opened status: %v", err)
		return fmt.Errorf("failed to bulk update opened status: %w", err)
	}

	if err := tx.Commit(); err != nil {
		log.Printf("Error committing transaction: %v", err)
		return fmt.Errorf("failed to commit transaction: %w", err)
	}
	return nil
}

// DeleteExpiredAssignedCardPacks deletes assigned card packs that were created before a certain date.
func (s *DBStore) DeleteExpiredAssignedCardPacks(expiryDate time.Time) (int64, error) {
	count, err := schema.AssignedCardPacks(
		qm.Where("created_at < ?", expiryDate),
	).DeleteAll(s.db)
	if err != nil {
		log.Printf("Error deleting expired assigned card packs: %v", err)
		return 0, fmt.Errorf("failed to delete expired assigned card packs: %w", err)
	}
	return count, nil
}

// GetAssignedCardPackStatistics generates statistics about assigned card packs.
func (s *DBStore) GetAssignedCardPackStatistics() (*AssignedCardPackStatistics, error) {
	// Example implementation - adjust queries as needed for your statistics requirements
	type statsResult struct {
		TotalCount    int64 `boil:"total_count"`
		OpenedCount   int64 `boil:"opened_count"`
		UnopenedCount int64 `boil:"unopened_count"`
	}
	var result statsResult

	query := `
	SELECT
	    COUNT(*) as total_count,
	    COUNT(*) FILTER (WHERE opened = true) as opened_count,
	    COUNT(*) FILTER (WHERE opened = false) as unopened_count
	FROM assigned_card_packs
	`
	err := s.db.QueryRow(query).Scan(&result.TotalCount, &result.OpenedCount, &result.UnopenedCount)
	if err != nil {
		log.Printf("Error getting assigned card pack statistics: %v", err)
		return nil, fmt.Errorf("failed to get assigned card pack statistics: %w", err)
	}

	stats := &AssignedCardPackStatistics{
		TotalCount:    result.TotalCount,
		OpenedCount:   result.OpenedCount,
		UnopenedCount: result.UnopenedCount,
	}

	return stats, nil
}

// AssignedCardPackStatistics struct to hold statistics data
type AssignedCardPackStatistics struct {
	TotalCount    int64
	OpenedCount   int64
	UnopenedCount int64
}

// SetAssignedCardPackRevealed sets the Revealed field to true.
func (s *DBStore) SetAssignedCardPackRevealed(id string) error {
	acp, err := schema.FindAssignedCardPack(s.db, id)
	if err != nil {
		log.Printf("Error finding assigned card pack to reveal: %v", err)
		return fmt.Errorf("failed to find assigned card pack to reveal with ID %s: %w", id, err)
	}

	acp.Revealed = true
	acp.RevealedAt = null.TimeFrom(time.Now())
	_, err = acp.Update(s.db, boil.Whitelist("revealed", "revealed_at", "updated_at"))
	if err != nil {
		log.Printf("Error setting assigned card pack as revealed: %v", err)
		return fmt.Errorf("failed to set assigned card pack as revealed with ID %s: %w", id, err)
	}
	return nil
}

// SetAssignedCardPackOpened sets the Opened field to true.
func (s *DBStore) SetAssignedCardPackOpened(id string) error {
	acp, err := schema.FindAssignedCardPack(s.db, id)
	if err != nil {
		log.Printf("Error finding assigned card pack to open: %v", err)
		return fmt.Errorf("failed to find assigned card pack to open with ID %s: %w", id, err)
	}

	acp.Opened = true
	acp.OpenedAt = null.TimeFrom(time.Now())
	_, err = acp.Update(s.db, boil.Whitelist("opened", "opened_at", "updated_at"))
	if err != nil {
		log.Printf("Error setting assigned card pack as opened: %v", err)
		return fmt.Errorf("failed to set assigned card pack as opened with ID %s: %w", id, err)
	}
	return nil
}

// UpdateAssignedCardPackCardIds updates the CardIds field of an assigned card pack.
func (s *DBStore) UpdateAssignedCardPackCardIds(id string, cardIds []string) error {
	acp, err := schema.FindAssignedCardPack(s.db, id)
	if err != nil {
		log.Printf("Error finding assigned card pack to update card IDs: %v", err)
		return fmt.Errorf("failed to find assigned card pack to update card IDs with ID %s: %w", id, err)
	}

	cardIdsJSON, err := json.Marshal(cardIds)
	if err != nil {
		log.Printf("Error marshaling card IDs to JSON: %v", err)
		return fmt.Errorf("failed to marshal card IDs to JSON: %w", err)
	}

	acp.CardIds = null.JSONFrom(cardIdsJSON)
	_, err = acp.Update(s.db, boil.Whitelist("card_ids", "updated_at"))
	if err != nil {
		log.Printf("Error updating card IDs of assigned card pack: %v", err)
		return fmt.Errorf("failed to update card IDs of assigned card pack with ID %s: %w", id, err)
	}
	return nil
}

func (s *DBStore) GetAssignedCardPackByStoreProductTransactionID(storeTransactionID string) (*schema.AssignedCardPack, error) {
	acp, err := schema.AssignedCardPacks(
		qm.Where("assigned_card_packs.store_transaction_id = ?", storeTransactionID),
	).One(s.db)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("Transaction does not exist")
		}
		return nil, errors.New("Transaction has an error")
	}
	return acp, nil
}

func (s *DBStore) RefundAssignedCardPackByID(id string) error {
	acp, err := schema.FindAssignedCardPack(s.db, id)
	if err != nil {
		log.Printf("Error finding assigned card pack to refund: %v", err)
		return fmt.Errorf("failed to find assigned card pack to refund with ID %s: %w", id, err)
	}

	// Update the assigned card pack's field "refunded" to true
	acp.Refunded = true
	if _, err := acp.Update(s.db, boil.Whitelist("refunded", "updated_at")); err != nil {
		log.Printf("Error refunding assigned card pack: %v", err)
		return fmt.Errorf("failed to refund assigned card pack with ID %s: %w", id, err)
	}

	return nil
}

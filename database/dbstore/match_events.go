package dbstore

import (
	"github.com/itemslabs/clubz-api/database/schema"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

// GetMatchEventsByMatchID retrieves all match events for a specific match ID.
func (s *DBStore) GetMatchEventsByMatchID(matchID string) (schema.MatchEventSlice, error) {
	return schema.MatchEvents(
		qm.Where("match_id = ?", matchID),
		qm.OrderBy("timestamp"),
	).All(s.db)
}

// GetMatchEventsByMatchIDAndType retrieves match events for a specific match ID and event type.
func (s *DBStore) GetMatchEventsByMatchIDAndType(matchID string, eventType int) (schema.MatchEventSlice, error) {
	return schema.MatchEvents(
		qm.Where("match_id = ? AND type = ?", matchID, eventType),
		qm.OrderBy("timestamp"),
	).All(s.db)
}

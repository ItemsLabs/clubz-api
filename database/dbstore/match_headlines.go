package dbstore

import (
	"github.com/itemslabs/clubz-api/database/schema"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func (s *DBStore) GetMatchHeadlines(matchID string, screenType int) (schema.MatchHeadlineSlice, error) {
	return schema.MatchHeadlines(
		qm.Where("match_id = ?", matchID),
		qm.Where("screen_type = ?", screenType),
	).All(s.db)
}

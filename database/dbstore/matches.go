package dbstore

import (
	"time"

	"github.com/gameon-app-inc/laliga-matchfantasy-api/database"
	"github.com/gameon-app-inc/laliga-matchfantasy-api/database/schema"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func (s *DBStore) GetMatchesInPeriod(from, to time.Time) (schema.MatchSlice, error) {
	return schema.Matches(
		qm.InnerJoin(schema.TableNames.Competitions+" c on c.id = matches.competition_id AND c.enabled = TRUE"),
		qm.Where("match_time between ? and ?", from, to),
		qm.Where("match_type != ?", database.MatchTypeUnknown),
		qm.Where(schema.TableNames.Matches+".enabled = ?", true),
		qm.OrderBy("match_time"),
		qm.Load("HomeTeam"),
		qm.Load("AwayTeam"),
		qm.Load("Competition"),
		qm.Load("MatchRewards"),
	).All(s.db)
}

func (s *DBStore) GetMatchesInGameWeek(gameWeek int) (schema.MatchSlice, error) {
	return schema.Matches(
		qm.InnerJoin(schema.TableNames.Competitions+" c on c.id = matches.competition_id AND c.enabled = TRUE"),
		qm.Where("week = ?", gameWeek),
		qm.Where("match_type != ?", database.MatchTypeUnknown),
		qm.Where(schema.TableNames.Matches+".enabled = ?", true),
		qm.OrderBy("match_time"),
		qm.Load("HomeTeam"),
		qm.Load("AwayTeam"),
		qm.Load("Competition"),
		qm.Load("MatchRewards"),
	).All(s.db)
}

func (s *DBStore) GetNextActiveMatch(t time.Time) (*schema.Match, error) {
	return schema.Matches(
		qm.InnerJoin(schema.TableNames.Competitions+" c on c.id = matches.competition_id AND c.enabled = TRUE"),
		qm.Where("match_time > ?", t),
		qm.Where("match_type != ?", database.MatchTypeUnknown),
		qm.Where(schema.TableNames.Matches+".enabled = ?", true),
		qm.OrderBy("match_time asc"),
		qm.Limit(1),
		qm.Load("HomeTeam"),
		qm.Load("AwayTeam"),
		qm.Load("Competition"),
		qm.Load("MatchRewards"),
	).One(s.db)
}

func (s *DBStore) GetMatchSquad(matchID string) (schema.MatchPlayerSlice, error) {
	return schema.MatchPlayers(
		qm.Where("match_id = ?"),
		qm.Load("Player"),
		qm.Load("Player.Team"),
	).All(s.db)
}

func (s *DBStore) GetMatchByID(matchID string) (*schema.Match, error) {
	return schema.Matches(
		qm.Where("id = ?", matchID),
		qm.Load("HomeTeam"),
		qm.Load("AwayTeam"),
		qm.Load("Competition"),
		qm.Load("MatchRewards"),
	).One(s.db)
}

func (s *DBStore) GetPlayerCount(matchID string) (int64, error) {
	cnt, err := schema.Games(
		qm.Where("match_id = ?", matchID),
	).Count(s.db)
	return cnt, err
}

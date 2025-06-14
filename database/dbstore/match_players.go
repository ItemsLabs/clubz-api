package dbstore

import (
	"github.com/gameon-app-inc/laliga-matchfantasy-api/database/schema"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func (s *DBStore) GetMatchPlayers(matchID string) (schema.MatchPlayerSlice, error) {
	return schema.MatchPlayers(
		qm.Where("match_id = ?", matchID),
		qm.OrderBy("avg_score desc nulls last"),
		qm.Load("Player"),
		qm.Load("Team"),
	).All(s.db)
}

func (s *DBStore) GetMatchPlayersPPG(matchID string) (map[string]float64, error) {
	matches, err := s.db.Query(`
SELECT id, ppg -- , import_id, full_name, total_goals
    FROM player_ppg
WHERE id IN
    (SELECT mp.player_id FROM matches m JOIN match_players mp ON mp.match_id = m.id where m.id = $1)`, matchID)
	if err != nil {
		return nil, err
	}
	defer matches.Close()
	ppgMap := make(map[string]float64)
	for matches.Next() {
		var id string
		// var importID string
		// var fullName string
		var ppg float64
		// var totalGoals int
		// if err := matches.Scan(&id, &importID, &fullName, &ppg, &totalGoals); err != nil {
		if err := matches.Scan(&id, &ppg); err != nil {
			return nil, err
		}
		ppgMap[id] = ppg
	}
	return ppgMap, nil
}

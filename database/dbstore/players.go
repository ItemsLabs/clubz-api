package dbstore

import (
	"fmt"
	"time"

	"github.com/gameon-app-inc/laliga-matchfantasy-api/database"
	"github.com/gameon-app-inc/laliga-matchfantasy-api/database/schema"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func (s *DBStore) GetPlayerByID(playerID string) (*schema.Player, error) {
	return schema.Players(
		qm.Where("id = ?", playerID),
	).One(s.db)
}

func (s *DBStore) GetLastMatchPlayers(playerID string, limit int, mods ...qm.QueryMod) (schema.MatchPlayerSlice, error) {
	mods = append(mods,
		qm.Where(`
exists(
select null
 from match_events mev
where mev.player_id = match_players.player_id
  and mev.match_id = match_players.match_id
)
`),
		qm.Where("player_id = ?", playerID),
		qm.Where("score is not null"),
		qm.InnerJoin("matches on matches.id = match_players.match_id"),
		qm.OrderBy("matches.match_time desc"),
		qm.Limit(limit),
	)

	return schema.MatchPlayers(mods...).All(s.db)
}

func (s *DBStore) GetAveragePointsDistribution(playerID string, matches []string) ([]*database.PointBucket, error) {
	query := fmt.Sprintf(`
with t as (
    select match_id,
           case
               when minute >= 0 and minute < 16 then '0-15'
               when minute >= 16 and minute < 31 then '16-30'
               when minute >= 31 and minute < 46 then '31-45'
               when minute >= 46 and minute < 61 then '46-60'
               when minute >= 61 and minute < 76 then '61-75'
               when minute >= 76 and minute <= 90 then '76-90'
               else 'none' end bucket_key,
           sum(points)         minute_points
    from match_events
    where status = $1
      and player_id = $2
      and match_id in (%s)
      and (points is not null and points != 0)
    group by match_id,
             case
                 when minute >= 0 and minute < 16 then '0-15'
                 when minute >= 16 and minute < 31 then '16-30'
                 when minute >= 31 and minute < 46 then '31-45'
                 when minute >= 46 and minute < 61 then '46-60'
                 when minute >= 61 and minute < 76 then '61-75'
                 when minute >= 76 and minute <= 90 then '76-90'
                 else 'none' end
    order by match_id, bucket_key
)
select low,
       high,
       avg(minute_points) points
from (values (0, 15, '0-15'),
             (16, 30, '16-30'),
             (31, 45, '31-45'),
             (46, 60, '46-60'),
             (61, 75, '61-75'),
             (76, 90, '76-90')) as v(low, high, key)
         left join t on v.key = t.bucket_key
group by key, low, high
order by low`, GenerateSubstituteParams(3, 2+len(matches)))

	args := make([]interface{}, 0, len(matches)+2)
	args = append(args, database.MatchEventStatusActive, playerID)
	for _, el := range matches {
		args = append(args, el)
	}
	var pointsInfo []*database.PointBucket

	if err := queries.Raw(query, args...).Bind(nil, s.db, &pointsInfo); err != nil {
		return nil, err
	}

	return pointsInfo, nil
}

func (s *DBStore) GetPercentOfPicks(matchID, playerID string) (float64, error) {
	query := `
select (select count(distinct games.id)
          from game_picks gp,
               games
         where gp.player_id = $1
           and gp.game_id = games.id
           and games.match_id = $2
       )::float /
       (select count(*) from games where match_id = $2)::float as value`

	var pickPercent struct {
		Value float64
	}

	if err := queries.Raw(query, playerID, matchID).Bind(nil, s.db, &pickPercent); err != nil {
		return 0, err
	}

	return pickPercent.Value * 100, nil
}

func (s *DBStore) GetPointsInInterval(matchID, playerID string, startTime, endTime time.Time) (int, error) {
	query := `
select sum(points)
  from match_events
 where match_id = $1
   and player_id = $2
   and status = $3
   and created_at between $4 and $5
   and points is not null
`

	var points struct {
		Value int
	}

	if err := queries.Raw(query, matchID, playerID, database.MatchEventStatusActive, startTime, endTime).Bind(nil, s.db, &points); err != nil {
		return 0, err
	}

	return points.Value, nil
}

func (s *DBStore) GetActionSummary(matchID, playerID string) ([]*database.ActionSummary, error) {
	query := `
select act.name,
       count(*) as "count",
       sum(mev.points) as "points"
  from match_events mev,
       actions act
 where mev.match_id = $1
   and mev.player_id = $2
   and mev.status = $3
   and mev.type = act.id
 group by act.id, act.name
 order by act.name
`
	var actions []*database.ActionSummary

	if err := queries.Raw(query, matchID, playerID, database.MatchEventStatusActive).Bind(nil, s.db, &actions); err != nil {
		return nil, err
	}

	return actions, nil
}

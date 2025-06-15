package dbstore

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/itemslabs/clubz-api/database"
	"github.com/itemslabs/clubz-api/database/schema"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func (s *DBStore) LockGame(gameID string) (*schema.Game, error) {
	return schema.Games(
		qm.Where("id = ?", gameID),
		qm.For("update"),
	).One(s.db)
}

func (s *DBStore) GetGamesLaterThan(userID string, from time.Time) (schema.GameSlice, error) {
	return schema.Games(
		qm.InnerJoin("matches on matches.id = games.match_id"),
		qm.Where("matches.match_time >= ?", from),
		qm.Where("games.user_id = ?", userID),
		qm.OrderBy("games.created_at desc"),
		qm.Load("Match.MatchPlayers"),
		qm.Load("GamePicks"),
		qm.Load("GamePicks.Player"),
		qm.Load("GamePowerups.Powerup"),
		qm.Load("User"),
	).All(s.db)
}

func (s *DBStore) GetGameByID(id string, userID string) (*schema.Game, error) {
	return schema.Games(
		qm.Where("id = ?", id),
		qm.Where("user_id = ?", userID),
		qm.Load("Match.MatchPlayers"),
		qm.Load("GamePicks"),
		qm.Load("GamePicks.Player"),
		qm.Load("GamePowerups.Powerup"),
		qm.Load("User"),
	).One(s.db)
}

func (s *DBStore) GetGameByUserIDMatchID(userID, matchID string) (*schema.Game, error) {
	return schema.Games(
		qm.Where("user_id = ?", userID),
		qm.Where("match_id = ?", matchID),
		qm.Load("Match.MatchPlayers"),
		qm.Load("GamePicks"),
		qm.Load("GamePicks.Player"),
		qm.Load("GamePowerups.Powerup"),
		qm.Load("User"),
	).One(s.db)
}

func (s *DBStore) GetActiveGameIDForMatch(matchID, userID string) (string, error) {
	game, err := schema.Games(
		qm.Where("match_id = ?", matchID),
		qm.Where("user_id = ?", userID),
	).One(s.db)
	if err != nil {
		return "", err
	}

	return game.ID, nil
}

func (s *DBStore) IsGameBelongToUser(gameID, userID string) (bool, error) {
	cnt, err := schema.Games(
		qm.Where("id = ?", gameID),
		qm.Where("user_id = ?", userID),
	).Count(s.db)
	if err != nil {
		return false, err
	}

	return cnt > 0, nil
}

func (s *DBStore) CreateGame(game *schema.Game) (*schema.Game, error) {
	return game, game.Insert(s.db, boil.Infer())
}

func (s *DBStore) GetNumberOfGames(userID string) (int, error) {
	query := `select coalesce(max(num),0) from games where user_id = $1;`

	row := s.db.QueryRow(query, userID)
	if row.Err() != nil {
		return 0, fmt.Errorf("cannot query max num from user: %w", row.Err())
	}
	var count int
	if err := row.Scan(&count); err != nil {
		return 0, fmt.Errorf("cannot scan count: %w", err)
	}

	return count, nil
}

func (s *DBStore) CreateGamePick(pick *schema.GamePick) (*schema.GamePick, error) {
	_ = pick.Insert(s.db, boil.Infer())
	return schema.GamePicks(
		qm.Where("id = ?", pick.ID),
		qm.Load("Player"),
	).One(s.db)
}

func (s *DBStore) GetGamePickByID(id, gameID string) (*schema.GamePick, error) {
	return schema.GamePicks(
		qm.Where("id = ?", id),
		qm.Where("game_id = ?", gameID),
		qm.Load("Player"),
	).One(s.db)
}

func (s *DBStore) GetActiveGamePicksWithPlayer(gameID, playerID string) (schema.GamePickSlice, error) {
	return schema.GamePicks(
		qm.Where("game_id = ?", gameID),
		qm.Where("player_id = ?", playerID),
		qm.Where("ended_at is null"),
	).All(s.db)
}

func (s *DBStore) UpdateGamePickEndedAt(pick *schema.GamePick) error {
	_, err := pick.Update(s.db, boil.Whitelist("ended_at", "updated_at"))
	return err
}

func (s *DBStore) GetGameEvents(gameID string, offset, limit int) (schema.GameEventSlice, error) {
	return schema.GameEvents(
		qm.Where("game_id = ?", gameID),
		qm.InnerJoin("match_events on game_events.match_event_id = match_events.id"),
		qm.OrderBy("match_events.period desc nulls last, match_events.minute desc, match_events.second desc"),
		qm.Load("Player"),
		qm.Offset(offset),
		qm.Limit(limit),
	).All(s.db)
}

func (s *DBStore) GetLeaderBoardForGame(gameID string) (*schema.MatchLeaderboard, error) {
	return schema.MatchLeaderboards(
		qm.Where("game_id = ?", gameID),
	).One(s.db)
}

func (s *DBStore) GetFollowingLeaderboard(matchID, userID string, limit int) (schema.MatchLeaderboardSlice, error) {
	return schema.MatchLeaderboards(
		qm.SQL(
			`
select m.id,
       m.score,
       row_number() over(order by m.score desc) as position,
       m.game_id,
       m.match_id,
       m.user_id
  from match_leaderboard m
 where m.match_id = $1
   and m.user_id in (
       select from_user_id
         from followers
        where followers.to_user_id = $2
       union
       select $2
     )
 order by m.score desc
 limit $3`, matchID, userID, limit,
		),
		qm.Load("User"),
	).All(s.db)
}

func (s *DBStore) GetLeaderboard(matchID string, limit int) (schema.MatchLeaderboardSlice, error) {
	return schema.MatchLeaderboards(
		qm.SQL(
			`
select m.id,
       m.score,
       row_number() over(order by m.score desc) as position,
       m.game_id,
       m.match_id,
       m.user_id,
       m.transaction_id
  from match_leaderboard m
 where m.match_id = $1
 order by m.score desc
 limit $2`, matchID, limit,
		),
		qm.Load("User"),
		qm.Load("Transaction"),
	).All(s.db)
}

func (s *DBStore) GetGameActivePowerUps(gameID string, position int) (schema.GamePowerupSlice, error) {
	return schema.GamePowerups(
		qm.Where("game_id = ?", gameID),
		qm.Where("ended_at is null"),
		qm.Where("position = ?", position),
	).All(s.db)
}

func (s *DBStore) GetGamePowerUps(gameID string) (schema.GamePowerupSlice, error) {
	return schema.GamePowerups(
		qm.Where("game_id = ?", gameID),
	).All(s.db)
}

func (s *DBStore) CreateGamePowerUp(pu *schema.GamePowerup) (*schema.GamePowerup, error) {
	return pu, pu.Insert(s.db, boil.Infer())
}

func (s *DBStore) GetUnnotifiedGamesByUserID(userID string) (schema.GameSlice, error) {
	games, err := schema.Games(
		qm.Where("user_id = ?", userID),
		qm.Where("notified = false"),
		qm.Load("Match.MatchPlayers"),
		qm.Load("GamePicks"),
		qm.Load("GamePicks.Player"),
		qm.Load("GamePowerups.Powerup"),
		qm.Load("User"),
	).All(s.db)
	if err != nil {
		return nil, err
	}
	return games, nil
}

func (s *DBStore) SetGameNotified(gameID string) error {
	// Ensure the game exists
	game, err := schema.Games(
		qm.Where("id = ?", gameID),
	).One(s.db)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return fmt.Errorf("no game found with id %s: %w", gameID, err)
		}
		return fmt.Errorf("error querying game with id %s: %w", gameID, err)
	}

	// Update the game to set notified to true
	game.Notified = true
	_, err = game.Update(s.db, boil.Whitelist("notified"))
	if err != nil {
		return fmt.Errorf("error updating notified status for game with id %s: %w", gameID, err)
	}

	return nil
}

func (s *DBStore) GetFinishedGames(userID string, limit int) (schema.GameSlice, error) {
	return schema.Games(
		qm.Where("user_id = ?", userID),
		qm.Where("status = ?", database.GameStatusFinished),
		qm.Limit(limit),
		qm.OrderBy("created_at desc"),
		qm.Load("Match.HomeTeam"),
		qm.Load("Match.AwayTeam"),
		qm.Load("GamePicks.Player"),
	).All(s.db)
}

// Function to delete all games and related data for a user
func (s *DBStore) DeleteAllGamesByUserID(userID string) error {
	tx, err := s.db.Begin()
	if err != nil {
		return fmt.Errorf("error starting transaction: %w", err)
	}
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		} else if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	query := "DELETE FROM games WHERE user_id = $1"
	_, err = tx.Exec(query, userID)
	if err != nil {
		return fmt.Errorf("error deleting games for user: %w", err)
	}

	return nil
}

// Corresponding DBStore method
func (s *DBStore) SetAllGamesNotified() error {
	result, err := s.db.Exec("UPDATE games SET notified = true WHERE notified = false")
	if err != nil {
		return fmt.Errorf("error updating games to notified: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error retrieving affected rows: %w", err)
	}

	if rowsAffected == 0 {
		return errors.New("no games updated, all games were already notified")
	}

	return nil
}

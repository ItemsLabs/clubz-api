package apiconv

import (
	"time"

	"github.com/itemslabs/clubz-api/database"

	"github.com/itemslabs/clubz-api/util"

	"github.com/itemslabs/clubz-api/config"
	"github.com/itemslabs/clubz-api/database/schema"
	"github.com/itemslabs/clubz-api/routes/model"
	"github.com/go-openapi/strfmt"
)

func ToGameSlice(games schema.GameSlice) []*model.Game {
	result := make([]*model.Game, 0, len(games))
	for _, game := range games {
		result = append(result, ToGame(game))
	}

	return result
}

func ToGame(game *schema.Game) *model.Game {
	if game == nil {
		return nil
	}

	allowedActions := new(model.GameAllowedActions)

	var getIdx = func(pos int) int {
		if pos < 1 || pos > config.MaxPicks() {
			return 0
		}

		return pos - 1
	}

	allowedActions.PowerupLeft = []int64{1, 1, 1, 1}
	var bonusPuUsed = false
	for _, pu := range game.R.GamePowerups {
		if pu == nil {
			continue
		}

		if pu.Bonus {
			bonusPuUsed = true
		}

		idx := getIdx(pu.Position)
		if allowedActions.PowerupLeft[idx] > 0 {
			allowedActions.PowerupLeft[idx] -= 1
		}
	}

	allowedActions.SwapsLeft = []int64{2, 2, 2, 2}
	for _, pick := range game.R.GamePicks {
		if pick == nil {
			continue
		}

		if pick.UserSwapped {
			idx := getIdx(pick.Position)
			if allowedActions.SwapsLeft[idx] > 0 {
				allowedActions.SwapsLeft[idx] -= 1
			}
		}
	}

	if !bonusPuUsed && (game.R.User.BonusPowerups > 0 ||
		game.SubscriptionTier == database.SubscriptionTierPremium ||
		game.SubscriptionTier == database.SubscriptionTierLite) {

		allowedActions.BonusPowerupLeft = 1
	}

	return &model.Game{
		ID:               game.ID,
		Version:          int64(game.Version),
		MatchID:          game.MatchID,
		UserID:           game.UserID,
		Status:           model.GameStatus(game.Status),
		Picks:            ToGamePickSlice(game.R.GamePicks, game.R.Match),
		Powerups:         ToGamePowerUpSlice(game.R.GamePowerups),
		Swaps:            []*model.PlayerSwap{},
		Rewards:          []*model.MatchReward{},
		AllowedActions:   allowedActions,
		Premium:          game.Premium,
		SubscriptionTier: ConvertSubscriptionTier(game.SubscriptionTier),
		Score:            ToFloatWithZero(game.Score),
	}
}

func ToGamePickSlice(picks schema.GamePickSlice, match *schema.Match) []*model.GamePick {
	if picks == nil || match == nil {
		return nil
	}

	result := make([]*model.GamePick, 0, len(picks))
	for _, pick := range picks {
		if pick == nil {
			continue
		}

		var include bool
		if match.FStart.Valid {
			include = !pick.EndedAt.Valid || pick.EndedAt.Time.After(match.FStart.Time)
		} else {
			include = !pick.EndedAt.Valid
		}

		if include {
			result = append(result, ToGamePick(pick, util.IsPlayerPlaying(pick.PlayerID, match.R.MatchPlayers)))
		}
	}

	return result
}

func ToGamePick(el *schema.GamePick, isPlaying bool) *model.GamePick {
	if el == nil || el.R == nil || el.R.Player == nil {
		return nil
	}

	return &model.GamePick{
		ID:                el.ID,
		GameID:            el.GameID,
		PlayerID:          el.PlayerID,
		ImageURL:          el.R.Player.ImageURL.String,
		PlayerName:        el.R.Player.FirstName.String,
		PlayerLastname:    el.R.Player.LastName.String,
		PlayerFullname:    el.R.Player.FullName.String,
		PlayerNickname:    el.R.Player.NickName.String,
		LineupPosition:    int64(el.Position),
		Start:             strfmt.DateTime(el.CreatedAt),
		End:               ToDateTime(el.EndedAt),
		Score:             ToFloatWithZero(el.Score),
		LastModifiedEvent: strfmt.DateTime(el.UpdatedAt),
		Minute:            int64(el.Minute),
		Second:            int64(el.Second),
		UserSwapped:       el.UserSwapped,
		IsPlaying:         isPlaying,
	}
}

func ToGamePowerUpSlice(arr schema.GamePowerupSlice) []*model.GamePowerUp {
	if arr == nil {
		return nil
	}

	result := make([]*model.GamePowerUp, 0, len(arr))
	for _, pu := range arr {
		if pu == nil {
			continue
		}

		result = append(result, ToGamePowerUp(pu))
	}

	return result
}

func ToGamePowerUp(el *schema.GamePowerup) *model.GamePowerUp {
	if el == nil || el.R == nil || el.R.Powerup == nil {
		return nil
	}

	return &model.GamePowerUp{
		ID:             el.ID,
		GameID:         el.GameID,
		PowerupID:      int64(el.PowerupID),
		PowerupName:    el.R.Powerup.Name,
		LineupPosition: int64(el.Position),
		Duration:       int64(el.Duration),
		Multiplier:     ToFloatWithZero(el.Multiplier),
		CreatedAt:      strfmt.DateTime(el.CreatedAt),
		EndedAt:        ToDateTime(el.EndedAt),
		Minute:         int64(el.Minute),
		Second:         int64(el.Second),
	}
}
func ToLeaderBoardPosition(pos *schema.MatchLeaderboard) *model.LeaderboardPosition {
	return &model.LeaderboardPosition{
		Position: ToInt64PrtFromIntPtr(pos.Position.Ptr()),
	}
}

func ToHistoricalGame(
	matchID string,
	matchName string,
	homeTeam string,
	awayTeam string,
	homeScore int,
	awayScore int,
	rank int,
	playerCount int,
	num int,
	points float64,
	date time.Time,
	prize float64,
	pickSlice schema.GamePickSlice,
	match *schema.Match) *model.HistoricalGame {

	picks := ToGamePickSlice(pickSlice, match)
	players := make([]*model.HistoricalGamePlayer, 0, len(picks))
	for _, pick := range picks {
		players = append(players, ToHistoricalGamePlayer(pick))
	}

	return &model.HistoricalGame{
		MatchID:     matchID,
		Points:      ToFloatWithZero(points),
		MatchName:   matchName,
		HomeTeam:    homeTeam,
		AwayTeam:    awayTeam,
		HomeScore:   int64(homeScore),
		AwayScore:   int64(awayScore),
		Num:         int64(num),
		Rank:        int64(rank),
		PlayerCount: int64(playerCount),
		Date:        strfmt.DateTime(date),
		Prize:       ToFloatWithZero(prize),
		Players:     players,
	}
}

func ToHistoricalGamePlayer(pick *model.GamePick) *model.HistoricalGamePlayer {
	return &model.HistoricalGamePlayer{
		ID:             pick.ID,
		Name:           pick.PlayerName,
		Lastname:       pick.PlayerLastname,
		Fullname:       pick.PlayerFullname,
		Nickname:       pick.PlayerNickname,
		ImageURL:       pick.ImageURL,
		Score:          float64(*pick.Score),
		LineupPosition: pick.LineupPosition,
		Start:          pick.Start,
		End:            pick.End,
	}
}

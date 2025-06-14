package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gameon-app-inc/laliga-matchfantasy-api/entity"
	"github.com/gameon-app-inc/laliga-matchfantasy-api/routes/apiconv"
	"github.com/gameon-app-inc/laliga-matchfantasy-api/routes/model"
	"github.com/gameon-app-inc/laliga-matchfantasy-api/types"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func (e *Env) GetWeekList(c echo.Context) error {
	limitParam := c.QueryParam("limit")
	if limitParam == "" {
		limitParam = "10"
	}
	limit, err := strconv.Atoi(limitParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	weeks, err := e.Store.GetGameWeeks(limit)
	if err != nil {
		return err
	}

	return e.RespondSuccess(c, apiconv.ToWeeks(weeks))

}

func (e *Env) GetCurrentWeekLeaderboard(c echo.Context) error {
	week, err := e.Store.GetCurrentGameWeek()
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	userID := userID(c)
	weekStartAt := week.StartAt
	weekEndAt := week.EndAt
	seasonStartAt := week.R.Season.StartAt.Time
	seasonEndAt := week.R.Season.EndAt.Time

	leaderboard, err := e.Store.GetCurrentWeekLeaderboard(seasonStartAt, seasonEndAt, weekStartAt, weekEndAt, userID)

	if err != nil {
		return err
	}

	rewardsMap, err := e.rewardsMap(week.ID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	progressionBar, err := e.calculateDivisionProgressionBar(week.ID, leaderboard)
	if err != nil {
		// ignore, but log error
		logrus.Errorf("Failed to calculate progression bar: %v", err)
	}

	return e.RespondSuccess(c, apiconv.ToLeaderboards(week, rewardsMap, leaderboard, progressionBar))
}

func (e *Env) GetWeekLeaderboard(c echo.Context) error {
	userID := userID(c)
	weekID := c.Param("id")
	week, err := e.Store.GetGameWeek(weekID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	weekStartAt := week.StartAt
	weekEndAt := week.EndAt
	seasonStartAt := week.R.Season.StartAt.Time
	seasonEndAt := week.R.Season.EndAt.Time
	var leaderboard []*types.LeaderboardEntry
	if week.Status == "c" {
		leaderboard, err = e.Store.GetConcludedWeekLeaderboard(weekID, userID)
		if err != nil {
			return err
		}
	} else {
		leaderboard, err = e.Store.GetCurrentWeekLeaderboard(seasonStartAt, seasonEndAt, weekStartAt, weekEndAt, userID)
		if err != nil {
			return err
		}
	}

	rewardsMap, err := e.rewardsMap(weekID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return e.RespondSuccess(c, apiconv.ToLeaderboards(week, rewardsMap, leaderboard, nil))

}

// GetUserGameWeekHistory retrieves user game week history by UUID.
func (e *Env) GetUserGameWeekHistory(c echo.Context) error {
	userID := userID(c)
	ID := c.Param("id")

	history, err := e.Store.GetUserGameWeekHistory(userID, ID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to fetch user game week history")
	}
	return e.RespondSuccess(c, apiconv.ToUserGameWeekHistory(history))
}

// ListUserGameWeekHistories retrieves list of user game week histories.
func (e *Env) ListUserGameWeekHistories(c echo.Context) error {
	userID := userID(c)

	histories, err := e.Store.ListUserGameWeekHistories(userID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to fetch user game week histories")
	}

	return e.RespondSuccess(c, apiconv.ToUserGameWeekHistories(histories))
}

// GetLatestUserGameWeekHistory retrieves the latest user game week history.
func (e *Env) GetLatestUserGameWeekHistory(c echo.Context) error {
	userID := userID(c)

	latestHistory, err := e.Store.GetLatestUserGameWeekHistory(userID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to fetch latest user game week history")
	}

	return e.RespondSuccess(c, apiconv.ToUserGameWeekHistory(latestHistory))
}

func (e *Env) rewardsMap(weekID string) (map[string][]*model.DivisionReward, error) {
	divisionRewards, err := e.Store.GetDivisionRewards(weekID)

	rewardsMap := make(map[string][]*model.DivisionReward)
	for _, r := range divisionRewards {
		reward := model.DivisionReward{
			Amount:       int64(r.R.Reward.Credits),
			MaxPosition:  int64(r.MaxPosition.Int),
			MinPosition:  int64(r.MinPosition),
			Game:         r.R.Reward.GameToken,
			Lapt:         r.R.Reward.LaptToken,
			Event:        int64(r.R.Reward.EventTickets),
			Balls:        int64(r.R.Reward.Ball),
			Shirts:       int64(r.R.Reward.Shirt),
			SignedBalls:  int64(r.R.Reward.SignedBall),
			SignedShirts: int64(r.R.Reward.SignedShirt),
			KickOffPack1: int64(r.R.Reward.KickoffPack1),
			KickOffPack2: int64(r.R.Reward.KickoffPack2),
			KickOffPack3: int64(r.R.Reward.KickoffPack3),
			SeasonPack1:  int64(r.R.Reward.SeasonPack1),
			SeasonPack2:  int64(r.R.Reward.SeasonPack2),
			SeasonPack3:  int64(r.R.Reward.SeasonPack3),
		}
		if r.DivisionID.Valid {
			rewardsMap[r.DivisionID.String] = append(rewardsMap[r.DivisionID.String], &reward)
		} else {
			rewardsMap["genesis"] = append(rewardsMap["genesis"], &reward)
		}
	}
	return rewardsMap, err
}

func (e *Env) calculateDivisionProgressionBar(weekID string, entries []*types.LeaderboardEntry) (
	*entity.ProgressionBar,
	error,
) {
	// find current user entry
	var currentEntry *types.LeaderboardEntry
	for _, entry := range entries {
		if entry.CurrentUser {
			currentEntry = entry
			break
		}
	}

	if currentEntry == nil {
		return nil, nil
	}

	// genesis
	if entity.IsGenesis(int(currentEntry.DivisionTier.Int64)) {
		return nil, nil
	}

	// find latest rank in division
	maxRank := int64(0)
	for _, entry := range entries {
		if entry.DivisionID == currentEntry.DivisionID && entry.Rank.Int64 > maxRank {
			maxRank = entry.Rank.Int64
		}
	}

	if maxRank == 0 {
		return nil, nil
	}

	// get division by id
	gwd, err := e.Store.GetGameWeekDivision(weekID, currentEntry.DivisionID.String)
	if err != nil {
		return nil, fmt.Errorf("cannot find game week division with id %s", currentEntry.DivisionID.String)
	}

	relegationPerc := gwd.RelegationCount / float64(gwd.Capacity)
	promotionPerc := gwd.PromotionCount / float64(gwd.Capacity)

	// calculate progression
	var progression float64

	if maxRank == 1 && currentEntry.Rank.Int64 == 1 {
		progression = 0.5
	} else {
		progression = 1 - float64(currentEntry.Rank.Int64)/float64(maxRank)
	}
	return &entity.ProgressionBar{
		Current:    progression,
		Relegation: relegationPerc,
		Promotion:  1 - promotionPerc,
	}, nil
}

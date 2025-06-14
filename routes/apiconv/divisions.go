package apiconv

import (
	"strconv"
	"strings"

	"github.com/gameon-app-inc/laliga-matchfantasy-api/database/schema"
	"github.com/gameon-app-inc/laliga-matchfantasy-api/entity"
	"github.com/gameon-app-inc/laliga-matchfantasy-api/routes/model"
	"github.com/gameon-app-inc/laliga-matchfantasy-api/types"
	"github.com/go-openapi/strfmt"
)

func ToWeeks(gameWeeks schema.GameWeekSlice) []*model.GameWeek {
	var weeks []*model.GameWeek

	for _, gw := range gameWeeks {
		weeks = append(weeks, toWeek(gw))
	}

	return weeks
}

func toWeek(week *schema.GameWeek) *model.GameWeek {
	return &model.GameWeek{
		ID:      week.ID,
		Name:    week.Name,
		StartAt: strfmt.DateTime(week.StartAt),
		EndAt:   strfmt.DateTime(week.EndAt),
		Status:  week.Status,
	}
}

func ToLeaderboards(
	week *schema.GameWeek,
	divisionRewards map[string][]*model.DivisionReward,
	leaderboard []*types.LeaderboardEntry,
	bar *entity.ProgressionBar,
) *model.NewLeaderboardResponse {
	response := &model.NewLeaderboardResponse{
		Week: &model.GameWeek{
			ID:      week.ID,
			Name:    week.Name,
			StartAt: strfmt.DateTime(week.StartAt),
			EndAt:   strfmt.DateTime(week.EndAt),
			Status:  week.Status,
		},
		DivisionLeaderboards: make([]*model.DivisionLeaderboard, 0),
	}

	leaderboardsByDivision := make(map[string][]*model.NewLeaderboardEntry)
	for _, lb := range leaderboard {
		entry := &model.NewLeaderboardEntry{
			UserID:            lb.UserID,
			UserName:          lb.UserName.String,
			TotalScore:        lb.TotalScore.Float64,
			Rank:              lb.Rank.Int64,
			CurrentUser:       lb.CurrentUser,
			WeekAverageRank:   &lb.WeekAverageRank.Float64,
			WeekMatchesPlayed: &lb.WeekMatchesPlayed.Int64,
		}
		tier := strconv.FormatInt(lb.DivisionTier.Int64, 10)

		var divisionID string
		if lb.DivisionID.Valid {
			divisionID = lb.DivisionID.String
		} else {
			divisionID = "genesis"
		}
		key := tier + "#" + divisionID

		if _, exists := leaderboardsByDivision[key]; !exists {
			leaderboardsByDivision[key] = make([]*model.NewLeaderboardEntry, 0)
		}
		leaderboardsByDivision[key] = append(leaderboardsByDivision[key], entry)
	}

	for divisionTierPlusID, entries := range leaderboardsByDivision {
		tierPlusID := strings.Split(divisionTierPlusID, "#")
		tier, _ := strconv.Atoi(tierPlusID[0])
		divisionID := tierPlusID[1]

		response.DivisionLeaderboards = append(
			response.DivisionLeaderboards, &model.DivisionLeaderboard{
				DivisionID:      divisionID,
				DivisionTier:    int64(tier),
				Leaderboard:     entries,
				DivisionRewards: divisionRewards[divisionID],
			},
		)
	}
	response.ProgressionBar = ToProgressionBar(bar)
	return response
}

func ToProgressionBar(bar *entity.ProgressionBar) *model.ProgressionBar {
	if bar == nil {
		return nil
	}
	return &model.ProgressionBar{
		Current:    bar.Current,
		Relegation: bar.Relegation,
		Promotion:  bar.Promotion,
	}
}

func ToUserGameWeekHistories(userGameWeekHistories schema.UserGameWeekHistorySlice) []*model.UserGameWeekHistory {
	var histories []*model.UserGameWeekHistory

	for _, ugwh := range userGameWeekHistories {
		histories = append(histories, ToUserGameWeekHistory(ugwh))
	}

	return histories
}

func ToUserGameWeekHistory(userGameWeekHistory *schema.UserGameWeekHistory) *model.UserGameWeekHistory {
	return &model.UserGameWeekHistory{
		GameWeek:             strfmt.UUID(userGameWeekHistory.GameWeekID),
		ID:                   strfmt.UUID(userGameWeekHistory.ID),
		NewDivision:          ToUUIDPtr(userGameWeekHistory.NewDivisionID.String),
		NewDivisionTier:      ToInt64Ptr(int64(userGameWeekHistory.NewDivisionTier.Int)),
		Status:               userGameWeekHistory.Status,
		User:                 strfmt.UUID(userGameWeekHistory.UserID),
		WeekCoins:            int64(userGameWeekHistory.WeekCoins),
		WeekDivision:         ToUUIDPtr(userGameWeekHistory.WeekDivisionID.String),
		WeekDivisionPosition: int64(userGameWeekHistory.WeekDivisionPosition),
		WeekDivisionTier:     ToInt64Ptr(int64(userGameWeekHistory.WeekDivisionTier.Int)),
		WeekPoints:           int64(userGameWeekHistory.WeekPoints),
		WeekAverageRank:      &userGameWeekHistory.WeekAverageRank.Float64,
		WeekMatchesPlayed:    ToInt64Ptr(int64(userGameWeekHistory.WeekMatchesPlayed)),
	}
}

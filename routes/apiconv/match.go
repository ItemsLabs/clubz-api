package apiconv

import (
	"fmt"
	"time"

	"github.com/itemslabs/clubz-api/database/schema"
	"github.com/itemslabs/clubz-api/routes/model"
	"github.com/go-openapi/strfmt"
)

func ToMatchSlice(matches schema.MatchSlice) []*model.Match {
	result := make([]*model.Match, 0, len(matches))
	for _, m := range matches {
		result = append(result, ToMatch(m))
	}

	return result
}

func ToMatch(match *schema.Match) *model.Match {
	result := &model.Match{
		ID:          match.ID,
		HomeTeam:    ToTeam(match.R.HomeTeam),
		AwayTeam:    ToTeam(match.R.AwayTeam),
		HomeScore:   int64(match.HomeScore),
		AwayScore:   int64(match.AwayScore),
		CreatedAt:   strfmt.DateTime(match.CreatedAt),
		UpdatedAt:   strfmt.DateTime(match.UpdatedAt),
		MatchTime:   strfmt.DateTime(match.MatchTime),
		HasLineups:  match.HasLineups,
		FStart:      ToDateTime(match.FStart),
		FEnd:        ToDateTime(match.FEnd),
		SStart:      ToDateTime(match.SStart),
		SEnd:        ToDateTime(match.SEnd),
		X1Start:     ToDateTime(match.X1Start),
		X1End:       ToDateTime(match.X1End),
		X2Start:     ToDateTime(match.X2Start),
		X2End:       ToDateTime(match.X2End),
		PStart:      ToDateTime(match.PStart),
		PEnd:        ToDateTime(match.MatchEnd),
		Version:     int64(match.Version),
		Status:      model.MatchStatus(match.Status),
		Period:      model.MatchPeriod(match.Period),
		Competition: ToCompetition(match.R.Competition),
		Rewards:     ToMatchRewardSlice(match.R.MatchRewards),
		Rules: &model.MatchRule{
			MaxStarPlayerPicks: 2,
			NumOfPicks:         4,
		},
		MatchType: model.MatchType(match.MatchType),
	}

	// Calculate matchTimer (soccer-specific logic)
	matchTimer := match.MatchTime.Format("15:04")
	now := time.Now()

	if !match.FStart.Valid {
		matchTimer = match.MatchTime.String()
	} else if now.Before(match.FStart.Time.Add(45 * time.Minute)) {
		// Calculate the minutes and seconds before the 45th minute
		minutes := int(now.Sub(match.FStart.Time).Minutes())
		seconds := int(now.Sub(match.FStart.Time).Seconds()) % 60
		matchTimer = fmt.Sprintf("%d'%02d\"", minutes, seconds)
	} else if match.Period == "1" && now.After(match.FStart.Time.Add(45*time.Minute)) {
		// Handling extra time in the first half
		addedTimeFirstHalf := int(now.Sub(match.FStart.Time.Add(45 * time.Minute)).Minutes())
		matchTimer = fmt.Sprintf("45+%d'", addedTimeFirstHalf)
	} else if match.FEnd.Valid && !match.SStart.Valid {
		// Handling half time if SStart is nil or empty
		matchTimer = "HT"
	} else if match.SStart.Valid && now.Before(match.SStart.Time.Add(45*time.Minute)) {
		// Calculate the minutes and seconds for the second half
		minutes := int(now.Sub(match.SStart.Time).Minutes())
		seconds := int(now.Sub(match.SStart.Time).Seconds()) % 60
		matchTimer = fmt.Sprintf("%d'%02d\"", minutes, seconds)
	} else if match.FEnd.Valid && now.After(match.FEnd.Time) {
		// Handle end of regular time
		matchTimer = "End of Regular Time"
	} else if match.X1Start.Valid && now.Before(match.X1End.Time) {
		// Extra time X1 logic
		minutes := int(now.Sub(match.X1Start.Time).Minutes())
		seconds := int(now.Sub(match.X1Start.Time).Seconds()) % 60
		matchTimer = fmt.Sprintf("ET1 %d'%02d\"", minutes, seconds)
	} else if match.X2Start.Valid && now.Before(match.X2End.Time) {
		// Extra time X2 logic
		minutes := int(now.Sub(match.X2Start.Time).Minutes())
		seconds := int(now.Sub(match.X2Start.Time).Seconds()) % 60
		matchTimer = fmt.Sprintf("ET2 %d'%02d\"", minutes, seconds)
	} else if match.PStart.Valid && now.Before(match.PEnd.Time) {
		// Handle penalties period
		matchTimer = "Penalties"
	} else {
		// If none of the above conditions are met, the match is finished
		matchTimer = "Finished"
	}

	result.MatchTimer = matchTimer
	return result

}

func ToMatchRewardSlice(rewards schema.MatchRewardSlice) []*model.MatchReward {
	result := make([]*model.MatchReward, 0, len(rewards))
	for _, el := range rewards {
		result = append(result, ToMatchReward(el))
	}

	return result
}

func ToMatchReward(el *schema.MatchReward) *model.MatchReward {
	return &model.MatchReward{
		Amount:       ToFloatWithZero(el.Amount),
		MinPosition:  int64(el.MinPosition),
		MaxPosition:  ToInt64PrtFromIntPtr(el.MaxPosition.Ptr()),
		Game:         el.Game,
		Lapt:         el.Lapt,
		Event:        int64(el.Event),
		Balls:        int64(el.Balls),
		Shirts:       int64(el.Shirts),
		SignedBalls:  int64(el.SignedBalls),
		SignedShirts: int64(el.SignedShirts),
		KickOffPack1: int64(el.KickoffPack1),
		KickOffPack2: int64(el.KickoffPack2),
		KickOffPack3: int64(el.KickoffPack3),
		SeasonPack1:  int64(el.SeasonPack1),
		SeasonPack2:  int64(el.SeasonPack2),
		SeasonPack3:  int64(el.SeasonPack3),
	}
}

func ToMatchSliceWithPlayerCount(matches schema.MatchSlice, playerCounts map[string]int64) []*model.Match {
	result := make([]*model.Match, 0, len(matches))
	for _, m := range matches {
		playerCount := playerCounts[m.ID]
		result = append(result, ToMatchWithPlayerCount(m, playerCount))
	}

	return result
}

func ToMatchWithPlayerCount(match *schema.Match, playerCount int64) *model.Match {
	result := &model.Match{
		ID:          match.ID,
		HomeTeam:    ToTeam(match.R.HomeTeam),
		AwayTeam:    ToTeam(match.R.AwayTeam),
		HomeScore:   int64(match.HomeScore),
		AwayScore:   int64(match.AwayScore),
		CreatedAt:   strfmt.DateTime(match.CreatedAt),
		UpdatedAt:   strfmt.DateTime(match.UpdatedAt),
		MatchTime:   strfmt.DateTime(match.MatchTime),
		HasLineups:  match.HasLineups,
		FStart:      ToDateTime(match.FStart),
		FEnd:        ToDateTime(match.FEnd),
		SStart:      ToDateTime(match.SStart),
		SEnd:        ToDateTime(match.SEnd),
		X1Start:     ToDateTime(match.X1Start),
		X1End:       ToDateTime(match.X1End),
		X2Start:     ToDateTime(match.X2Start),
		X2End:       ToDateTime(match.X2End),
		PStart:      ToDateTime(match.PStart),
		PEnd:        ToDateTime(match.MatchEnd),
		Version:     int64(match.Version),
		Status:      model.MatchStatus(match.Status),
		Period:      model.MatchPeriod(match.Period),
		Competition: ToCompetition(match.R.Competition),
		Rewards:     ToMatchRewardSlice(match.R.MatchRewards),
		Rules: &model.MatchRule{
			MaxStarPlayerPicks: 2,
			NumOfPicks:         4,
		},
		MatchType:   model.MatchType(match.MatchType),
		PlayerCount: playerCount, // Set player count here
	}

	// Calculate matchTimer (soccer-specific logic)
	matchTimer := match.MatchTime.Format("15:04")
	now := time.Now()

	if !match.FStart.Valid {
		matchTimer = match.MatchTime.String()
	} else if now.Before(match.FStart.Time.Add(45 * time.Minute)) {
		// Calculate the minutes and seconds before the 45th minute
		minutes := int(now.Sub(match.FStart.Time).Minutes())
		seconds := int(now.Sub(match.FStart.Time).Seconds()) % 60
		matchTimer = fmt.Sprintf("%d'%02d\"", minutes, seconds)
	} else if match.Period == "1" && now.After(match.FStart.Time.Add(45*time.Minute)) {
		// Handling extra time in the first half
		addedTimeFirstHalf := int(now.Sub(match.FStart.Time.Add(45 * time.Minute)).Minutes())
		matchTimer = fmt.Sprintf("45+%d'", addedTimeFirstHalf)
	} else if match.FEnd.Valid && !match.SStart.Valid {
		// Handling half time if SStart is nil or empty
		matchTimer = "HT"
	} else if match.SStart.Valid && now.Before(match.SStart.Time.Add(45*time.Minute)) {
		// Calculate the minutes and seconds for the second half
		minutes := int(now.Sub(match.SStart.Time).Minutes())
		seconds := int(now.Sub(match.SStart.Time).Seconds()) % 60
		matchTimer = fmt.Sprintf("%d'%02d\"", minutes, seconds)
	} else if match.FEnd.Valid && now.After(match.FEnd.Time) {
		// Handle end of regular time
		matchTimer = "End of Regular Time"
	} else if match.X1Start.Valid && now.Before(match.X1End.Time) {
		// Extra time X1 logic
		minutes := int(now.Sub(match.X1Start.Time).Minutes())
		seconds := int(now.Sub(match.X1Start.Time).Seconds()) % 60
		matchTimer = fmt.Sprintf("ET1 %d'%02d\"", minutes, seconds)
	} else if match.X2Start.Valid && now.Before(match.X2End.Time) {
		// Extra time X2 logic
		minutes := int(now.Sub(match.X2Start.Time).Minutes())
		seconds := int(now.Sub(match.X2Start.Time).Seconds()) % 60
		matchTimer = fmt.Sprintf("ET2 %d'%02d\"", minutes, seconds)
	} else if match.PStart.Valid && now.Before(match.PEnd.Time) {
		// Handle penalties period
		matchTimer = "Penalties"
	} else {
		// If none of the above conditions are met, the match is finished
		matchTimer = "Finished"
	}

	result.MatchTimer = matchTimer
	return result
}

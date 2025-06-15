package apiconv

import (
	"github.com/itemslabs/clubz-api/database"
	"github.com/itemslabs/clubz-api/database/schema"
	"github.com/itemslabs/clubz-api/routes/model"
)

func ToPregamePlayerStat(latestMatchPlayers schema.MatchPlayerSlice, selectedPercent float64, pointDistribution []*database.PointBucket) *model.PregamePlayerStat {
	matches := make([]*model.LatestMatchStat, 0, len(latestMatchPlayers))
	for _, mp := range latestMatchPlayers {
		matches = append(matches, &model.LatestMatchStat{
			TeamName:     GetTeamDisplayName(mp.R.Team),
			TeamCrestURL: mp.R.Team.CrestURL.String,
			MinPlayed:    int64(mp.PlayedSeconds.Int / 60),
			Points:       int64(mp.Score.Float64),
		})
	}

	points := make([]*model.PointBucket, 0, len(pointDistribution))
	for _, p := range pointDistribution {
		points = append(points, &model.PointBucket{
			Low:   int64(p.Low),
			High:  int64(p.High),
			Value: p.Points.Float64,
		})
	}

	return &model.PregamePlayerStat{
		LatestMatches:     matches,
		SelectedPercent:   selectedPercent,
		PointDistribution: points,
	}
}

func ToLiveGamePlayerStat(selectedPercent float64, pointDistribution []*database.PointBucket, last10MinPoints int, actionSummary []*database.ActionSummary) *model.LiveGamePlayerStat {
	points := make([]*model.PointBucket, 0, len(pointDistribution))
	for _, p := range pointDistribution {
		points = append(points, &model.PointBucket{
			Low:   int64(p.Low),
			High:  int64(p.High),
			Value: p.Points.Float64,
		})
	}

	actions := make([]*model.ActionSummary, 0, len(actionSummary))
	for _, el := range actionSummary {
		actions = append(actions, &model.ActionSummary{
			Name:   el.Name,
			Count:  int64(el.Count),
			Points: el.Points,
		})
	}

	return &model.LiveGamePlayerStat{
		SelectedPercent:   selectedPercent,
		PointDistribution: points,
		Last10MinPoints:   float64(last10MinPoints),
		ActionSummary:     actions,
	}
}

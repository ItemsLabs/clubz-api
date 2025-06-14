package apiconv

import (
	"github.com/gameon-app-inc/laliga-matchfantasy-api/database"
	"github.com/gameon-app-inc/laliga-matchfantasy-api/database/schema"
	"github.com/gameon-app-inc/laliga-matchfantasy-api/routes/model"
)

func ToMatchPlayerSlice(players schema.MatchPlayerSlice, matchPlayersPPG map[string]float64) []*model.MatchPlayer {
	result := make([]*model.MatchPlayer, 0, len(players))
	for _, pl := range players {
		ppg := matchPlayersPPG[pl.PlayerID]
		result = append(result, ToMatchPlayer(pl, ppg))
	}

	return result
}

func ToMatchPlayer(mp *schema.MatchPlayer, ppg float64) *model.MatchPlayer {
	// change player avg score from match player avg score
	mp.R.Player.AvgScore = mp.AvgScore

	return &model.MatchPlayer{
		Position:     mp.Position.Ptr(),
		JerseyNumber: ToInt64PrtFromIntPtr(mp.JerseyNumber.Ptr()),
		TeamID:       mp.TeamID,
		Player:       ToPlayer(mp.R.Player),
		IsPlaying: mp.FromLineups &&
			mp.Position.String != database.PositionUnknown &&
			mp.Position.String != database.PositionSubstitute,
		IsStar:     false,
		MatchScore: ToFloatWithZeroPtr(mp.Score.Ptr()),
		TeamCrest:  mp.R.Team.CrestURL.String,
		TeamName:   mp.R.Team.Name,
		Ppg:        ppg,
	}
}

package util

import (
	"github.com/itemslabs/clubz-api/database"
	"github.com/itemslabs/clubz-api/database/schema"
)

func IsPlayerPlaying(playerID string, matchPlayers schema.MatchPlayerSlice) bool {
	// find corresponding match players
	for _, mp := range matchPlayers {
		if mp.PlayerID == playerID {
			return mp.FromLineups &&
				mp.Position.String != database.PositionUnknown &&
				mp.Position.String != database.PositionSubstitute
		}
	}

	return false
}

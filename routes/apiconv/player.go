package apiconv

import (
	"strings"

	"github.com/gameon-app-inc/laliga-matchfantasy-api/database/schema"
	"github.com/gameon-app-inc/laliga-matchfantasy-api/routes/model"
)

func ToPlayer(player *schema.Player) *model.Player {
	fullName := player.FullName.String
	if fullName == "" {
		fullName = player.FirstName.String + " " + player.LastName.String
	}
	player.NickName.String = strings.TrimSpace(player.NickName.String)
	if player.NickName.String != "" {
		fullName = player.NickName.String
	}

	return &model.Player{
		ID:             player.ID,
		FullName:       fullName,
		ImageURL:       player.ImageURL.Ptr(),
		AvgScore:       ToFloatWithZeroPtr(player.AvgScore.Ptr()),
		NormalizedName: player.NormalizedName.String,
		ImportID:       ToStringPtr(player.ImportID),
	}
}

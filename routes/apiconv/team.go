package apiconv

import (
	"strings"

	"github.com/itemslabs/clubz-api/database/schema"
	"github.com/itemslabs/clubz-api/routes/model"
)

func ToTeam(team *schema.Team) *model.Team {
	return &model.Team{
		ID:       team.ID,
		Name:     GetTeamDisplayName(team),
		Abbr:     ToStringPtr(team.Abbr),
		CrestURL: ToStringPtr(team.CrestURL),
	}
}

func GetTeamDisplayName(team *schema.Team) string {
	var displayName = team.Name
	if team.ShortName.Valid && strings.TrimSpace(team.ShortName.String) != "" {
		displayName = team.ShortName.String
	}

	return displayName
}

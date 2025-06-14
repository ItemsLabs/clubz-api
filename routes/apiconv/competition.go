package apiconv

import (
	"github.com/gameon-app-inc/laliga-matchfantasy-api/database/schema"
	"github.com/gameon-app-inc/laliga-matchfantasy-api/routes/model"
)

func ToCompetition(c *schema.Competition) *model.Competition {
	var displayName = c.Name
	if c.ShortName.Valid {
		displayName = c.ShortName.String
	}
	return &model.Competition{
		ID:   c.ID,
		Name: displayName,
	}
}

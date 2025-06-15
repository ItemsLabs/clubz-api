package apiconv

import (
	"github.com/itemslabs/clubz-api/database/schema"
	"github.com/itemslabs/clubz-api/routes/model"
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

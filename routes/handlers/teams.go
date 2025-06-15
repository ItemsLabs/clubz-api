package handlers

import (
	"github.com/itemslabs/clubz-api/database/schema"
	"github.com/itemslabs/clubz-api/routes/apiconv"
	"github.com/labstack/echo/v4"
)

type TeamListResponse struct {
	Team schema.Team
}

// TeamList godoc
// @Summary List all teams
// @Description List all teams
// @ID teams-list
// @Produce json
// @Success 200 {object} schema.Team
// @Router /teams [get]
func (e *Env) GetTeamByID(c echo.Context) error {
	team, err := e.Store.GetTeamByID(c.Param("id"))
	if err != nil {
		return err
	}

	return e.RespondSuccess(c, apiconv.ToTeam(team))
}

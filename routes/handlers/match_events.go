package handlers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// GetMatchEvents godoc
// @Summary Get all events for a match
// @Description Get all events for a specific match ID
// @ID get-match-events
// @Produce json
// @Param id path string true "Match ID"
// @Success 200 {array} schema.MatchEvent
// @Router /matches/{id}/events [get]
func (e *Env) GetMatchEvents(c echo.Context) error {
	matchID := c.Param("id")
	events, err := e.Store.GetMatchEventsByMatchID(matchID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return e.RespondSuccess(c, events)
}

// GetMatchEventsByType godoc
// @Summary Get events of a specific type for a match
// @Description Get events of a specific type for a match ID
// @ID get-match-events-by-type
// @Produce json
// @Param id path string true "Match ID"
// @Param type query int true "Event Type"
// @Success 200 {object} schema.MatchEvent
// @Router /matches/{id}/events/type [get]
func (e *Env) GetMatchEventsByType(c echo.Context) error {
	matchID := c.Param("id")
	eventType, err := strconv.Atoi(c.QueryParam("type"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid event type")
	}
	events, err := e.Store.GetMatchEventsByMatchIDAndType(matchID, eventType)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return e.RespondSuccess(c, events)
}

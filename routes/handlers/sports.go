package handlers

import (
	"github.com/labstack/echo/v4"
)

type SportsService struct {
}

// SportList godoc
// @Summary List all sports
// @Description List all sports
// @ID sports-list
// @Produce json
// @Success 200 {object} model.Sport
// @Router /sports [get]
func (e *Env) SportList(c echo.Context) error {
	sportList, err := e.Store.GetSports()
	if err != nil {
		return err
	}

	return e.RespondSuccess(c, sportList)
}

// GetSportByID godoc
// @Summary Get sport by ID
// @Description Get sport by ID
// @ID get-sport-by-id
// @Produce json
// @Param id path string true "Sport ID"
// @Success 200 {object} model.Sport
// @Router /sports/{id} [get]
func (e *Env) GetSportByID(c echo.Context) error {
	sport, err := e.Store.GetSportByID(c.Param("id"))
	if err != nil {
		return err
	}

	return e.RespondSuccess(c, sport)
}

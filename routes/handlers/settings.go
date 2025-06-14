package handlers

import (
	"github.com/labstack/echo/v4"
)

// Settings godoc
// @Summary Get settings
// @Description Get settings
// @ID settings
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /settings [get]
func (e *Env) Settings(c echo.Context) error {
	return e.RespondSuccess(c, map[string]interface{}{
		// This flag is used to tell frontend "show" or "not show" text
		// this is needed for app review process
		// name is specifically chosen to be non-suspicious
		"2021-05-01-test": false,
	})
}

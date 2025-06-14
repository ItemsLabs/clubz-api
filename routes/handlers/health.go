package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Readiness godoc
// @Summary Check if the service is ready
// @Description Check if the service is ready
// @ID readiness
// @Produce json
// @Success 200 {string} string "OK"
// @Router /readiness [get]
func (e *Env) Readiness(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}

// Health godoc
// @Summary Check if the service is healthy
// @Description Check if the service is healthy
// @ID health
// @Produce json
// @Success 200 {string} string "OK"
// @Router /health [get]
func (e *Env) Health(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}

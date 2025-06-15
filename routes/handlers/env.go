package handlers

import (
	"net/http"

	"github.com/go-redis/redis/v8"

	"github.com/itemslabs/clubz-api/database"
	"github.com/itemslabs/clubz-api/types"
	"github.com/labstack/echo/v4"
)

type Env struct {
	Store            database.Store
	ActionStore      database.ActionStore
	RedisClient      *redis.Client
	WebsocketConfigs map[string]types.WebsocketConfig
}

// RespondSuccess godoc
// @Summary Responds with a 200 status code and the given object
// @Description Responds with a 200 status code and the given object
// @Accept  json
// @Produce  json
// @Param obj body interface{} true "Object to be returned"
// @Success 200 {object} interface{}
func (e *Env) RespondSuccess(c echo.Context, obj interface{}) error {
	return c.JSON(http.StatusOK, obj)
}

// RespondNoContent godoc
// @Summary Responds with a 204 status code
// @Description Responds with a 204 status code
// @Accept  json
// @Produce  json
// @Success 204
func (e *Env) RespondNoContent(c echo.Context) error {
	return c.NoContent(http.StatusNoContent)
}

// RespondNotFound godoc
// @Summary Responds with a 404 status code
// @Description Responds with a 404 status code
// @Accept  json
// @Produce  json
// @NotFound 404
func (e *Env) RespondNotFound(c echo.Context) error {
	return c.NoContent(http.StatusNotFound)
}

// RespondSuccess godoc
// @Summary Responds with a 200 status code and the given object
// @Description Responds with a 200 status code and the given object
// @Accept  json
// @Produce  json
// @Param obj body interface{} true "Object to be returned"
// @Success 200 {object} interface{}
func RespondSuccess(c echo.Context, obj interface{}) error {
	return c.JSON(http.StatusOK, obj)
}

//	RespondNoContent godoc
//
// @Summary Responds with a 204 status code
// @Description Responds with a 204 status code
// @Accept  json
// @Produce  json
// @Success 204
func RespondNoContent(c echo.Context) error {
	return c.NoContent(http.StatusNoContent)
}

//	RespondNotFound godoc
//
// @Summary Responds with a 404 status code
// @Description Responds with a 404 status code
// @Accept  json
// @Produce  no-content
// @Success 404
func RespondNotFound(c echo.Context) error {
	return c.NoContent(http.StatusNotFound)
}

package handlers

import (
	"github.com/itemslabs/clubz-api/routes/apiconv"
	"github.com/itemslabs/clubz-api/routes/model"
	"github.com/labstack/echo/v4"
)

func (e *Env) ActionsList(c echo.Context) error {
	actions, err := e.Store.GetActions()
	if err != nil {
		return err
	}

	return e.RespondSuccess(c, apiconv.ToActionSlice(actions))
}

func (e *Env) UpdateAction(c echo.Context) error {
	actionID := c.Param("id")

	var in model.UpdateActionRequest

	if err := e.ParseBody(c, &in); err != nil {
		return err
	}

	action, err := e.Store.UpdateAction(actionID, &in)
	if err != nil {
		return err
	}
	return e.RespondSuccess(c, apiconv.ToAction(action))
}

func (e *Env) CreateActions(c echo.Context) error {
	var in model.CreateActionsRequest

	if err := e.ParseBody(c, &in); err != nil {
		return err
	}
	actions, err := e.Store.CreateActions(&in)
	if err != nil {
		return err
	}

	return e.RespondSuccess(c, actions)
}

func (e *Env) DeleteAction(c echo.Context) error {
	actionID := c.Param("id")

	err := e.Store.DeleteAction(actionID)
	if err != nil {
		return err
	}

	return e.RespondSuccess(c, 1)
}

// func (e *Env) ClearActions(c echo.Context) error {
// 	actions, err := e.Store.GetActions()
// 	if err != nil {
// 		return err
// 	}

// 	return e.RespondSuccess(c, apiconv.ToActionSlice(actions))
// }

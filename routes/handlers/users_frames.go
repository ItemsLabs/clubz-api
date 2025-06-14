package handlers

import (
	"net/http"
	"strconv"

	"github.com/gameon-app-inc/laliga-matchfantasy-api/routes/apiconv"
	"github.com/labstack/echo/v4"
)

// GetUserFrames returns a list of frames associated with a user.
func (e *Env) GetUserFrames(c echo.Context) error {
	userID := userID(c)

	userFrames, err := e.Store.GetUserFrames(userID)
	if err != nil {
		return err
	}

	return e.RespondSuccess(c, apiconv.ToUserFrameSlice(userFrames))
}

// CreateUserFrame assigns a frame to a user.
func (e *Env) CreateUserFrame(c echo.Context) error {
	userID := userID(c)
	frameIDStr := c.Param("frame_id")
	frameID, err := strconv.Atoi(frameIDStr)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid frame ID")
	}

	userFrame, err := e.Store.CreateUserFrame(userID, frameID)
	if err != nil {
		return err
	}

	return e.RespondSuccess(c, apiconv.ToUserFrame(userFrame))
}

// DeleteUserFrame removes a user frame by its ID.
func (e *Env) DeleteUserFrame(c echo.Context) error {
	userFrameIDStr := c.Param("id")
	userFrameID, err := strconv.Atoi(userFrameIDStr)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid user frame ID")
	}

	err = e.Store.DeleteUserFrame(userFrameID)
	if err != nil {
		return err
	}

	return e.RespondSuccess(c, 1)
}

// ChangeAllUsersFramesToUnselected changes all user frames to unselected and then select a specific user frame.
func (e *Env) ChangeAllUsersFramesToUnselected(c echo.Context) error {
	err := e.Store.ChangeAllUsersFramesToUnselected()
	if err != nil {
		return err
	}

	return e.RespondSuccess(c, 1)
}

// ChangeUserFrameToSelected changes all user frames to unselected and then select a specific user frame.
func (e *Env) SelectUserFrame(c echo.Context) error {
	userFrameIDStr := c.Param("id")

	err := e.Store.ChangeAllUsersFramesToUnselected()
	if err != nil {
		return err
	}

	userFrameID, err := strconv.Atoi(userFrameIDStr)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid user frame ID")
	}

	err = e.Store.ChangeUserFrameToSelected(userFrameID)
	if err != nil {
		return err
	}

	return e.RespondSuccess(c, 1)
}

package handlers

import (
	"github.com/gameon-app-inc/laliga-matchfantasy-api/routes/apiconv"
	"github.com/gameon-app-inc/laliga-matchfantasy-api/routes/model"
	"github.com/labstack/echo/v4"
)

// FramesList returns a list of frames.
func (e *Env) FramesList(c echo.Context) error {
	frames, err := e.Store.GetFrames()
	if err != nil {
		return err
	}
	return e.RespondSuccess(c, apiconv.ToFrameSlice(frames))
}

func (e *Env) GetFrame(c echo.Context) error {
	frameID := c.Param("id")

	frame, err := e.Store.GetFrameByID(frameID)
	if err != nil {
		return err
	}

	return e.RespondSuccess(c, apiconv.ToFrame(frame))
}

// UpdateFrame updates a frame by its ID.
func (e *Env) UpdateFrame(c echo.Context) error {
	frameID := c.Param("id")

	var in model.UpdateFrameRequest

	if err := e.ParseBody(c, &in); err != nil {
		return err
	}

	frame, err := e.Store.UpdateFrame(frameID, &in)
	if err != nil {
		return err
	}
	return e.RespondSuccess(c, apiconv.ToFrame(frame))
}

// CreateFrames creates multiple frames.
func (e *Env) CreateFrames(c echo.Context) error {
	var in model.CreateFramesRequest

	if err := e.ParseBody(c, &in); err != nil {
		return err
	}

	frames, err := e.Store.CreateFrames(&in)
	if err != nil {
		return err
	}

	return e.RespondSuccess(c, frames)
}

// DeleteFrame deletes a frame by its ID.
func (e *Env) DeleteFrame(c echo.Context) error {
	frameID := c.Param("id")

	err := e.Store.DeleteFrame(frameID)
	if err != nil {
		return err
	}

	return e.RespondSuccess(c, 1)
}

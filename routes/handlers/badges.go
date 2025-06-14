package handlers

import (
	"github.com/gameon-app-inc/laliga-matchfantasy-api/routes/apiconv"
	"github.com/gameon-app-inc/laliga-matchfantasy-api/routes/model"
	"github.com/labstack/echo/v4"
)

// BadgesList returns a list of badges.
func (e *Env) BadgesList(c echo.Context) error {
	badges, err := e.Store.GetBadges()
	if err != nil {
		return err
	}
	return e.RespondSuccess(c, apiconv.ToBadgeSlice(badges))
}

func (e *Env) GetBadge(c echo.Context) error {
	badgeID := c.Param("id")

	badge, err := e.Store.GetBadgeByID(badgeID)
	if err != nil {
		return err
	}

	return e.RespondSuccess(c, apiconv.ToBadge(badge))
}

// UpdateBadge updates a badge by its ID.
func (e *Env) UpdateBadge(c echo.Context) error {
	badgeID := c.Param("id")

	var in model.UpdateBadgeRequest

	if err := e.ParseBody(c, &in); err != nil {
		return err
	}

	badge, err := e.Store.UpdateBadge(badgeID, &in)
	if err != nil {
		return err
	}
	return e.RespondSuccess(c, apiconv.ToBadge(badge))
}

// CreateBadges creates multiple badges.
func (e *Env) CreateBadges(c echo.Context) error {
	var in model.CreateBadgesRequest

	if err := e.ParseBody(c, &in); err != nil {
		return err
	}

	badges, err := e.Store.CreateBadges(&in)
	if err != nil {
		return err
	}

	return e.RespondSuccess(c, badges)
}

// DeleteBadge deletes a badge by its ID.
func (e *Env) DeleteBadge(c echo.Context) error {
	badgeID := c.Param("id")

	err := e.Store.DeleteBadge(badgeID)
	if err != nil {
		return err
	}

	return e.RespondSuccess(c, 1)
}

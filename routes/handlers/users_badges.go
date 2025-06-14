package handlers

import (
	"strconv"

	"github.com/gameon-app-inc/laliga-matchfantasy-api/routes/apiconv"
	"github.com/labstack/echo/v4"
)

// GetUserBadges returns a list of badges associated with a user.
func (e *Env) GetUserBadges(c echo.Context) error {
	userID := userID(c)

	userBadges, err := e.Store.GetUserBadges(userID)
	if err != nil {
		return err
	}

	return e.RespondSuccess(c, apiconv.ToUserBadgeSlice(userBadges))
}

// CreateUserBadge assigns a badge to a user.
func (e *Env) CreateUserBadge(c echo.Context) error {
	userID := userID(c)
	badgeID := c.Param("badge_id")
	badgeIDInt, err := strconv.Atoi(badgeID)
	if err != nil {
		return err
	}

	userBadge, err := e.Store.CreateUserBadge(userID, badgeIDInt)
	if err != nil {
		return err
	}

	return e.RespondSuccess(c, apiconv.ToUserBadge(userBadge))
}

// DeleteUserBadge removes a user badge by its ID.
func (e *Env) DeleteUserBadge(c echo.Context) error {
	userBadgeID := c.Param("id")
	id, err := strconv.Atoi(userBadgeID)
	if err != nil {
		return err
	}

	err = e.Store.DeleteUserBadge(id)
	if err != nil {
		return err
	}

	return e.RespondSuccess(c, 1)
}

// ChangeAllUsersBadgesToUnselected changes all user badges to unselected and then select a specific user badge.
func (e *Env) ChangeAllUsersBadgesToUnselected(c echo.Context) error {
	err := e.Store.ChangeAllUsersBadgesToUnselected()
	if err != nil {
		return err
	}

	return e.RespondSuccess(c, 1)
}

// ChangeUserBadgeToSelected changes a user badge to selected.
func (e *Env) SelectUserBadge(c echo.Context) error {
	userBadgeID := c.Param("id")

	err := e.Store.ChangeAllUsersBadgesToUnselected()
	if err != nil {
		return err
	}

	id, err := strconv.Atoi(userBadgeID)
	if err != nil {
		return err
	}

	err = e.Store.ChangeUserBadgeToSelected(id)
	if err != nil {
		return err
	}

	return e.RespondSuccess(c, 1)
}

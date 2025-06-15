package handlers

import (
	"strconv"

	"github.com/itemslabs/clubz-api/routes/apiconv"
	"github.com/labstack/echo/v4"
)

// GetUserBanners returns a list of banners associated with a user.
func (e *Env) GetUserBanners(c echo.Context) error {
	userID := userID(c)

	userBanners, err := e.Store.GetUserBanners(userID)
	if err != nil {
		return err
	}

	return e.RespondSuccess(c, apiconv.ToUserBannerSlice(userBanners))
}

// CreateUserBanner assigns a banner to a user.
func (e *Env) CreateUserBanner(c echo.Context) error {
	userID := userID(c)
	bannerIDStr := c.Param("banner_id")
	bannerID, err := strconv.Atoi(bannerIDStr)
	if err != nil {
		return err
	}

	userBanner, err := e.Store.CreateUserBanner(userID, bannerID)
	if err != nil {
		return err
	}

	return e.RespondSuccess(c, apiconv.ToUserBanner(userBanner))
}

// DeleteUserBanner removes a user banner by its ID.
func (e *Env) DeleteUserBanner(c echo.Context) error {
	userBannerIDStr := c.Param("id")
	userBannerID, err := strconv.Atoi(userBannerIDStr)
	if err != nil {
		return err
	}

	err = e.Store.DeleteUserBanner(userBannerID)
	if err != nil {
		return err
	}

	return e.RespondSuccess(c, 1)
}

// ChangeAllUsersBannersToUnselected changes all user banners to unselected and then select a specific user banner.
func (e *Env) ChangeAllUsersBannersToUnselected(c echo.Context) error {
	err := e.Store.ChangeAllUsersBannersToUnselected()
	if err != nil {
		return err
	}
	return e.RespondSuccess(c, 1)
}

// SelectUserBanner selects a user banner by its ID.
func (e *Env) SelectUserBanner(c echo.Context) error {
	userBannerIDStr := c.Param("id")

	err := e.Store.ChangeAllUsersBannersToUnselected()
	if err != nil {
		return err
	}

	userBannerID, err := strconv.Atoi(userBannerIDStr)
	if err != nil {
		return err
	}

	err = e.Store.SelectUserBanner(userBannerID)
	if err != nil {
		return err
	}

	return e.RespondSuccess(c, 1)
}

package handlers

import (
	"github.com/gameon-app-inc/laliga-matchfantasy-api/routes/apiconv"
	"github.com/gameon-app-inc/laliga-matchfantasy-api/routes/model"
	"github.com/labstack/echo/v4"
)

// BannersList returns a list of banners.
func (e *Env) BannersList(c echo.Context) error {
	banners, err := e.Store.GetBanners()
	if err != nil {
		return err
	}
	return e.RespondSuccess(c, apiconv.ToBannerSlice(banners))
}

func (e *Env) GetBanner(c echo.Context) error {
	bannerID := c.Param("id")

	banner, err := e.Store.GetBannerByID(bannerID)
	if err != nil {
		return err
	}

	return e.RespondSuccess(c, apiconv.ToBanner(banner))
}

// UpdateBanner updates a banner by its ID.
func (e *Env) UpdateBanner(c echo.Context) error {
	bannerID := c.Param("id")

	var in model.UpdateBannerRequest

	if err := e.ParseBody(c, &in); err != nil {
		return err
	}

	banner, err := e.Store.UpdateBanner(bannerID, &in)
	if err != nil {
		return err
	}
	return e.RespondSuccess(c, apiconv.ToBanner(banner))
}

// CreateBanners creates multiple banners.
func (e *Env) CreateBanners(c echo.Context) error {
	var in model.CreateBannersRequest

	if err := e.ParseBody(c, &in); err != nil {
		return err
	}

	banners, err := e.Store.CreateBanners(&in)
	if err != nil {
		return err
	}

	return e.RespondSuccess(c, banners)
}

// DeleteBanner deletes a banner by its ID.
func (e *Env) DeleteBanner(c echo.Context) error {
	bannerID := c.Param("id")

	err := e.Store.DeleteBanner(bannerID)
	if err != nil {
		return err
	}

	return e.RespondSuccess(c, 1)
}

package handlers

import (
	"github.com/itemslabs/clubz-api/routes/apiconv"
	"github.com/labstack/echo/v4"
)

// GetOrderByID godoc
// @Summary Get Order by ID
// @Description Get Order by ID
// @ID get-order-by-id
// @Produce json
// @Param id path string true "Order ID"
// @Success 200 {object} Order
// @Router /purchase/{id} [get]

func (e *Env) GetItemByID(c echo.Context) error {
	item, err := e.Store.GetItemByID(c.Param("id"))

	if err != nil {
		return err
	}

	return e.RespondSuccess(c, apiconv.ToItem(item))
}

func (e *Env) GetItems(c echo.Context) error {
	item, err := e.Store.GetItemByID(c.Param("id"))

	if err != nil {
		return err
	}

	return e.RespondSuccess(c, apiconv.ToItem(item))
}

func (e *Env) ItemList(c echo.Context) error {
	actions, err := e.Store.GetItems()
	if err != nil {
		return err
	}

	return e.RespondSuccess(c, apiconv.ToItemSlice(actions))
}

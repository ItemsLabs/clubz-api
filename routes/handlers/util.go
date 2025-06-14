package handlers

import (
	"strconv"

	"github.com/go-openapi/strfmt"
	"github.com/labstack/echo/v4"
)

type validator interface {
	Validate(formats strfmt.Registry) error
}

func (e *Env) ParseBody(c echo.Context, i interface{}) error {
	if err := c.Bind(i); err != nil {
		return err
	}

	v, ok := i.(validator)
	if !ok {
		return nil
	}

	return v.Validate(nil)
}

func (e *Env) GetOffsetAndLimit(page, pageSize string) (int, int) {
	pageSizeVal, _ := strconv.Atoi(pageSize)
	pageVal, _ := strconv.Atoi(page)

	if pageSizeVal > 0 && pageVal > 0 {
		return (pageVal - 1) * pageSizeVal, pageSizeVal
	}
	return 0, 999999
}

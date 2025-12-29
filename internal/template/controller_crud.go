package template

import (
	"github.com/drewjya/forge/internal/util"
)

func ControllerCRUD(name string) string {
	nn := util.NewName(name)
	N := nn.Title()
	name = nn.Package()
	module := goModule()
	p := name

	return `package ` + p + `controller

import (
	"net/http"

	"` + module + `/internal/request"
	"` + module + `/internal/types"

	"github.com/labstack/echo/v4"
)

func (h *` + N + `Controller) GetAll(c echo.Context) error {
	rows, err := h.Service.GetAll(c.Request().Context())
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, rows)
}

func (h *` + N + `Controller) Create(c echo.Context) error {
	var req request.` + N + `Request
	if err := c.Bind(&req); err != nil {
		return err
	}
	if err := c.Validate(&req); err != nil {
		return err
	}
	return h.Service.Create(c.Request().Context(), req)
}

func (h *` + N + `Controller) Update(c echo.Context) error {
	id, err := types.ParseULID(c.Param("id"))
	if err != nil {
		return err
	}
	var req request.` + N + `Request
	if err := c.Bind(&req); err != nil {
		return err
	}
	if err := c.Validate(&req); err != nil {
		return err
	}
	return h.Service.Update(c.Request().Context(), id, req)
}

func (h *` + N + `Controller) Delete(c echo.Context) error {
	id, err := types.ParseULID(c.Param("id"))
	if err != nil {
		return err
	}
	return h.Service.Delete(c.Request().Context(), id)
}
`
}

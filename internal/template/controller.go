package template

import (
	"github.com/drewjya/forge/internal/util"
)

func ControllerInterface(name string, crud bool) string {
	nn := util.NewName(name)
	N := nn.Title()
	name = nn.Package()

	if crud {
		return `package ` + name + `controller

import "github.com/labstack/echo/v4"

type I` + N + `Controller interface {
	GetAll(c echo.Context) error
	Create(c echo.Context) error
	Update(c echo.Context) error
	Delete(c echo.Context) error
}
`

	}
	return `package ` + name + `controller

import "github.com/labstack/echo/v4"

type I` + N + `Controller interface {
	Get` + N + `s(c echo.Context) error
	Create` + N + `(c echo.Context) error
	Update` + N + `(c echo.Context) error
}
`
}

func Controller(name string, crud bool) string {

	nn := util.NewName(name)
	N := nn.Title()
	name = nn.Package()
	module := goModule()
	if crud {
		return `package ` + name + `controller

import (
	"` + module + `/internal/modules/` + name + `/` + name + `service"
)

type ` + N + `Controller struct {
	Service ` + name + `service.I` + N + `Service
}

func New` + N + `Controller(
	service ` + name + `service.I` + N + `Service) I` + N + `Controller {
	return &` + N + `Controller{
		Service: service,
	}
}
`
	}
	return `package ` + name + `controller

import "` + goModule() + `/internal/modules/` + name + `/` + name + `service"

type ` + N + `Controller struct {
	Service ` + name + `service.I` + N + `Service
}

func New` + N + `Controller(
	service ` + name + `service.I` + N + `Service,
) I` + N + `Controller {
	return &` + N + `Controller{
		Service: service,
	}
}
`
}

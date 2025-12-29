package template

import (
	"github.com/drewjya/forge/internal/util"
)

func Router(name string, crud bool) string {
	N := util.Title(name)
	module := goModule()
	if crud {

		p := name

		return `package ` + p + `routes

import (
		"` + module + `/internal/modules/` + p + `/` + p + `controller"
		"github.com/labstack/echo/v4"
)

type ` + N + `Router struct {
		handler ` + p + `controller.I` + N + `Controller
}

func New` + N + `Router(
		handler ` + p + `controller.I` + N + `Controller,
) *` + N + `Router {
		return &` + N + `Router{
			handler: handler,
		}
}

func (r *` + N + `Router) RegisterRoutes(group *echo.Group) {
		res := group.Group("/` + p + `s")

		res.GET("", r.handler.GetAll)
		res.POST("", r.handler.Create)
		res.PUT("/:id", r.handler.Update)
		res.DELETE("/:id", r.handler.Delete)
}
`
	}
	return `package ` + name + `routes

import (
	"` + goModule() + `/internal/modules/` + name + `/` + name + `controller"
	"github.com/labstack/echo/v4"
)

type ` + N + `Router struct {
	handler ` + name + `controller.I` + N + `Controller
}

func New` + N + `Router(
	handler ` + name + `controller.I` + N + `Controller,
) *` + N + `Router {
	return &` + N + `Router{
		handler: handler,
	}
}

func (r *` + N + `Router) RegisterRoutes(group *echo.Group) {
}
`
}

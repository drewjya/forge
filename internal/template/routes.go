package template

func Routes() string {
	return `package routes

import (
	"github.com/labstack/echo/v4"

	// ::forge:imports::
)

type RouteDeps struct {
	fx.In
	// ::forge:deps::
}

func RegisterRoutes(
	e *echo.Echo,
	deps RouteDeps,
) {
	api := e.Group("/api")

	// ::forge:routes::
}
`
}

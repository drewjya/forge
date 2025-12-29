package template

func RoutesModule() string {
	return `package routes

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Invoke(RegisterRoutes),
)
`
}

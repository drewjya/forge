package template

func BootstrapModule(module string) string {
	return `package bootstrap

import (
	"go.uber.org/fx"
	"` + module + `/internal/routes"
)

var Module = fx.Options(
	routes.Module,
)
`
}

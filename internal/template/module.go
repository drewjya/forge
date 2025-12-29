package template

import (
	"github.com/drewjya/forge/internal/util"
)

func Module(name string) string {
	nn := util.NewName(name)
	n := nn.Title()
	name = nn.Package()
	m := goModule()

	return `package ` + name + `

import (
	"go.uber.org/fx"

	"` + m + `/internal/modules/` + name + `/` + name + `controller"
	"` + m + `/internal/modules/` + name + `/` + name + `repository"
	"` + m + `/internal/modules/` + name + `/` + name + `service"
)

var Module = fx.Options(
	fx.Provide(
		fx.Annotate(
			` + name + `repository.New` + n + `Repository,
			fx.As(new(` + name + `repository.I` + n + `Repository)),
		),
		fx.Annotate(
			` + name + `service.New` + n + `Service,
			fx.As(new(` + name + `service.I` + n + `Service)),
		),
		fx.Annotate(
			` + name + `controller.New` + n + `Controller,
			fx.As(new(` + name + `controller.I` + n + `Controller)),
		),
	),
)
`
}

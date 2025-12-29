package template

func LogModule() string {
	return `package log

import "go.uber.org/fx"

var Module = fx.Provide(NewLogger)
`
}

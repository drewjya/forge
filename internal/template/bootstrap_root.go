package template

func BootstrapRoot(module string) string {
	return `package bootstrap

import (
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"

	"` + module + `/bootstrap/config"
	"` + module + `/bootstrap/db"
	"` + module + `/bootstrap/log"
	"` + module + `/bootstrap/server"

	//do not remove
	// ::forge:imports::
)

var Module = fx.Options(
	//do not remove
	// ::forge:modules::
	config.Module,
	log.Module,
	fx.WithLogger(func(log *zap.Logger) fxevent.Logger {
		return &fxevent.ZapLogger{
			Logger: log.Named("fx").
				WithOptions(zap.IncreaseLevel(zap.ErrorLevel)),
		}
	}),
	db.Module,
	server.Module,
)
`
}

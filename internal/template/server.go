package template

func Server(module string) string {
	return `package server

import (
	"context"
	"fmt"

	"` + module + `/bootstrap/config"
	"` + module + `/bootstrap/validation"
	"` + module + `/internal/routes"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i any) error {
	return cv.validator.Struct(i)
}

func NewEcho() *echo.Echo {
	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}
	e.HTTPErrorHandler = validation.ErrorHandler

	e.Use(middleware.Recover())
	e.Use(middleware.RequestID())
	e.Use(middleware.CORS())

	return e
}

func StartServer(
	lc fx.Lifecycle,
	e *echo.Echo,
	logger *zap.Logger,
	cfg *config.Config,
) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				logger.Info(fmt.Sprintf(
					"🚀 Echo server starting on :%s",
					cfg.Port,
				))
				if err := e.Start(":" + cfg.Port); err != nil {
					logger.Fatal("Echo failed", zap.Error(err))
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return e.Shutdown(ctx)
		},
	})
}

var Module = fx.Options(
	routes.Module,
	fx.Provide(NewEcho),
	fx.Invoke(StartServer),
)
`
}

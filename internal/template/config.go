package template

func Config() string {
	return `package config

type Config struct {
	Port         string
	DatabaseDSN string
}
`
}

func ConfigModule() string {
	return `package config

import (
	"os"

	"go.uber.org/fx"
)

func NewConfig() *Config {
	return &Config{
		Port:         env("PORT", "8080"),
		DatabaseDSN: env("DATABASE_DSN", ""),
	}
}

func env(k, d string) string {
	if v := os.Getenv(k); v != "" {
		return v
	}
	return d
}

var Module = fx.Provide(NewConfig)
`
}

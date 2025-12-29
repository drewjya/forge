package template

func Log() string {
	return `package log

import "go.uber.org/zap"

func NewLogger() (*zap.Logger, error) {
	return zap.NewProduction()
}
`
}

package template

func Main(module string) string {
	return `package main

import (
	"go.uber.org/fx"
	"` + module + `/bootstrap"
)

func main() {
	fx.New(
		bootstrap.Module,
	).Run()
}
`
}

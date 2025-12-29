package template

func GoMod(module string) string {
	return `module ` + module + `

go 1.22

require (
	go.uber.org/fx v1.23.0
	github.com/labstack/echo/v4 v4.12.0
	github.com/jackc/pgx/v5 v5.7.1
)
`
}

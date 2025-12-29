package template

func BobgenYAML() string {
	return `psql:
  dsn: ${DATABASE_URL}
  output: "gen/models"
  driver_name: "github.com/jackc/pgx/v5"

  models:
    generate: true

tags: ["json"]
struct_tag_casing: "camel"

no_tests: true
plugins_preset: "all" # Valid values are "default", "all" or "none".
plugins:
  dbinfo:
    disabled: false
    pkgname: "dbinfo"
    destination: "gen/dbinfo"
  enums:
    disabled: false
    pkgname: "enums"
    destination: "gen/enums"
  models:
    disabled: false
    pkgname: "models"
    destination: "gen/models"
  factory:
    disabled: false
    pkgname: "factory"
    destination: "gen/factory"
  dberrors:
    disabled: false
    pkgname: "dberrors"
    destination: "gen/dberrors"
  where:
    disabled: false
  loaders:
    disabled: false
  joins:
    disabled: false

`
}

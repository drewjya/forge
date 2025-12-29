package template

func MakeFile() string {
	return `load:
	migrate -path migrations -database "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable" up
bobgen:
	go run github.com/stephenafamo/bob/gen/bobgen-psql@latest -c bobgen.yaml
run:
	go run cmd/api/main.go
build:
	go build cmd/api/main.go

`
}

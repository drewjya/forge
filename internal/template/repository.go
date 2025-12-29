package template

import (
	"github.com/andre/forge/internal/util"
)

func RepositoryInterface(name string, crud bool) string {
	N := util.Title(name)
	module := goModule()
	if crud {
		return `package ` + name + `repository

import (
		"context"

	"` + module + `/gen/models"
	"` + module + `/internal/types"

	"github.com/stephenafamo/bob"
	"github.com/stephenafamo/bob/dialect/psql/dialect"
)

type I` + N + `Repository interface {
	GetAll(ctx context.Context, mods ...bob.Mod[*dialect.SelectQuery]) (models.` + N + `Slice, error)
	GetByID(ctx context.Context, id types.ULID) *models.` + N + `
	Create(ctx context.Context, setter models.` + N + `Setter) error
	Update(ctx context.Context, id types.ULID, setter models.` + N + `Setter) error
	Delete(ctx context.Context, id types.ULID) error
}
`

	}

	return `package ` + name + `repository

type I` + N + `Repository interface {
}
`
}

func Repository(name string, crud bool) string {
	N := util.Title(name)
	if crud {
		return `package ` + name + `repository

import (
	"github.com/stephenafamo/bob"
)

type ` + N + `Repository struct {
	Executor bob.Executor
}

func New` + N + `Repository(
	executor bob.Executor,
) I` + N + `Repository {
	return &` + N + `Repository{
		Executor: executor,
	}
}
`

	}
	return `package ` + name + `repository

type ` + N + `Repository struct {
}

func New` + N + `Repository() I` + N + `Repository {
	return &` + N + `Repository{}
}
`
}

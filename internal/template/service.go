package template

import (
	"github.com/drewjya/forge/internal/util"
)

func ServiceInterface(name string, crud bool) string {
	nn := util.NewName(name)
	N := nn.Title()
	name = nn.Package()
	module := goModule()
	if crud {
		return `package ` + name + `service

import (
	"context"

	"` + module + `/gen/models"
	"` + module + `/internal/request"
	"` + module + `/internal/types"
)

type I` + N + `Service interface {
	GetAll(ctx context.Context) (models.` + N + `Slice, error)
	Create(ctx context.Context, req request.` + N + `Request) error
	Update(ctx context.Context, id types.ULID, req request.` + N + `Request) error
	Delete(ctx context.Context, id types.ULID) error
}
`
	}
	return `package ` + name + `service

type I` + N + `Service interface {
}
`
}

func Service(name string, crud bool) string {
	nn := util.NewName(name)
	N := nn.Title()
	name = nn.Package()
	module := goModule()
	if crud {
		return `package ` + name + `service

import (
	"` + module + `/internal/modules/` + name + `/` + name + `repository"
)

type ` + N + `Service struct {
	Repository ` + name + `repository.I` + N + `Repository
}

func New` + N + `Service(
	repo ` + name + `repository.I` + N + `Repository,
) I` + N + `Service {
	return &` + N + `Service{
		Repository: repo,
	}
}
`
	}
	return `package ` + name + `service

import "` + goModule() + `/internal/modules/` + name + `/` + name + `repository"

type ` + N + `Service struct {
	Repository ` + name + `repository.I` + N + `Repository
}

func New` + N + `Service(
	repo ` + name + `repository.I` + N + `Repository,
) I` + N + `Service {
	return &` + N + `Service{
		Repository: repo,
	}
}
`
}

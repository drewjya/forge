package template

import (
	"github.com/andre/forge/internal/util"
)

func ServiceCRUD(name string) string {
	N := util.Title(name)
	module := goModule()
	p := name

	return `package ` + p + `service

import (
	"context"

	"` + module + `/gen/models"
	"` + module + `/internal/request"
	"` + module + `/internal/types"
)

func (s *` + N + `Service) GetAll(ctx context.Context) (models.` + N + `Slice, error) {
	return s.Repository.GetAll(ctx)
}

func (s *` + N + `Service) Create(ctx context.Context, req request.` + N + `Request) error {
	setter := models.` + N + `Setter{
		Name: req.Name,
	}
	return s.Repository.Create(ctx, setter)
}

func (s *` + N + `Service) Update(ctx context.Context, id types.ULID, req request.` + N + `Request) error {
	setter := models.` + N + `Setter{
		Name: req.Name,
	}
	return s.Repository.Update(ctx, id, setter)
}

func (s *` + N + `Service) Delete(ctx context.Context, id types.ULID) error {
	return s.Repository.Delete(ctx, id)
}
`
}

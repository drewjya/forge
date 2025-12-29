package template

import (
	"github.com/drewjya/forge/internal/util"
)

func RepositoryCRUD(name string) string {
	module := goModule()
	nn := util.NewName(name)
	N := nn.Title()
	name = nn.Package()
	P := nn.TitlePlural()

	p := name

	return `package ` + p + `repository

import (
	"context"

	"` + module + `/gen/models"
	"` + module + `/internal/types"

	"github.com/stephenafamo/bob"
	"github.com/stephenafamo/bob/dialect/psql/dialect"
)

func (r *` + N + `Repository) GetAll(ctx context.Context, mods ...bob.Mod[*dialect.SelectQuery]) (models.` + N + `Slice, error) {
	return models.` + P + `.Query(mods...).All(ctx, r.Executor)
}

func (r *` + N + `Repository) GetByID(ctx context.Context, id types.ULID) *models.` + N + ` {
	row, _ := models.` + P + `.Query(
		models.SelectWhere.` + P + `.ID.EQ(id.String()),
	).One(ctx, r.Executor)
	return row
}

func (r *` + N + `Repository) Create(ctx context.Context, setter models.` + N + `Setter) error {
	_, err := models.` + P + `.Insert(&setter).One(ctx, r.Executor)
	return err
}

func (r *` + N + `Repository) Update(ctx context.Context, id types.ULID, setter models.` + N + `Setter) error {
	_, err := models.` + P + `.Update(
		models.UpdateWhere.` + P + `.ID.EQ(id.String()),
		setter.UpdateMod(),
	).One(ctx, r.Executor)
	return err
}

func (r *` + N + `Repository) Delete(ctx context.Context, id types.ULID) error {
	_, err := models.` + P + `.Delete(
		models.DeleteWhere.` + P + `.ID.EQ(id.String()),
	).Exec(ctx, r.Executor)
	return err
}
`
}

package template

func DB(module string) string {

	return `package db

import (
	"context"
	"log"
	"time"

	"` + module + `/bootstrap/config"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/stephenafamo/bob"
	"go.uber.org/fx"
)

func NewPGXPool(
	lc fx.Lifecycle,
	cfg *config.Config,
) (*pgxpool.Pool, error) {
	start := time.Now()

	pool, err := pgxpool.New(context.Background(), cfg.DatabaseDSN)
	if err != nil {
		return nil, err
	}

	log.Println("DB connected in", time.Since(start))

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			return pool.Ping(ctx)
		},
		OnStop: func(ctx context.Context) error {
			pool.Close()
			return nil
		},
	})

	return pool, nil
}

func NewBobExecutor(pool *pgxpool.Pool) bob.Executor {
	conn := stdlib.OpenDBFromPool(pool)
	return bob.Debug(bob.NewDB(conn))
}

var Module = fx.Module(
	"db",
	fx.Provide(
		NewPGXPool,
		NewBobExecutor,
	),
)
`
}

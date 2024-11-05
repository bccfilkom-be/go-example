package postgresql

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func NewConn(ctx context.Context, cfg *pgx.ConnConfig) (DBTX, error) {
	conn, err := pgx.ConnectConfig(context.Background(), cfg)
	if err != nil {
		return nil, err
	}
	if err := conn.Ping(ctx); err != nil {
		return nil, err
	}
	return conn, nil
}

func NewPool(ctx context.Context, cfg *pgxpool.Config) (DBTX, error) {
	pool, err := pgxpool.NewWithConfig(ctx, cfg)
	if err != nil {
		return nil, err
	}
	if err := pool.Ping(ctx); err != nil {
		return nil, err
	}
	return pool, nil
}

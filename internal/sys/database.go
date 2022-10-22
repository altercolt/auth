package sys

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
)

func Postgres(ctx context.Context, DSN string) (*pgxpool.Pool, error) {
	db, err := pgxpool.New(ctx, DSN)
	if err != nil {
		return nil, err
	}

	return db, nil
}

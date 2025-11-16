package postgres

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewDB(ctx context.Context, conn string) (*pgxpool.Pool, error) {
	return pgxpool.New(ctx, conn)
}

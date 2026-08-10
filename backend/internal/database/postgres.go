package database

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewPostgresPool(ctx context.Context, databaseURL string) (*pgxpool.Pool, error) {
	if databaseURL == "" {
		return nil, fmt.Errorf("DATABASE_URL is not set")
	}

	pool, poolError := pgxpool.New(ctx, databaseURL)
	if poolError != nil {
		return nil, fmt.Errorf("Create db pool: %w", poolError)
	}

	if pingError := pool.Ping(ctx); pingError != nil {
		pool.Close()
		return nil, fmt.Errorf("Ping db: %w", pingError)
	}

	return pool, nil
}

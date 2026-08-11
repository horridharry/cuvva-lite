package policy

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5/pgxpool"
)

var ErrAlreadyPurchased = errors.New("quote already purchased")

type Repository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) Create(
	ctx context.Context,
	p Policy,
) (Policy, error) {
	const query = `
		INSERT INTO policies (
			quote_id,
			vehicle_id,
			starts_at,
			ends_at
		)
		VALUES ($1, $2, $3, $4)
		RETURNING
			id,
			created_at
	`

	err := r.db.QueryRow(
		ctx,
		query,
		p.QuoteID,
		p.VehicleID,
		p.StartsAt,
		p.EndsAt,
	).Scan(
		&p.ID,
		&p.CreatedAt,
	)

	if err != nil {
		return Policy{}, err
	}

	return p, nil
}

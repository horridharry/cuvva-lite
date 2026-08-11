package payment

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{
		db: db,
	}
}

func (repository *Repository) Create(
	ctx context.Context,
	p Payment,
) (Payment, error) {
	const query = `
		INSERT INTO payments (
			quote_id,
			amount_pence,
			status
		)
		VALUES ($1, $2, $3)
		RETURNING
			id,
			created_at
	`

	err := repository.db.QueryRow(
		ctx,
		query,
		p.QuoteID,
		p.AmountPence,
		p.Status,
	).Scan(
		&p.ID,
		&p.CreatedAt,
	)

	if err != nil {
		return Payment{}, err
	}

	return p, nil
}

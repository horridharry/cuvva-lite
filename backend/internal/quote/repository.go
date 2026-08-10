package quote

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
	q Quote,
) (Quote, error) {
	const query = `
		INSERT INTO quotes (
			vehicle_id,
			driver_age,
			years_licensed,
			penalty_points,
			duration_minutes,
			price_pence,
			expires_at
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING
			id,
			created_at
	`

	queryError := repository.db.QueryRow(
		ctx,
		query,
		q.VehicleID,
		q.DriverAge,
		q.YearsLicensed,
		q.PenaltyPoints,
		q.DurationMinutes,
		q.PricePence,
		q.ExpiresAt,
	).Scan(
		&q.ID,
		&q.CreatedAt,
	)

	if queryError != nil {
		return Quote{}, queryError
	}

	return q, nil
}

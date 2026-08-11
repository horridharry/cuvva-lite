package quote

import (
	"context"
	"errors"
	"log"

	"github.com/jackc/pgx/v5"
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
			starts_at,
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
		&q.StartsAt,
		&q.CreatedAt,
	)

	if queryError != nil {
		return Quote{}, queryError
	}

	return q, nil
}

func (r *Repository) FindByID(
	ctx context.Context,
	id int64,
) (Quote, error) {
	const query = `
		SELECT
			id,
			vehicle_id,
			driver_age,
			years_licensed,
			penalty_points,
			duration_minutes,
			price_pence,
			starts_at,
			expires_at,
			created_at
		FROM quotes
		WHERE id = $1
	`

	var q Quote

	err := r.db.QueryRow(ctx, query, id).Scan(
		&q.ID,
		&q.VehicleID,
		&q.DriverAge,
		&q.YearsLicensed,
		&q.PenaltyPoints,
		&q.DurationMinutes,
		&q.PricePence,
		&q.StartsAt,
		&q.ExpiresAt,
		&q.CreatedAt,
	)

	log.Printf("FindByID id=%d err=%v quote=%+v", id, err, q)

	if errors.Is(err, pgx.ErrNoRows) {
		return Quote{}, InvalidQuoteRequest
	}

	if err != nil {
		return Quote{}, err
	}

	return q, nil
}

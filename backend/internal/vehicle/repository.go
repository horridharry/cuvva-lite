package vehicle

import (
	"context"
	"errors"
	"strings"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

var VehicleNotFound = errors.New("Vehicle not found")

type Repository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) FindByRegistration(
	ctx context.Context,
	registration string,
) (Vehicle, error) {
	registration = strings.ToUpper(
		strings.ReplaceAll(registration, " ", ""),
	)

	const query = `
		SELECT
			id,
			registration,
			make,
			model,
			year,
			fuel_type,
			engine_size_cc
		FROM vehicles
		WHERE reg = $1
	`

	var vehicle Vehicle

	queryError := r.db.QueryRow(ctx, query, registration).Scan(
		&vehicle.ID,
		&vehicle.Registration,
		&vehicle.Make,
		&vehicle.Model,
		&vehicle.Year,
		&vehicle.FuelType,
		&vehicle.EngineSizeCC,
	)

	if errors.Is(queryError, pgx.ErrNoRows) {
		return Vehicle{}, VehicleNotFound
	}

	if queryError != nil {
		return Vehicle{}, queryError
	}

	return vehicle, nil
}

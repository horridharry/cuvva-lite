package quote

import (
	"context"
	"errors"
	"time"

	"github.com/horridharry/cuvva-lite/backend/internal/pricing"
	"github.com/horridharry/cuvva-lite/backend/internal/vehicle"
)

var InvalidQuoteRequest = errors.New("Invalid quote request")

type Service struct {
	repository        *Repository
	vehicleRepository *vehicle.Repository
}

func NewService(
	repository *Repository,
	vehicleRepository *vehicle.Repository,
) *Service {
	return &Service{
		repository:        repository,
		vehicleRepository: vehicleRepository,
	}
}

func (s *Service) Create(
	ctx context.Context,
	request CreateRequest,
) (Quote, error) {
	if request.DriverAge < 18 ||
		request.DriverAge > 80 ||
		request.YearsLicensed < 0 ||
		request.PenaltyPoints < 0 {
		return Quote{}, InvalidQuoteRequest
	}

	if !validDuration(request.DurationMinutes) {
		return Quote{}, InvalidQuoteRequest
	}

	v, err := s.vehicleRepository.FindByID(
		ctx,
		request.VehicleID,
	)

	if err != nil {
		return Quote{}, err
	}

	pricePence := pricing.Calculate(pricing.Input{
		DriverAge:       request.DriverAge,
		YearsLicensed:   request.YearsLicensed,
		PenaltyPoints:   request.PenaltyPoints,
		EngineSizeCC:    v.EngineSizeCC,
		DurationMinutes: request.DurationMinutes,
	})

	q := Quote{
		VehicleID:       v.ID,
		DriverAge:       request.DriverAge,
		YearsLicensed:   request.YearsLicensed,
		PenaltyPoints:   request.PenaltyPoints,
		DurationMinutes: request.DurationMinutes,
		PricePence:      pricePence,

		// Quote is purchasable for 15 minutes.
		ExpiresAt: time.Now().Add(15 * time.Minute),
	}

	return s.repository.Create(ctx, q)
}

func validDuration(minutes int) bool {
	switch minutes {
	case 60, 180, 360, 1440:
		return true
	default:
		return false
	}
}

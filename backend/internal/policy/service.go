package policy

import (
	"context"
	"errors"
	"time"

	"github.com/horridharry/cuvva-lite/backend/internal/payment"
	"github.com/horridharry/cuvva-lite/backend/internal/quote"
)

var ErrQuoteExpired = errors.New("quote expired")
var ErrPaymentDeclined = errors.New("payment declined")

type Service struct {
	repository     *Repository
	paymentService *payment.Service
}

func NewService(
	repository *Repository,
	paymentService *payment.Service,
) *Service {
	return &Service{
		repository:     repository,
		paymentService: paymentService,
	}
}

func (service *Service) CreateFromQuote(
	ctx context.Context,
	q quote.Quote,
	paymentMethod string,
) (Policy, error) {
	if time.Now().After(q.ExpiresAt) {
		return Policy{}, ErrQuoteExpired
	}

	paymentStatus := service.paymentService.Authorise(paymentMethod)

	if paymentStatus != payment.StatusSucceeded {
		return Policy{}, ErrPaymentDeclined
	}

	startsAt := q.StartsAt

	endsAt := startsAt.Add(
		time.Duration(q.DurationMinutes) * time.Minute,
	)

	p := Policy{
		QuoteID:   q.ID,
		VehicleID: q.VehicleID,
		StartsAt:  startsAt,
		EndsAt:    endsAt,
	}

	return service.repository.Create(ctx, p)
}

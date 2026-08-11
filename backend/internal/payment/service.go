package payment

import "context"

type Service struct {
	repository *Repository
}

func NewService(repository *Repository) *Service {
	return &Service{
		repository: repository,
	}
}

func (service *Service) Authorise(
	ctx context.Context,
	quoteID int64,
	amountPence int,
	paymentMethod string,
) (Payment, error) {
	status := StatusDeclined

	if paymentMethod == "4242" {
		status = StatusSucceeded
	}

	p := Payment{
		QuoteID:     quoteID,
		AmountPence: amountPence,
		Status:      status,
	}

	return service.repository.Create(ctx, p)
}

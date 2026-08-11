package payment

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) Authorise(paymentMethod string) Status {
	switch paymentMethod {
	case "4242":
		return StatusSucceeded
	default:
		return StatusDeclined
	}
}

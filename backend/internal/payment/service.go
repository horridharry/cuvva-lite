package payment

func Authorise(method string) Status {
	switch method {
	case "4242":
		return StatusSucceeded

	case "0002":
		return StatusDeclined

	default:
		return StatusDeclined
	}
}

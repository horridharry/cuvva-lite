package payment

type Status string

const (
	StatusSucceeded Status = "SUCCEEDED"
	StatusDeclined  Status = "DECLINED"
)

type Payment struct {
	ID          int64  `json:"id"`
	QuoteID     int64  `json:"quoteId"`
	AmountPence int    `json:"amountPence"`
	Status      Status `json:"status"`
}

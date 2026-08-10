package policy

import "time"

type Policy struct {
	ID        int64     `json:"id"`
	QuoteID   int64     `json:"quoteId"`
	VehicleID int64     `json:"vehicleId"`
	StartsAt  time.Time `json:"startsAt"`
	EndsAt    time.Time `json:"endsAt"`
	CreatedAt time.Time `json:"createdAt"`
}

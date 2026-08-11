package quote

import "time"

type Quote struct {
	ID              int64     `json:"id"`
	VehicleID       int64     `json:"vehicleId"`
	DriverAge       int       `json:"driverAge"`
	YearsLicensed   int       `json:"yearsLicensed"`
	PenaltyPoints   int       `json:"penaltyPoints"`
	DurationMinutes int       `json:"durationMinutes"`
	PricePence      int       `json:"pricePence"`
	StartsAt        time.Time `json:"startsAt"`
	ExpiresAt       time.Time `json:"expiresAt"`
	CreatedAt       time.Time `json:"createdAt"`
}

type CreateRequest struct {
	VehicleID       int64 `json:"vehicleId"`
	DriverAge       int   `json:"driverAge"`
	YearsLicensed   int   `json:"yearsLicensed"`
	PenaltyPoints   int   `json:"penaltyPoints"`
	DurationMinutes int   `json:"durationMinutes"`
}

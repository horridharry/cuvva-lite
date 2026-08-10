package pricing

import "math"

type Input struct {
	DriverAge       int
	YearsLicensed   int
	PenaltyPoints   int
	EngineSizeCC    int
	DurationMinutes int
}

func Calculate(input Input) int {
	basePence := 500

	hours := int(math.Ceil(float64(input.DurationMinutes) / 60))
	basePence += hours * 250

	multiplier := 1.0

	switch {
	case input.DriverAge < 21:
		multiplier *= 1.8
	case input.DriverAge <= 25:
		multiplier *= 1.4
	}

	switch {
	case input.YearsLicensed < 2:
		multiplier *= 1.5
	case input.YearsLicensed < 5:
		multiplier *= 1.2
	}

	switch {
	case input.PenaltyPoints >= 6:
		multiplier *= 1.5
	case input.PenaltyPoints >= 3:
		multiplier *= 1.2
	}

	switch {
	case input.EngineSizeCC > 2000:
		multiplier *= 1.25
	case input.EngineSizeCC > 1600:
		multiplier *= 1.1
	}

	price := float64(basePence) * multiplier

	return int(math.Round(price))
}

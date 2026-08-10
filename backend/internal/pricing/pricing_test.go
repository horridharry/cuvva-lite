package pricing

import "testing"

func TestCalculate_HigherRiskDriverCostsMore(t *testing.T) {
	lowRisk := Calculate(Input{
		DriverAge:       35,
		YearsLicensed:   15,
		PenaltyPoints:   0,
		EngineSizeCC:    1200,
		DurationMinutes: 180,
	})

	highRisk := Calculate(Input{
		DriverAge:       20,
		YearsLicensed:   1,
		PenaltyPoints:   6,
		EngineSizeCC:    2500,
		DurationMinutes: 180,
	})

	if highRisk <= lowRisk {
		t.Fatalf(
			"expected high-risk quote %d to exceed low-risk quote %d",
			highRisk,
			lowRisk,
		)
	}
}

func TestCalculate_IsDeterministic(t *testing.T) {
	input := Input{
		DriverAge:       24,
		YearsLicensed:   4,
		PenaltyPoints:   0,
		EngineSizeCC:    1798,
		DurationMinutes: 180,
	}

	first := Calculate(input)
	second := Calculate(input)

	if first != second {
		t.Fatalf(
			"expected same input to produce same price: %d != %d",
			first,
			second,
		)
	}
}

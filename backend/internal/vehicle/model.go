package vehicle

type Vehicle struct {
	ID           int64  `json:"id"`
	Registration string `json:"Registration"`
	Make         string `json:"make"`
	Model        string `json:"model"`
	Year         int    `json:"year"`
	FuelType     string `json:"fuelType"`
	EngineSizeCC int    `json:"engineSizeCc"`
}

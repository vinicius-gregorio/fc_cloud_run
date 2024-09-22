package response

type TemperatureResponse struct {
	TempC float64 `json:"temp_C"`
	TempF float64 `json:"temp_F"`
	TempK float64 `json:"temp_K"`
}

// ErrorResponse is a struct to capture error messages from the API response.
type ErrorResponse struct {
	Erro string `json:"erro"`
}

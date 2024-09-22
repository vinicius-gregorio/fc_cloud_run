package entity

import "github.com/vinicius-gregorio/fc_cloud_run/internal/failures"

type CurrentWeather struct {
	LastUpdatedEpoch int     `json:"last_updated_epoch"`
	LastUpdated      string  `json:"last_updated"`
	TempC            float64 `json:"temp_c"`
	TempF            float64 `json:"temp_f"`
	TempK            float64 `json:"temp_k"`
}

func NewCurrentWeather(lastUpdatedEpoch int, lastUpdated string, tempC, tempF float64) (*CurrentWeather, error) {
	tk := tempC + 273.15

	cw := &CurrentWeather{
		LastUpdatedEpoch: lastUpdatedEpoch,
		LastUpdated:      lastUpdated,
		TempC:            tempC,
		TempF:            tempF,
		TempK:            tk,
	}
	if err := cw.validate(); err != nil {
		return nil, err
	}

	return cw, nil

}

func (cw *CurrentWeather) validate() error {
	if cw.TempC < -100 {
		return failures.ErrInvalidTemperatureCelsius_LessThan100
	}
	if cw.TempC > 100 {
		return failures.ErrInvalidTemperatureCelsius_GreaterThan100
	}
	if cw.TempF < -148 {
		return failures.ErrInvalidTemperatureFahrenheit_LessThan148
	}
	if cw.TempF > 212 {
		return failures.ErrInvalidTemperatureFahrenheit_GreaterThan212
	}
	if cw.TempK < 0 {
		return failures.ErrInvalidTemperatureKelvin_LessThan0
	}
	if cw.TempK > 373 {
		return failures.ErrInvalidTemperatureKelvin_GreaterThan373
	}

	return nil

}

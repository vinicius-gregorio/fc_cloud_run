package entity

import "github.com/vinicius-gregorio/fc_cloud_run/internal/failures"

type LocationWeather struct {
	Name      string `json:"name"`
	Region    string `json:"region"`
	Country   string `json:"country"`
	LocalTime string `json:"localtime"`
	CEP       string `json:"cep"`
}

func NewLocationWheather(name, region, country, localtime string) (*LocationWeather, error) {
	lw := &LocationWeather{
		Name:      name,
		Region:    region,
		Country:   country,
		LocalTime: localtime,
	}

	if err := lw.validate(); err != nil {
		return nil, err
	}

	return lw, nil
}

func (lw *LocationWeather) validate() error {
	if lw.Name == "" {
		return failures.ErrNameIsRequired
	}
	if lw.Region == "" {
		return failures.ErrRegionIsRequired
	}
	if lw.Country == "" {
		return failures.ErrCountryIsRequired
	}
	if lw.LocalTime == "" {
		return failures.ErrLocalTimeIsRequired
	}

	return nil
}

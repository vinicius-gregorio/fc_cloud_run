package external_repository

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/vinicius-gregorio/fc_cloud_run/config"
	"github.com/vinicius-gregorio/fc_cloud_run/internal/entity"
)

// WeatherRepositoryImpl is the implementation of the WeatherRepository interface.
type WeatherRepositoryImpl struct {
	config *config.EnvConfig
}

// NewWeatherRepository creates a new instance of WeatherRepositoryImpl with the given config.
func NewWeatherRepositoryImpl(cfg *config.EnvConfig) *WeatherRepositoryImpl {
	return &WeatherRepositoryImpl{config: cfg}
}

// GetLocationInfoByCep fetches location information using the CEP API.
func (repo *WeatherRepositoryImpl) GetLocationInfoByCep(cep string) (*entity.Location, error) {
	var loc entity.Location
	url := fmt.Sprintf("%s/%s/json/", repo.config.CEPAPIURL, cep)
	fmt.Println(url)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if err := json.NewDecoder(resp.Body).Decode(&loc); err != nil {
		return nil, err
	}

	return &loc, nil

}

// GetWeatherByLocation fetches weather information using the Weather API.
func (repo *WeatherRepositoryImpl) GetWeatherByLocation(location entity.Location) (*entity.Weather, error) {
	var Weather entity.Weather

	// URL-encode each component of the location query string
	localidade := url.QueryEscape(location.Localidade)
	uf := url.QueryEscape(location.Estado)
	country := url.QueryEscape("Brazil")

	// Construct the URL using the encoded parameters
	url := fmt.Sprintf("%s/v1/current.json?q=%s,%s,%s&key=%s",
		repo.config.WeatherAPIURL, localidade, uf, country, repo.config.WeatherAPIKey)

	fmt.Println(url)

	// Create the HTTP request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	// Perform the HTTP request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Decode the JSON response into the Weather struct
	if err := json.NewDecoder(resp.Body).Decode(&Weather); err != nil {
		return nil, err
	}

	// Process the weather data to create entities
	lw, err := entity.NewLocationWheather(Weather.Location.Country, Weather.Location.Region, "Brazil", Weather.Location.LocalTime)
	if err != nil {
		return nil, err
	}
	cw, err := entity.NewCurrentWeather(Weather.Current.LastUpdatedEpoch, Weather.Current.LastUpdated, Weather.Current.TempC, Weather.Current.TempF)
	if err != nil {
		return nil, err
	}
	w, err := entity.NewWeather(*lw, *cw)
	if err != nil {
		return nil, err
	}

	return w, nil
}

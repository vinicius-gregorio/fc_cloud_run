package external_repository

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/vinicius-gregorio/fc_cloud_run/config"
	"github.com/vinicius-gregorio/fc_cloud_run/external/response"
	"github.com/vinicius-gregorio/fc_cloud_run/internal/entity"
	"github.com/vinicius-gregorio/fc_cloud_run/internal/failures"
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
// GetLocationInfoByCep fetches location information using the CEP API.
func (repo *WeatherRepositoryImpl) GetLocationInfoByCep(cep string) (*entity.Location, error) {
	var loc entity.Location
	var errorResp response.ErrorResponse

	// Construct the request URL
	url := fmt.Sprintf("%s/%s/json/", repo.config.CEPAPIURL, cep)

	// Create the HTTP request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	// Perform the HTTP request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Read the entire response body to handle potential errors and decoding issues
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %v", err)
	}

	// Print the response body for debugging purposes
	bodyString := string(bodyBytes)
	fmt.Println("Response Body:", bodyString)

	// Attempt to unmarshal the response into the error struct first
	if err := json.Unmarshal(bodyBytes, &errorResp); err == nil && errorResp.Erro == "true" {
		return nil, failures.ErrCepNotFound
	}

	// Attempt to unmarshal the response into the Location struct if no errors were detected
	if err := json.Unmarshal(bodyBytes, &loc); err != nil {
		return nil, fmt.Errorf("failed to decode location response: %v", err)
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

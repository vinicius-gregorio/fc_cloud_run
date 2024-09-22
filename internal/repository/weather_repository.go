package repository

import "github.com/vinicius-gregorio/fc_cloud_run/internal/entity"

type WeatherRepository interface {
	GetLocationInfoByCep(cep string) (*entity.Location, error)
	GetWeatherByLocation(location entity.Location) (*entity.Weather, error)
}

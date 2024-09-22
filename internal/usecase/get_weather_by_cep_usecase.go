package usecase

import (
	"github.com/vinicius-gregorio/fc_cloud_run/internal/entity"
	"github.com/vinicius-gregorio/fc_cloud_run/internal/repository"
)

type GetWeatherInputDTO struct {
	Cep string
}

type GetWeatherOutputDTO struct {
	Weather *entity.Weather
}

type GetWeatherButCEPUsecase interface {
	Call(input GetWeatherInputDTO) (*GetWeatherOutputDTO, error)
}

// WeatherUseCase struct holds the repository that will be used in the use case.
type GetWeatherUseCase struct {
	repo repository.WeatherRepository
}

// NewWeatherUseCase creates a new instance of WeatherUseCase with the given repository.
func NewGetWeatherUseCase(repo repository.WeatherRepository) *GetWeatherUseCase {
	return &GetWeatherUseCase{repo: repo}
}

// Call implements the GetWeatherButCEPUsecase interface.
func (uc *GetWeatherUseCase) Call(input GetWeatherInputDTO) (*GetWeatherOutputDTO, error) {
	// Step 1: Get location information by CEP
	location, err := uc.repo.GetLocationInfoByCep(input.Cep)
	if err != nil {
		return nil, err
	}

	// Step 2: Get weather information by location
	weather, err := uc.repo.GetWeatherByLocation(*location)
	if err != nil {
		return nil, err
	}

	// Step 3: Return the weather information in the output DTO
	return &GetWeatherOutputDTO{Weather: weather}, nil
}

package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/vinicius-gregorio/fc_cloud_run/config"
	external_repository "github.com/vinicius-gregorio/fc_cloud_run/external/repository"
	"github.com/vinicius-gregorio/fc_cloud_run/infra"
	"github.com/vinicius-gregorio/fc_cloud_run/internal/entity"
	"github.com/vinicius-gregorio/fc_cloud_run/internal/usecase"
)

func main() {

	config, err := config.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	println(config.WeatherAPIKey)
	println(config.WeatherAPIURL)
	println(config.CEPAPIURL)

	weatherRepo := external_repository.NewWeatherRepositoryImpl(config)
	weatherUseCase := usecase.NewGetWeatherUseCase(weatherRepo)

	infra.StartHTTPServer(getRoutes(weatherUseCase), config.WebServerPort)

}

func getRoutes(getWeatherUsecase usecase.GetWeatherButCEPUsecase) []infra.HTTPRoute {
	return []infra.HTTPRoute{
		{
			Path:   "/temp/{cep}",
			Method: "GET",
			Handler: func(w http.ResponseWriter, r *http.Request) {
				// get the CEP from the URL
				cep := chi.URLParam(r, "cep")

				fmt.Println("http request cep is:   |", cep, "|")
				nl, err := entity.NewLocationByCEP(cep)
				if err != nil {
					w.WriteHeader(http.StatusBadRequest)
					w.Write([]byte(err.Error()))
					return
				}
				input := usecase.GetWeatherInputDTO{
					Cep: nl.Cep,
				}
				output, err := getWeatherUsecase.Call(input)
				if err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					w.Write([]byte(err.Error()))
					return
				}
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				json.NewEncoder(w).Encode(output)

			},
		},
		{
			Path:   "/cep/{cep}",
			Method: "GET",
			Handler: func(w http.ResponseWriter, r *http.Request) {
				w.Write([]byte("Hello World"))
				w.WriteHeader(http.StatusOK)

			},
		},
		{
			Path:   "/weather",
			Method: "GET",
			Handler: func(w http.ResponseWriter, r *http.Request) {
				w.Write([]byte("Hello World"))
				w.WriteHeader(http.StatusOK)

			},
		},
	}
}

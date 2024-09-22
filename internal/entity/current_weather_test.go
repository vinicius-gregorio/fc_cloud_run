package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vinicius-gregorio/fc_cloud_run/internal/failures"
)

func TestNewCurrentWeather(t *testing.T) {
	t.Run("should return a new CurrentWeather", func(t *testing.T) {
		currentWeather, err := NewCurrentWeather(1679808000, "2023-03-25 12:00", 25.0, 77.0)
		assert.Nil(t, err)
		assert.NotNil(t, currentWeather)
		assert.Equal(t, 1679808000, currentWeather.LastUpdatedEpoch)
		assert.Equal(t, "2023-03-25 12:00", currentWeather.LastUpdated)
		assert.Equal(t, 25.0, currentWeather.TempC)
		assert.Equal(t, 77.0, currentWeather.TempF)
		assert.Equal(t, 298.15, currentWeather.TempK)
	})

	t.Run("should return an error if tempC is less than -100", func(t *testing.T) {
		currentWeather, err := NewCurrentWeather(1679808000, "2023-03-25 12:00", -101.0, 77.0)
		assert.NotNil(t, err)
		assert.Equal(t, failures.ErrInvalidTemperatureCelsius_LessThan100, err)
		assert.Nil(t, currentWeather)
	})

	t.Run("should return an error if tempC is greater than 100", func(t *testing.T) {
		currentWeather, err := NewCurrentWeather(1679808000, "2023-03-25 12:00", 101.0, 77.0)
		assert.NotNil(t, err)
		assert.Equal(t, failures.ErrInvalidTemperatureCelsius_GreaterThan100, err)
		assert.Nil(t, currentWeather)
	})

	t.Run("should return an error if tempF is less than -148", func(t *testing.T) {
		currentWeather, err := NewCurrentWeather(1679808000, "2023-03-25 12:00", 25.0, -149.0)
		assert.NotNil(t, err)
		assert.Equal(t, failures.ErrInvalidTemperatureFahrenheit_LessThan148, err)
		assert.Nil(t, currentWeather)
	})

	t.Run("should return an error if tempF is greater than 212", func(t *testing.T) {
		currentWeather, err := NewCurrentWeather(1679808000, "2023-03-25 12:00", 25.0, 213.0)
		assert.NotNil(t, err)
		assert.Equal(t, failures.ErrInvalidTemperatureFahrenheit_GreaterThan212, err)
		assert.Nil(t, currentWeather)
	})

	t.Run("should return an error if tempK is less than 0", func(t *testing.T) {
		currentWeather, err := NewCurrentWeather(1679808000, "2023-03-25 12:00", -274.15, -460.67)
		assert.NotNil(t, err)
		assert.Equal(t, failures.ErrInvalidTemperatureCelsius_LessThan100, err)
		assert.Nil(t, currentWeather)
	})

	t.Run("should return an error if tempK is greater than 373", func(t *testing.T) {
		currentWeather, err := NewCurrentWeather(1679808000, "2023-03-25 12:00", 101, 213.8)
		assert.NotNil(t, err)
		assert.Equal(t, failures.ErrInvalidTemperatureCelsius_GreaterThan100, err)
		assert.Nil(t, currentWeather)
	})
}

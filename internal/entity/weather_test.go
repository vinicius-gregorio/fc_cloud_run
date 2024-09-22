package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewWeather(t *testing.T) {
	t.Run("should return a new Weather", func(t *testing.T) {
		locationWeather, _ := NewLocationWheather("São Paulo", "São Paulo", "Brazil", "2023-03-25 12:00")
		currentWeather, _ := NewCurrentWeather(1679808000, "2023-03-25 12:00", 25.0, 77.0)
		weather, err := NewWeather(*locationWeather, *currentWeather)
		assert.Nil(t, err)
		assert.NotNil(t, weather)
		assert.Equal(t, *locationWeather, weather.Location)
		assert.Equal(t, *currentWeather, weather.Current)
	})
}

func TestWeather_Validate(t *testing.T) {
	t.Run("should return nil if weather is valid", func(t *testing.T) {
		locationWeather, _ := NewLocationWheather("São Paulo", "São Paulo", "Brazil", "2023-03-25 12:00")
		currentWeather, _ := NewCurrentWeather(1679808000, "2023-03-25 12:00", 25.0, 77.0)
		weather, err := NewWeather(*locationWeather, *currentWeather)
		assert.Nil(t, err)
		err = weather.validate()
		assert.Nil(t, err)
	})
}

package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vinicius-gregorio/fc_cloud_run/internal/failures"
)

func TestNewLocationWheather(t *testing.T) {
	t.Run("should return a new LocationWeather", func(t *testing.T) {
		locationWeather, err := NewLocationWheather("São Paulo", "São Paulo", "Brazil", "2023-03-25 12:00")
		assert.Nil(t, err)
		assert.NotNil(t, locationWeather)
		assert.Equal(t, "São Paulo", locationWeather.Name)
		assert.Equal(t, "São Paulo", locationWeather.Region)
		assert.Equal(t, "Brazil", locationWeather.Country)
		assert.Equal(t, "2023-03-25 12:00", locationWeather.LocalTime)
	})

	t.Run("should return an error if name is empty", func(t *testing.T) {
		locationWeather, err := NewLocationWheather("", "São Paulo", "Brazil", "2023-03-25 12:00")
		assert.NotNil(t, err)
		assert.Equal(t, failures.ErrNameIsRequired, err)
		assert.Nil(t, locationWeather)
	})

	t.Run("should return an error if region is empty", func(t *testing.T) {
		locationWeather, err := NewLocationWheather("São Paulo", "", "Brazil", "2023-03-25 12:00")
		assert.NotNil(t, err)
		assert.Equal(t, failures.ErrRegionIsRequired, err)
		assert.Nil(t, locationWeather)
	})

	t.Run("should return an error if country is empty", func(t *testing.T) {
		locationWeather, err := NewLocationWheather("São Paulo", "São Paulo", "", "2023-03-25 12:00")
		assert.NotNil(t, err)
		assert.Equal(t, failures.ErrCountryIsRequired, err)
		assert.Nil(t, locationWeather)
	})

	t.Run("should return an error if localtime is empty", func(t *testing.T) {
		locationWeather, err := NewLocationWheather("São Paulo", "São Paulo", "Brazil", "")
		assert.NotNil(t, err)
		assert.Equal(t, failures.ErrLocalTimeIsRequired, err)
		assert.Nil(t, locationWeather)
	})
}

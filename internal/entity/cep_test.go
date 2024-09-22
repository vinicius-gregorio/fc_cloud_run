package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vinicius-gregorio/fc_cloud_run/internal/failures"
)

func TestNewLocation(t *testing.T) {
	t.Run("should return a new Location", func(t *testing.T) {
		location, err := NewLocation("01001000", "Praça da Sé", "lado ímpar", "10", "Sé", "São Paulo", "SP", "SP", "Centro", "3550308", "1004", "11", "7107")
		assert.Nil(t, err)
		assert.NotNil(t, location)
		assert.Equal(t, "01001000", location.Cep)
		assert.Equal(t, "Praça da Sé", location.Logradouro)
		assert.Equal(t, "lado ímpar", location.Complemento)
		assert.Equal(t, "10", location.Unidade)
		assert.Equal(t, "Sé", location.Bairro)
		assert.Equal(t, "São Paulo", location.Localidade)
		assert.Equal(t, "SP", location.UF)
		assert.Equal(t, "SP", location.Estado)
		assert.Equal(t, "Centro", location.Regiao)
		assert.Equal(t, "3550308", location.IBGE)
		assert.Equal(t, "1004", location.GIA)
		assert.Equal(t, "11", location.DDD)
		assert.Equal(t, "7107", location.SIAFI)
	})

	t.Run("should return an error if cep is empty", func(t *testing.T) {
		location, err := NewLocation("", "Praça da Sé", "lado ímpar", "10", "Sé", "São Paulo", "SP", "SP", "Centro", "3550308", "1004", "11", "7107")
		assert.NotNil(t, err)
		assert.Equal(t, failures.ErrCepIsRequired, err)
		assert.Nil(t, location)
	})

	t.Run("should return an error if cep is invalid - digits", func(t *testing.T) {
		location, err := NewLocation("01001A00", "Praça da Sé", "lado ímpar", "10", "Sé", "São Paulo", "SP", "SP", "Centro", "3550308", "1004", "11", "7107")
		assert.NotNil(t, err)
		assert.Equal(t, failures.ErrCepInvalid_Digits, err)
		assert.Nil(t, location)
	})

	t.Run("should return an error if cep is invalid - length", func(t *testing.T) {
		location, err := NewLocation("0100100", "Praça da Sé", "lado ímpar", "10", "Sé", "São Paulo", "SP", "SP", "Centro", "3550308", "1004", "11", "7107")
		assert.NotNil(t, err)
		assert.Equal(t, failures.ErrCepInvalid_Length, err)
		assert.Nil(t, location)
	})

	t.Run("should return an error if state is empty", func(t *testing.T) {
		location, err := NewLocation("01001000", "Praça da Sé", "lado ímpar", "10", "Sé", "São Paulo", "SP", "", "Centro", "3550308", "1004", "11", "7107")
		assert.NotNil(t, err)
		assert.Equal(t, failures.ErrStateIsRequired, err)
		assert.Nil(t, location)
	})

	t.Run("should return an error if locality is empty", func(t *testing.T) {
		location, err := NewLocation("01001000", "Praça da Sé", "lado ímpar", "10", "Sé", "", "SP", "SP", "Centro", "3550308", "1004", "11", "7107")
		assert.NotNil(t, err)
		assert.Equal(t, failures.ErrLocalityIsRequired, err)
		assert.Nil(t, location)
	})
}

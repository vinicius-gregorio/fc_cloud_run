package entities

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vinicius-gregorio/fc_cloud_run/internal/failures"
)

func TestNewCep(t *testing.T) {
	t.Run("should return a new Cep", func(t *testing.T) {
		cep, err := NewCep("01001000")
		assert.Nil(t, err)
		assert.NotNil(t, cep)
		assert.Equal(t, "01001000", cep.Cep)
	})

	t.Run("should return an error if cep is empty", func(t *testing.T) {
		cep, err := NewCep("")
		assert.NotNil(t, err)
		assert.Equal(t, failures.ErrCepIsRequired, err)
		assert.Nil(t, cep)
	})
}

func TestCep_Validate(t *testing.T) {
	t.Run("should return nil if cep is valid", func(t *testing.T) {
		cep := &Cep{
			Cep: "01001000",
		}
		err := cep.validate()
		assert.Nil(t, err)
	})

	t.Run("should return an error if cep is empty", func(t *testing.T) {
		cep := &Cep{}
		err := cep.validate()
		assert.NotNil(t, err)
		assert.Equal(t, failures.ErrCepIsRequired, err)
	})
}

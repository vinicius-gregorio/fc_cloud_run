package entities

import "github.com/vinicius-gregorio/fc_cloud_run/internal/failures"

// Cep is a struct that represents a CEP entity
type Cep struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Unidade     string `json:"unidade"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	UF          string `json:"uf"`
	Estado      string `json:"estado"`
	Regiao      string `json:"regiao"`
	IBGE        string `json:"ibge"`
	GIA         string `json:"gia"`
	DDD         string `json:"ddd"`
	SIAFI       string `json:"siafi"`
}

func NewCep(cep string) (*Cep, error) {
	c := &Cep{
		Cep: cep,
	}

	if err := c.validate(); err != nil {
		return nil, err
	}

	return c, nil
}

func (c *Cep) validate() error {
	if c.Cep == "" {
		return failures.ErrCepIsRequired
	}

	return nil
}

package entity

import (
	"github.com/vinicius-gregorio/fc_cloud_run/internal/failures"
)

// Location is a struct that represents a CEP entity
type Location struct {
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

func NewLocation(cep, logradouro, complemento, unidade, bairro, localidade, uf, estado, regiao, ibge, gia, ddd, siafi string) (*Location, error) {
	l := &Location{
		Cep:         cep,
		Logradouro:  logradouro,
		Complemento: complemento,
		Unidade:     unidade,
		Bairro:      bairro,
		Localidade:  localidade,
		UF:          uf,
		Estado:      estado,
		Regiao:      regiao,
		IBGE:        ibge,
		GIA:         gia,
		DDD:         ddd,
		SIAFI:       siafi,
	}

	if err := l.validate(); err != nil {
		return nil, err
	}

	return l, nil
}

func NewLocationByCEP(cep string) (*Location, error) {
	l := &Location{
		Cep: cep,
	}

	if err := l.validateCEPDigits(); err != nil {
		return nil, err
	}

	return l, nil
}

func (l *Location) validate() error {
	if l.Cep == "" {
		return failures.ErrCepIsRequired
	}
	if err := l.validateCEPDigits(); err != nil {
		return err
	}
	if err := l.validateFields(); err != nil {
		return err
	}

	return nil
}

func (l *Location) validateCEPDigits() error {
	if len(l.Cep) != 8 {
		return failures.ErrCepInvalid_Length
	}
	//if the cep contains anything besides numbers, it is invalid
	for _, digit := range l.Cep {
		if digit < '0' || digit > '9' {
			return failures.ErrCepInvalid_Digits
		}
	}

	return nil
}

func (l *Location) validateFields() error {
	if l.Estado == "" {
		return failures.ErrStateIsRequired
	}

	if l.Localidade == "" {
		return failures.ErrLocalityIsRequired
	}

	return nil
}

package services

import (
	"github.com/matheusportoo/busca-cep/services/brasilapi"
	"github.com/matheusportoo/busca-cep/services/correios"
	"github.com/matheusportoo/busca-cep/services/viacep"
	"github.com/matheusportoo/busca-cep/services/wsapicep"
)

type Services struct {
	BrasilApi *brasilapi.BrasilApi
	Correios *correios.Correios
	ViaCep *viacep.ViaCep
	WSApiCep *wsapicep.WSApiCep
}

func New(cep string) *Services {
	return &Services{
		BrasilApi: brasilapi.New(cep),
		Correios: correios.New(cep),
		ViaCep: viacep.New(cep),
		WSApiCep: wsapicep.New(cep),
	}
}
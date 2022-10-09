package core

import "github.com/matheusportoo/busca-cep/core/cep"

type Core struct {
	Cep *cep.Cep
}

func New() *Core {
	return &Core{
		Cep: &cep.Cep{},
	}
}
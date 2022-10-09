package main

import (
	"github.com/matheusportoo/busca-cep/core"
	"github.com/matheusportoo/busca-cep/services"
)

func main() {
	cep := "14406-346"

	services := services.New(cep)
	core := core.New()

	core.Cep.Get(services, cep)
}


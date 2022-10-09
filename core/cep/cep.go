package cep

import (
	"context"
	"fmt"
	"time"

	"github.com/matheusportoo/busca-cep/services"
)

type Cep struct {}

func (s Cep) Get(services *services.Services, cep string) {
	ctx, cancel := context.WithCancel(context.Background())
	t := time.Now()
	fmt.Printf("fetching cep %s \n", cep)

	go getBy(services, "viacep", cancel)
	go getBy(services, "wsapicep", cancel)
	go getBy(services, "brasilapi", cancel)
	go getBy(services, "correios", cancel)

	wait(ctx, t)

	fmt.Println("finished :D")
}

func wait(ctx context.Context, t time.Time) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println(time.Until(t))
			return
		}
	}
}

func getBy(services *services.Services, serviceName string,  cancel context.CancelFunc) {
	defer cancel()

	if serviceName == "viacep" {
		viacepResponse := services.ViaCep.GetCep()
		services.ViaCep.Print(viacepResponse)
		return
	}

	if serviceName == "wsapicep" {
		wsapicepResponse := services.WSApiCep.GetCep()
		services.WSApiCep.Print(wsapicepResponse)
		return
	}

	if serviceName == "brasilapi" {
		brasilapiResponse := services.BrasilApi.GetCep()
		services.BrasilApi.Print(brasilapiResponse)
		return
	}

	if serviceName == "correios" {
		correiosResponse := services.Correios.GetCep()
		services.Correios.Print(correiosResponse)
		return
	}
}

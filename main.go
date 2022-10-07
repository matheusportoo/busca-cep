package main

import (
	"context"
	"fmt"
	"time"

	"github.com/matheusportoo/busca-cep/services"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	cep := "14406-346"
	t := time.Now()

	services := services.New(cep)

	fmt.Printf("fetching cep %s \n", cep)

	go getCepBy(services, "viacep", cep, cancel)
	go getCepBy(services, "wsapicep", cep, cancel)
	go getCepBy(services, "brasilapi", cep, cancel)
	go getCepBy(services, "correios", cep, cancel)

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

func getCepBy(services *services.Services, serviceName string, cep string, cancel context.CancelFunc) {
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

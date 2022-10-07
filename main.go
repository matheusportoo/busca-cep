package main

import (
	"context"
	"fmt"
	"time"

	"github.com/matheusportoo/busca-cep/services/brasilapi"
	"github.com/matheusportoo/busca-cep/services/correios"
	"github.com/matheusportoo/busca-cep/services/viacep"
	"github.com/matheusportoo/busca-cep/services/wsapicep"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	cep := "14406-346"
	t := time.Now()

	fmt.Printf("fetching cep %s \n", cep)

	go getCepBy("viacep", cep, cancel)
	go getCepBy("wsapicep", cep, cancel)
	go getCepBy("brasilapi", cep, cancel)
	go getCepBy("correios", cep, cancel)

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

func getCepBy(service string, cep string, cancel context.CancelFunc) {
	defer cancel()

	if (service == "viacep") {
		viacepResponse := viacep.GetCep(cep)
		viacep.Print(viacepResponse)
	}

	if (service == "wsapicep") {
		wsapicepResponse := wsapicep.GetCep(cep)
		wsapicep.Print(wsapicepResponse)
	}

	if (service == "brasilapi") {
		brasilapiResponse := brasilapi.GetCep(cep)
		brasilapi.Print(brasilapiResponse)
	}

	if (service == "correios") {
		correiosResponse := correios.GetCep(cep)
		correios.Print(correiosResponse)
	}
}
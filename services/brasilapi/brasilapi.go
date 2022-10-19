package brasilapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type BrasilApi struct {
	Cep string
}

func New(cep string) *BrasilApi {
	return &BrasilApi{cep}
}

func (s *BrasilApi) GetCep() Response {
	var url string = fmt.Sprintf("https://brasilapi.com.br/api/cep/v1/%s", s.Cep)
	resp, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	cepResponse := Response{}
	json.Unmarshal([]byte(body), &cepResponse)

	return cepResponse
}

func (s BrasilApi) Print(r Response) {
	fmt.Println()
	fmt.Println("************************")
	fmt.Printf("Response: %s \n", "Brasil Api")
	fmt.Println("************************")
	fmt.Println()
	fmt.Printf("cep: %s \n", r.Cep)
	fmt.Printf("state: %s \n", r.State)
	fmt.Printf("city: %s \n", r.City)
	fmt.Printf("neighborhood: %s \n", r.Neighborhood)
	fmt.Printf("street: %s \n", r.Street)
	fmt.Printf("service: %s \n", r.Service)
	fmt.Println("-------------------------")
}
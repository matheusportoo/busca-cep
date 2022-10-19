package viacep

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type ViaCep struct {
	Cep string
}

func New(cep string) *ViaCep {
	return &ViaCep{cep}
}

func (s *ViaCep) GetCep() Response {
	var url string = fmt.Sprintf("https://viacep.com.br/ws/%s/json", s.Cep)
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

func (s ViaCep) Print(r Response) {
	fmt.Println("************************")
	fmt.Printf("Response: %s \n", "Viacep")
	fmt.Println("************************")
	fmt.Println()
	fmt.Printf("cep: %s \n", r.Cep)
	fmt.Printf("logradouro: %s \n", r.Logradouro)
	fmt.Printf("complemento: %s \n", r.Complemento)
	fmt.Printf("bairro: %s \n", r.Bairro)
	fmt.Printf("localidade: %s \n", r.Localidade)
	fmt.Printf("uf: %s \n", r.Uf)
	fmt.Printf("ibge: %s \n", r.Ibge)
	fmt.Printf("gia: %s \n", r.Gia)
	fmt.Printf("ddd: %s \n", r.Ddd)
	fmt.Printf("siafi: %s \n", r.Siafi)
	fmt.Println("-------------------------")
}
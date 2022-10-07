package brasilapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetCep(cep string) Response {
	var url string = fmt.Sprintf("https://brasilapi.com.br/api/cep/v1/%s", cep)
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

func Print(r Response) {
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
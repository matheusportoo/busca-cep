package wsapicep

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type WSApiCep struct {
	Cep string
}

func New(cep string) *WSApiCep {
	return &WSApiCep{cep}
}

func (s *WSApiCep) GetCep() Response {
	var url string = fmt.Sprintf("https://ws.apicep.com/busca-cep/api/cep/%s.json", s.Cep)
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

func (s WSApiCep) Print(r Response) {
	fmt.Println()
	fmt.Println("************************")
	fmt.Printf("Response: %s \n", "WS api cep")
	fmt.Println("************************")
	fmt.Println()
	fmt.Printf("status: %d \n", r.Status)
	fmt.Printf("ok: %v \n", r.Ok)
	fmt.Printf("code: %s \n", r.Code)
	fmt.Printf("state: %s \n", r.State)
	fmt.Printf("city: %s \n", r.City)
	fmt.Printf("district: %s \n", r.District)
	fmt.Printf("address: %s \n", r.Address)
	fmt.Printf("statusText: %s \n", r.StatusText)
	fmt.Println("-------------------------")
}
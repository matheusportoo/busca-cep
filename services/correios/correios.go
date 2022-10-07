package correios

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type Correios struct {
	Cep string
}

func New(cep string) *Correios {
	return &Correios{cep}
}

func (s *Correios) GetCep() Response {
	formData := url.Values{
		"cep": {s.Cep},
	}

	resp, err := http.PostForm("https://buscacepinter.correios.com.br/app/cep/carrega-cep.php", formData)

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

func (s Correios) Print(r Response) {
	fmt.Println()
	fmt.Println("************************")
	fmt.Printf("Response: %s \n", "Correios")
	fmt.Println("************************")
	fmt.Println()
	fmt.Printf("uf: %s \n", r.Dados[0].Uf)
	fmt.Printf("localidade: %s \n", r.Dados[0].Localidade)
	fmt.Printf("locNoSem: %s \n", r.Dados[0].LocNoSem)
	fmt.Printf("locNu: %s \n", r.Dados[0].LocNu)
	fmt.Printf("localidadeSubordinada: %s \n", r.Dados[0].LocalidadeSubordinada)
	fmt.Printf("logradouroDNEC: %s \n", r.Dados[0].LogradouroDNEC)
	fmt.Printf("logradouroTextoAdicional: %s \n", r.Dados[0].LogradouroTextoAdicional)
	fmt.Printf("logradouroTexto: %s \n", r.Dados[0].LogradouroTexto)
	fmt.Printf("bairro: %s \n", r.Dados[0].Bairro)
	fmt.Printf("baiNu: %s \n", r.Dados[0].BaiNu)
	fmt.Printf("nomeUnidade: %s \n", r.Dados[0].NomeUnidade)
	fmt.Printf("cep: %s \n", r.Dados[0].Cep)
	fmt.Printf("tipoCep: %s \n", r.Dados[0].TipoCep)
	fmt.Printf("numeroLocalidade: %s \n", r.Dados[0].NumeroLocalidade)
	fmt.Printf("situacao: %s \n", r.Dados[0].Situacao)
	fmt.Printf("faixasCaixaPostal: %s \n", r.Dados[0].FaixasCaixaPostal)
	fmt.Printf("faixasCep: %s \n", r.Dados[0].FaixasCep)
	fmt.Println("-------------------------")
}
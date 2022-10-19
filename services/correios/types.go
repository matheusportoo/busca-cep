package correios

type Response struct {
	Erro bool
	Mensagem string
	Total int
	Dados []Cep
}

type Cep struct {
	Uf string `json:"uf"`
	Localidade string `json:"localidade"`
	LocNoSem string `json:"locNoSem"`
	LocNu string `json:"locNu"`
	LocalidadeSubordinada string `json:"localidadeSubordinada"`
	LogradouroDNEC string `json:"logradouroDNEC"`
	LogradouroTextoAdicional string `json:"logradouroTextoAdicional"`
	LogradouroTexto string `json:"logradouroTexto"`
	Bairro string `json:"bairro"`
	BaiNu string `json:"baiNu"`
	NomeUnidade string `json:"nomeUnidade"`
	Cep string `json:"cep"`
	TipoCep string `json:"tipoCep"`
	NumeroLocalidade string `json:"numeroLocalidade"`
	Situacao string `json:"situacao"`
	FaixasCaixaPostal []string `json:"faixasCaixaPostal"`
	FaixasCep []string `json:"faixasCep"`
}
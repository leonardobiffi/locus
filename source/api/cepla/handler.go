package cepla

type ResponseState struct {
	Uf   string `json:"id"`
	Name string `json:"nome"`
}

type ResponseCity struct {
	Name string `json:"nome"`
}

type Response struct {
	Cep      string `json:"cep"`
	Uf       string `json:"uf"`
	City     string `json:"cidade"`
	District string `json:"bairro"`
	Address  string `json:"logradouro"`
}

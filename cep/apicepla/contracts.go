package apicepla

type Response struct {
	Cep      string `json:"cep"`
	Uf       string `json:"uf"`
	City     string `json:"cidade"`
	District string `json:"bairro"`
	Address  string `json:"logradouro"`
}

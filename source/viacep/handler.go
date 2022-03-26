package apiviacep

type Response struct {
	Cep      string `json:"cep"`
	Address  string `json:"logradouro"`
	District string `json:"bairro"`
	City     string `json:"localidade"`
	State    string `json:"uf"`
	DDD      string `json:"ddd"`
}

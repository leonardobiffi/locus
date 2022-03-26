package opencep

type Response struct {
	Cep       string `json:"cep"`
	Uf        string `json:"uf"`
	City      string `json:"localidade"`
	District  string `json:"bairro"`
	Address   string `json:"logradouro"`
	SourceApi string `json:"source"`
}

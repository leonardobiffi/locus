package apivercel

type Response struct {
	Cep      string
	Uf       string
	City     string
	District string
	Address  string
}

type ResponseVercel struct {
	Date string `json:"date"`
	Info *Info  `json:"info"`
}

type Info struct {
	Cep      string  `json:"cep"`
	Address  string  `json:"address"`
	State    string  `json:"state"`
	District string  `json:"district"`
	City     string  `json:"city"`
	Status   *int    `json:"status"`
	Message  *string `json:"message"`
}

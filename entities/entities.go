package entities

type Response struct {
	Cep          string
	Uf           string
	City         string
	District     string
	Address      string
	SourceApi    string
	ResponseTime string
}

type API interface {
	Get() (Response, error)
}

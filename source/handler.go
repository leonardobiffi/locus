package source

type Response struct {
	Cep       string
	Uf        string
	City      string
	District  string
	Address   string
	SourceApi string
}

func GetSources() []string {
	return []string{
		"cepla",
		"vercel",
		"viacep",
		"opencep",
	}
}

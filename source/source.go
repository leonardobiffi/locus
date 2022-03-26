package source

import (
	"locus/entities"
	"locus/source/api/cepla"
	"locus/source/api/opencep"
	"locus/source/api/vercel"
)

func New(source, cep string) entities.API {
	switch source {
	case cepla.SourceApi:
		return cepla.New(cep)
	case opencep.SourceApi:
		return opencep.New(cep)
	case vercel.SourceApi:
		return vercel.New(cep)

	default:
		return newDefault(cep)
	}
}

type defaults struct {
	cep string
}

func newDefault(cep string) *defaults {
	return &defaults{cep: cep}
}

func (s *defaults) Get() (entities.Response, error) {
	message := make(chan entities.Response)

	go execute(cepla.New(s.cep), message)
	go execute(opencep.New(s.cep), message)
	go execute(vercel.New(s.cep), message)

	response := <-message
	if response.Cep != "" {
		return response, nil
	}

	return entities.Response{}, nil
}

func execute(api entities.API, message chan entities.Response) {
	resp, err := api.Get()
	if err != nil {
		return
	}

	message <- resp
}

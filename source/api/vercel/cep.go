package vercel

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"locus/entities"
	"locus/utils"
	"net/http"
	"time"
)

const SourceApi = "vercel"

type service struct {
	cep       string
	response  Response
	sourceApi string
}

func New(cep string) *service {
	return &service{
		cep:       cep,
		sourceApi: SourceApi,
	}
}

// Get return CEP info using => https://cep-api.vercel.app/api/{cep}
func (s *service) Get() (response entities.Response, err error) {
	url := fmt.Sprintf("https://cep-api.vercel.app/api/%s", s.cep)

	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	jsonErr := json.Unmarshal(body, &s.response)
	if jsonErr != nil {
		return
	}

	if s.response.Info.Status != nil {
		return response, fmt.Errorf("cep %s not found", s.cep)
	}

	return entities.Response{
		Cep:          s.response.Info.Cep,
		Uf:           s.response.Info.State,
		City:         s.response.Info.City,
		District:     s.response.Info.District,
		Address:      s.response.Info.Address,
		SourceApi:    s.sourceApi,
		ResponseTime: utils.FormatResponseTime(start),
	}, nil
}

package opencep

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"locus/entities"
	"locus/utils"
	"net/http"
	"strings"
	"time"
)

const SourceApi = "opencep"

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

// Get return CEP info using => https://opencep.com/v1/{cep}
func (s *service) Get() (response entities.Response, err error) {
	url := fmt.Sprintf("https://opencep.com/v1/%s", s.cep)

	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		return response, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return response, err
	}

	jsonErr := json.Unmarshal(body, &s.response)
	if jsonErr != nil {
		return response, fmt.Errorf("cep %s not found", s.cep)
	}

	return entities.Response{
		Cep:          strings.ReplaceAll(s.response.Cep, "-", ""),
		Uf:           s.response.Uf,
		City:         s.response.City,
		District:     s.response.District,
		Address:      s.response.Address,
		SourceApi:    s.sourceApi,
		ResponseTime: utils.FormatResponseTime(start),
	}, nil
}

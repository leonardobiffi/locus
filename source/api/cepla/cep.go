package cepla

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"locus/config"
	"locus/utils"
	"log"
	"net/http"
	"strings"
	"time"

	"locus/entities"
)

const SourceApi = "cepla"

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

// Get return CEP info using => http://cep.la/{cep}
func (s *service) Get() (response entities.Response, err error) {
	url := fmt.Sprintf("http://cep.la/%s", s.cep)

	client := &http.Client{}
	_, err = client.Get(url)
	if err != nil {
		return
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return
	}

	req.Header.Add("Accept", "application/json")

	start := time.Now()
	resp, err := client.Do(req)
	if err != nil {
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	jsonErr := json.Unmarshal(body, &s.response)
	if jsonErr != nil {
		return response, fmt.Errorf("cep %s not found", s.cep)
	}

	return entities.Response{
		Cep:          s.response.Cep,
		Uf:           s.response.Uf,
		City:         s.response.City,
		District:     s.response.District,
		Address:      s.response.Address,
		SourceApi:    s.sourceApi,
		ResponseTime: utils.FormatResponseTime(start),
	}, nil
}

func ListState() (response []ResponseState) {
	url := "http://cep.la"

	client := &http.Client{}
	_, err := client.Get(url)
	if err != nil {
		log.Fatalln(err)
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatalln(err)
	}
	req.Header.Add("Accept", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	jsonErr := json.Unmarshal(body, &response)
	if jsonErr != nil {
		fmt.Println(jsonErr)
		return
	}

	return response
}

func ListCity(state, page string) (response []ResponseCity) {
	state = strings.ToUpper(state)
	if page == "all" {
		fmt.Println("Buscando lista de cidades ...")
	}

	i := 1
	var responseAll []ResponseCity
	for {
		url := fmt.Sprintf("http://cep.la/%s/%s", state, fmt.Sprint(i))

		client := &http.Client{}
		_, err := client.Get(url)
		if err != nil {
			log.Fatalln(err)
		}

		req, err := http.NewRequest(http.MethodGet, url, nil)
		if err != nil {
			log.Fatalln(err)
		}
		req.Header.Add("Accept", "application/json")

		resp, err := client.Do(req)
		if err != nil {
			log.Fatalln(err)
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}

		jsonErr := json.Unmarshal(body, &response)
		if jsonErr != nil {
			fmt.Println(jsonErr)
			return
		}

		responseAll = append(responseAll, response...)

		if len(response) <= 0 || page != "all" {
			return responseAll
		}

		i++
	}

}

func ListCep(uf, city, district string) (responseApi []Response) {
	var response []Response
	url := fmt.Sprintf("http://cep.la/%s/%s/%s", uf, city, district)

	client := &http.Client{}
	_, err := client.Get(url)
	if err != nil {
		log.Fatalln(err)
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatalln(err)
	}

	req.Header.Add("Accept", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	jsonErr := json.Unmarshal(body, &response)
	if jsonErr != nil {
		fmt.Println(jsonErr)
		fmt.Println(config.ColorRed, fmt.Sprintf("UF %s not found!", uf))
		return
	}

	return response
}

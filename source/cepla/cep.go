package apicepla

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"locus/config"
	"locus/source"
	"log"
	"net/http"
	"strings"
)

const (
	SourceApi = "cepla"
)

// GetCep return CEP info using => http://cep.la/api
func GetCep(findCep string, messages chan source.Response) {
	response := Response{}

	url := fmt.Sprintf("http://cep.la/%s", findCep)

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
		fmt.Println(config.ColorRed, fmt.Sprintf("CEP %s not found!", findCep))
		messages <- source.Response{}
	}

	messages <- source.Response{
		Cep:       response.Cep,
		Uf:        response.Uf,
		City:      response.City,
		District:  response.District,
		Address:   response.Address,
		SourceApi: SourceApi,
	}
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

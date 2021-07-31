package apicepla

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"locus-cli/cep"
	"locus-cli/config"
	"log"
	"net/http"
	"strings"
)

// GetCep return CEP info using => http://cep.la/api
func GetCep(findCep string, messages chan cep.Response) {
	response := Response{}

	url := fmt.Sprintf("http://cep.la/%s", findCep)

	client := &http.Client{}
	resp, err := client.Get(url)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	req.Header.Add("Accept", "application/json")

	resp, err = client.Do(req)
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
		messages <- cep.Response{}
	}

	messages <- cep.Response{
		Cep:       response.Cep,
		Uf:        response.Uf,
		City:      response.City,
		District:  response.District,
		Address:   response.Address,
		ApiSource: "cepla",
	}
}

func ListState() (response []ResponseState) {
	url := fmt.Sprintf("http://cep.la")

	client := &http.Client{}
	resp, err := client.Get(url)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	req.Header.Add("Accept", "application/json")

	resp, err = client.Do(req)
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
		resp, err := client.Get(url)

		req, err := http.NewRequest(http.MethodGet, url, nil)
		req.Header.Add("Accept", "application/json")

		resp, err = client.Do(req)
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
	resp, err := client.Get(url)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	req.Header.Add("Accept", "application/json")

	resp, err = client.Do(req)
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

package apiviacep

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
	SourceApi = "viacep"
)

// GetCep return CEP info using => https://viacep.com.br/ws
func GetCep(findCep string, messages chan source.Response) {
	response := Response{}
	url := fmt.Sprintf("https://viacep.com.br/ws/%s/json/", findCep)

	resp, err := http.Get(url)
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
		log.Fatal(jsonErr)
	}

	messages <- source.Response{
		Cep:       strings.ReplaceAll(response.Cep, "-", ""),
		Uf:        response.State,
		City:      response.City,
		District:  response.District,
		Address:   response.Address,
		SourceApi: SourceApi,
	}
}

func ListCep(uf, city, address string) (responseApi []Response) {
	var response []Response
	address = strings.ReplaceAll(address, " ", "+")

	url := fmt.Sprintf("https://viacep.com.br/ws/%s/%s/%s/json", uf, city, address)

	resp, err := http.Get(url)
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

	for _, r := range response {
		responseApi = append(responseApi,
			Response{
				Cep:      strings.ReplaceAll(r.Cep, "-", ""),
				Address:  r.Address,
				District: r.District,
				City:     r.City,
				State:    r.State,
				DDD:      r.DDD,
			},
		)
	}

	return responseApi
}

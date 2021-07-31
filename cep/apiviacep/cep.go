package apiviacep

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

// GetCep return CEP info using => https://viacep.com.br/ws
func GetCep(findCep string, messages chan cep.Response) {
	response := Response{}
	url := fmt.Sprintf("https://viacep.com.br/ws/%s/json", findCep)

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

	messages <- cep.Response{
		Cep:       strings.ReplaceAll(response.Cep, "-", ""),
		Uf:        response.State,
		City:      response.City,
		District:  response.District,
		Address:   response.Address,
		ApiSource: "viacep",
	}
}

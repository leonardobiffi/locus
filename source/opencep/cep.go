package opencep

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
	SourceApi = "opencep"
)

// GetCep return CEP info using => https://opencep.com/v1/cep
func GetCep(findCep string, messages chan source.Response) {
	response := Response{}
	url := fmt.Sprintf("https://opencep.com/v1/%s", findCep)

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
		Uf:        response.Uf,
		City:      response.City,
		District:  response.District,
		Address:   response.Address,
		SourceApi: SourceApi,
	}
}

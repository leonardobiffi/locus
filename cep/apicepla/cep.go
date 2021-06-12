package apicepla

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"locus-cli/config"
	"log"
	"net/http"
)

// GetCep return CEP info using => http://cep.la/api
func GetCep(cep string) (response Response) {
	url := fmt.Sprintf("http://cep.la/%s", cep)

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
		fmt.Println(config.ColorRed, fmt.Sprintf("CEP %s not found!", cep))
		return Response{}
	}

	return response
}

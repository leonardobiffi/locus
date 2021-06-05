/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"log"
	"net/http"

	"github.com/jedib0t/go-pretty/table"
	"github.com/jedib0t/go-pretty/text"
	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:     "get",
	Version: rootCmd.Version,
	Short:   "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: getCep,
}

var (
	CepFlag    string
	colorRed   = "\033[31m"
	colorGreen = "\033[32m"
)

type Response struct {
	Date string `json:"date"`
	Info *Info  `json:"info"`
}

type Info struct {
	Cep      string  `json:"cep"`
	Address  string  `json:"address"`
	State    string  `json:"state"`
	District string  `json:"district"`
	City     string  `json:"city"`
	Status   *int    `json:"status"`
	Message  *string `json:"message"`
}

func init() {
	rootCmd.AddCommand(getCmd)

	getCmd.Flags().StringVarP(&CepFlag, "cep", "c", "", "Set CEP [required]")
	getCmd.MarkFlagRequired("cep")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func getCep(cmd *cobra.Command, args []string) {

	url := fmt.Sprintf("https://cep-api.vercel.app/api/%s", CepFlag)

	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	response := Response{}
	jsonErr := json.Unmarshal(body, &response)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	if response.Info.Status != nil {
		fmt.Println(colorRed, *response.Info.Message)
		return
	}

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.SetCaption(fmt.Sprintf("Informações do CEP: %s", CepFlag))
	t.AppendHeader(table.Row{"Cidade", "CEP", "Endereço", "Estado", "Bairro"})

	t.AppendRow(table.Row{
		response.Info.City,
		response.Info.Cep,
		response.Info.Address,
		response.Info.State,
		response.Info.District,
	})

	t.SetStyle(table.StyleLight)
	t.Style().Color = table.ColorOptions{
		IndexColumn:  nil,
		Footer:       text.Colors{text.FgHiBlue, text.FgHiBlue},
		Header:       text.Colors{text.FgHiBlue, text.FgHiBlue},
		Row:          text.Colors{text.FgHiBlue, text.FgHiBlue},
		RowAlternate: text.Colors{text.FgHiBlue, text.FgHiBlue},
	}

	t.Render()
}

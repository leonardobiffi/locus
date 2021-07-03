package cmd

import (
	"locus-cli/cep"
	"locus-cli/cep/apicepla"
	"locus-cli/cep/apivercel"
	"locus-cli/utils"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get Info about the CEP",
	Long:  `Get Info about the CEP.`,
	RunE:  getCep,
}

var (
	CepFlag     string
	PrintPretty bool
	PrintJson   bool
)

func init() {
	rootCmd.AddCommand(getCmd)

	getCmd.Flags().StringVarP(&CepFlag, "cep", "c", "", "Set CEP [required]")
	getCmd.MarkFlagRequired("cep")

	getCmd.Flags().BoolVarP(&PrintPretty, "pretty", "p", false, "Print Pretty Table output")
	getCmd.Flags().BoolVarP(&PrintJson, "json", "j", false, "Print Pretty JSON output")

}

func getCep(cmd *cobra.Command, args []string) error {

	messages := make(chan cep.Response)

	go apicepla.GetCep(CepFlag, messages)
	go apivercel.GetCep(CepFlag, messages)

	response := <-messages

	if response.Cep != "" {
		header := table.Row{"Cidade", "CEP", "Endereço", "Estado", "Bairro"}
		row := table.Row{
			response.City,
			response.Cep,
			response.Address,
			response.Uf,
			response.District,
		}

		if PrintPretty {
			utils.PrintTablePretty(header, row)
			return nil
		}

		if PrintJson {
			utils.PrintJson(response)
			return nil
		}

		utils.PrintTable(header, row)
	}

	return nil
}

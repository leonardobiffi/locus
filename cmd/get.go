package cmd

import (
	"locus/source"
	cepla "locus/source/cepla"
	"locus/source/opencep"
	vercel "locus/source/vercel"
	viacep "locus/source/viacep"
	"locus/utils"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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
	messages := make(chan source.Response)

	sourceApi := viper.GetString("source")
	getFromSource(CepFlag, sourceApi, messages)

	response := <-messages

	if response.Cep != "" {
		header := table.Row{"Cidade", "CEP", "Endereço", "Estado", "Bairro", "API"}
		row := table.Row{
			response.City,
			response.Cep,
			response.Address,
			response.Uf,
			response.District,
			response.SourceApi,
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

func getFromSource(cep, sourceApi string, messages chan source.Response) {
	switch sourceApi {
	case cepla.SourceApi:
		go cepla.GetCep(CepFlag, messages)
	case vercel.SourceApi:
		go vercel.GetCep(CepFlag, messages)
	case viacep.SourceApi:
		go viacep.GetCep(CepFlag, messages)
	case opencep.SourceApi:
		go opencep.GetCep(CepFlag, messages)
	default:
		go cepla.GetCep(cep, messages)
		go vercel.GetCep(cep, messages)
		go viacep.GetCep(cep, messages)
		go opencep.GetCep(cep, messages)
	}
}

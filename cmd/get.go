package cmd

import (
	"fmt"
	"locus-cli/cep/apicepla"
	"os"

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
	CepFlag string
)

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

	// response := apivercel.GetCep(CepFlag)
	response := apicepla.GetCep(CepFlag)

	if response.Cep != "" {
		t := table.NewWriter()
		t.SetOutputMirror(os.Stdout)
		t.SetCaption(fmt.Sprintf("Informações do CEP: %s", CepFlag))
		t.AppendHeader(table.Row{"Cidade", "CEP", "Endereço", "Estado", "Bairro"})

		t.AppendRow(table.Row{
			response.City,
			response.Cep,
			response.Address,
			response.Uf,
			response.District,
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
}

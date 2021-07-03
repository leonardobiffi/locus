package cmd

import (
	"locus-cli/cep/apicepla"
	"locus-cli/utils"

	"github.com/spf13/cobra"
)

var (
	UfFlag       string
	CityFlag     string
	DistrictFlag string
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List CEPs by location",
	Long:  `List CEPs by location.`,
	RunE:  listCep,
}

func init() {
	rootCmd.AddCommand(listCmd)

	listCmd.Flags().StringVarP(&UfFlag, "uf", "u", "", "Set UF [required]")
	listCmd.Flags().StringVarP(&CityFlag, "city", "i", "", "Set City [required]")
	listCmd.Flags().StringVarP(&DistrictFlag, "district", "d", "", "Set District [required]")

	listCmd.MarkFlagRequired("uf")
	listCmd.MarkFlagRequired("city")
	listCmd.MarkFlagRequired("district")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func listCep(cmd *cobra.Command, args []string) error {

	response := apicepla.ListCep(UfFlag, CityFlag, DistrictFlag)
	utils.PrintJson(response)

	return nil

	// t := table.NewWriter()
	// t.SetOutputMirror(os.Stdout)
	// t.AppendHeader(table.Row{"CEP", "Endereço"})

	// for _, r := range response {
	// 	t.AppendRow(table.Row{
	// 		r.Cep,
	// 		r.Address,
	// 	})
	// }

	// t.SetStyle(table.StyleLight)
	// t.Style().Color = table.ColorOptions{
	// 	IndexColumn:  nil,
	// 	Footer:       text.Colors{text.FgHiBlue, text.FgHiBlue},
	// 	Header:       text.Colors{text.FgHiBlue, text.FgHiBlue},
	// 	Row:          text.Colors{text.FgHiBlue, text.FgHiBlue},
	// 	RowAlternate: text.Colors{text.FgHiBlue, text.FgHiBlue},
	// }

	// t.Render()
}

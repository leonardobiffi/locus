package cmd

// import (
// 	apicepla "locus/source/cepla"
// 	apiviacep "locus/source/viacep"
// 	"locus/utils"

// 	"github.com/jedib0t/go-pretty/v6/table"
// 	"github.com/spf13/cobra"
// )

// var (
// 	UfFlag       string
// 	PageFlag     string
// 	CityFlag     string
// 	DistrictFlag string
// 	AddressFlag  string
// )

// // listCmd represents the list command
// var listCmd = &cobra.Command{
// 	Use:   "list",
// 	Short: "List CEPs by location",
// 	Long:  `List CEPs by location.`,
// 	RunE:  listCep,
// }

// func init() {
// 	rootCmd.AddCommand(listCmd)

// 	listCmd.Flags().StringVarP(&UfFlag, "uf", "u", "", "Set UF")
// 	listCmd.Flags().StringVarP(&CityFlag, "city", "i", "", "Set City")
// 	listCmd.Flags().StringVarP(&DistrictFlag, "district", "d", "", "Set District")
// 	listCmd.Flags().StringVarP(&AddressFlag, "address", "a", "", "Set Address")

// 	listCmd.Flags().StringVarP(&PageFlag, "page", "p", "", "List Paginate")
// }

// func listCep(cmd *cobra.Command, args []string) error {

// 	// response := apicepla.ListCep(UfFlag, CityFlag, DistrictFlag)
// 	// response := apiviacep.ListCep(UfFlag, CityFlag, AddressFlag)

// 	if UfFlag != "" {
// 		response := apicepla.ListCity(UfFlag, PageFlag)

// 		header := table.Row{"Cidade"}
// 		var rows []table.Row
// 		for _, r := range response {

// 			rows = append(rows, table.Row{
// 				r.Name,
// 			})
// 		}

// 		utils.PrintTableRows(header, rows)
// 		return nil
// 	}

// 	if AddressFlag != "" {
// 		response := apiviacep.ListCep(UfFlag, CityFlag, AddressFlag)

// 		header := table.Row{"Cep", "Rua", "Bairro", "Cidade"}
// 		var rows []table.Row
// 		for _, r := range response {

// 			rows = append(rows, table.Row{
// 				r.Cep,
// 				r.Address,
// 				r.District,
// 				r.City,
// 			})
// 		}

// 		utils.PrintTableRows(header, rows)
// 		return nil
// 	}

// 	response := apicepla.ListState()

// 	header := table.Row{"UF", "Estado"}
// 	var rows []table.Row
// 	for _, r := range response {

// 		rows = append(rows, table.Row{
// 			r.Uf,
// 			r.Name,
// 		})
// 	}

// 	utils.PrintTableRows(header, rows)

// 	return nil
// }

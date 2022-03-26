package cmd

import (
	"locus/source"
	"locus/utils"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
)

// sourceCmd represents the get command
var sourceCmd = &cobra.Command{
	Use:   "source",
	Short: "Info about source APIs",
	Long:  `Info about source APIs.`,
	RunE:  runSourceApi,
}

var (
	SourceList bool
)

func init() {
	rootCmd.AddCommand(sourceCmd)

	sourceCmd.Flags().BoolVarP(&SourceList, "list", "l", false, "List source APIs")
}

func runSourceApi(cmd *cobra.Command, args []string) error {
	header := table.Row{"Source"}
	var rows []table.Row
	for _, source := range source.GetSources() {
		rows = append(rows, table.Row{source})
	}

	utils.PrintTableRows(header, rows)
	return nil
}

package utils

import (
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
)

func PrintTable(header table.Row, row table.Row) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)

	t.AppendHeader(header)
	t.AppendRow(row)

	t.Render()
}

func PrintTablePretty(header table.Row, row table.Row) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)

	t.AppendHeader(header)
	t.AppendRow(row)

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

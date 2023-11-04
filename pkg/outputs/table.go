package outputs

import (
	"os"

	"github.com/olekukonko/tablewriter"
)

// TableOutput renders a table in terminal
func TableOutput(headers []string, outputData [][]string) {
	table := tablewriter.NewWriter(os.Stdout)

	// enables autoMerge only for nodepool infos
	if headers[0] == "NODEPOOL" {
		table.SetAutoMergeCells(true)
	}

	// default configs for the table-writer
	table.SetRowLine(false)
	table.SetBorder(false)
	table.SetAutoWrapText(false)
	table.SetAutoFormatHeaders(true)
	table.SetHeaderLine(false)
	table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetAutoWrapText(false)
	table.SetCenterSeparator("")
	table.SetColumnSeparator("")
	table.SetRowSeparator("")
	table.SetTablePadding("\t")

	// set headers and add the outputData
	table.SetHeader(headers)
	table.AppendBulk(outputData)

	table.Render()
}

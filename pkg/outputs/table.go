package outputs

import (
	"os"

	"github.com/olekukonko/tablewriter"
)

// TableOutput renders a table in terminal
func TableOutput(headers []string, outputData [][]string) {
	table := tablewriter.NewWriter(os.Stdout)

	// misc configs for the table-writer
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
	table.SetNoWhiteSpace(true)

	// set headers and add the outputData
	table.SetHeader(headers)
	table.AppendBulk(outputData)

	table.Render()
}

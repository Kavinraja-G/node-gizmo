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
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetAutoWrapText(true)
	table.SetColumnSeparator("")
	table.SetHeaderLine(false)

	// set headers and add the outputData
	table.SetHeader(headers)
	table.AppendBulk(outputData)

	table.Render()
}

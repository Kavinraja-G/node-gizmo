package outputs

import (
	"github.com/olekukonko/tablewriter"
	"os"
)

// TableOutput renders a table in terminal
func TableOutput(headers []string, outputData [][]string) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetRowLine(false)
	table.SetBorder(false)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetAutoWrapText(true)
	table.SetColumnSeparator("")
	table.SetHeaderLine(false)

	table.SetHeader(headers)
	table.AppendBulk(outputData)

	table.Render()
}

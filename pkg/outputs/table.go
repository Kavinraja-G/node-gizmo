package outputs

import (
	"os"
	"strings"

	"github.com/Kavinraja-G/node-gizmo/pkg"
	"github.com/olekukonko/tablewriter"
)

func OutputGenericNodeInfo(genericNodeInfos []pkg.GenericNodeInfo, outputOpts pkg.OutputOptsForGenericNodeInfo) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetRowLine(false)
	table.SetBorder(false)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetAutoWrapText(true)
	table.SetColumnSeparator("")
	table.SetHeaderLine(false)

	// default headers
	headers := []string{"NAME", "VERSION", "IMAGE", "OS", "ARCHITECTURE", "STATUS"}
	if outputOpts.ShowTaints {
		headers = append(headers, "TAINTS")
	}
	table.SetHeader(headers)

	for _, nodeInfo := range genericNodeInfos {
		lineItems := []string{
			nodeInfo.NodeName,
			nodeInfo.K8sVersion,
			nodeInfo.Image,
			nodeInfo.Os,
			nodeInfo.OsArch,
			nodeInfo.NodeStatus,
		}
		if outputOpts.ShowTaints {
			lineItems = append(lineItems, strings.Join(nodeInfo.Taints, "\n"))
		}
		table.Append(lineItems)
	}

	table.Render()
}

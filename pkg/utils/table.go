package utils

import (
	"fmt"
	"os"
	"strings"
	"text/tabwriter"

	"github.com/Kavinraja-G/node-gizmo/pkg"
)

func printLine(w *tabwriter.Writer, lineItems []string) {
	fmt.Fprintf(w, strings.Join(lineItems[:], "\t ")+"\n")
}

func OutputGenericNodeInfo(genericNodeInfos []pkg.GenericNodeInfo) {
	w := tabwriter.NewWriter(os.Stdout, 10, 0, 2, ' ', 0)

	printLine(w, []string{"NAME", "VERSION", "OS", "ARCHITECTURE", "STATUS"})

	for _, nodeInfo := range genericNodeInfos {
		var lineItems []string
		lineItems = append(lineItems, nodeInfo.NodeName)
		lineItems = append(lineItems, nodeInfo.K8sVersion)
		lineItems = append(lineItems, nodeInfo.Os)
		lineItems = append(lineItems, nodeInfo.OsArch)

		if nodeInfo.NodeStatus {
			lineItems = append(lineItems, "Ready")
		} else {
			lineItems = append(lineItems, "NotReady")
		}

		printLine(w, lineItems)
	}

	if err := w.Flush(); err != nil {
		fmt.Printf("Error displaying the output %v", err)
	}
}

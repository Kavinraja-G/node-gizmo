package nodes

import (
	"context"

	"github.com/Kavinraja-G/node-gizmo/pkg"
	"github.com/Kavinraja-G/node-gizmo/pkg/outputs"
	"github.com/Kavinraja-G/node-gizmo/utils"
	"github.com/spf13/cobra"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// NewCmdNodeCapacityInfo initializes the 'capacity' command
func NewCmdNodeCapacityInfo() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "capacity",
		Short:   "Displays Node capacity related information",
		Aliases: []string{"capacities", "cp"},
		RunE: func(cmd *cobra.Command, args []string) error {
			return showNodeCapacities(cmd, args)
		},
	}

	return cmd
}

// showNodeCapacities driver function for 'node capacity' command
func showNodeCapacities(cmd *cobra.Command, args []string) error {
	var nodeCapacityInfo []pkg.NodeCapacities

	nodes, err := utils.Cfg.Clientset.CoreV1().Nodes().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return err
	}

	for _, node := range nodes.Items {
		nodeCapacity := pkg.NodeCapacities{
			NodeName:         node.Name,
			CPU:              node.Status.Capacity.Cpu().String(),
			Memory:           node.Status.Capacity.Memory().Value(),
			Disk:             node.Status.Capacity.Storage().Value(),
			EphemeralStorage: node.Status.Capacity.StorageEphemeral().Value(),
			PodCapacity:      node.Status.Capacity.Pods().String(),
		}

		nodeCapacityInfo = append(nodeCapacityInfo, nodeCapacity)
	}

	outputHeaders, outputData := generateNodeCapacityOutputData(nodeCapacityInfo)
	outputs.SortOutputBasedOnHeader(outputHeaders, outputData, sortByHeader)
	outputs.TableOutput(outputHeaders, outputData)

	return nil
}

// generateNodeCapacityOutputData generates the NodeCapacity outputs and the required headers for table-writer
func generateNodeCapacityOutputData(nodeCapacityInfo []pkg.NodeCapacities) ([]string, [][]string) {
	var headers = []string{"NAME", "CPU", "MEMORY", "DISK", "EPHEMERAL", "POD CAPACITY"}
	var outputData [][]string

	for _, nodeCapacity := range nodeCapacityInfo {
		lineItems := []string{
			nodeCapacity.NodeName,
			nodeCapacity.CPU,
			utils.PrettyByteSize(nodeCapacity.Memory),
			utils.PrettyByteSize(nodeCapacity.Disk),
			utils.PrettyByteSize(nodeCapacity.EphemeralStorage),
			nodeCapacity.PodCapacity,
		}
		outputData = append(outputData, lineItems)
	}

	return headers, outputData
}

package nodes

import (
	"context"
	"fmt"
	"github.com/Kavinraja-G/node-gizmo/utils"
	"strings"

	"github.com/Kavinraja-G/node-gizmo/pkg/outputs"

	"github.com/Kavinraja-G/node-gizmo/pkg"
	"github.com/spf13/cobra"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var (
	showTaints           bool
	showNodeProviderInfo bool
	showNodeTopologyInfo bool
)

// NewCmdNodeInfo initializes the 'node' command
func NewCmdNodeInfo() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "node",
		Short:   "Displays generic node related information in the cluster",
		Aliases: []string{"nodes"},
		RunE: func(cmd *cobra.Command, args []string) error {
			return showNodeInfo(cmd, args)
		},
	}

	// additional local flags
	cmd.Flags().BoolVarP(&showTaints, "show-taints", "t", false, "Shows taints added on a node")
	cmd.Flags().BoolVarP(&showNodeProviderInfo, "show-providers", "p", false, "Shows cloud provider name for a node")
	cmd.Flags().BoolVarP(&showNodeTopologyInfo, "show-topology", "T", false, "Shows node topology info like region & zones for a node")

	// additional sub-commands
	cmd.AddCommand(NewCmdNodeCapacityInfo())
	cmd.AddCommand(NewCmdNodeExec())

	return cmd
}

// showNodeInfo driver function for generic node info command
func showNodeInfo(cmd *cobra.Command, args []string) error {
	var nodeInfos []pkg.GenericNodeInfo
	var outputOpts = pkg.OutputOptsForGenericNodeInfo{
		ShowTaints:           showTaints,
		ShowNodeProviderInfo: showNodeProviderInfo,
		ShowNodeTopologyInfo: showNodeTopologyInfo,
	}

	nodes, err := utils.Cfg.Clientset.CoreV1().Nodes().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return err
	}

	for _, node := range nodes.Items {
		genericNodeInfo := pkg.GenericNodeInfo{
			NodeName:   node.Name,
			K8sVersion: node.Status.NodeInfo.KubeletVersion,
			Image:      node.Status.NodeInfo.OSImage,
			Os:         node.Status.NodeInfo.OperatingSystem,
			OsArch:     node.Status.NodeInfo.Architecture,
			NodeStatus: getNodeStatus(node.Status.Conditions),
		}

		if ok, _ := cmd.Flags().GetBool("show-taints"); ok {
			genericNodeInfo.Taints = getNodeTaints(node.Spec.Taints)
		}
		if ok, _ := cmd.Flags().GetBool("show-providers"); ok {
			genericNodeInfo.NodeProvider = getNodeProviderName(node.Spec.ProviderID)
		}
		if ok, _ := cmd.Flags().GetBool("show-topology"); ok {
			genericNodeInfo.NodeTopologyRegion, genericNodeInfo.NodeTopologyZone = pkg.GetNodeTopologyInfo(node.Labels)
		}

		nodeInfos = append(nodeInfos, genericNodeInfo)
	}

	outputHeaders, outputData := generateNodeInfoOutputData(nodeInfos, outputOpts)
	outputs.TableOutput(outputHeaders, outputData)

	return nil
}

// getNodeProviderName returns the providerName stripped from the providerID in the spec
// providerID = aws://someRandomNodeID => "aws"
// providerID = nil || "" => "others"
func getNodeProviderName(providerID string) string {
	// providerID format <ProviderName>://<ProviderSpecificNodeID>
	if providerID != "" {
		return strings.Split(providerID, ":")[0]
	}
	return "others"
}

// getNodeTaints returns the taints that are added to the node
func getNodeTaints(rawTaints []corev1.Taint) []string {
	var taints []string
	for _, taint := range rawTaints {
		taints = append(taints, fmt.Sprintf("%v=%v:%v", taint.Key, taint.Value, taint.Effect))
	}
	return taints
}

// getNodeStatus returns if the provided node is 'Ready' or 'NotReady'
func getNodeStatus(nodeConditions []corev1.NodeCondition) string {
	for _, nodeCondition := range nodeConditions {
		if nodeCondition.Type == corev1.NodeReady {
			return "Ready"
		}
	}
	return "NotReady"
}

// generateNodeInfoOutputData generates the NodeInfo outputs and the required headers for table-writer
func generateNodeInfoOutputData(genericNodeInfos []pkg.GenericNodeInfo, outputOpts pkg.OutputOptsForGenericNodeInfo) ([]string, [][]string) {
	var headers = []string{"NAME", "VERSION", "IMAGE", "OS", "ARCHITECTURE", "STATUS"}
	var outputData [][]string

	if outputOpts.ShowTaints {
		headers = append(headers, "TAINTS")
	}
	if outputOpts.ShowNodeProviderInfo {
		headers = append(headers, "PROVIDER")
	}
	if outputOpts.ShowNodeTopologyInfo {
		headers = append(headers, "REGION", "ZONE")
	}

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
		if outputOpts.ShowNodeProviderInfo {
			lineItems = append(lineItems, nodeInfo.NodeProvider)
		}
		if outputOpts.ShowNodeTopologyInfo {
			lineItems = append(lineItems, nodeInfo.NodeTopologyRegion, nodeInfo.NodeTopologyZone)
		}
		outputData = append(outputData, lineItems)
	}
	return headers, outputData
}

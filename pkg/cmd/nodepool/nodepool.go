package nodepool

import (
	"context"

	"github.com/Kavinraja-G/node-gizmo/pkg/outputs"
	"github.com/Kavinraja-G/node-gizmo/utils"

	"github.com/Kavinraja-G/node-gizmo/pkg"
	"github.com/spf13/cobra"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var sortByHeader string

func NewCmdNodepoolInfo() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "nodepool",
		Short:   "Displays detailed information about Nodepool",
		Aliases: []string{"np", "ng"},
		RunE: func(cmd *cobra.Command, args []string) error {
			return showNodePoolInfo(cmd, args)
		},
	}

	// additional local flags
	cmd.Flags().StringVarP(&sortByHeader, "sort-by", "", "nodepool", "Sorts output using a valid Column name. Defaults to 'nodepool' if the column name is not valid")

	return cmd
}

// showNodePoolInfo driver function for the 'nodepool' command
func showNodePoolInfo(cmd *cobra.Command, args []string) error {
	var genericNodepoolInfos []pkg.GenericNodepoolInfo

	nodes, err := utils.Cfg.Clientset.CoreV1().Nodes().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return err
	}

	for _, node := range nodes.Items {
		cloudProvider, nodepoolID := getNodepoolIDAndProvider(node.Labels)
		region, zone := pkg.GetNodeTopologyInfo(node.Labels)

		genericNodepoolInfos = append(genericNodepoolInfos, pkg.GenericNodepoolInfo{
			NodepoolID:   nodepoolID,
			Node:         node.Name,
			Provider:     cloudProvider,
			InstanceType: getNodeInstanceType(node.Labels),
			Region:       region,
			Zone:         zone,
			K8sVersion:   node.Status.NodeInfo.KubeletVersion,
		})
	}

	outputHeaders, outputData := generateNodepoolInfoData(genericNodepoolInfos)
	outputs.SortOutputBasedOnHeader(outputHeaders, outputData, sortByHeader)
	outputs.TableOutput(outputHeaders, outputData)

	return nil
}

// getNodepoolIDAndProvider returns the cloud provider type for the nodepool (EKS/Karpenter, GKE, AKS, can be Unknown)
func getNodepoolIDAndProvider(labels map[string]string) (string, string) {
	if id, ok := labels[pkg.AwsNodepoolLabel]; ok {
		return "EKS", id
	}
	if id, ok := labels[pkg.GkeNodepoolLabel]; ok {
		return "GKE", id
	}
	if id, ok := labels[pkg.AksNodepoolLabel]; ok {
		return "AKS", id
	}
	if id, ok := labels[pkg.KarpenterNodepool]; ok {
		return "Karpenter", id
	}
	if id, ok := labels[pkg.KarpenterNodepoolV1]; ok {
		return "Karpenter", id
	}
	if id, ok := labels[pkg.OpenshiftMachinepool]; ok {
		return "Openshift", id
	}

	return "Unknown", "Unknown"
}

// getNodeInstanceType returns the node instanceType based on the instance-type label
func getNodeInstanceType(labels map[string]string) string {
	if val, ok := labels[pkg.NodeInstanceTypeLabel]; ok {
		return val
	}

	return "Unknown"
}

// generateNodepoolInfoData generates the Nodepool info outputs and the required headers for table-writer
func generateNodepoolInfoData(genericNodepoolInfos []pkg.GenericNodepoolInfo) ([]string, [][]string) {
	var headers = []string{"NODEPOOL", "PROVIDER", "REGION", "ZONE", "INSTANCE-TYPE", "VERSION", "NODES"}
	var outputData [][]string

	for _, nodepoolInfo := range genericNodepoolInfos {
		lineItems := []string{
			nodepoolInfo.NodepoolID,
			nodepoolInfo.Provider,
			nodepoolInfo.Region,
			nodepoolInfo.Zone,
			nodepoolInfo.InstanceType,
			nodepoolInfo.K8sVersion,
			nodepoolInfo.Node,
		}

		outputData = append(outputData, lineItems)
	}

	return headers, outputData
}

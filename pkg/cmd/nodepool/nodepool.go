package nodepool

import (
	"context"
	"github.com/Kavinraja-G/node-gizmo/utils"
	"strings"

	"github.com/Kavinraja-G/node-gizmo/pkg/outputs"

	"github.com/Kavinraja-G/node-gizmo/pkg"
	"github.com/spf13/cobra"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func NewCmdNodepoolInfo() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "nodepool",
		Short:   "Displays detailed information about Nodepool/Nodegroup",
		Aliases: []string{"np", "ng"},
		RunE: func(cmd *cobra.Command, args []string) error {
			return showNodePoolInfo(cmd, args)
		},
	}

	return cmd
}

// showNodePoolInfo driver function for the 'nodepool' command
func showNodePoolInfo(cmd *cobra.Command, args []string) error {
	var genericNodepoolInfos = make(map[string]pkg.GenericNodepoolInfo)

	nodes, err := utils.Cfg.Clientset.CoreV1().Nodes().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return err
	}

	for _, node := range nodes.Items {
		var genericNodepoolInfo pkg.GenericNodepoolInfo

		cloudProvider, nodepoolID := getNodepoolIDAndProvider(node.Labels)
		if _, ok := genericNodepoolInfos[nodepoolID]; !ok {
			genericNodepoolInfo.NodepoolID = nodepoolID
			genericNodepoolInfo.Nodes = append(genericNodepoolInfo.Nodes, node.Name)
			genericNodepoolInfo.Provider = cloudProvider
			genericNodepoolInfo.InstanceType = getNodeInstanceType(node.Labels)
			genericNodepoolInfo.Region, genericNodepoolInfo.Zone = pkg.GetNodeTopologyInfo(node.Labels)

			// finally add the genericNodepoolInfo data to the genericNodepoolInfos
			genericNodepoolInfos[nodepoolID] = genericNodepoolInfo
		} else {
			var currentNodepoolInfo = genericNodepoolInfos[nodepoolID]
			currentNodepoolInfo.Nodes = append(currentNodepoolInfo.Nodes, node.Name)
			genericNodepoolInfos[nodepoolID] = currentNodepoolInfo
		}
	}

	outputHeaders, outputData := generateNodepoolInfoData(genericNodepoolInfos)
	outputs.TableOutput(outputHeaders, outputData)

	return nil
}

// getNodeAddresses returns the respective nodeAddress values for 'Hostname', 'InternalIP', 'ExternalIP', 'ExternalDNS'
// TODO: Add nodeAddress fields for detailed node info using flags
func getNodeAddresses(addresses []corev1.NodeAddress, addressType string) string {
	for _, address := range addresses {
		if (address.Type == corev1.NodeHostName) && (addressType == string(corev1.NodeHostName)) {
			return address.Address
		}
		if (address.Type == corev1.NodeInternalIP) && (addressType == string(corev1.NodeInternalIP)) {
			return address.Address
		}
		if (address.Type == corev1.NodeExternalIP) && (addressType == string(corev1.NodeExternalIP)) {
			return address.Address
		}
		if (address.Type == corev1.NodeExternalDNS) && (addressType == string(corev1.NodeExternalDNS)) {
			return address.Address
		}
	}
	return "Unknown"
}

// getNodepoolIDAndProvider returns the cloud provider type for the nodepool (EKS, GKE, AKS, can be Unknown)
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
func generateNodepoolInfoData(genericNodepoolInfos map[string]pkg.GenericNodepoolInfo) ([]string, [][]string) {
	var headers = []string{"NODEPOOL", "PROVIDER", "REGION", "ZONE", "INSTANCE-TYPE", "NODES"}
	var outputData [][]string

	for _, nodepoolInfo := range genericNodepoolInfos {
		lineItems := []string{
			nodepoolInfo.NodepoolID,
			nodepoolInfo.Provider,
			nodepoolInfo.Region,
			nodepoolInfo.Zone,
			nodepoolInfo.InstanceType,
			strings.Join(nodepoolInfo.Nodes, "\n"),
		}

		outputData = append(outputData, lineItems)
	}

	return headers, outputData
}

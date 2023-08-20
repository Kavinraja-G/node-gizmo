package cmd

import (
	"context"
	"log"

	"github.com/Kavinraja-G/node-gizmo/pkg"
	"github.com/Kavinraja-G/node-gizmo/pkg/auth"
	"github.com/Kavinraja-G/node-gizmo/pkg/utils"
	"github.com/spf13/cobra"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func NewCmdNodeInfo() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "node",
		Short:   "Generic node information",
		Aliases: []string{"nodes"},
		RunE: func(cmd *cobra.Command, args []string) error {
			return showNodeInfo(cmd, args)
		},
	}

	return cmd
}

func showNodeInfo(cmd *cobra.Command, args []string) error {
	var nodeInfos []pkg.GenericNodeInfo

	clientset, err := auth.K8sAuth()
	if err != nil {
		log.Fatalf("Error while authenticating to kubernetes: %v", err)
	}

	nodes, err := clientset.CoreV1().Nodes().List(context.TODO(), metav1.ListOptions{})
	for _, node := range nodes.Items {
		nodeInfos = append(nodeInfos, pkg.GenericNodeInfo{
			NodeName:   node.Name,
			K8sVersion: node.Status.NodeInfo.KubeletVersion,
			Image:      node.Status.NodeInfo.OSImage,
			Os:         node.Status.NodeInfo.OperatingSystem,
			OsArch:     node.Status.NodeInfo.Architecture,
			NodeStatus: getNodeStatus(node.Status.Conditions),
		})
	}

	utils.OutputGenericNodeInfo(nodeInfos)

	return nil
}

func getNodeStatus(nodeConditions []corev1.NodeCondition) string {
	for _, nodeCondition := range nodeConditions {
		if nodeCondition.Type == corev1.NodeReady {
			return "Ready"
		}
	}
	return "NotReady"
}

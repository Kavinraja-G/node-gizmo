package cmd

import (
	"context"
	"fmt"
	"log"

	"github.com/Kavinraja-G/node-gizmo/pkg/outputs"

	"github.com/Kavinraja-G/node-gizmo/pkg"
	"github.com/Kavinraja-G/node-gizmo/pkg/auth"
	"github.com/spf13/cobra"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var showTaints bool

func NewCmdNodeInfo() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "node",
		Short:   "Generic node information",
		Aliases: []string{"nodes"},
		RunE: func(cmd *cobra.Command, args []string) error {
			return showNodeInfo(cmd, args)
		},
	}

	// flags
	cmd.Flags().BoolVarP(&showTaints, "show-taints", "t", false, "Shows taints on Nodes in the output")

	return cmd
}

func showNodeInfo(cmd *cobra.Command, args []string) error {
	var nodeInfos []pkg.GenericNodeInfo
	var outputOpts = pkg.OutputOptsForGenericNodeInfo{
		ShowTaints: showTaints,
	}

	clientset, err := auth.K8sAuth()
	if err != nil {
		log.Fatalf("Error while authenticating to kubernetes: %v", err)
	}

	nodes, err := clientset.CoreV1().Nodes().List(context.TODO(), metav1.ListOptions{})
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

		nodeInfos = append(nodeInfos, genericNodeInfo)
	}

	outputs.OutputGenericNodeInfo(nodeInfos, outputOpts)

	return nil
}

func getNodeTaints(rawTaints []corev1.Taint) []string {
	var taints []string
	for _, taint := range rawTaints {
		taints = append(taints, fmt.Sprintf("%v=%v:%v", taint.Key, taint.Value, taint.Effect))
	}
	return taints
}

func getNodeStatus(nodeConditions []corev1.NodeCondition) string {
	for _, nodeCondition := range nodeConditions {
		if nodeCondition.Type == corev1.NodeReady {
			return "Ready"
		}
	}
	return "NotReady"
}

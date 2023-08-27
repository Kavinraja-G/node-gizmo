package nodes

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/Kavinraja-G/node-gizmo/pkg/auth"
	"github.com/spf13/cobra"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func NewCmdNodeExec() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "exec",
		Short:   "Spawns a nsenter pod to exec into the respective node",
		Aliases: []string{"ex"},
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return errors.New("please provide a nodeName to exec into")
			}
			if !isValidNode(args[0]) {
				return errors.New(fmt.Sprintf("%v is not a valid node", args[0]))
			}
			return execIntoNode(cmd, args[0])
		},
	}
	return cmd
}

func isValidNode(nodeName string) bool {
	clientset, err := auth.K8sAuth()
	if err != nil {
		log.Fatalf("Error while authenticating to kubernetes: %v", err)
	}

	nodes, err := clientset.CoreV1().Nodes().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Fatalf("Error while listing the nodes in the cluster: %v", err)
	}

	for _, node := range nodes.Items {
		if node.Name == nodeName {
			return true
		}
	}

	return false
}

func execIntoNode(cmd *cobra.Command, nodeName string) error {
	var nodeshellPodName = fmt.Sprintf("nodeshell-%v", nodeName)
	var nodeshellPodNamespace = "default"
	var podSCPrivileged = true
	var podTerminationGracePeriodSeconds = int64(0)

	pod := &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      nodeshellPodName,
			Namespace: nodeshellPodNamespace,
			Labels: map[string]string{
				"app.kubernetes.io/name":       nodeshellPodName,
				"app.kubernetes.io/version":    cmd.Version,
				"app.kubernetes.io/component":  "exec",
				"app.kubernetes.io/managed-by": "node-gizmo",
			},
		},
		Spec: corev1.PodSpec{
			Containers: []corev1.Container{
				{
					Name:    "nodeshell",
					Image:   "docker.io/alpine:3.18",
					Command: []string{"nsenter"},
					Args:    []string{"-t", "1", "-m", "-u", "-i", "-n", "sleep", "14000"},
					SecurityContext: &corev1.SecurityContext{
						Privileged: &podSCPrivileged,
					},
				},
			},
			RestartPolicy:                 corev1.RestartPolicyNever,
			TerminationGracePeriodSeconds: &podTerminationGracePeriodSeconds,
			HostNetwork:                   true,
			HostPID:                       true,
			HostIPC:                       true,
			Tolerations: []corev1.Toleration{
				{
					Operator: corev1.TolerationOpExists, // this will attract any taints added to nodes
				},
			},
			NodeSelector: map[string]string{
				"kubernetes.io/hostname": nodeName,
			},
			NodeName: nodeName,
		},
	}

	clientset, err := auth.K8sAuth()
	if err != nil {
		log.Fatalf("Error while authenticating to kubernetes: %v", err)
	}

	_, err = clientset.CoreV1().Pods(nodeshellPodNamespace).Create(context.TODO(), pod, metav1.CreateOptions{})
	if err != nil {
		log.Fatalf("Error while creating the nodeshell pod: %v", err)
	}

	return nil
}

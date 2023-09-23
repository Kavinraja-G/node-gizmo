package nodes

import (
	"context"
	"errors"
	"fmt"
	"github.com/Kavinraja-G/node-gizmo/utils"
	"log"
	"os"
	"time"

	k8errors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/util/wait"

	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/tools/remotecommand"

	"github.com/spf13/cobra"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var (
	nodeshellPodNamespace            string
	nodeshellPodImage                string
	nodeshellPodNamePrefix           = "nodeshell-"
	podSCPrivileged                  = true
	podTerminationGracePeriodSeconds = int64(0)
)

// NewCmdNodeExec initialises the 'exec' command
func NewCmdNodeExec() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "exec nodeName",
		Short:   "Spawns a 'nsenter' pod to exec into the provided node",
		Aliases: []string{"ex"},
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return errors.New("please provide a nodeName to exec")
			}
			nodeName := args[0]

			if !isValidNode(nodeName) {
				return errors.New(fmt.Sprintf("%v is not a valid node", args[0]))
			}

			err := createExecPodInTargetedNode(nodeName)
			if err != nil {
				return err
			}

			return execIntoNode(cmd, args[0])
		},
		PostRunE: func(cmd *cobra.Command, args []string) error {
			return cleanUpNodeshellPods(cmd, args[0])
		},
	}

	// additional local flags
	cmd.Flags().StringVarP(&nodeshellPodNamespace, "namespace", "n", "kube-system", "Namespace where nsenter pod to be created")
	cmd.Flags().StringVarP(&nodeshellPodImage, "image", "i", "docker.io/alpine:3.18", "Image used by nsenter pod")

	return cmd
}

// isValidNode validates the given node is available in the cluster or not
func isValidNode(nodeName string) bool {
	nodes, err := utils.Cfg.Clientset.CoreV1().Nodes().List(context.TODO(), metav1.ListOptions{})
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

// createExecPodInTargetedNode creates the nsenter pod in the given node
func createExecPodInTargetedNode(nodeName string) error {
	var nodeshellPodName = fmt.Sprintf("nodeshell-%v", nodeName)

	pod := &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      nodeshellPodName,
			Namespace: nodeshellPodNamespace,
			Labels: map[string]string{
				"app.kubernetes.io/name":       nodeshellPodName,
				"app.kubernetes.io/component":  "exec",
				"app.kubernetes.io/managed-by": "node-gizmo",
			},
		},
		Spec: corev1.PodSpec{
			Containers: []corev1.Container{
				{
					Name:    "nodeshell",
					Image:   nodeshellPodImage,
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

	_, err := utils.Cfg.Clientset.CoreV1().Pods(nodeshellPodNamespace).Create(context.TODO(), pod, metav1.CreateOptions{})
	if err != nil {
		// validates if the pod already exists
		if k8errors.IsAlreadyExists(err) {
			return nil
		}
		return err
	}

	// wait for exec pod to get RUNNING
	checkExecPodRunningStatus := func() (bool, error) {
		pod, err := utils.Cfg.Clientset.CoreV1().Pods(nodeshellPodNamespace).Get(context.TODO(), nodeshellPodName, metav1.GetOptions{})
		if err != nil {
			return false, err
		}

		if pod.Status.Phase == corev1.PodRunning {
			return true, nil
		}

		return false, nil
	}

	backoff := wait.Backoff{
		Steps:    5,                // Number of retry steps.
		Duration: 10 * time.Second, // Initial backoff duration.
		Factor:   1.0,              // Multiplier for each step's duration.
		Jitter:   0.1,              // Jitter to randomize the duration slightly.
	}

	// waits exponentially for exec pod to be RUNNING
	startTime := time.Now()
	err = wait.ExponentialBackoff(backoff, func() (bool, error) {
		elapsedWaitTime := time.Since(startTime)
		log.Printf("Waiting for exec pod %s to be RUNNING. Elapsed time: %v", nodeshellPodName, elapsedWaitTime)
		return checkExecPodRunningStatus()
	})
	if err != nil {
		log.Fatalf("exec pod did not reached the RUNNING state: %v", err)
	}

	return err
}

// execIntoNode is the driver function used to exec into the nsenter pod deployed in the targeted node
func execIntoNode(cmd *cobra.Command, nodeName string) error {
	var nodeshellPodName = nodeshellPodNamePrefix + nodeName

	var podExecCmd = []string{"sh", "-c", "(bash || ash || sh)"}
	req := utils.Cfg.Clientset.CoreV1().RESTClient().Post().Resource("pods").Name(nodeshellPodName).Namespace(nodeshellPodNamespace).SubResource("exec")
	opts := &corev1.PodExecOptions{
		Command: podExecCmd,
		Stdin:   true,
		Stdout:  true,
		Stderr:  true,
		TTY:     true,
	}

	req.VersionedParams(opts, scheme.ParameterCodec)

	//TODO: Check if there is any way we can fetch the config from the clientset itself
	k8sConfig, err := utils.GetKubeConfig()
	if err != nil {
		log.Fatalf("Error while getting Kubeconfig: %v", err)
		return err
	}

	// initiate the exec command on the nsenter pod and creates a bidirectional connection to the server
	exec, err := remotecommand.NewSPDYExecutor(k8sConfig, "POST", req.URL())
	if err != nil {
		log.Fatalf("Error while running exec on nodeshell pod: %v", err)
		return err
	}

	err = exec.StreamWithContext(context.TODO(), remotecommand.StreamOptions{
		Stdin:  os.Stdin,
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	})

	return err
}

// cleanUpNodeshellPods cleans up the nsenter pod once the shell is exited
func cleanUpNodeshellPods(cmd *cobra.Command, nodeName string) error {
	var nodeshellPodName = nodeshellPodNamePrefix + nodeName

	err := utils.Cfg.Clientset.CoreV1().Pods(nodeshellPodNamespace).Delete(context.TODO(), nodeshellPodName, metav1.DeleteOptions{})
	if err != nil {
		log.Fatalf("Error while creating the nodeshell pod: %v", err)
		return err
	}

	return err
}

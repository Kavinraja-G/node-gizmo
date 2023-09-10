package auth

import (
	"log"
	"path/filepath"

	"github.com/Kavinraja-G/node-gizmo/pkg/utils"

	"k8s.io/client-go/rest"
	"k8s.io/client-go/util/homedir"

	k8s "k8s.io/client-go/kubernetes"

	_ "k8s.io/client-go/plugin/pkg/client/auth/oidc"
	"k8s.io/client-go/tools/clientcmd"
)

// GetKubeConfig is used to fetch the kubeConfig based on the KUBECONFIG env or '~/.kube/config' location
func GetKubeConfig() (*rest.Config, error) {
	var kubeConfigPath string
	if home := homedir.HomeDir(); home != "" {
		kubeConfigPath = filepath.Join(home, ".kube", "config")
	} else {
		kubeConfigPath = utils.GetEnv("KUBECONFIG", "~/.kube/config")
	}

	k8sConfig, err := clientcmd.BuildConfigFromFlags("", kubeConfigPath)
	return k8sConfig, err
}

// K8sAuth is used to get the Kubernetes clientset from the config
func K8sAuth() (*k8s.Clientset, error) {
	k8sConfig, err := GetKubeConfig()
	clientset, err := k8s.NewForConfig(k8sConfig)
	if err != nil {
		log.Fatalf("Error while creating the clientset: %v", err)
	}

	return clientset, err
}

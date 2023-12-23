package utils

import (
	"log"
	"path/filepath"

	k8s "k8s.io/client-go/kubernetes"
	_ "k8s.io/client-go/plugin/pkg/client/auth/oidc"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

type Config struct {
	Clientset *k8s.Clientset
}

var Cfg Config

// GetKubeConfig is used to fetch the kubeConfig based on the KUBECONFIG env or '~/.kube/config' location
func GetKubeConfig() (*rest.Config, error) {
	var kubeConfigPath string
	var defaultKubeConfigPath = "~/.kube/config"

	if home := homedir.HomeDir(); home != "" {
		defaultKubeConfigPath = filepath.Join(home, ".kube", "config")
	}
	kubeConfigPath = GetEnv("KUBECONFIG", defaultKubeConfigPath)

	k8sConfig, err := clientcmd.BuildConfigFromFlags("", kubeConfigPath)
	return k8sConfig, err
}

// k8sAuth is used to get the Kubernetes clientset from the config
func k8sAuth() (*k8s.Clientset, error) {
	k8sConfig, err := GetKubeConfig()
	if err != nil {
		log.Fatalf("Error while getting kubeconfig: %v", err)
	}

	clientset, err := k8s.NewForConfig(k8sConfig)
	if err != nil {
		log.Fatalf("Error while creating the clientset: %v", err)
	}

	return clientset, err
}

// InitConfig initiates a kubernetes clientset & other generic configs with the current context
func InitConfig() {
	var err error
	Cfg.Clientset, err = k8sAuth()
	if err != nil {
		log.Fatalf("Error while authenticating to kubernetes: %v", err)
	}
}

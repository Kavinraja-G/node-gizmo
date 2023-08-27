package auth

import (
	"k8s.io/client-go/rest"
	"k8s.io/client-go/util/homedir"
	"log"
	"path/filepath"
)

import (
	"github.com/Kavinraja-G/kube-bouncer/pkg/utils"
	k8s "k8s.io/client-go/kubernetes"
	_ "k8s.io/client-go/plugin/pkg/client/auth/oidc"
	"k8s.io/client-go/tools/clientcmd"
)

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

func K8sAuth() (*k8s.Clientset, error) {
	k8sConfig, err := GetKubeConfig()
	clientset, err := k8s.NewForConfig(k8sConfig)
	if err != nil {
		log.Fatalf("Error while creating the clientset: %v", err)
	}

	return clientset, err
}

package auth

import (
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

func K8sAuth() (*k8s.Clientset, error) {
	var kubeConfigPath string
	if home := homedir.HomeDir(); home != "" {
		kubeConfigPath = filepath.Join(home, ".kube", "config")
	} else {
		kubeConfigPath = utils.GetEnv("KUBECONFIG", "~/.kube/config")
	}

	k8cfg, err := clientcmd.BuildConfigFromFlags("", kubeConfigPath)

	clientset, err := k8s.NewForConfig(k8cfg)
	if err != nil {
		log.Fatalf("Error while creating the clientset: %v", err)
	}

	return clientset, err
}

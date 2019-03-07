package util

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	_ "k8s.io/client-go/plugin/pkg/client/auth"
	"k8s.io/client-go/tools/clientcmd"
)

var clientset *kubernetes.Clientset

// GetNamespaces is accessing current context and returning all namespace name
func GetNamespaces() []string {
	clientset, err := newClient()
	if err != nil {
		panic(err.Error())
	}
	ns, err := clientset.CoreV1().Namespaces().List(metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	namespaces := []string{}
	for _, item := range ns.Items {
		namespaces = append(namespaces, item.GetName())
	}
	return namespaces
}

// newClient return kubectl client for accessing kubernetes with out-of-box
func newClient() (kubernetes.Interface, error) {
	loadingRules := clientcmd.NewDefaultClientConfigLoadingRules()
	config, err := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(loadingRules, &clientcmd.ConfigOverrides{}).ClientConfig()
	if err != nil {
		return nil, err
	}
	return kubernetes.NewForConfig(config)
}

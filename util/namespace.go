package util

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	_ "k8s.io/client-go/plugin/pkg/client/auth"
	"k8s.io/client-go/tools/clientcmd"
)

var clientset *kubernetes.Clientset

func init() {
	config, err := clientcmd.BuildConfigFromFlags("", configFile)
	if err != nil {
		panic(err.Error())
	}
	// create the clientset
	clientset, err = kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
}

func GetNamespaces() []string {

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

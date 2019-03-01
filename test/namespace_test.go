package cmd

import (
	"fmt"
	"os"
	"testing"

	"github.com/onuryartasi/context-manager/util"
)

func TestChangeNamespace(t *testing.T) {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	os.Setenv("KUBECONFIG", fmt.Sprintf("%s/mocks/config1:%s/mocks/config2", wd, wd))
	config := util.GetRawConfig()
	namespaces := []string{"default", "kube-public", "kube-system"}
	contexts := util.GetContexts(config)
	for _, context := range contexts {
		util.SetContext(context)
		config2 := util.GetRawConfig()
		for _, namespace := range namespaces {
			util.SetNamespace(config2, namespace)
			config3 := util.GetRawConfig()
			if namespace != config3.Contexts[context].Namespace {
				t.Errorf("Expected: %s\nActual: %s\ninput:%s", namespace, config3.Contexts[context].Namespace, namespace)
			}
		}
	}

}
func TestPrintCurrentNamespace(t *testing.T) {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	os.Setenv("KUBECONFIG", fmt.Sprintf("%s/mocks/config1:%s/mocks/config2", wd, wd))
	config := util.GetRawConfig()
	namespaces := []string{"default", "kube-public", "kube-system"}
	contexts := util.GetContexts(config)
	for _, context := range contexts {
		util.SetContext(context)
		config2 := util.GetRawConfig()
		for _, namespace := range namespaces {
			util.SetNamespace(config2, namespace)
			_, actual := util.GetCurrentContext()
			if actual != namespace {
				t.Errorf("Expected: %s\nActual: %s\n", namespace, actual)
			}
		}
	}
}

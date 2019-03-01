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
	os.Setenv("KUBECONFIG", fmt.Sprintf("%s/test/mocks/config1:%s/test/mocks/config2", wd, wd))
	config := util.GetRawConfig()
	namespaces := []string{"default", "kube-public", "kube-system"}
	contexts := util.GetContexts(config)
	for _, context := range contexts {
		util.SetContext(context)
		for _, namespace := range namespaces {
			util.SetNamespace(config, namespace)
			config2 := util.GetRawConfig()
			if namespace != config2.Contexts[context].Namespace {
				t.Errorf("Expected: %s\nActual: %s\n,input:%s", namespace, config2.Contexts[context].Namespace, namespace)
			}
		}
	}

}

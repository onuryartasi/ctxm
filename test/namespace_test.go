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
	input := namespaces[0]

	err = util.SetNamespace(config, input)
	if err != nil {
		panic(err)
	}
	config2 := util.GetRawConfig()

	if input != config2.Contexts[config2.CurrentContext].Namespace {
		t.Errorf("Doesn't matches %s, %s", input, config2.Contexts[config2.CurrentContext].Namespace)
	}

}

package cmd

import (
	"fmt"
	"os"
	"testing"

	"github.com/onuryartasi/context-manager/util"
)

func TestGetContextNames(t *testing.T) {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	os.Setenv("KUBECONFIG", fmt.Sprintf("%s/mocks/config1:%s/mocks/config2", wd, wd))
	config := util.GetRawConfig()
	contexts := util.GetContexts(config)
	if !(len(contexts) > 0) {
		t.Errorf("Getting any contexts, %v", contexts)
	}
}

func TestChangeContext(t *testing.T) {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	os.Setenv("KUBECONFIG", fmt.Sprintf("%s/mocks/config1:%s/mocks/config2", wd, wd))
	config := util.GetRawConfig()
	contexts := util.GetContexts(config)

	for _, context := range contexts {
		util.SetContext(context)
		config2 := util.GetRawConfig()
		if context != config2.CurrentContext {
			t.Errorf("Expected: %s\nActual: %s\n,input:%s", context, config2.CurrentContext, context)
		}
	}

}

func TestPrintCurrentContext(t *testing.T) {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	os.Setenv("KUBECONFIG", fmt.Sprintf("%s/mocks/config1:%s/mocks/config2", wd, wd))
	config := util.GetRawConfig()
	contexts := util.GetContexts(config)
	for _, context := range contexts {
		util.SetContext(context)
		actual, _ := util.GetCurrentContext()
		if actual != context {
			t.Errorf("Expected: %s\nActual: %s\n", context, actual)
		}
	}
}

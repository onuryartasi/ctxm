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
	os.Setenv("KUBECONFIG", fmt.Sprintf("%s//test/mocks/config1:%s/test/mocks/config2", wd, wd))
	config := util.GetRawConfig()
	contexts := util.GetContexts(config)
	fmt.Println(contexts)
}

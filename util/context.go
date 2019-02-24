package util

import (
	"io/ioutil"
	"log"
	"path/filepath"

	"github.com/ghodss/yaml"
	"github.com/mitchellh/go-homedir"
	"github.com/onuryartasi/context-manager/types"
)

var kubeConfig types.KubeConfig

func GetContexts() []string {
	contextNames := []string{}
	home, err := homedir.Dir()
	if err != nil {
		log.Fatalf("Homedir error %s", err)
	}
	configFile := filepath.Join(home, ".kube", "config")
	config, err := ioutil.ReadFile(configFile)
	if err != nil {
		log.Fatalf("Can't read config %s, Error: %s", configFile, err)
	}
	err = yaml.Unmarshal(config, &kubeConfig)
	if err != nil {
		log.Fatalf("Can't parse configfile %s", err)
	}
	for _, context := range kubeConfig.Contexts {
		contextNames = append(contextNames, context.Name)
	}
	return contextNames
}

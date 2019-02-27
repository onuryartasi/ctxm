package util

import (
	"io/ioutil"
	"os"
	"strings"

	"github.com/ghodss/yaml"
	"github.com/onuryartasi/context-manager/types"
	"k8s.io/client-go/tools/clientcmd"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"
)

func GetConfigFilePath() string {
	kubeConfigEnv := os.Getenv("KUBECONFIG")

	if len(kubeConfigEnv) > 0 {
		return kubeConfigEnv
	}

	return clientcmd.RecommendedHomeFile
}

func GetRawConfig() clientcmdapi.Config {
	loadingRules := clientcmd.NewDefaultClientConfigLoadingRules()
	config, err := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(loadingRules, &clientcmd.ConfigOverrides{}).RawConfig()
	if err != nil {
		panic(err)
	}

	return config
}

func GetStructConfig(configPath string) types.KubeConfig {
	var config types.KubeConfig
	configData, err := ioutil.ReadFile(configPath)
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(configData, &config)
	if err != nil {
		panic(err)
	}
	return config
}

func SetNamespace(config clientcmdapi.Config, namespace string) error {

	kubeConfigEnv := os.Getenv("KUBECONFIG")
	if len(kubeConfigEnv) > 0 {
		configPaths := strings.Split(kubeConfigEnv, ":")
		if len(configPaths) > 1 {
			for _, configPath := range configPaths {
				configBase, _ := clientcmd.LoadFromFile(configPath)
				_, ok := configBase.Contexts[config.CurrentContext]
				if ok {
					configBase.Contexts[config.CurrentContext].Namespace = namespace
					err := clientcmd.WriteToFile(*configBase, configPath)
					return err
				}

			}
		}

	}

	config.Contexts[config.CurrentContext].Namespace = namespace
	err := clientcmd.WriteToFile(config, clientcmd.RecommendedHomeFile)
	return err
}

func GetContexts(config clientcmdapi.Config) []string {
	contexts := []string{}
	for key := range config.Contexts {
		contexts = append(contexts, key)
	}
	return contexts

}

func SetContext(contex string) {

	kubeConfigEnv := os.Getenv("KUBECONFIG")
	if len(kubeConfigEnv) > 0 {
		configPaths := strings.Split(kubeConfigEnv, ":")
		config, err := clientcmd.LoadFromFile(configPaths[0])
		if err != nil {
			panic(err)
		}

		config.CurrentContext = contex
		err = clientcmd.WriteToFile(*config, configPaths[0])
		if err != nil {
			panic(err)
		}
		return
	}

	config, err := clientcmd.LoadFromFile(clientcmd.RecommendedHomeFile)
	if err != nil {
		panic(err)
	}
	config.CurrentContext = contex
	err = clientcmd.WriteToFile(*config, clientcmd.RecommendedHomeFile)
	if err != nil {
		panic(err)
	}

}

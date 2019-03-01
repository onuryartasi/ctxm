package util

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	homedir "github.com/mitchellh/go-homedir"
	"gopkg.in/yaml.v2"
	"k8s.io/client-go/tools/clientcmd"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"
)

type PrevContextConfig struct {
	PrevContext   string `yaml:"PrevContext"`
	PrevNamespace string `yaml:"PrevNamespace"`
}

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

func SetNamespace(config clientcmdapi.Config, namespace string) error {

	kubeConfigEnv := os.Getenv("KUBECONFIG")
	if len(kubeConfigEnv) > 0 {
		configPaths := strings.Split(kubeConfigEnv, ":")
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

func GetPrevContextConfig() PrevContextConfig {
	var config PrevContextConfig
	configPath, ok := IsExistsPrevContext()
	if !ok {
		config = PrevContextConfig{PrevContext: "", PrevNamespace: ""}
		data, err := yaml.Marshal(&config)
		if err != nil {
			panic(err)
		}

		err = ioutil.WriteFile(configPath, data, 0644)
		return config
	}

	reading, err := ioutil.ReadFile(configPath)
	if err != nil {
		panic(err)
	}
	yaml.Unmarshal(reading, &config)
	return config
}

func IsExistsPrevContext() (string, bool) {
	dir, err := homedir.Dir()
	if err != nil {
		panic(err)
	}
	configPath := filepath.Join(dir, ".context-manager", "config")
	_, err = os.Stat(configPath)
	if os.IsNotExist(err) {
		return configPath, false
	}
	return configPath, true
}
func (config *PrevContextConfig) SetContextPrevContextConfig(context string) {
	config.PrevContext = context
	config.PrevNamespace = ""

}

func (config *PrevContextConfig) SetNamespacePrevContextConfig(namespace string) {
	config.PrevNamespace = namespace

}

func (config *PrevContextConfig) WriteFile() {
	data, err := yaml.Marshal(&config)
	if err != nil {
		panic(err)
	}
	configPath, ok := IsExistsPrevContext()

	if ok {
		err := ioutil.WriteFile(configPath, data, 0644)
		if err != nil {
			panic(err)
		}
	}

}

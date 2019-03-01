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

//PrevConfig Previous Context and Namespace storage type
type PrevConfig struct {
	PrevContext   string `yaml:"PrevContext"`
	PrevNamespace string `yaml:"PrevNamespace"`
}

// GetConfigFilePath if set KUBECONFING than return this or return RecomendedHomeFile (ex. /home/$USER/.kube/config)
func GetConfigFilePath() string {
	kubeConfigEnv := os.Getenv(clientcmd.RecommendedConfigPathEnvVar)

	if len(kubeConfigEnv) > 0 {
		return kubeConfigEnv
	}

	return clientcmd.RecommendedHomeFile
}

// GetRawConfig is return kubeconfig struct, if have a multiple kubeconfig before merged later return struct
func GetRawConfig() clientcmdapi.Config {
	loadingRules := clientcmd.NewDefaultClientConfigLoadingRules()
	config, err := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(loadingRules, &clientcmd.ConfigOverrides{}).RawConfig()
	if err != nil {
		panic(err)
	}

	return config
}

// SetNamespace is changed current namespace for current context, if have a multiple kubeconfig, searching context name in KUBECONFIG env later writing  to founded ConfigPath
func SetNamespace(config clientcmdapi.Config, namespace string) {

	configFilePath := GetConfigFilePath()

	configPaths := strings.Split(configFilePath, ":")
	for _, configPath := range configPaths {
		configBase, _ := clientcmd.LoadFromFile(configPath)
		_, ok := configBase.Contexts[config.CurrentContext]
		if ok {
			configBase.Contexts[config.CurrentContext].Namespace = namespace
			err := clientcmd.WriteToFile(*configBase, configPath)
			if err != nil {
				panic(err)
			}
		}

	}

}

// GetContexts return context names in kubeconfig struct
func GetContexts(config clientcmdapi.Config) []string {
	contexts := []string{}
	for key := range config.Contexts {
		contexts = append(contexts, key)
	}
	return contexts

}

// SetContext Change current context in first file if setting KUBECONFIG env,because kubectl look first file in KUBECONFIG env ,if not set env than looking home config
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

// GetPrevConfig return PrevConfig struct, if PrevConfig not exists than create empty config
func GetPrevConfig() PrevConfig {
	var config PrevConfig
	configPath, ok := IsExistsPrevConfig()
	if !ok {
		config = PrevConfig{PrevContext: "", PrevNamespace: ""}
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

// IsExistsPrevConfig checker Prevconfig exists and return bool,configPath
func IsExistsPrevConfig() (string, bool) {
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

// SetContextPrevConfig is change prev context and empty namespace
func (config *PrevConfig) SetContextPrevConfig(context string) {
	config.PrevContext = context
	config.PrevNamespace = ""

}

// SetNamespacePrevConfig is changer prev namespace
func (config *PrevConfig) SetNamespacePrevConfig(namespace string) {
	config.PrevNamespace = namespace

}

// WriteFile is save current PrevConfig struct to file
func (config *PrevConfig) WriteFile() {
	data, err := yaml.Marshal(&config)
	if err != nil {
		panic(err)
	}
	configPath, ok := IsExistsPrevConfig()

	if ok {
		err := ioutil.WriteFile(configPath, data, 0644)
		if err != nil {
			panic(err)
		}
	}

}

package types

import (
	"time"
)

type KubeConfig struct {
	APIVersion string `yaml:"apiVersion"`
	Clusters   []struct {
		Cluster struct {
			CertificateAuthorityData string `yaml:"certificate-authority-data,omitempty"`
			Server                   string `yaml:"server"`
		} `yaml:"cluster"`
		Name string `yaml:"name"`
	} `yaml:"clusters"`
	Contexts []struct {
		Context struct {
			Cluster   string `yaml:"cluster"`
			User      string `yaml:"user"`
			Namespace string `yaml:"namespace,omitempty"`
		} `yaml:"context"`
		Name string `yaml:"name"`
	} `yaml:"contexts"`
	CurrentContext string `yaml:"current-context"`
	Kind           string `yaml:"kind"`
	Preferences    struct {
	} `yaml:"preferences"`
	Users []struct {
		Name string `yaml:"name"`
		User struct {
			ClientCertificateData string `yaml:"client-certificate-data,omitempty"`
			ClientKeyData         string `yaml:"client-key-data,omitempty"`
			ClientCertificate     string `yaml:"client-certificate,omitempty"`
			ClientKey             string `yaml:"client-key,omitempty"`
			AuthProvider          struct {
				Config struct {
					AccessToken string    `yaml:"access-token,omitempty"`
					CmdArgs     string    `yaml:"cmd-args,omitempty"`
					CmdPath     string    `yaml:"cmd-path,omitempty"`
					Expiry      time.Time `yaml:"expiry,omitempty"`
					ExpiryKey   string    `yaml:"expiry-key,omitempty"`
					TokenKey    string    `yaml:"token-key,omitempty"`
				} `yaml:"config"`
				Name string `yaml:"name"`
			} `yaml:"auth-provider,omitempty"`
		} `yaml:"user"`
	} `yaml:"users"`
}

func (config *KubeConfig) GetContextNames() []string {
	contextNames := []string{}
	for _, context := range config.Contexts {
		contextNames = append(contextNames, context.Name)
	}
	return contextNames
}

func (config *KubeConfig) SetContext(contextName string) {
	config.CurrentContext = contextName
}

func (config *KubeConfig) SetNamespace(namespace string) {
	for id, value := range config.Contexts {
		if value.Name == config.CurrentContext {
			config.Contexts[id].Context.Namespace = namespace
		}
	}
}

// Copyright © 2019 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"github.com/onuryartasi/context-manager/util"
	"github.com/spf13/cobra"
	"gopkg.in/AlecAivazis/survey.v1"
)

// namespaceCmd represents the namespace command
var namespaceCmd = &cobra.Command{
	Use:   "namespace",
	Short: "Change Current-context's namespace",
	Long:  `Change  current context's namespce with multiple kubeconfig or Home config`,
	Run: func(cmd *cobra.Command, args []string) {
		ChangeNamespace(args...)

	},
	Aliases: []string{"ns", "namespaces"},
}

func init() {
	rootCmd.AddCommand(namespaceCmd)

}

func ChangeNamespace(args ...string) {
	var selectedNamespace string
	config := util.GetRawConfig()
	namespaces := util.GetNamespaces()
	contextQuestion := &survey.Select{
		Message: "Choose a namespace:",
		Options: namespaces,
	}
	survey.AskOne(contextQuestion, &selectedNamespace, nil)
	if len(selectedNamespace) > 0 {

		util.SetNamespace(config, selectedNamespace)
		if config.Contexts[config.CurrentContext].Namespace != selectedNamespace {
			prevConfig := util.GetPrevConfig()
			prevConfig.SetNamespacePrevConfig(config.Contexts[config.CurrentContext].Namespace)
			prevConfig.WriteFile()
		}

	}

}

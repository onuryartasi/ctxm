// Copyright © 2019 Onur YARTAŞI <onuryartasi@live.com>
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
	"fmt"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/onuryartasi/context-manager/util"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gopkg.in/AlecAivazis/survey.v1"
)

var cfgFile string
var current bool

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ctxm",
	Short: "Easy to use kubernetes contexts",
	Run: func(cmd *cobra.Command, args []string) {
		if current {
			context, _ := util.GetCurrentContext()
			fmt.Printf("%s\n", context)
		} else {
			ChangeContext(args...)
		}

	},
}

// currentCmd represent current context
var currentCmd = &cobra.Command{
	Use:   "current",
	Short: "Get Current context and namespace",
	Long:  `Current context and current namespace printer`,
	Run: func(cmd *cobra.Command, args []string) {
		context, _ := util.GetCurrentContext()
		fmt.Printf("%s\n", context)
	},
	Aliases: []string{"."},
}

// previousCmd represents the previous command
var previousContextCmd = &cobra.Command{
	Use:   "previous",
	Short: "Change context to previous context ",
	Run: func(cmd *cobra.Command, args []string) {
		output := PreviousContext(args...)
		fmt.Println(output)
	},
	Aliases: []string{"prev", ".."},
}

func PreviousContext(args ...string) string {
	prevConfig := util.GetPrevConfig()
	config := util.GetRawConfig()
	if prevConfig.PrevContext != "" {
		prevContext := prevConfig.PrevContext
		util.SetContext(prevConfig.PrevContext)
		prevConfig.SetContextPrevConfig(config.CurrentContext)
		prevConfig.WriteFile()
		return fmt.Sprintf("%s", prevContext)
	} else {
		return fmt.Sprintf("Not found previous Context")
	}
}

// ChangeContext is changer current-context
func ChangeContext(args ...string) {
	var selectedContext string
	config := util.GetRawConfig()
	contexts := util.GetContexts(config)
	contextQuestion := &survey.Select{
		Message: "Choose a context:",
		Options: contexts,
	}
	survey.AskOne(contextQuestion, &selectedContext, nil)
	if len(selectedContext) > 0 {
		util.SetContext(selectedContext)

		if selectedContext != config.CurrentContext {
			prevConfig := util.GetPrevConfig()
			prevConfig.SetContextPrevConfig(config.CurrentContext)
			prevConfig.WriteFile()
		}

	}

}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.context-manager.yaml)")
	rootCmd.Flags().BoolVarP(&current, "current", "c", false, "Get Current Context")
	rootCmd.AddCommand(previousContextCmd)
	rootCmd.AddCommand(currentCmd)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".context-manager" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".context-manager")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

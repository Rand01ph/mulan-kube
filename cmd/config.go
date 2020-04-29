/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"github.com/cuisongliu/drone-kube/tools"
	"gitlab.mulan.com/root/mulan-kube/config"

	"github.com/spf13/cobra"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("config called")
		config.Main()
	},
}

func init() {
	rootCmd.AddCommand(configCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// configCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// configCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	kubeServer := tools.Env("KUBE_SERVER", "PLUGIN_SERVER", "PLUGIN_KUBE_SERVER")
	kubeCa := tools.Env("KUBE_CA", "PLUGIN_CA", "PLUGIN_KUBE_CA")
	kubeAdmin := tools.Env("KUBE_ADMIN", "PLUGIN_ADMIN", "PLUGIN_KUBE_ADMIN")
	kubeAdminKey := tools.Env("KUBE_ADMIN_KEY", "PLUGIN_ADMIN_KEY", "PLUGIN_KUBE_ADMIN_KEY")

	configCmd.Flags().StringVarP(&config.KubeServer, "server", "s", kubeServer, "~/.kube/config  server")
	//certificate-authority ca.pem
	configCmd.Flags().StringVarP(&config.KubeCa, "ca", "c", kubeCa, "~/.kube/config certificate-authority-data")
	//client-certificate admin.pem
	configCmd.Flags().StringVarP(&config.KubeAdmin, "admin", "n", kubeAdmin, "~/.kube/config client-certificate-data")
	//client-key admin-key.pem
	configCmd.Flags().StringVarP(&config.KubeAdminKey, "admin-key", "k", kubeAdminKey, "~/.kube/config client-key-data")
}

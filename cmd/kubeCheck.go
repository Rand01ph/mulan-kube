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
	"os"
	"path/filepath"
	"time"

	"github.com/spf13/cobra"
	"gitlab.mulan.com/root/mulan-kube/kubeCheck"
)

// kubeCheckCmd represents the kubeCheck command
var kubeCheckCmd = &cobra.Command{
	Use:   "kubeCheck",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("kubeCheck called")
		kubeCheck.Run()
	},
}

func init() {
	rootCmd.AddCommand(kubeCheckCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// kubeCheckCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// kubeCheckCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	kubeCheckCmd.Flags().DurationVarP(&kubeCheck.Duration, "duration", "d", 10*time.Minute, "Time wait for complate")

	kubeCheckCmd.Flags().StringSliceVarP(&kubeCheck.Labels, "labels", "l", []string{}, "deployment status check label")

	kubeCheckCmd.Flags().StringVarP(&kubeCheck.Namespace, "namespace", "n", "default", "deployment status check namespace")

	if home := homeDir(); home != "" {
		kubeCheckCmd.Flags().StringVarP(&kubeCheck.KubeConfig, "kubeconfig", "k", filepath.Join(home, ".kube", "config"), "Path to kubeconfig")
	} else {
		kubeCheckCmd.Flags().StringVarP(&kubeCheck.KubeConfig, "kubeconfig", "k", "", "Path to kubeconfig")
	}

}

func homeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE") // windows
}

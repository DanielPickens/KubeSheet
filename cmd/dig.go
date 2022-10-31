/*
Copyright © 2022 Daniel M. Pickens

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
	"metav1"

	"github.com/spf13/cobra"
	"k8s.io/client-go/kubernetes"
)

// digCmd represents the dig command
var digCmd = &cobra.Command{
	Use:   "dig",
	Short: "Dig into a resource",
	Long: `Dig into a resource. For example:
	dig pod <pod-name>,

.`,
	Run: func(cmd *cobra.Command, args []string) {
		clientset, err := getClient()
		if err != nil {
			panic(err.Error())
		}
		pods, err := clientset.CoreV1().Pods("").List(metav1.ListOptions{})
		if err != nil {
			panic(err.Error())
		}
		fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))
	},
}

//Using the kubernetes client-go library to list pods in a cluster:

func ListPods(clientset *kubernetes.Clientset) {
	pods, err := clientset.CoreV1().Pods("").List(metav1.ListOptions{}, metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))
}

var (
	kubeconfig string
)

func init() {
	rootCmd.AddCommand(digCmd)

}

var ActivityLengthCmd = &cobra.Command{
	Use:   "ActivityLength",
	Short: "Checks the time a pod has been in use",
	Long: `Checks ammount of time a pod has been in use. For example:
	activitylength <pod-name>,

.`,
	Run: func(cmd *cobra.Command, args []string) {
		clientset, err := getClient()
		if err != nil {
			panic(err.Error())
		}
		pods, err := clientset.CoreV1().Pods("").List(metav1.ListOptions{})
		if err != nil {
			panic(err.Error())
		}
		fmt.Printf("There are %d pods actively in use in this cluster\n", len(pods.Items))
	},
}

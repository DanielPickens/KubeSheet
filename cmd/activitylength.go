package cmd

import (
	"fmt"
	"metav1"

	"github.com/spf13/cobra"
)

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

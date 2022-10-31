package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// activityCmd represents the activity command
var activityCmd = &cobra.Command{
	Use:   "activity",
	Short: "Activity of a resource",
	Long: `Activity of a resource. For example:
	activity pod <pod-name>,
	activity node <node-name>,
	activity namespace <namespace-name>,
	activity deployment <deployment-name>,
	activity service <service-name>,
	activity ingress <ingress-name>,
	activity configmap <configmap-name>,
	activity secret <secret-name>,
	activity daemonset <daemonset-name>,
	activity job <job-name>,
	activity cronjob <cronjob-name>,
	activity statefulset <statefulset-name>,
	
.`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("activity called")
	},
}

func init() {
	rootCmd.AddCommand(activityCmd)
}

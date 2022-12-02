package cmd

import (
	"fmt"
	"metav1"
	"context"
	"time"
	cobra "github.com/spf13/cobra"
)

var IngressCmd = &cobra.Command{
	Use:   "ingress",
	Short: "Ingress of a resource",
	Long: `Determines Ingress of a resource. For example:
	ingress pod <pod-name>`

	Run: func(cmd *cobra.Command, args []string) {
		pods, err := getPods()
		if err != nil {
			panic(err.Error())
		}
		fmt.Printf("There are %d pods in the cluster/n", len(pods.Items))
		for _, pod := range pods.Items {
		fmt.Printf("Pod %s is running on node %s) /n", pod.Name, pod.Spec.NodeName)
		}
	},
}


func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(IngressCmd)
}

var (
	kubeconfig string
	
)
func getClient() (*kubernetes.Clientset, error) {
	config, err := rest.InClusterConfig()
	if err != nil {
		return nil, err
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}
	return clientset, nil
}

func getPods() (*v1.PodList, error) {
	clientset, err := getClient()
	if err != nil {
		return nil, err
	}
	pods, err := clientset.CoreV1().Pods("").List(metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	return pods, nil
}

func getPodsByNode() (map[string][]v1.Pod, error) {
	pods, err := getPods()
	if err != nil {
		return nil, err
	}
	podsByNode := make(map[string][]v1.Pod)
	for _, pod := range pods.Items {
		podsByNode[pod.Spec.NodeName] = append(podsByNode[pod.Spec.NodeName], pod)
	}
	return podsByNode, nil
}

func getPodsByNamespace() (map[string][]v1.Pod, error) {
	pods, err := getPods()
	if err != nil {
		return nil, err
	}
	podsByNamespace := make(map[string][]v1.Pod)
	for _, pod := range pods.Items {
		podsByNamespace[pod.Namespace] = append(podsByNamespace[pod.Namespace], pod)
	}
	return podsByNamespace, nil
}

func getPodsByDeployment() (map[string][]v1.Pod, error) {
	pods, err := getPods()
	if err != nil {
		return nil, err
	}
	podsByDeployment := make(map[string][]v1.Pod)
	for _, pod := range pods.Items {
		podsByDeployment[pod.Labels["app"]] = append(podsByDeployment[pod.Labels["app"]], pod)
	}
	return podsByDeployment, nil
}

func getPodsByService() (map[string][]v1.Pod, error) {
	pods, err := getPods()
	if err != nil {
		return nil, err
	}
	podsByService := make(map[string][]v1.Pod)
	for _, pod := range pods.Items {
		podsByService[pod.Labels["app"]] = append(podsByService[pod.Labels["app"]], pod)
	}
	return podsByService, nil
}

func getPodsByIngress() (map[string][]v1.Pod, error) {
	pods, err := getPods()
	if err != nil {
		return nil, err
	}
	podsByIngress := make(map[string][]v1.Pod)
	for _, pod := range pods.Items {
		podsByIngress[pod.Labels["app"]] = append(podsByIngress[pod.Labels["app"]], pod)
	}
	return podsByIngress, nil
}
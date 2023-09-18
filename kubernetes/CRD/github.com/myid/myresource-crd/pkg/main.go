package main

import (
	"context"
	"fmt"
	clientset2 "github.com/myid/myresource-crd/pkg/clientset/clientset"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/clientcmd"
)

// Demo using the clientset
func main() {
	config, err := clientcmd.BuildConfigFromFlags("", "/Users/hamzamasood/.kube")

	if err != nil {
		fmt.Println("Couldn't get config", err)
	}
	clientset, err := clientset2.NewForConfig(config)

	if err != nil {
		fmt.Println("Couldn't get config", err)
	}
	list, err := clientset.MygroupV1alpha1().MyResources("default").List(context.Background(), metav1.ListOptions{})
	if err != nil {
		fmt.Println("Couldn't get config", err)
	}

	for _, res := range list.Items {
		fmt.Printf("%s\n", res.GetName())
	}
}

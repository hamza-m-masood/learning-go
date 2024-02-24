package main

import (
	"context"
	"fmt"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func RunClient() {
	//create client
	config, err := clientcmd.BuildConfigFromFlags("", "/Users/hamzamasood/.kube/config")
	if err != nil {
		fmt.Println("Could not get kubeconfig", err)
	}
	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		fmt.Println("Could not generate client", err)
	}
	pod := corev1.Pod{
		Spec: corev1.PodSpec{
			Containers: []corev1.Container{
				{
					Name:  "nginx",
					Image: "nginx",
				},
			},
		},
	}
	pod.SetName("nginx")
	//create pod from client
	_, err = clientSet.CoreV1().Pods("default").Create(context.TODO(), &pod, metav1.CreateOptions{})
	if err != nil {
		fmt.Println("Could not create pod", err)
	}

}

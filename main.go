package main

import (
	"context"
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"l8s.io/client-go/tools/clientcmd"
)

func main(){
	config, err := clientcmd.BuildConfigFromFlags("",clientcmd.RecommendedHomeFile)
	if err != nil {
		panic(err)
	}

	clientset,err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	pods,err := clientset.CoreV1().Pods("default").List(context.Background(),metav1.ListOptions{})
	if err != nil{
		panic(err)
	}

	for _,pod := range pods.Items{
		fmt.Println(pod.Name)
	}
}

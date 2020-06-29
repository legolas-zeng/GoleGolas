//Author: zwa
//Date: 2020/6/24

package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"

	appsv1 "k8s.io/api/apps/v1"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

//k8s.io/apimachinery v0.18.2
//k8s.io/client-go v0.18.2
//k8s.io/api v0.18.2

func printTime(msg string) {
	fmt.Println(msg, time.Now().Format("15:04:05"))
}

func createDeployment(clientset *kubernetes.Clientset, ns string, name string) {
	deploymentsClient := clientset.AppsV1().Deployments(ns)
	deployment := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: int32Ptr(2),
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"application": name,
				},
			},
			Template: apiv1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"application": name,
					},
				},
				Spec: apiv1.PodSpec{
					Containers: []apiv1.Container{
						{
							Name:  name,
							Image: "nginx:1.18",
							Ports: []apiv1.ContainerPort{
								{
									Name:          "http",
									Protocol:      apiv1.ProtocolTCP,
									ContainerPort: 80,
								},
							},
						},
					},
				},
			},
		},
	}

	result, err := deploymentsClient.Create(context.TODO(), deployment, metav1.CreateOptions{})

	if err != nil {
		panic(err)
	}

	printTime("Create deployment ---->>>")
	printTime(result.GetName())

}

func deleteDeployment(clientset *kubernetes.Clientset, ns string, name string) {

	deletePolicy := metav1.DeletePropagationForeground

	deploymentsClient := clientset.AppsV1().Deployments(ns)

	if err := deploymentsClient.Delete(context.TODO(), name, metav1.DeleteOptions{PropagationPolicy: &deletePolicy}); err != nil {
		panic(err)
	}

	printTime("Delete deployment ---->>>")
	printTime(name)

}

func runForever(wg *sync.WaitGroup, clientset *kubernetes.Clientset, name string, timerun int, ns string) {

	printTime("runForever() start")

	for {

		fmt.Println("Create deployment", name)

		createDeployment(clientset, ns, name)
		time.Sleep(time.Duration(timerun) * time.Minute)

		fmt.Println("Delete deployment", name)

		deleteDeployment(clientset, ns, name)
		time.Sleep(5 * time.Minute)

	}
	wg.Done()
}

func main() {

	printTime("main() started")

	//homePath := "/home/test"
	var wg sync.WaitGroup
	namePod := map[string]int{"vo-deploy": 15, "qa-deploy": 10}
	nameSpace := "default"

	config, err := rest.InClusterConfig()
	if err != nil {
		kubeconfig := filepath.Join("D:\\GoleGolas\\kube\\KubeCommand", "admin.conf")
		if envvar := os.Getenv("KUBECONFIG"); len(envvar) > 0 {
			kubeconfig = envvar
		}
		config, err = clientcmd.BuildConfigFromFlags("", kubeconfig)
		if err != nil {
			fmt.Printf("The kubeconfig was not  loaded: %v\n", err)
			os.Exit(1)
		}
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	i := 0
	for k, v := range namePod {
		i++

		wg.Add(i)
		go runForever(&wg, clientset, k, v, nameSpace)

	}
	wg.Wait()
	printTime("The goroutine completed ---->>>")
}

func int32Ptr(i int32) *int32 {
	return &i
}

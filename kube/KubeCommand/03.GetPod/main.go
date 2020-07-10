//Author: zwa
//Date: 2020/6/24
package main

import (
	"context"
	"flag"
	"fmt"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	"k8s.io/client-go/tools/clientcmd"
	"log"
	"os"
	"sync"
	//	"k8s.io/apimachinery/pkg/watch"
	"path/filepath"
)

//k8s.io/apimachinery v0.17.0
//k8s.io/client-go v0.17.0

type Clientset struct {
	*kubernetes.Clientset
}

type Deployer struct {
	ctx        context.Context
	clients    sync.Map
	configPath string
}

func GetPodList(clientset *kubernetes.Clientset, nsname string) {
	// 获取pod列表 pod为名称空间级别资源需指定名称空间
	pods, err := clientset.CoreV1().Pods(nsname).List(metav1.ListOptions{})
	if err != nil {
		panic(err)
	}
	// 循环打印pod的信息
	for _, pod := range pods.Items {
		fmt.Println(pod.ObjectMeta.Name, pod.Status.Phase)
	}
}

func GetPodInfo(clientset *kubernetes.Clientset, podname string) {
	pod, err := clientset.CoreV1().Pods("monitoring").Get(podname, metav1.GetOptions{})
	if err != nil {
		panic(err)
	}
	fmt.Println(pod)
}

func (clientset *Clientset) CreatePods(namespace string, pod corev1.Pod) (err error) {
	_, err = clientset.CoreV1().Pods(namespace).Create(&pod)
	return
}

func DeletePod(clientset *kubernetes.Clientset, podname string) {
	err := clientset.CoreV1().Pods("default").Delete(podname, &metav1.DeleteOptions{})
	if err != nil {
		panic(err)
	}
}

func main() {
	var ns, label, field string
	//	kubeconfig := os.Getenv("KUBECONFIG")
	kubeconfig := filepath.Join(os.Getenv("HOME"), "D:\\GoleGolas\\kube\\KubeCommand", "admin.conf")
	fmt.Println(kubeconfig)
	flag.StringVar(&ns, "namespace", "", "provide the namespace")
	flag.StringVar(&field, "f", "", "Field selector")
	flag.StringVar(&label, "l", "", "give the label")
	flag.StringVar(&kubeconfig, "kubeconfig", kubeconfig, "kubeconfig file")
	flag.Parse()
	// bootstrap config
	fmt.Println("Using kubeconfig:", kubeconfig)
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	GetPodList(clientset, "monitoring")
	GetPodInfo(clientset, "alertmanager-rrjc-prometheus-prometheus-alertmanager-0")

	if err != nil {
		log.Fatal(err)
	}
	api := clientset.CoreV1()

	listOptions := metav1.ListOptions{LabelSelector: label, FieldSelector: field}
	//listOptions := metav1.ListOptions{LabelSelector: "", FieldSelector: ""}
	pd, err := api.Pods(ns).List(listOptions)
	if err != nil {
		log.Fatal(err)
	}
	for _, pod := range pd.Items {
		fmt.Fprintf(os.Stdout, "pod name: %v\n", pod.Name)
	}
	//	fmt.Println(pd)

}

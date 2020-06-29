//Author: zwa
//Date: 2020/6/28

package main

import (
	"GoleGolas/kube/KubeCommand/common"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

//k8s.io/apimachinery v0.17.0
//k8s.io/client-go v0.17.0
//k8s.io/api v0.17.0

func GetNodeList(clientset *kubernetes.Clientset) {
	nodes, err := clientset.CoreV1().Nodes().List(metav1.ListOptions{})
	if err != nil {
		panic(err)
	}
	for _, nds := range nodes.Items {
		fmt.Printf("NodeName: %s\n", nds.Name)
	}
}

func GetNode(clientset *kubernetes.Clientset, nodename string) {
	//获取 指定NODE 的详细信息
	fmt.Println("\n ####### node详细信息 ######")
	nodeRel, err := clientset.CoreV1().Nodes().Get(nodename, metav1.GetOptions{})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Name: %s \n", nodeRel.Name)
	fmt.Printf("CreateTime: %s \n", nodeRel.CreationTimestamp)
	fmt.Printf("NowTime: %s \n", nodeRel.Status.Conditions[0].LastHeartbeatTime)
	fmt.Printf("kernelVersion: %s \n", nodeRel.Status.NodeInfo.KernelVersion)
	fmt.Printf("SystemOs: %s \n", nodeRel.Status.NodeInfo.OSImage)
	fmt.Printf("Cpu: %s \n", nodeRel.Status.Capacity.Cpu())
	fmt.Printf("docker: %s \n", nodeRel.Status.NodeInfo.ContainerRuntimeVersion)
	// fmt.Printf("Status: %s \n", nodeRel.Status.Conditions[len(nodes.Items[0].Status.Conditions)-1].Type)
	fmt.Printf("Status: %s \n", nodeRel.Status.Conditions[len(nodeRel.Status.Conditions)-1].Type)
	fmt.Printf("Mem: %s \n", nodeRel.Status.Allocatable.Memory().String())
}

func main() {
	clientset := common.ConKube()
	GetNodeList(clientset)
	GetNode(clientset, "192.168.3.12")

}

//Author: zwa
//Date: 2020/6/28

package main

import (
	"GoleGolas/kube/KubeCommand/common"
	"fmt"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func getnamespacelist(clientset *kubernetes.Clientset) {
	// 获取所有 namespace
	namespaceList, err := clientset.CoreV1().Namespaces().List(metav1.ListOptions{})
	if err != nil {
		fmt.Println(err.Error())
	} else {
		for _, nsList := range namespaceList.Items {
			fmt.Printf("NsName: %s \n", nsList.Name)
		}
	}
}

func getnamespace(clientset *kubernetes.Clientset, nsname string) {
	// 获取指定namespace的详细信息
	result, err := clientset.CoreV1().Namespaces().Get(nsname, metav1.GetOptions{})
	if err != nil {
		panic(err)
	}
	fmt.Printf("\n Name: %s \n Status: %s \n CreateTime: %s \n", result.ObjectMeta.Name, result.Status.Phase, result.CreationTimestamp)

}

func createnamesapce(clientset *kubernetes.Clientset, nsname string) {
	// 创建 namesapce
	// 通过实现 clientset 的 CoreV1Interface 接口列表中的 NamespacesGetter 接口方法 Namespaces 返回 NamespaceInterface
	// NamespaceInterface 接口拥有操作 Namespace 资源的方法，例如 Create、Update、Get、List 等方法
	name := nsname
	namespaceClient := clientset.CoreV1().Namespaces()
	namespace := &apiv1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
		},
	}
	// 创建 namespace
	result, err := namespaceClient.Create(namespace)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("Create ns %s SuccessFul !", result.ObjectMeta.Name)
	}

}

func deletenamespace(clientset *kubernetes.Clientset, nsname string) {
	// 删除指定的 namespace
	name := nsname
	deletePolicy := metav1.DeletePropagationForeground
	err := clientset.CoreV1().Namespaces().Delete(name, &metav1.DeleteOptions{
		PropagationPolicy: &deletePolicy,
	})

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("Deleted Namespaces %s\n", name)
	}
}

func main() {
	clientset := common.ConKube()
	getnamespacelist(clientset)
	getnamespace(clientset, "default")

	deletenamespace(clientset, "rrjc")

}

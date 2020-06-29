//Author: zwa
//Date: 2020/6/29

package main

import (
	"GoleGolas/kube/KubeCommand/common"
	"errors"
	"fmt"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

//k8s.io/apimachinery v0.17.0
//k8s.io/client-go v0.17.0
//k8s.io/api v0.17.0

// GetConfigmap func
func GetConfigmap(clientset *kubernetes.Clientset, namespace string) ([]corev1.ConfigMap, error) {
	configmaps, err := clientset.CoreV1().ConfigMaps(namespace).List(metav1.ListOptions{})
	if err != nil {
		return []corev1.ConfigMap{}, err
	}
	return configmaps.Items, nil
}

// GetConfigmapByName func
func GetConfigmapByName(clientset *kubernetes.Clientset, namespace string, labelSelector string, name string) (corev1.ConfigMap, error) {
	configmaps, err := clientset.CoreV1().ConfigMaps(namespace).List(metav1.ListOptions{LabelSelector: labelSelector})
	var configmap corev1.ConfigMap
	if err != nil {
		return configmap, err
	}
	for _, Item := range configmaps.Items {
		if Item.Name == name {
			configmap = Item
		}
	}
	if configmap.Name == "" {
		return configmap, errors.New("not match")
	}
	return configmap, nil
}

// CreateConfigmap func
func CreateConfigmap(clientset *kubernetes.Clientset, namespace string, configmap corev1.ConfigMap) (err error) {
	_, err = clientset.CoreV1().ConfigMaps(namespace).Create(&configmap)
	return
}

// UpdateConfigmap func
func UpdateConfigmap(clientset *kubernetes.Clientset, namespace string, configmap corev1.ConfigMap) (err error) {
	_, err = clientset.CoreV1().ConfigMaps(namespace).Update(&configmap)
	return
}

// DeleteConfigmap func
func DeleteConfigmap(clientset *kubernetes.Clientset, namespace string, name string) (err error) {
	err = clientset.CoreV1().ConfigMaps(namespace).Delete(name, &metav1.DeleteOptions{})
	return
}

func main() {
	clientset := common.ConKube()
	fmt.Println(GetConfigmap(clientset, "default"))
}

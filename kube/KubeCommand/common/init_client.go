//Author: zwa
//Date: 2020/6/22

package common

import (
	"flag"
	"fmt"
	"io/ioutil"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"os"
	"path/filepath"
)

//k8s.io/apimachinery v0.17.0
//k8s.io/client-go v0.17.0
//k8s.io/api v0.17.0

// 初始化k8s客户端
func InitClient() (clientset *kubernetes.Clientset, err error) {
	var (
		restConf *rest.Config
	)

	if restConf, err = GetRestConf(); err != nil {
		return
	}

	// 生成clientset配置
	if clientset, err = kubernetes.NewForConfig(restConf); err != nil {
		goto END
	}
END:
	return
}

func ConKube() (clientset *kubernetes.Clientset) {
	var kubeconfig *string
	kubeconfig = flag.String("kubeconfig", filepath.Join(os.Getenv("HOME"), "D:\\GoleGolas\\kube\\KubeCommand", "admin.conf"), "(optional) absolute path to the kubeconfig file")
	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		fmt.Println(err.Error())
	}

	clientset, err = kubernetes.NewForConfig(config)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("Connect kubernetes Successful！！！")
	return clientset
}

// 获取k8s restful client配置
func GetRestConf() (restConf *rest.Config, err error) {
	var (
		kubeconfig []byte
	)

	// 读kubeconfig文件
	if kubeconfig, err = ioutil.ReadFile("D:\\GoleGolas\\kube\\KubeCommand\\admin.conf"); err != nil {
		goto END
	}
	// 生成rest client配置
	if restConf, err = clientcmd.RESTConfigFromKubeConfig(kubeconfig); err != nil {
		goto END
	}
END:
	return
}

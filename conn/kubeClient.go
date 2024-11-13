package conn

import (
	beego "github.com/beego/beego/v2/server/web"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"log"
)

/*
获取客户端
*/
func CreateClient() (*kubernetes.Clientset, error) {
	configPath, _ := beego.AppConfig.String("configPath")
	config, err := clientcmd.BuildConfigFromFlags("",configPath)
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
	}
	return clientset,nil
}

//  封装 dynamicclient 客户端
func ForDynamicClient() dynamic.Interface {
	configPath, _ := beego.AppConfig.String("configPath")
	config, err := clientcmd.BuildConfigFromFlags("",configPath)
	dynamicClient, err := dynamic.NewForConfig(config)
	if err != nil {
		log.Fatalln(err)
	}
	return dynamicClient
}

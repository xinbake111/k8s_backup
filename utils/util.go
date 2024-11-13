package utils

import (
	"WorkPro/models"
	"context"
	"io/ioutil"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"log"
	"net/http"
	"os"
)

/*判断文件或文件夹是否存在*/
func isExist(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		if os.IsNotExist(err) {
			return false
		}
		log.Fatalln(err)
		return false
	}
	return true

}

/*
 读取文件 记录文件名称
*/
func Read(filepath string) (filenames []string) {
	var s []string
	fileList, err := ioutil.ReadDir(filepath)
	if err != nil {
		log.Fatalln(err)
	}
	for _, file := range fileList {
		if !file.IsDir() {
			filename := filepath + "/" + file.Name()
			s = append(s, filename)
		}
	}
	return s
}
/*
	判断该资源是否已经存在
*/
func In(target string,str_array []string) bool  {
	for _,element :=range str_array {
		if target==element {
			return true
		}
	}
	return false
}
/*
	获取get
*/
func GetOpt(url string) string{
	resp, err := http.Get(url)
	if err != nil {
		panic(err)

	}
	defer resp.Body.Close()
	s,err:=ioutil.ReadAll(resp.Body)
	return string(s)
}
/*
  获取节点资源mem cpu
*/
func GetNodeRes()  {
	mod,err:=models.GetClusterList()
	if err!=nil {
		log.Println(err)
	}
	for _, n := range mod.NodeList.Items {
		nodename:=n.GetName()
		nodeRel,err:=mod.ClientSet.CoreV1().Nodes().Get(context.TODO(),nodename,metav1.GetOptions{})
		if err != nil {
			log.Println(err)
		}
		nodeRel.Status.Allocatable.Memory().String()
		nodeRel.Status.Capacity.Cpu()
	}
}



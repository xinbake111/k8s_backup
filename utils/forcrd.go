package utils

import (
	"WorkPro/conn"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	crdv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"log"
	"os"
	"sync"
	"time"
)

type crds struct {
	Crdgroup    string
	Crdvresion  []string
	Crdresource string
}
var mutex sync.Mutex

//整体备份crd资源
func BackupAllCrd(namespace string)  {
	BackupCrd()
	BackupCrdResources(namespace)
}
//整体备份crdresources资源
func RestoreAllCrd(namespace string){
	go RestoreCrd()
	time.Sleep(time.Duration(2) * time.Second)
	RestoreCrdResource(namespace)
}
/*备份crd流程*/
func BackupCrd()  {
	GetClusterCRD()
}
func BackupCrdResources(namespace string)  {
	GetCrdResources(namespace)
}
/*恢复crd流程*/
func RestoreCrd() {
	mutex.Lock()
	RestoreClusterCRD()
	mutex.Unlock()
}
func RestoreCrdResource(namespace string) {
	ReadCrdResource(namespace)
}

//查询集群crd的资源
func ForCRD() ([]crdv1.CustomResourceDefinition, error) {
	dynamicClient:= conn.ForDynamicClient()
	//查询集群中创建的crd资源
	crdd := schema.GroupVersionResource{Group: "apiextensions.k8s.io", Version: "v1", Resource: "customresourcedefinitions"}
	crdlist, err := dynamicClient.Resource(crdd).List(context.TODO(), metav1.ListOptions{})

	if err != nil {
		log.Fatalln(err)
	}
	//var forcrd []crds
	var crdstru []crdv1.CustomResourceDefinition
	for _, crd := range crdlist.Items {
		// var crdstru crdv1.CustomResourceDefinition
		//对无序列化数据进行序列化
		var stru crdv1.CustomResourceDefinition
		runtime.DefaultUnstructuredConverter.FromUnstructured(crd.UnstructuredContent(), &stru)
		crdstru = append(crdstru, stru)
	}
	return crdstru, err
}

//对gvr进行填充
func ForGVR() []crds {
	dynamicClient := conn.ForDynamicClient()
	//查询集群中创建的crd资源
	crdd := schema.GroupVersionResource{Group: "apiextensions.k8s.io", Version: "v1", Resource: "customresourcedefinitions"}
	crdlist, err := dynamicClient.Resource(crdd).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Fatalln(err)
	}
	var forcrd []crds
	for _, crd := range crdlist.Items {
		// var crdstru crdv1.CustomResourceDefinition
		//对无序列化数据进行序列化
		var stru crdv1.CustomResourceDefinition
		runtime.DefaultUnstructuredConverter.FromUnstructured(crd.UnstructuredContent(), &stru)
		var crdversion []string
		for _, value := range stru.Spec.Versions {
			v := value.Name
			crdversion = append(crdversion, v)
		}
		//对集群中所有的crd进行一个存储
		forcrd = append(forcrd, crds{
			Crdgroup:    stru.Spec.Group,
			Crdresource: stru.Spec.Names.Plural,
			Crdvresion:  crdversion,
		})
	}
	return forcrd
}

//获取集群中的crd
func GetClusterCRD() {
	crdstru, err := ForCRD()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("____________写入ClusterCRD开始_____________")
	//fileMkdir := "./b_clustercrd"  //需要进行该

	fileMkdir :="./BackupFile/b_clustercrdlist"

	//判断crd文件夹是否存在
	if !isExist(fileMkdir) {
		//如果不存在进行创建
		os.MkdirAll(fileMkdir, os.ModePerm)
	}
	//filepath := fileMkdir + "/" + "clusterCrd.txt"
	filepath := fileMkdir + "/clustercrd.txt"
	bytes, err := json.Marshal(crdstru)
	if err != nil {
		log.Fatalln(err)
	}
	err1 := ioutil.WriteFile(filepath, bytes, 0666)
	if err1 != nil {
		log.Fatalln(err1)
	}
	fmt.Println("____________写入ClusterCRD结束_____________")

}

//从已经备份的crd文件中信息恢复到集群中去
func RestoreClusterCRD() {

	filepath := "./BackupFile/b_clustercrdlist"
	fmt.Println("____________恢复ClusterCRD开始_____________")
	filenames := Read(filepath)
	for i := 0; i < len(filenames); i++ {
		var crdlist []crdv1.CustomResourceDefinition
		bytes, err := ioutil.ReadFile(filenames[i])
		if err != nil {
			log.Fatalln(err)
		}
		err1 := json.Unmarshal(bytes, &crdlist)
		if err1 != nil {
			log.Fatalln(err1)
		}
		for _, crd := range crdlist {
			//查询集群中是否有crd
			names := CrdList()
			if !In(crd.Name, names) {
				crd.ObjectMeta.ResourceVersion = " "
				dynamicClient := conn.ForDynamicClient()
				//反序列化
				// obj ->  []byte
				bytedata, err := json.Marshal(&crd)
				if err != nil {
					log.Fatalln(err)
				}

				// []byte -> Unstructured
				utd := &unstructured.Unstructured{}
				err2 := json.Unmarshal(bytedata, &utd.Object)
				if err2 != nil {
					log.Fatalln(err2)
				}

				// utd, err := runtime.UnstructuredConverter.ToUnstructured(crd)
				//查询集群中创建的crd资源
				crdd := schema.GroupVersionResource{Group: "apiextensions.k8s.io", Version: "v1", Resource: "customresourcedefinitions"}
				dynamicClient.Resource(crdd).Create(context.TODO(), utd, metav1.CreateOptions{})
			}
		}
	}
	fmt.Println("____________恢复ClusterCRD结束_____________")

}

type backPersonCrd struct {
	Personcrd []unstructured.Unstructured
	Crd       []crds
}

/*
需要对namespace下的进行一个判断
和pod相同crd也是存在namespace下的
需要在文件中写入crd的gvr方便恢复的时候进行是否在集群中进行判断
*/
/*
写入crd下的资源
*/
func GetCrdResources(namespace string) {
	dynamicClient := conn.ForDynamicClient()
	forcrd := ForGVR()
	// var bpcrd backPersonCrd
	if namespace == "" {
		namespace = "default"
	}
	var crdResource []unstructured.Unstructured
	for _, crd := range forcrd {
		//只获取最新版本的crdresources
		// for _, crdversion := range crd.Crdvresion {

		//找出每个crd下面定义的资源
		crdd2 := schema.GroupVersionResource{Group: crd.Crdgroup, Version: crd.Crdvresion[len(crd.Crdvresion)-1], Resource: crd.Crdresource}
		crdlist2, err := dynamicClient.Resource(crdd2).Namespace(namespace).List(context.TODO(), metav1.ListOptions{})
		if err != nil {
			log.Fatalln(err)
		}
		for _, c := range crdlist2.Items {
			fmt.Println(c)
			crdResource = append(crdResource, c)
		}
		WriteCrdResources(crd, namespace, crdResource)
	}
}

//go语言中小写为私有属性，不可以被json访问到
type backupCrdResoures struct {
	CrdResources []unstructured.Unstructured
	Crd          crds
}

func WriteCrdResources(crd crds, namespace string, crdResources []unstructured.Unstructured) {
	spaceResource :="crdresources"
	if namespace == "" {
		namespace = "default"
	}
	var CResource backupCrdResoures
	CResource.Crd = crd
	CResource.CrdResources = crdResources
	fmt.Println("____________写入crdResources开始_____________")
	fileMkdir :="./BackupFile/"+namespace+"/b_"+spaceResource+"list"
	//fileMkdir := "./b_crdresources"
	if !isExist(fileMkdir) {
		os.MkdirAll(fileMkdir, os.ModePerm)
	}
	filepath := fileMkdir + "/" + namespace+"_"+spaceResource+".txt"
	//filepath := fileMkdir + "/" + namespace + "_CrdResources.txt"
	bytes, err := json.Marshal(CResource)
	if err != nil {
		log.Fatalln(err)
	}
	err1 := ioutil.WriteFile(filepath, bytes, 0666)
	if err1 != nil {
		log.Fatalln(err1)
	}
	fmt.Println("____________写入crdResources结束_____________")
}

/*恢复自定义crd类下的资源*/
func ReadCrdResource(namespace string) {
	spaceResource :="crdresources"
	dynamicClient := conn.ForDynamicClient()
	fmt.Println("____________恢复CrdResource开始_____________")
	fileMkdir :="./BackupFile/"+namespace+"/b_"+spaceResource+"list"
	//filepath := "./b_crdresources"
	file, err := ioutil.ReadDir(fileMkdir)
	var s []string
	if err != nil {
		log.Fatalln(err)
	}
	for _, fi := range file {
		if !fi.IsDir() {
			filename := fileMkdir + "/" + fi.Name()
			s = append(s, filename)
		}
	}
	for i := 0; i < len(s); i++ {
		var bpcrdResource backupCrdResoures
		bytes, err := ioutil.ReadFile(s[i])
		if err != nil {
			log.Fatalln(err)
		}
		err1 := json.Unmarshal(bytes, &bpcrdResource)
		if err1 != nil {
			log.Fatalln(err)
		}
		for _, value := range bpcrdResource.CrdResources {
			fmt.Println(value)

			// 需要对集群内的crd进行一次判断避免已经创建

			//进行创建个人的crd
			//需要对personcrd的资源进行序列化
			names := ResourcePersonCrd(dynamicClient, bpcrdResource, "default")
			if !In(value.GetName(), names) {
				// value.ObjectMeta.ResourceVersion = " "
				// value.Spec.ClusterIPs = nil
				value.SetResourceVersion(" ")
				//创建personcrd 根据特定的crd资源进行创建
				crdd2 := schema.GroupVersionResource{Group: bpcrdResource.Crd.Crdgroup, Version: bpcrdResource.Crd.Crdvresion[len(bpcrdResource.Crd.Crdvresion)-1], Resource: bpcrdResource.Crd.Crdresource}
				crdResource, err := dynamicClient.Resource(crdd2).Namespace(value.GetNamespace()).Create(context.TODO(), &value, metav1.CreateOptions{})

				if err != nil {
					log.Fatalln(err)
				}
				fmt.Printf("%s  %s,\n", crdResource.GetName(), crdResource.GetNamespace())
			}

		}
	}
	fmt.Println("____________恢复CrdResource结束_____________")
}

/*获取集群中的自定义资源*/
func CrdList( ) (nameList []string) {
	dynamicClient := conn.ForDynamicClient()
	//查询集群中创建的crd资源
	crdd := schema.GroupVersionResource{Group: "apiextensions.k8s.io", Version: "v1", Resource: "customresourcedefinitions"}
	crdlist, err := dynamicClient.Resource(crdd).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Println(err)
	}
	var names []string
	for _, crd := range crdlist.Items {
		names = append(names, crd.GetName())
	}
	return names
}

// 查询当前集群中的自定义crd
func ForPersonCrd(dynamicClient dynamic.Interface, forcrd []crds, namespace string) (nameList []string) {
	// //查询集群中的crdd资源
	// crdd := schema.GroupVersionResource{Group: , Version: "v1", Resource: "customresourcedefinitions"}
	// crdlist, err := dynamicClient.Resource(crdd).List(context.TODO(), metav1.ListOptions{})
	// //查询集群中定义的crd
	// pcrdd := schema.GroupVersionResource{}
	//查询集群中的crd资源
	var names []string
	for _, crd := range forcrd {
		for _, crdversion := range crd.Crdvresion {
			//找出每个crd下面定义的资源
			crdd2 := schema.GroupVersionResource{Group: crd.Crdgroup, Version: crdversion, Resource: crd.Crdresource}
			crdlist, err := dynamicClient.Resource(crdd2).Namespace(namespace).List(context.TODO(), metav1.ListOptions{})
			if err != nil {
				log.Fatalln(err)
			}
			for _, pcrdname := range crdlist.Items {
				names = append(names, pcrdname.GetName())
			}
		}
	}
	return names
}

//查询当前集群用户自定义的crd所创造的资源
func ResourcePersonCrd(dynamicClient dynamic.Interface, bpcrd backupCrdResoures, namespace string) (nameList []string) {
	var names []string
	crd := schema.GroupVersionResource{Group: bpcrd.Crd.Crdgroup, Version: bpcrd.Crd.Crdvresion[len(bpcrd.Crd.Crdvresion)-1], Resource: bpcrd.Crd.Crdresource}
	crdList, err := dynamicClient.Resource(crd).Namespace(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Println(err)
	}
	for _, crdResource := range crdList.Items {
		names = append(names, crdResource.GetName())
	}
	return names

}



package utils

import (
	"WorkPro/models"
	"context"
	"encoding/json"
	"fmt"
	appv1 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"log"
)



/*
	通过恢复deployment
	deplyoment恢复方法与恢复pod类似
	但是由于deployment对pod可进行创建 所以要先进行deployment的恢复 如果deployment已经恢复则pod恢复后会自动死掉
*/
func RestoreSpaceDeploy(namespace string)  {
	//先进行判断判断当前集权中是否有该namespace
	get,err:=models.GetClusterList()
	if err!=nil {
		log.Println(err)
	}
	namespaceList :=get.NamespaceList
	var items []string
	for _,i :=range namespaceList.Items{
		items=append(items, i.GetName())
	}
	if !In(namespace,items) {
		//如果不存在
		log.Println("namespace不存在，请先进行namespace的恢复")
		return
	}else {
		//进行pod恢复
		spacelist,err:=models.GetSpaceList(namespace)
		if err!=nil {
			log.Println(err)
		}
		deploylist :=spacelist.DeploymentList //已经存在的pod
		var deployItems []string
		for _,i :=range deploylist.Items{
			deployItems=append(deployItems, i.GetName())
		}
		//读取pod pod存储是按文件夹进行存储所以 进行文件夹读取
		fileDir :="./BackupFile/"+namespace+"/b_deploylist"
		list :=RestoreOutPutFile(fileDir)

		resources :=list[0]
		//转化结构体
		//转为bytes 转为结构体
		bytes,err:=json.Marshal(resources)
		var res appv1.DeploymentList
		json.Unmarshal(bytes,&res)
		fmt.Println(res)
		for _,deploy:=range res.Items   {
			if !In(deploy.Name,deployItems) {
				deploy.ObjectMeta.ResourceVersion = " "
				d,err:=spacelist.ClientSet.AppsV1().Deployments(namespace).Create(context.TODO(),&deploy,metav1.CreateOptions{})
				if err != nil {
					log.Println(err)
				}
				fmt.Printf("restore Pod%s \n", d.GetName())
				ReadPvcForName(d.Name,namespace,d.Spec.Template.Spec.Volumes)
			}
		}
	}
}
/*
	通过namespace恢复下列的pod
	对pod进行一个统一的恢复 如果是deploy 或 job等创建的pod会自己死亡
*/
func RestoreSpacePod(namespace string)  {
	//先进行判断判断当前集权中是否有该namespace
	get,err:=models.GetClusterList()
	if err!=nil {
		log.Fatalln(err)
	}
	namespaceList :=get.NamespaceList
	var items []string
	for _,i :=range namespaceList.Items{
		items=append(items, i.GetName())
	}
	if !In(namespace,items) {
		//如果不存在
		log.Fatalln("namespace不存在，请先进行namespace的恢复")
		return
	}else {
		//进行pod恢复
		spacelist,err:=models.GetSpaceList(namespace)
		if err!=nil {
			log.Fatalln(err)
		}
		podlist :=spacelist.PodList //已经存在的pod
		var podItems []string
		for _,i :=range podlist.Items{
			podItems=append(podItems, i.GetName())
		}
		//读取pod pod存储是按文件夹进行存储所以 进行文件夹读取
		fileDir :="./BackupFile/"+namespace+"/b_podlist"
		list :=RestoreOutPutFile(fileDir)

		resources :=list[0]
		//转化结构体
		//转为bytes 转为结构体
		bytes,err:=json.Marshal(resources)
		var res v1.PodList
		json.Unmarshal(bytes,&res)

		for _,pod:=range res.Items   {
			if !In(pod.Name,podItems) {
				pod.ObjectMeta.ResourceVersion = " "
				po,err:=spacelist.ClientSet.CoreV1().Pods(namespace).Create(context.TODO(),&pod,metav1.CreateOptions{})
				if err != nil {
					log.Fatalln(err)
				}
				fmt.Printf("restore Pod%s \n", po.GetName())
				ReadPvcForName(pod.Name,namespace,pod.Spec.Volumes)
			}
		}
	}
}
func RestoreSpaceService(namespace string)  {
	//先进行判断判断当前集权中是否有该namespace
	get,err:=models.GetClusterList()
	if err!=nil {
		log.Fatalln(err)
	}
	namespaceList :=get.NamespaceList
	var items []string
	for _,i :=range namespaceList.Items{
		items=append(items, i.GetName())
	}
	if !In(namespace,items) {
		//如果不存在
		log.Fatalln("namespace不存在，请先进行namespace的恢复")
		return
	}else {
		//进行pod恢复
		spacelist,err:=models.GetSpaceList(namespace)
		if err!=nil {
			log.Fatalln(err)
		}
		svclist :=spacelist.ServiceList //已经存在的svc
		var svcItems []string
		for _,i :=range svclist.Items{
			svcItems=append(svcItems, i.GetName())
		}
		//读取pod pod存储是按文件夹进行存储所以 进行文件夹读取
		fileDir :="./BackupFile/"+namespace+"/b_servicelist"
		list :=RestoreOutPutFile(fileDir)

		resources :=list[0]
		//转化结构体
		//转为bytes 转为结构体
		bytes,err:=json.Marshal(resources)
		var res v1.ServiceList
		json.Unmarshal(bytes,&res)

		for _,svc:=range res.Items   {
			if !In(svc.Name,svcItems) {
				svc.ObjectMeta.ResourceVersion = " "
				sc,err:=spacelist.ClientSet.CoreV1().Services(namespace).Create(context.TODO(),&svc,metav1.CreateOptions{})
				if err != nil {
					log.Fatalln(err)
				}
				fmt.Printf("restore Pod%s \n", sc.GetName())
			}
		}
	}
}

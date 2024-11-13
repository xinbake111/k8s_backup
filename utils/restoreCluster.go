package utils

import (
	"WorkPro/models"
	"context"
	"encoding/json"
	"fmt"
	v1 "k8s.io/api/core/v1"
	rbacv1 "k8s.io/api/rbac/v1"
	storagev1 "k8s.io/api/storage/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"log"
)
/*
	读取数据文件
*/
func RestoreOutPutFile(fileDir string)([]interface{})  {
	return RestoreECCFileOutPutFile(fileDir)
}

/*
	`读取namespace数据文件
*/
func RestoreNameSpace(clusterResource *string) {
	clusterList,err :=models.GetClusterList()
	if err!=nil {
		log.Fatalln(err)
	}
	namespaceList :=clusterList.NamespaceList
	var nameItems []string
	for _,namespace :=range namespaceList.Items {
		nameItems = append(nameItems, namespace.GetName())
	}
	fileDir :="./BackupFile/b_"+*clusterResource+"list"
	//获取备份数据
	list :=RestoreOutPutFile(fileDir)
	resources :=list[0]
	//转化结构体
	//转为bytes 转为结构体
	bytes,err:=json.Marshal(resources)
	var res v1.NamespaceList
	json.Unmarshal(bytes,&res)


	for _,namespace:=range res.Items   {
		if !In(namespace.Name,nameItems) {
			ns,err:=clusterList.ClientSet.CoreV1().Namespaces().Create(context.TODO(), &namespace, metav1.CreateOptions{})
			if err != nil {
				log.Fatalln(err)
			}
			fmt.Printf("restore namespace %s \n", ns.GetName())
		}

	}
}
/*
	恢复 storageclass 资源
*/
func RestoreStorageClass(clusterResource *string)  {
	clusterList,err  :=models.GetClusterList()
	if err!=nil {
		log.Fatalln(err)
	}
	storageClassList :=clusterList.StorageClassList
	var storageClassItems []string
	//获取当前集群中所存在的资源
	for _,storageClass :=range storageClassList.Items {
		storageClassItems = append(storageClassItems, storageClass.GetName())
	}
	fileDir :="./BackupFile/b_"+*clusterResource+"list"
	//获取备份数据
	list :=RestoreOutPutFile(fileDir)
	resources :=list[0]
	//转化结构体
	//转为bytes 转为结构体
	bytes,err:=json.Marshal(resources)
	var res storagev1.StorageClassList
	json.Unmarshal(bytes,&res)


	for _,storageClass:=range res.Items   {
		if !In(storageClass.Name,storageClassItems) {
			storageClass.ObjectMeta.ResourceVersion = " "
			sc,err:=clusterList.ClientSet.StorageV1().StorageClasses().Create(context.TODO(),&storageClass,metav1.CreateOptions{})
			if err != nil {
				log.Fatalln(err)
			}
			fmt.Printf("restore storageclass%s \n", sc.GetName())
		}
	}
}
/*
	恢复clusterRole
*/
func RestoreClusterRole(clusterResource *string)  {
	clusterList,err  :=models.GetClusterList()
	if err!=nil {
		log.Fatalln(err)
	}
	List :=clusterList.ClusterRoleList
	var Items []string
	//获取当前集群中所存在的资源
	for _,i :=range List.Items {
		Items = append(Items, i.GetName())
	}
	fileDir :="./BackupFile/b_"+*clusterResource+"list"
	//获取备份数据
	list :=RestoreOutPutFile(fileDir)
	resources :=list[0]
	//转化结构体
	//转为bytes 转为结构体
	bytes,err:=json.Marshal(resources)
	var res rbacv1.ClusterRoleList
	json.Unmarshal(bytes,&res)

	for _,clusterRole:=range res.Items   {
		if !In(clusterRole.Name,Items) {
			cr,err:=clusterList.ClientSet.RbacV1().ClusterRoles().Create(context.TODO(),&clusterRole,metav1.CreateOptions{})
			if err != nil {
				log.Fatalln(err)
			}
			fmt.Printf("restore clusterRole%s \n", cr.GetName())
		}
	}
}

/*
	恢复 clusterRoleBind
*/
func RestoreClusterRoleBind(clusterResource *string)  {
	clusterList,err  :=models.GetClusterList()
	if err!=nil {
		log.Fatalln(err)
	}
	List :=clusterList.ClusterRoleBindingList
	var Items []string
	//获取当前集群中所存在的资源
	for _,i :=range List.Items {
		Items = append(Items, i.GetName())
	}
	fileDir :="./BackupFile/b_"+*clusterResource+"list"
	//获取备份数据
	list :=RestoreOutPutFile(fileDir)
	resources :=list[0]
	//转化结构体
	//转为bytes 转为结构体
	bytes,err:=json.Marshal(resources)
	var res *rbacv1.ClusterRoleBindingList
	json.Unmarshal(bytes,&res)

	for _,clusterRoleBind:=range res.Items  {
		if !In(clusterRoleBind.Name,Items) {
			crb,err:=clusterList.ClientSet.RbacV1().ClusterRoleBindings().Create(context.TODO(),&clusterRoleBind,metav1.CreateOptions{})
			if err != nil {
				log.Fatalln(err)
			}
			fmt.Printf("restore clusterRoleBind%s \n", crb.GetName())
		}
	}
}
/*
	恢复 pv
*/
func RestorePV(clusterResource *string)  {
	clusterList,err  :=models.GetClusterList()
	if err!=nil {
		log.Fatalln(err)
	}
	List :=clusterList.PvList
	var Items []string
	//获取当前集群中所存在的资源
	for _,i :=range List.Items {
		Items = append(Items, i.GetName())
	}
	fileDir :="./BackupFile/b_"+*clusterResource+"list"
	//获取备份数据
	list :=RestoreOutPutFile(fileDir)
	resources :=list[0]
	//转化结构体
	//转为bytes 转为结构体
	bytes,err:=json.Marshal(resources)
	var res v1.PersistentVolumeList
	json.Unmarshal(bytes,&res)

	for _,pv:=range res.Items  {
		if !In(pv.Name,Items) {
			pv.ObjectMeta.ResourceVersion = " "
			p,err:=clusterList.ClientSet.CoreV1().PersistentVolumes().Create(context.TODO(),&pv,metav1.CreateOptions{})
			if err != nil {
				log.Fatalln(err)
			}
			fmt.Printf("restore pv%s \n", p.GetName())
		}
	}
}



package models

import (
	"WorkPro/conn"
	"context"
	v1 "k8s.io/api/core/v1"
	rbacv1 "k8s.io/api/rbac/v1"
	storagev1 "k8s.io/api/storage/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"log"
)

/*
获取集群级别资源集群级别资源
*/
type cluster struct {
	ClientSet *kubernetes.Clientset
	NodeList *v1.NodeList
	NamespaceList *v1.NamespaceList
	PvList *v1.PersistentVolumeList
	StorageClassList *storagev1.StorageClassList
	ClusterRoleList *rbacv1.ClusterRoleList
	ClusterRoleBindingList *rbacv1.ClusterRoleBindingList
}

func GetClusterList() (*cluster,error) {
	clientSet,err := conn.CreateClient()
	if err!=nil {
		log.Fatalln(err)
	}
	nodelist,err :=clientSet.CoreV1().Nodes().List(context.TODO(),metav1.ListOptions{})
	if err!=nil {
		log.Fatalln(err)
	}
	namespacelist, err := clientSet.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})
	if err!=nil {
		log.Fatalln(err)
	}
	storageclassList, err := clientSet.StorageV1().StorageClasses().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Fatalln(err)
	}
	pvList, err := clientSet.CoreV1().PersistentVolumes().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Fatalln(err)
	}
	crList ,err :=clientSet.RbacV1().ClusterRoles().List(context.TODO(),metav1.ListOptions{})
	if err != nil {
		log.Fatalln(err)
	}
	crbList, err := clientSet.RbacV1().ClusterRoleBindings().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Fatalln(err)
	}
	clu :=&cluster{
		NodeList: nodelist,
		NamespaceList: namespacelist,
		StorageClassList: storageclassList,
		PvList: pvList,
		ClusterRoleList: crList,
		ClusterRoleBindingList: crbList,
	}

	return clu,nil
}
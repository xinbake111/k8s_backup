package models

import (
	"WorkPro/conn"
	"context"
	appv1 "k8s.io/api/apps/v1"
	batchv1 "k8s.io/api/batch/v1"
	v1 "k8s.io/api/core/v1"
	rbacv1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"log"
)

type Space struct {
	ClientSet *kubernetes.Clientset
	PodList *v1.PodList
	DeploymentList *appv1.DeploymentList
	DaemonSetList *appv1.DaemonSetList
	JobList *batchv1.JobList
	CornJobList *batchv1.CronJobList
	ServiceList  *v1.ServiceList
	PvcList *v1.PersistentVolumeClaimList
	ConfigMapList *v1.ConfigMapList
	RoleList *rbacv1.RoleList
	RoleBindList *rbacv1.RoleBindingList
}

func GetSpaceList(namespace string)  (*Space,error) {
	clientSet,err := conn.CreateClient()
	if err!=nil {
		log.Fatalln(err)
	}
	PodList,err :=clientSet.CoreV1().Pods(namespace).List(context.TODO(),metav1.ListOptions{})
	if err!=nil {
		log.Fatalln(err)
	}
	DeploymentList,err :=clientSet.AppsV1().Deployments(namespace).List(context.TODO(),metav1.ListOptions{})
	if err!=nil {
		log.Fatalln(err)
	}
	DaemonSetList,err :=clientSet.AppsV1().DaemonSets(namespace).List(context.TODO(),metav1.ListOptions{})
	if err!=nil {
		log.Fatalln(err)
	}
	JobList,err:=clientSet.BatchV1().Jobs(namespace).List(context.TODO(),metav1.ListOptions{})
	if err!=nil {
		log.Fatalln(err)
	}
	CornJobList,err:=clientSet.BatchV1().CronJobs(namespace).List(context.TODO(),metav1.ListOptions{})
	if err!=nil {
		log.Fatalln(err)
	}
	ServiceList ,err :=clientSet.CoreV1().Services(namespace).List(context.TODO(),metav1.ListOptions{})
	if err!=nil {
		log.Fatalln(err)
	}
	PvcList,err:=clientSet.CoreV1().PersistentVolumeClaims(namespace).List(context.TODO(),metav1.ListOptions{})
	if err!=nil {
		log.Fatalln(err)
	}
	ConfigMapList,err :=clientSet.CoreV1().ConfigMaps(namespace).List(context.TODO(),metav1.ListOptions{})
	if err!=nil {
		log.Fatalln(err)
	}
	RoleList,err :=clientSet.RbacV1().Roles(namespace).List(context.TODO(),metav1.ListOptions{})
	if err!=nil {
		log.Fatalln(err)
	}
	RoleBindList,err :=clientSet.RbacV1().RoleBindings(namespace).List(context.TODO(),metav1.ListOptions{})
	if err!=nil {
		log.Fatalln(err)
	}
	space :=&Space{
		ClientSet: clientSet,
		PodList: PodList,
		DeploymentList: DeploymentList,
		DaemonSetList: DaemonSetList,
		JobList: JobList,
		CornJobList: CornJobList,
		ServiceList: ServiceList,
		PvcList: PvcList,
		ConfigMapList: ConfigMapList,
		RoleList: RoleList,
		RoleBindList: RoleBindList,
	}
	return space,err
}
package utils

import (
	"WorkPro/models"
	"context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"log"
)

//备份结构体

type CsiBackupInfo struct {
	PvcName string
	Pvc_StorageClass map[string]string
	Storageclass_Csiplugins map[string]string
	Pvc_SnapClassName map[string]string
	SnapClass_SnapName map[string]string
}

func CsiBackup(namespace string,podname string)  {
	cru,err:=models.GetSpaceList(namespace)
	if err!=nil {
		log.Println(err)
	}
	var pvcname []string
	var pvc_storageclass=make(map[string]string)
	var storageclass_csiplugins=make(map[string]string)
	podres,err:=cru.ClientSet.CoreV1().Pods(namespace).Get(context.TODO(),podname,metav1.GetOptions{})
	for _, volume := range podres.Spec.Volumes {
		pvcstr :=volume.PersistentVolumeClaim.ClaimName
		pvcname =append(pvcname, pvcstr)
		//读取pvc字段
		pvc,err :=cru.ClientSet.CoreV1().PersistentVolumeClaims(namespace).Get(context.TODO(),pvcstr,metav1.GetOptions{})
		if err!=nil {
			log.Println(err)
		}
		storageclassName :=pvc.Spec.StorageClassName
		if storageclassName!=nil {
			storageclass,err:=cru.ClientSet.StorageV1beta1().StorageClasses().Get(context.TODO(),*storageclassName,metav1.GetOptions{})
			if err!=nil {
				log.Println(err)
			}
			pvc_storageclass[pvcstr]=*storageclassName
			csiplugin :=storageclass.Provisioner
			storageclass_csiplugins[*storageclassName]=csiplugin
		}
	}

}

//创建snapclass 进行存储


//创建snap 进行存储
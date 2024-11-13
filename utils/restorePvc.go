package utils

import (
	"WorkPro/models"
	"context"
	"encoding/json"
	"fmt"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"log"
)

/*----------------------------------根据具体的pvc名称进行pvc的恢复-------------------------------------------------------*/
func ReadPvcForName(podname string, namespace string,pvcNames []v1.Volume) {
	fmt.Println("____________恢复podpvc开始_____________")
	var pvcVolumes []string
	for i := 0; i < len(pvcNames); i++ {
		pvcVolumes =append(pvcVolumes,pvcNames[i].Name)
	}
	//查询存放pvc的路径
	filename :=models.QueryResources(namespace,podname)
	//如果路径存在
	if filename != "" {
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
			//进行pvc恢复
			spacelist,err:=models.GetSpaceList(namespace)
			if err!=nil {
				log.Fatalln(err)
			}
			pvclist:=spacelist.PvcList //已经存在的pvc
			var pvcItems []string
			for _,i :=range pvclist.Items{
				pvcItems=append(pvcItems,i.GetName())
			}
			list :=RestoreOutPutFile(filename)
			//序列化
			for i := 0; i < len(list); i++ {
				resources :=list[i]
				//转化结构体
				//转为bytes 转为结构体
				bytes,err:=json.Marshal(resources)
				if err!=nil {
					log.Println(err)
				}
				var pvcList v1.PersistentVolumeClaimList
				json.Unmarshal(bytes,&pvcList)


				for _, pvc := range pvcList.Items {
					if pvc.Spec.StorageClassName != nil {
						//pvc进行对比
						if In(pvc.GetName(), pvcVolumes) && !In(pvc.GetName(), pvcItems) {
							pvc.ObjectMeta.ResourceVersion = " "
							newpvc := &v1.PersistentVolumeClaim{
								ObjectMeta: metav1.ObjectMeta{
									Name: pvc.Name,
								},
								Spec: v1.PersistentVolumeClaimSpec{
									StorageClassName: pvc.Spec.StorageClassName,
									Resources:        pvc.Spec.Resources,
									AccessModes:      pvc.Spec.AccessModes,
									VolumeMode:       pvc.Spec.VolumeMode,
								},
							}
							p, err := spacelist.ClientSet.CoreV1().PersistentVolumeClaims(namespace).Create(context.TODO(), newpvc, metav1.CreateOptions{})
							if err != nil {
								log.Fatalln(err)
							}
							fmt.Println(p.Name)

						}
					} else {
						if In(pvc.GetName(), pvcVolumes) && !In(pvc.GetName(), pvcItems) {
							pvc.ObjectMeta.ResourceVersion = " "
							newpvc := &v1.PersistentVolumeClaim{
								ObjectMeta: metav1.ObjectMeta{
									Name: pvc.Name,
								},
								Spec: v1.PersistentVolumeClaimSpec{
									Resources:   pvc.Spec.Resources,
									AccessModes: pvc.Spec.AccessModes,
									VolumeMode:  pvc.Spec.VolumeMode,
								},
							}
							p, err := spacelist.ClientSet.CoreV1().PersistentVolumeClaims(namespace).Create(context.TODO(), newpvc, metav1.CreateOptions{})
							if err != nil {
								log.Fatalln(err)
							}
							fmt.Println(p.Name)
						}

					}

				}
			}
		}
		fmt.Println("____________恢复podpvc结束_____________")
	}
}

func ReadDeployPvcName(deployname string, namespace string,pvcNames []v1.Volume) {
	fmt.Println("____________恢复podpvc开始_____________")
	var pvcVolumes []string
	for i := 0; i < len(pvcNames); i++ {
		pvcVolumes =append(pvcVolumes,pvcNames[i].Name)
	}
	//查询存放pvc的路径
	filename :=models.QueryDeployResources(namespace,deployname)
	//如果路径存在
	if filename != "" {
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
			//进行pvc恢复
			spacelist,err:=models.GetSpaceList(namespace)
			if err!=nil {
				log.Fatalln(err)
			}
			deploylist:=spacelist.DeploymentList //已经存在的pvc
			var pvcItems []string
			for _,i :=range deploylist.Items{
				pvcItems=append(pvcItems,i.GetName())
			}
			list :=RestoreOutPutFile(filename)
			//序列化
			for i := 0; i < len(list); i++ {
				resources :=list[i]
				//转化结构体
				//转为bytes 转为结构体
				bytes,err:=json.Marshal(resources)
				if err != nil {
					log.Println(err)
				}
				var pvcList v1.PersistentVolumeClaimList
				json.Unmarshal(bytes,&pvcList)


				for _, pvc := range pvcList.Items {
					if pvc.Spec.StorageClassName != nil {
						//pvc进行对比
						if In(pvc.GetName(), pvcVolumes) && !In(pvc.GetName(), pvcItems) {
							pvc.ObjectMeta.ResourceVersion = " "
							newpvc := &v1.PersistentVolumeClaim{
								ObjectMeta: metav1.ObjectMeta{
									Name: pvc.Name,
								},
								Spec: v1.PersistentVolumeClaimSpec{
									StorageClassName: pvc.Spec.StorageClassName,
									Resources:        pvc.Spec.Resources,
									AccessModes:      pvc.Spec.AccessModes,
									VolumeMode:       pvc.Spec.VolumeMode,
								},
							}
							p, err := spacelist.ClientSet.CoreV1().PersistentVolumeClaims(namespace).Create(context.TODO(), newpvc, metav1.CreateOptions{})
							if err != nil {
								log.Fatalln(err)
							}
							fmt.Println(p.Name)

						}
					} else {
						if In(pvc.GetName(), pvcVolumes) && !In(pvc.GetName(), pvcItems) {
							pvc.ObjectMeta.ResourceVersion = " "
							newpvc := &v1.PersistentVolumeClaim{
								ObjectMeta: metav1.ObjectMeta{
									Name: pvc.Name,
								},
								Spec: v1.PersistentVolumeClaimSpec{
									Resources:   pvc.Spec.Resources,
									AccessModes: pvc.Spec.AccessModes,
									VolumeMode:  pvc.Spec.VolumeMode,
								},
							}
							p, err := spacelist.ClientSet.CoreV1().PersistentVolumeClaims(namespace).Create(context.TODO(), newpvc, metav1.CreateOptions{})
							if err != nil {
								log.Fatalln(err)
							}
							fmt.Println(p.Name)
						}

					}

				}
			}
		}
		fmt.Println("____________恢复deploypvc结束_____________")
	}
}


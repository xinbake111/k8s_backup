package utils

import (
	"WorkPro/models"
	"encoding/json"
	"fmt"
	appv1 "k8s.io/api/apps/v1"
	batchv1 "k8s.io/api/batch/v1"
	v1 "k8s.io/api/core/v1"
	"log"
)
/*
	创建一个结构体 里面存放pod 与 pod下的pvc
*/
type BackupSource struct {
	
}
type SpaceResForPVC struct {
	NameSpace string
	PodName string
	DeployName string
	DaemonSetName string
	JobName string
	CornJobName string
	PvcName string
	Path	string
}
/*
	K8s资源写入为数据文件
*/

func BackUpInPutFile(v interface{},fileMkdir string,filePath string)  {
	BackUpInPutECCFileFile(v,fileMkdir,filePath)
}

/**
 读取namespace数据并存入到文件之中
**/
func BackUpNameSpace()  {
	clusterList,err :=models.GetClusterList()
	if err!=nil {
		log.Fatalln(err)
	}
	namespaceList :=clusterList.NamespaceList
	for _,namespace :=range namespaceList.Items  {
		//排除集群默认命名空间
		if namespace.Name != "kube-node-lease" &&
			namespace.Name != "kube-public" && namespace.Name != "kube-system" {
			//创建资源类型
			var spaceResource = []string {"pod","deploy", "daemonset", "job","cornjob",
				"configmap","service","pvc","role","rolebind"}
			for i := 0; i < len(spaceResource); i++ {
				//进行当前Namespace下资源获取
				BackupSpaceResource(namespace.Name,&spaceResource[i])
			}
		}
	}
}
/*
		通过对传入 cluster资源类别 进行备份
*/
func BackupClusterResource(clusterResource *string)  {
	clusterList,err :=models.GetClusterList()
	if err!=nil {
		log.Fatalln(err)
	}
	var List interface{}
	switch  {
	case *clusterResource=="namespace":
		List =clusterList.NamespaceList
	case *clusterResource=="storageclass":
		List =clusterList.StorageClassList
	case *clusterResource=="pv":
		List =clusterList.PvList
	case *clusterResource=="clusterrole":
		List =clusterList.ClusterRoleList
	case *clusterResource=="clusterrolebind":
		List =clusterList.ClusterRoleBindingList
	}
	fmt.Println("_________进行"+*clusterResource+"备份_________")
	fileMkdir :="./BackupFile/b_"+*clusterResource+"list"
	filepath := fileMkdir + "/" + *clusterResource+".txt"
	BackUpInPutFile(List,fileMkdir,filepath)
	fmt.Println("_________结束ClusterRoleBindingList备份_________")

	//fmt.Println("_________进行元数据"+*clusterResource+"备份_________")
	//fileMkdirA :="./BackupFileLI/b_"+*clusterResource+"list"
	//filepathA := fileMkdirA + "/" + *clusterResource+".txt"
	//BackUpInPutFileA(List,fileMkdirA,filepathA)
	//fmt.Println("_________结束ClusterRoleBindingList备份_________")
}
/*
	通过对传入资源种类进行备份 和namespace
*/
func BackupSpaceResource(namespace string,spaceResource *string)  {
	spaceList ,err :=models.GetSpaceList(namespace)
	if err!=nil {
		log.Fatalln(err)
	}
	var List interface{}
	var pdc []SpaceResForPVC
	switch  {
	case *spaceResource=="pod":
		List =spaceList.PodList
		//如果对pod进行备份则需要进行关系绑定 创建结构体进行pod和pvc的映射关系以便于后期恢复
		//可以进行数据库及逆行改变，如果备份则再数据中添加映射关系，如果恢复则把此映射进行删除
		pdc=GetPodPvcList(namespace,List)
		fmt.Println("_________进行"+*spaceResource+"备份_________")
		fileMkdir :="./BackupFile/"+namespace+"/b_"+*spaceResource+"list"
		filepath := fileMkdir + "/" + namespace+"_"+*spaceResource+".txt"
		BackUpInPutFile(List,fileMkdir,filepath)
		fmt.Println("_________结束"+*spaceResource+"备份_________")
		if len(pdc)!=0 {
			for i := 0; i < len(pdc); i++ {
				BackUpPvcResources(namespace,pdc[i])
			}
		}else {
			log.Println("阶段备份结束")
		}

	case *spaceResource=="deploy":
		fmt.Println("------deploy阶段开始----")
		List =spaceList.DeploymentList
		pdc=GetDeployPvcList(namespace,List)
		fmt.Println("_________进行"+*spaceResource+"备份_________")
		fileMkdir :="./BackupFile/"+namespace+"/b_"+*spaceResource+"list"
		filepath := fileMkdir + "/" + namespace+"_"+*spaceResource+".txt"
		BackUpInPutFile(List,fileMkdir,filepath)
		fmt.Println("_________结束"+*spaceResource+"备份_________")
		if len(pdc)!=0 {
			for i := 0; i < len(pdc); i++ {
				BackUpPvcResources(namespace,pdc[i])
			}
		}else {
			log.Println("阶段备份结束")
		}
	case *spaceResource=="daemonset":
		List =spaceList.DaemonSetList
		pdc=GetDaemonSetPvcList(namespace,List)
		fmt.Println("_________进行"+*spaceResource+"备份_________")
		fileMkdir :="./BackupFile/"+namespace+"/b_"+*spaceResource+"list"
		filepath := fileMkdir + "/" + namespace+"_"+*spaceResource+".txt"
		BackUpInPutFile(List,fileMkdir,filepath)
		fmt.Println("_________结束"+*spaceResource+"备份_________")
		if len(pdc)!=0 {
			for i := 0; i < len(pdc); i++ {
				BackUpPvcResources(namespace,pdc[i])
			}
		}else {
			log.Println("阶段备份结束")
		}
	case *spaceResource=="job":
		List =spaceList.JobList
		pdc=GetJobPvcList(namespace,List)
		fmt.Println("_________进行"+*spaceResource+"备份_________")
		fileMkdir :="./BackupFile/"+namespace+"/b_"+*spaceResource+"list"
		filepath := fileMkdir + "/" + namespace+"_"+*spaceResource+".txt"
		BackUpInPutFile(List,fileMkdir,filepath)
		fmt.Println("_________结束"+*spaceResource+"备份_________")
		if len(pdc)!=0 {
			for i := 0; i < len(pdc); i++ {
				BackUpPvcResources(namespace,pdc[i])
			}
		}else {
			log.Println("阶段备份结束")
		}
	case *spaceResource=="cornjob":
		List =spaceList.CornJobList
		pdc=GetCornJobPvcList(namespace,List)
		fmt.Println("_________进行"+*spaceResource+"备份_________")
		fileMkdir :="./BackupFile/"+namespace+"/b_"+*spaceResource+"list"
		filepath := fileMkdir + "/" + namespace+"_"+*spaceResource+".txt"
		BackUpInPutFile(List,fileMkdir,filepath)
		fmt.Println("_________结束"+*spaceResource+"备份_________")
		if len(pdc)!=0 {
			for i := 0; i < len(pdc); i++ {
				BackUpPvcResources(namespace,pdc[i])
			}
		}else {
			log.Println("阶段备份结束")
		}
	case *spaceResource=="configmap":
		List =spaceList.ConfigMapList
		fmt.Println("_________进行"+*spaceResource+"备份_________")
		fileMkdir :="./BackupFile/"+namespace+"/b_"+*spaceResource+"list"
		filepath := fileMkdir + "/" + namespace+"_"+*spaceResource+".txt"
		BackUpInPutFile(List,fileMkdir,filepath)
		fmt.Println("_________结束"+*spaceResource+"备份_________")
	case *spaceResource=="service":
		List =spaceList.ServiceList
		fmt.Println("_________进行"+*spaceResource+"备份_________")
		fileMkdir :="./BackupFile/"+namespace+"/b_"+*spaceResource+"list"
		filepath := fileMkdir + "/" + namespace+"_"+*spaceResource+".txt"
		BackUpInPutFile(List,fileMkdir,filepath)
		fmt.Println("_________结束"+*spaceResource+"备份_________")
	case *spaceResource=="pvc":
		List =spaceList.PvcList
		fmt.Println("_________进行"+*spaceResource+"备份_________")
		fileMkdir :="./BackupFile/"+namespace+"/b_"+*spaceResource+"list"
		filepath := fileMkdir + "/" + namespace+"_"+*spaceResource+".txt"
		BackUpInPutFile(List,fileMkdir,filepath)
		fmt.Println("_________结束"+*spaceResource+"备份_________")
	case *spaceResource=="role":
		List =spaceList.RoleList
		fmt.Println("_________进行"+*spaceResource+"备份_________")
		fileMkdir :="./BackupFile/"+namespace+"/b_"+*spaceResource+"list"
		filepath := fileMkdir + "/" + namespace+"_"+*spaceResource+".txt"
		BackUpInPutFile(List,fileMkdir,filepath)
		fmt.Println("_________结束"+*spaceResource+"备份_________")
	case *spaceResource=="rolebind":
		List =spaceList.RoleBindList
		fmt.Println("_________进行"+*spaceResource+"备份_________")
		fileMkdir :="./BackupFile/"+namespace+"/b_"+*spaceResource+"list"
		filepath := fileMkdir + "/" + namespace+"_"+*spaceResource+".txt"
		BackUpInPutFile(List,fileMkdir,filepath)
		fmt.Println("_________结束"+*spaceResource+"备份_________")
	}
	//if *spaceResource!= "pod" ||*spaceResource!= "deploy" ||
	//	*spaceResource!= "daemonset" ||*spaceResource!= "job" || *spaceResource!="cornjob"  {
	//	fmt.Println("_________进行"+*spaceResource+"备份_________")
	//	fileMkdir :="./BackupFile/"+namespace+"/b_"+*spaceResource+"list"
	//	filepath := fileMkdir + "/" + namespace+"_"+*spaceResource+".txt"
	//	BackUpInPutFile(List,fileMkdir,filepath)
	//	fmt.Println("_________结束"+*spaceResource+"备份_________")
	//}
}
func GetPodPvcList(namespace string, v interface{}) ([]SpaceResForPVC) {
	list :=v.(*v1.PodList)
	var podpvc []SpaceResForPVC
	for _,pod :=range list.Items {
		if len(pod.Spec.Volumes)==0 {
			log.Fatalln("此pod未使用挂在卷")
		}else {
			for _,volume :=range pod.Spec.Volumes{
				if volume.PersistentVolumeClaim!=nil&& volume.PersistentVolumeClaim.ClaimName !=" " {
					podpvc =append(podpvc,SpaceResForPVC{
						NameSpace: namespace,
						PodName: pod.Name,
						PvcName: volume.PersistentVolumeClaim.ClaimName,
					})
				}
			}
		}
	}
	return podpvc
}

func GetDeployPvcList(namespace string, v interface{}) ([]SpaceResForPVC) {
	list :=v.(*appv1.DeploymentList)
	var SpaceResPvc []SpaceResForPVC
	for _,res :=range list.Items {
		if len(res.Spec.Template.Spec.Volumes)==0 {
			log.Println("此pod未使用挂在卷")
		}else {
			for _,volume :=range res.Spec.Template.Spec.Volumes{
				if volume.PersistentVolumeClaim!=nil&& volume.PersistentVolumeClaim.ClaimName !=" " {
					SpaceResPvc =append(SpaceResPvc,SpaceResForPVC{
						NameSpace: namespace,
						DeployName: res.Name,
						PvcName: volume.PersistentVolumeClaim.ClaimName,
					})
				}
			}
		}
	}
	return SpaceResPvc
}
func GetDaemonSetPvcList(namespace string, v interface{}) ([]SpaceResForPVC) {
	list :=v.(*appv1.DaemonSetList)
	var SpaceResPvc []SpaceResForPVC
	for _,res :=range list.Items {
		if len(res.Spec.Template.Spec.Volumes)==0 {
			log.Println("此pod未使用挂在卷")
		}else {
			for _,volume :=range res.Spec.Template.Spec.Volumes{
				if volume.PersistentVolumeClaim!=nil&& volume.PersistentVolumeClaim.ClaimName !=" " {
					SpaceResPvc =append(SpaceResPvc,SpaceResForPVC{
						NameSpace: namespace,
						DaemonSetName: res.Name,
						PvcName: volume.PersistentVolumeClaim.ClaimName,
					})
				}
			}
		}
	}
	return SpaceResPvc
}
func GetJobPvcList(namespace string, v interface{}) ([]SpaceResForPVC) {
	list :=v.(*batchv1.JobList)
	var SpaceResPvc []SpaceResForPVC
	for _,res :=range list.Items {
		if len(res.Spec.Template.Spec.Volumes)==0 {
			log.Println("此pod未使用挂在卷")
		}else {
			for _,volume :=range res.Spec.Template.Spec.Volumes{
				if volume.PersistentVolumeClaim!=nil&& volume.PersistentVolumeClaim.ClaimName !=" " {
					SpaceResPvc =append(SpaceResPvc,SpaceResForPVC{
						NameSpace: namespace,
						JobName: res.Name,
						PvcName: volume.PersistentVolumeClaim.ClaimName,
					})
				}
			}
		}
	}
	return SpaceResPvc
}
func GetCornJobPvcList(namespace string, v interface{}) ([]SpaceResForPVC) {
	list :=v.(*batchv1.CronJobList)
	var SpaceResPvc []SpaceResForPVC
	for _,res :=range list.Items {
		if len(res.Spec.JobTemplate.Spec.Template.Spec.Volumes)==0 {
			log.Println("此pod未使用挂在卷")
		}else {
			for _,volume :=range res.Spec.JobTemplate.Spec.Template.Spec.Volumes{
				if volume.PersistentVolumeClaim!=nil&& volume.PersistentVolumeClaim.ClaimName !=" " {
					SpaceResPvc =append(SpaceResPvc,SpaceResForPVC{
						NameSpace: namespace,
						CornJobName: res.Name,
						PvcName: volume.PersistentVolumeClaim.ClaimName,
					})
				}
			}
		}
	}
	return SpaceResPvc
}

func BackUpPvcResources(namespace string,SpaceForPvc SpaceResForPVC)  {
	spaceList ,err :=models.GetSpaceList(namespace)
	if err!=nil {
		log.Fatalln(err)
	}
	var List interface{}
	//查询出来pvc的内容
	List =spaceList.PvcList
	fmt.Println("_________进行pvc备份_________")
	fileMkdir :="./BackupFile/"+namespace+"/b_pvclist"
	filepath := fileMkdir + "/" + namespace+"_pvc.txt"
	BackUpInPutFile(List,fileMkdir,filepath)
	pvclist:=List.(*v1.PersistentVolumeClaimList)
	//进行关系映射
	for _,pvc := range pvclist.Items{
		if pvc.GetName()==SpaceForPvc.PvcName {
			SpaceForPvc.Path=filepath
		}
	}
	bytes, err := json.Marshal(SpaceForPvc)
	if err!=nil {
		log.Fatalln(err)
	}
	var newrfp = models.SpaceResForPVC{}
	json.Unmarshal(bytes,newrfp)
	models.InsertRelation(newrfp)
	fmt.Println("_________结束pvc备份_________")
}
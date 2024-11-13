package utils

import (
	"WorkPro/models"
	"fmt"
	"log"
)
func BackUpStorageClass()  {
	clusterList,err :=models.GetClusterList()
	if err!=nil {
		log.Fatalln(err)
	}
	storageClassList :=clusterList.StorageClassList
	fmt.Println("_________进行StorageClass备份_________")
	fileMkdir :="./BackupFile/b_storageclass"
	filepath := fileMkdir + "/" + "StroageclassResources.txt"
	BackUpInPutFile(storageClassList,fileMkdir,filepath)
	fmt.Println("_________结束StorageClass备份_________")
}
func BackUpPv()  {
	clusterList,err :=models.GetClusterList()
	if err!=nil {
		log.Fatalln(err)
	}
	PvList :=clusterList.PvList
	fmt.Println("_________进行PV备份_________")
	fileMkdir :="./BackupFile/b_pv"
	filepath := fileMkdir + "/" + "pv.txt"
	BackUpInPutFile(PvList,fileMkdir,filepath)
	fmt.Println("_________结束PV备份_________")
}
func BackUpClusterRole()  {
	clusterList,err :=models.GetClusterList()
	if err!=nil {
		log.Fatalln(err)
	}
	ClusterRoleList :=clusterList.ClusterRoleList
	fmt.Println("_________进行ClusterRoleList备份_________")
	fileMkdir :="./BackupFile/b_clusterroles"
	filepath := fileMkdir + "/" + "clusterroles.txt"
	BackUpInPutFile(ClusterRoleList,fileMkdir,filepath)
	fmt.Println("_________结束ClusterRoleList备份_________")
}
func BackUpClusterRoleBind()  {
	clusterList,err :=models.GetClusterList()
	if err!=nil {
		log.Fatalln(err)
	}
	ClusterRoleBindingList :=clusterList.ClusterRoleBindingList
	fmt.Println("_________进行ClusterRoleBindingList备份_________")
	fileMkdir :="./BackupFile/b_clusterrolebinds"
	filepath := fileMkdir + "/" + "clusterrolebinds.txt"
	BackUpInPutFile(ClusterRoleBindingList,fileMkdir,filepath)
	fmt.Println("_________结束ClusterRoleBindingList备份_________")
}
/*
	获取某个namespace下的deployment资源
*/
func BackDeployment(namespace string)  {
	spaceList,err :=models.GetSpaceList(namespace)
	if err!=nil {
		log.Fatalln(err)
	}
	deployList :=spaceList.DeploymentList
	fmt.Println("_________进行ClusterRoleBindingList备份_________")
	fileMkdir :="./BackupFile/"+namespace+"b_deploylist/"
	filepath := fileMkdir + "/" + namespace+"_deploylist.txt"
	BackUpInPutFile(deployList,fileMkdir,filepath)
	fmt.Println("_________结束ClusterRoleBindingList备份_________")

}
/*
	获取某个namespace下的pod资源
*/
func BackUpPod(namespace string)  {
	spaceList,err :=models.GetSpaceList(namespace)
	if err!=nil {
		log.Fatalln(err)
	}
	podlist :=spaceList.PodList
	fmt.Println("_________进行ClusterRoleBindingList备份_________")
	fileMkdir :="./BackupFile/"+namespace+"/b_podlist"
	filepath := fileMkdir + "/" + namespace+"_podlist.txt"
	BackUpInPutFile(podlist,fileMkdir,filepath)
	fmt.Println("_________结束ClusterRoleBindingList备份_________")
}
func BackUpPVC(namespace string){
	spaceList,err :=models.GetSpaceList(namespace)
	if err!=nil {
		log.Fatalln(err)
	}
	pvclist :=spaceList.PvcList
	fmt.Println("_________进行ClusterRoleBindingList备份_________")
	fileMkdir :="./BackupFile/"+namespace+"/b_pvclist"
	filepath := fileMkdir + "/" + namespace+"_pvclist.txt"
	BackUpInPutFile(pvclist,fileMkdir,filepath)
	fmt.Println("_________结束ClusterRoleBindingList备份_________")
}
func BackUpService(namespace string)  {
	spaceList ,err :=models.GetSpaceList(namespace)
	if err!=nil {
		log.Fatalln(err)
	}
	serviceList :=spaceList.ServiceList
	fmt.Println("_________进行ServiceList备份_________")
	fileMkdir :="./BackupFile/"+namespace+"/b_servicelist"
	filepath := fileMkdir + "/" + namespace+"_servicelist.txt"
	BackUpInPutFile(serviceList,fileMkdir,filepath)
	fmt.Println("_________结束ServiceList备份_________")
}
func BackUpDaemonSet(namespace string)  {
	spaceList ,err :=models.GetSpaceList(namespace)
	if err!=nil {
		log.Fatalln(err)
	}
	daemonSetList:=spaceList.DaemonSetList
	fmt.Println("_________进行ServiceList备份_________")
	fileMkdir :="./BackupFile/"+namespace+"/b_daemonlist"
	filepath := fileMkdir + "/" + namespace+"_daemonlist.txt"
	BackUpInPutFile(daemonSetList,fileMkdir,filepath)
	fmt.Println("_________结束ServiceList备份_________")
}
func BackUpJob(namespace string)  {
	spaceList ,err :=models.GetSpaceList(namespace)
	if err!=nil {
		log.Fatalln(err)
	}
	jobList:=spaceList.JobList
	fmt.Println("_________进行ServiceList备份_________")
	fileMkdir :="./BackupFile/"+namespace+"/b_joblist"
	filepath := fileMkdir + "/" + namespace+"_joblist.txt"
	BackUpInPutFile(jobList,fileMkdir,filepath)
	fmt.Println("_________结束ServiceList备份_________")
}
func BackUpCornJob(namespace string)  {
	spaceList ,err :=models.GetSpaceList(namespace)
	if err!=nil {
		log.Fatalln(err)
	}
	cornjobList:=spaceList.CornJobList
	fmt.Println("_________进行ServiceList备份_________")
	fileMkdir :="./BackupFile/"+namespace+"/b_cornjoblist"
	filepath := fileMkdir + "/" + namespace+"_cornjoblist.txt"
	BackUpInPutFile(cornjobList,fileMkdir,filepath)
	fmt.Println("_________结束ServiceList备份_________")
}



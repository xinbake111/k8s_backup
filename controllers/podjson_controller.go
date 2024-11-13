package controllers

import (
	"WorkPro/models"
	//"encoding/json"
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	"log"
	"time"
)

type PodJsonController struct {
	beego.Controller
}
type PodCon struct {
	Id string
	Namespace string
	Name string
	Ready bool
	Status string
	Restart int
	IP string
	Node string

	Age        time.Duration
}

func (c *PodJsonController) Get() {
	//集群内信息
	//namespace :=c.GetString("namespace")
	test :="default"
	sp ,err :=models.GetSpaceList(test)
	if err!=nil {
		log.Fatalln(err)
	}
	count :=len(sp.PodList.Items)
	data :=make ([]PodCon,count)
	for i,res := range sp.PodList.Items{
		data[i].Id=string(res.UID)
		data[i].Name=string(res.Name)
		data[i].Namespace= res.Namespace
		data[i].Ready=res.Status.ContainerStatuses[0].Ready
		data[i].Status=string(res.Status.Phase)
		data[i].Restart=int(res.Status.ContainerStatuses[0].RestartCount)
		data[i].IP=res.Status.PodIP
		data[i].Node=res.Spec.NodeName
		now:=time.Now()
		creattime:=res.CreationTimestamp
		age :=now.Sub(creattime.Time)
		data[i].Age=age
	}
	fmt.Println("text-----------")
	c.Data["json"]=map[string]interface{}{"code": 0,"msg": "","count": count,"data": data}
	c.ServeJSON()
}


package controllers

import (
	"WorkPro/models"
	"encoding/json"
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	"log"
	"time"
)

type DeployJsonController struct {
	beego.Controller
}
type DeployCon struct {
	Id string
	Namespace string
	Name string
	Ready int
	Available int
	UpToDate int
	Images string
	Selector string
	Age        time.Duration
}

func (c *DeployJsonController) Get() {
	//集群内信息
	//namespace :=c.GetString("namespace")
	test :="default"
	sp ,err :=models.GetSpaceList(test)
	if err!=nil {
		log.Fatalln(err)
	}
	count :=len(sp.DeploymentList.Items)
	data :=make ([]DeployCon,count)
	for i,res := range sp.DeploymentList.Items{
		data[i].Id=string(res.UID)
		data[i].Name=string(res.Name)
		data[i].Ready=int(res.Status.ReadyReplicas)
		data[i].Available=int(res.Status.AvailableReplicas)
		data[i].UpToDate=int(res.Status.UpdatedReplicas)
		data[i].Images=res.Spec.Template.Spec.Containers[0].Image
		for j := 1; j < len(res.Spec.Template.Spec.Containers); j++ {
			data[i].Images =data[i].Images+","+res.Spec.Template.Spec.Containers[j].Image
		}
		dataType , _ := json.Marshal(res.Spec.Selector.MatchLabels)
		data[i].Selector = string(dataType)
		now:=time.Now()
		creattime:=res.CreationTimestamp
		age :=now.Sub(creattime.Time)
		data[i].Age=age
	}
	fmt.Println("text-----------")
	c.Data["json"]=map[string]interface{}{"code": 0,"msg": "","count": count,"data": data}
	c.ServeJSON()
}


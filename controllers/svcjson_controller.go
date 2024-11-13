package controllers

import (
	"WorkPro/models"
	"encoding/json"

	//"encoding/json"
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	"log"
	"time"
)

type SvcJsonController struct {
	beego.Controller
}
type SvcCon struct {
	Id string
	Namespace string
	Name string
	Type string
	ClusterIp string
	Port string
	Selector string
	Age        time.Duration
}

func (c *SvcJsonController) Get() {
	//集群内信息
	//namespace :=c.GetString("namespace")
	test :="default"
	sp ,err :=models.GetSpaceList(test)
	if err!=nil {
		log.Fatalln(err)
	}
	count :=len(sp.ServiceList.Items)
	data :=make ([]SvcCon,count)
	for i,res := range sp.ServiceList.Items{
		data[i].Id=string(res.UID)
		data[i].Name=string(res.Name)
		data[i].Namespace= res.Namespace
		data[i].Type=string(res.Spec.Type)
		data[i].ClusterIp=res.Spec.ClusterIP
		port , _ := json.Marshal(res.Spec.Ports)
		data[i].Port = string(port)
		labels , _ := json.Marshal(res.Labels)
		data[i].Selector =string(labels)
		now:=time.Now()
		creattime:=res.CreationTimestamp
		age :=now.Sub(creattime.Time)
		data[i].Age=age
	}
	fmt.Println("text-----------")
	c.Data["json"]=map[string]interface{}{"code": 0,"msg": "","count": count,"data": data}
	c.ServeJSON()
}


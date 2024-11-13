package controllers

import (
	"WorkPro/models"
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	"log"
	"time"
)

type NamespaceJsonController struct {
	beego.Controller
}
type NamespaceCon struct {
	Id         string
	Name       string
	Status     string
	Age        time.Duration
}

func (c *NamespaceJsonController) Get() {
	//集群内信息
	clu ,err :=models.GetClusterList()
	if err!=nil {
		log.Fatalln(err)
	}
	nums :=len(clu.NamespaceList.Items)
	data :=make ([]NamespaceCon,nums)
	count :=0;
	for i,namespace := range clu.NamespaceList.Items{
		data[i].Id=string(namespace.UID)
		data[i].Name=string(namespace.Name)
		data[i].Status=string(namespace.Status.Phase)
		now:=time.Now()
		creattime:=namespace.CreationTimestamp
		age :=now.Sub(creattime.Time)
		data[i].Age=age
		count++
	}
	fmt.Println("text-----------")
	c.Data["json"]=map[string]interface{}{"code": 0,"msg": "","count": count,"data": data}
	c.ServeJSON()
	//c.Data["NamespaceList"] =models.MakeNameSpaceBlocks(clu.NamespaceList.Items)
	//c.TplName = "page/space.html"
}

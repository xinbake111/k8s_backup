package controllers

import (
	"WorkPro/models"
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	"log"
)

type ClusterController struct {
	beego.Controller
}

func (c *ClusterController) Get() {
	fmt.Println("_______________获取集群中的cluster信息__________________")
	clu ,err :=models.GetClusterList()
	if err!=nil {
		log.Fatalln(err)
	}
	//c.Data["data"]=
	fmt.Println("____________________________________________________")

	c.Data["ClusterList"] = clu
	c.TplName = "cluster.html"
}

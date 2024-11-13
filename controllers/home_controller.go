package controllers

import (
	"WorkPro/models"
	"WorkPro/utils"
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	"log"
	"runtime"
)

type HomeController struct {
	beego.Controller
}


func (c *HomeController) Index() {
	// 系统配置信息
	c.Data["OS"], _ = beego.AppConfig.String("os")
	c.Data["Author"], _ = beego.AppConfig.String("author")
	c.Data["GOVersion"] 	= runtime.Version()
	c.Data["Version"], _ = beego.AppConfig.String("version")
	//集群内信息
	clu ,err :=models.GetClusterList()
	if err!=nil {
		log.Fatalln(err)
	}
	var nodesize int
	for _,i := range clu.NodeList.Items{
		fmt.Println(i.GetName())
		nodesize++
	}
	c.Data["NodeSize"] = nodesize
	var namespaceszie int
	var podsize int
	var servicesize int
	for _,i :=range clu.NamespaceList.Items{
		sl,err:=models.GetSpaceList(i.GetName())
		if err!=nil {
			log.Fatalln(err)
		}
		for _,j :=range sl.PodList.Items{
			j.GetName()
			podsize++
		}
		for _,j :=range sl.ServiceList.Items{
			j.GetName()
			servicesize++
		}
		namespaceszie++
	}
	c.Data["NamespaceSize"] =namespaceszie
	c.Data["PodSize"] =podsize
	c.Data["ServiceSize"] =servicesize
	//请求监控
	utils.GetOpt("http://10.147.18.73:7091/api/v1/rules")
	c.TplName = "page/welcome-1.html"
}


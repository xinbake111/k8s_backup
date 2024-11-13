package controllers

import (
	"WorkPro/utils"
	beego "github.com/beego/beego/v2/server/web"
)

type BackupTestController struct {
	beego.Controller
}

func (c * BackupTestController) Get() {
	var clusterResource = [5]string {"namespace","storageclass", "pv", "clusterrole",
		"clusterrolebind"}
	for i := 0; i < 5; i++ {
		utils.BackupClusterResource(&clusterResource[i])
	}
	utils.BackUpNameSpace()
	//
	utils.Bucket()

	c.TplName = "page/404.html"
}

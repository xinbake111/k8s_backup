package controllers

import (
	"WorkPro/utils"
	beego "github.com/beego/beego/v2/server/web"
)

type RestoreTestController struct {
	beego.Controller
}

func (c * RestoreTestController) Get() {
	utils.RestoreSpaceDeploy("default")
	utils.RestoreSpaceDeploy("flytest")
	utils.RestoreSpacePod("flytest")
	c.TplName = "page/404.html"
}

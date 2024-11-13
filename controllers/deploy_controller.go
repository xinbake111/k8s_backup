package controllers

import beego "github.com/beego/beego/v2/server/web"

type DeploySpaceController struct {
	beego.Controller
}

func (c *DeploySpaceController) Get() {
	c.TplName = "page/deploytable.html"
}

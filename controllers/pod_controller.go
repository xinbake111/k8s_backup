package controllers

import beego "github.com/beego/beego/v2/server/web"

type PodController struct {
	beego.Controller
}

func (c *PodController) Get() {
	c.TplName = "page/podtable.html"
}

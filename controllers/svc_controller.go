package controllers

import beego "github.com/beego/beego/v2/server/web"

type SvcController struct {
	beego.Controller
}

func (c *SvcController) Get() {
	c.TplName = "page/svctable.html"
}

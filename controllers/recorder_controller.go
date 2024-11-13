package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
)

type RecController struct {
	beego.Controller
}

func (c *RecController) Get() {
	c.TplName = "page/recordertable.html"
}

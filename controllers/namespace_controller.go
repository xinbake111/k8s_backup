package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
)

type NameSpaceController struct {
	beego.Controller
}

func (c *NameSpaceController) Get() {
	c.TplName = "page/namespacetable.html"
}

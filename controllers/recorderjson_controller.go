package controllers

import (
	"WorkPro/models"
	beego "github.com/beego/beego/v2/server/web"
	"time"
)

type RecJsonController struct {
	beego.Controller
}
type ReceCon struct {
	Id         string
	Name       string
	Status     string
	Age        time.Duration
}

func (c *RecJsonController) Get() {
	m:=models.QueryTimeBackup()


	c.Data["json"]=map[string]interface{}{"code": 0,"msg": "","count": len(m),"data": m}
	c.ServeJSON()

}

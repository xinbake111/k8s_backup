package main

import (
	"WorkPro/conn"
	_ "WorkPro/routers"
	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	conn.InitMysql()
	beego.Run()
}


package routers

import (
	"WorkPro/controllers"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	beego.Router("/cluster", &controllers.ClusterController{})
	beego.Router("/page/welcome", &controllers.HomeController{},"Get:Index")
    //namespace数据
	beego.Router("/page/namespace",&controllers.NameSpaceController{})
	beego.Router("/page/namespace/json",&controllers.NamespaceJsonController{})
	//deploy数据
	beego.Router("/page/deploy",&controllers.DeploySpaceController{})
	beego.Router("/page/deploy/json",&controllers.DeployJsonController{})
	//pod数据
	beego.Router("/page/pod",&controllers.PodController{})
	beego.Router("/page/pod/json",&controllers.PodJsonController{})
	//svc数据
	beego.Router("/page/svc",&controllers.SvcController{})
	beego.Router("/page/svc/json",&controllers.SvcJsonController{})
	//备份测试
	beego.Router("/backuptest",&controllers.BackupTestController{})
	//恢复测试
	beego.Router("/restore",&controllers.RestoreTestController{})
	//备份记录
	beego.Router("/page/recorder",&controllers.RecController{})
	beego.Router("/page/recorder/json",&controllers.RecJsonController{})
}

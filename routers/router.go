// @APIVersion 1.0.0
// @Title mobile API
// @Description mobile has every tool to get any job done, so codename for the new mobile APIs.
// @Contact astaxie@gmail.com

package routers

import (
	"github.com/astaxie/beego"
	"github.com/xuchengzhi/apimonitor/controllers"
)

func init() {
	// beego.Router("/", &controllers.MainController{})
	beego.Router("/", &controllers.MainController{})
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),
		beego.NSNamespace("/app",
			beego.NSInclude(
				&controllers.AppController{},
			),
		),
		beego.NSNamespace("/thrift",
			beego.NSInclude(
				&controllers.ThriftController{},
			),
		),
	)

	beego.AddNamespace(ns)
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.BConfig.WebConfig.Session.SessionName = "sessionID"
}

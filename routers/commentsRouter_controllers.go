package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["github.com/xuchengzhi/apimonitor/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/xuchengzhi/apimonitor/controllers:UserController"],
		beego.ControllerComments{
			Method:           "Auth",
			Router:           `/auth`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/xuchengzhi/apimonitor/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/xuchengzhi/apimonitor/controllers:UserController"],
		beego.ControllerComments{
			Method:           "Login",
			Router:           `/login`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/xuchengzhi/apimonitor/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/xuchengzhi/apimonitor/controllers:UserController"],
		beego.ControllerComments{
			Method:           "Registered",
			Router:           `/reg`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/xuchengzhi/apimonitor/controllers:AppController"] = append(beego.GlobalControllerRouter["github.com/xuchengzhi/apimonitor/controllers:AppController"],
		beego.ControllerComments{
			Method:           "FileUp",
			Router:           `/fileup`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/xuchengzhi/apimonitor/controllers:AppController"] = append(beego.GlobalControllerRouter["github.com/xuchengzhi/apimonitor/controllers:AppController"],
		beego.ControllerComments{
			Method:           "Applist",
			Router:           `/applist`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/xuchengzhi/apimonitor/controllers:AppController"] = append(beego.GlobalControllerRouter["github.com/xuchengzhi/apimonitor/controllers:AppController"],
		beego.ControllerComments{
			Method:           "TestApis",
			Router:           `/test`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/xuchengzhi/apimonitor/controllers:ThriftController"] = append(beego.GlobalControllerRouter["github.com/xuchengzhi/apimonitor/controllers:ThriftController"],
		beego.ControllerComments{
			Method:           "ActRun",
			Router:           `/act_run`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/xuchengzhi/apimonitor/controllers:ThriftController"] = append(beego.GlobalControllerRouter["github.com/xuchengzhi/apimonitor/controllers:ThriftController"],
		beego.ControllerComments{
			Method:           "ADB",
			Router:           `/adb`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/xuchengzhi/apimonitor/controllers:ThriftController"] = append(beego.GlobalControllerRouter["github.com/xuchengzhi/apimonitor/controllers:ThriftController"],
		beego.ControllerComments{
			Method:           "Atx",
			Router:           `/atx`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/xuchengzhi/apimonitor/controllers:ThriftController"] = append(beego.GlobalControllerRouter["github.com/xuchengzhi/apimonitor/controllers:ThriftController"],
		beego.ControllerComments{
			Method:           "Order",
			Router:           `/order`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/xuchengzhi/apimonitor/controllers:ThriftController"] = append(beego.GlobalControllerRouter["github.com/xuchengzhi/apimonitor/controllers:ThriftController"],
		beego.ControllerComments{
			Method:           "Device",
			Router:           `/device`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/xuchengzhi/apimonitor/controllers:ThriftController"] = append(beego.GlobalControllerRouter["github.com/xuchengzhi/apimonitor/controllers:ThriftController"],
		beego.ControllerComments{
			Method:           "Dev_up",
			Router:           `/dev_up`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/xuchengzhi/apimonitor/controllers:ThriftController"] = append(beego.GlobalControllerRouter["github.com/xuchengzhi/apimonitor/controllers:ThriftController"],
		beego.ControllerComments{
			Method:           "ThriftAtxServer",
			Router:           `/ThriftAtxServer`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/xuchengzhi/apimonitor/controllers:ThriftController"] = append(beego.GlobalControllerRouter["github.com/xuchengzhi/apimonitor/controllers:ThriftController"],
		beego.ControllerComments{
			Method:           "ThriftAdbServer",
			Router:           `/ThriftAdbServer`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/xuchengzhi/apimonitor/controllers:ThriftController"] = append(beego.GlobalControllerRouter["github.com/xuchengzhi/apimonitor/controllers:ThriftController"],
		beego.ControllerComments{
			Method:           "ThriftActServer",
			Router:           `/ThriftActServer`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})
}

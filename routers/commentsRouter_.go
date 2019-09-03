package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["github.com/xuchengzhi/apimonitor/controllers:ThriftController"] = append(beego.GlobalControllerRouter["github.com/xuchengzhi/apimonitor/controllers:ThriftController"],
        beego.ControllerComments{
            Method: "ActRun",
            Router: `/paas/:paasid`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/xuchengzhi/apimonitor/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/xuchengzhi/apimonitor/controllers:UserController"],
        beego.ControllerComments{
            Method: "Auth",
            Router: `/auth`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/xuchengzhi/apimonitor/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/xuchengzhi/apimonitor/controllers:UserController"],
        beego.ControllerComments{
            Method: "Login",
            Router: `/login`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/xuchengzhi/apimonitor/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/xuchengzhi/apimonitor/controllers:UserController"],
        beego.ControllerComments{
            Method: "Registered",
            Router: `/reg`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}

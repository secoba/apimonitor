package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["apimonitor/controllers:UserController"] = append(beego.GlobalControllerRouter["apimonitor/controllers:UserController"],
        beego.ControllerComments{
            Method: "Auth",
            Router: `/auth`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apimonitor/controllers:UserController"] = append(beego.GlobalControllerRouter["apimonitor/controllers:UserController"],
        beego.ControllerComments{
            Method: "Login",
            Router: `/login`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apimonitor/controllers:UserController"] = append(beego.GlobalControllerRouter["apimonitor/controllers:UserController"],
        beego.ControllerComments{
            Method: "Registered",
            Router: `/reg`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}

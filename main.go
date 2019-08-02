package main

import (
	"github.com/xuchengzhi/apimonitor/models"
	// "github.com/xuchengzhi/apimonitor/controllers"
	_ "github.com/xuchengzhi/apimonitor/routers"
	// "context"
	// "apimonitor/Libs"
	"github.com/astaxie/beego"
	// "net/http"
	// "strings"
)

func main() {
	// go rpcserver.Run()
	models.Init()
	models.UserInit()
	models.ApiInit()
	beego.Run()
	// beego.SetStaticPath("/static", "./static")
	beego.SetStaticPath("/static", "./img")
	// beego.BConfig.RouterCaseSensitive = false
	// beego.BConfig.WebConfig.ViewsPath="views"

}

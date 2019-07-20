package main

import (
	"apimonitor/models"
	_ "apimonitor/routers"
	// "context"
	// "apimonitor/Libs"
	"github.com/astaxie/beego"
	// "net/http"
	// "strings"
)

func main() {
	// go rpcserver.Run()
	models.Init()
	beego.Run()
	beego.SetStaticPath("/static", "./static")
	beego.SetStaticPath("/image", "./image")
	beego.BConfig.RouterCaseSensitive = false
	beego.BConfig.Listen.Graceful = true
}

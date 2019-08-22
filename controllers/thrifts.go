package controllers

import (
	"github.com/xuchengzhi/apimonitor/ThriftTest"
)

type ThriftController struct {
	BaseController
}

func (this *ThriftController) ThriftTest() {
	msg := this.GetString("msg")
	res := Thrift.Run(msg)
	this.Data["json"] = Response{200, "success.", res}
	this.ServeJSON()
}

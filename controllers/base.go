package controllers

import (
	"github.com/astaxie/beego"
	// "strconv"
	"strings"
)

type BaseController struct {
	beego.Controller
}

//Response 结构体
type Response struct {
	Errcode int         `json:"code"`
	Errmsg  string      `json:"msg"`
	Data    interface{} `json:"data"`
}

//Response 结构体
type ErrResponse struct {
	Errcode int         `json:"code"`
	Errmsg  interface{} `json:"msg"`
}

func (this *BaseController) getClientIp() string {
	s := strings.Split(this.Ctx.Request.RemoteAddr, ":")
	return s[0]
}

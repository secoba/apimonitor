package controllers

import (
	"github.com/xuchengzhi/apimonitor/models"
	"github.com/xuchengzhi/Library/AppInfo"
	"github.com/satori/go.uuid"
	"path"
	"path/filepath"
	"strings"
	"fmt"
	// "github.com/astaxie/beego"
	"log"
)


var (
	ErrAppTypeErr     = ErrResponse{423001, "上传文件非app类型"}
	ErrAppSaveErr  = ErrResponse{423002, "文件保存失败"}
)


type AppController struct {
	BaseController
}
// type LoginToken struct {
// 	User  models.User `json:"user"`
// 	Token string      `json:"token"`
// }

type AppJson struct {
	Name     string
	Version  string
	AppName  string
	MVersion string
	OldName  string
	Apptype  string
	Describe string
}


func (this *AppController) FileUp() {
	file,head,err:=this.GetFile("file")
	u1 := uuid.Must(uuid.NewV4())
	filename := head.Filename
	
	apptype := path.Ext(filename)
	log.Println(filename,file,apptype)
	if err!=nil {
		log.Println(err.Error())
        this.Ctx.ResponseWriter.WriteHeader(423)
		this.Data["json"] = ErrResponse{403001, "msg"}
		this.ServeJSON()
        return
    }
    defer file.Close()
    var AllowExtMap map[string]bool = map[string]bool{
        ".apk":true,
        ".ipa":true,
        ".IPA":true,
    }
    if _,ok:=AllowExtMap[apptype];!ok{
    	this.Ctx.ResponseWriter.WriteHeader(423)
		this.Data["json"] = ErrAppTypeErr
		this.ServeJSON()
        return 
    }

	var app_info AppJson
	var appname string

	if apptype == ".apk" {
		appname = fmt.Sprintf("android/%v%v", u1, apptype)
		stat, tmp := CheckApp.Adr("./App/" + appname)
		
		if stat {
			app_info.AppName = tmp.Name
			app_info.Version = tmp.Version
			app_info.OldName = filename
			app_info.MVersion = tmp.VCode
			app_info.Apptype = "android"
		}

		if tmp.Name == "com.handwriting.makefont" {
			app_info.Name = "手迹造字"
		} else if tmp.Name == "com.font" {
			app_info.Name = "写字先生"
		}
	} else {
		appname = fmt.Sprintf("IOS/%v%v", u1, apptype)
		fullName, _ := filepath.Abs(filepath.Dir("App/" + appname))
		fullName = strings.Replace(fullName, "\\", "/", -1)
		fullName = strings.Replace(fullName, "/IOS", "", -1)
		stat, tmps := CheckApp.IOS(fullName + "/" + appname)

		if stat {
			app_info.AppName = tmps.Name
			app_info.Version = tmps.Version
			app_info.OldName = filename
			app_info.MVersion = tmps.VCode
			app_info.Apptype = "ios"
		}

		if tmps.Name == "me.myfont.HandFontMaker" {
			app_info.Name = "手迹造字"
		} else if tmps.Name == "com.founder.MrWrite" {
			app_info.Name = "写字先生"
		}
	}


	err =this.SaveToFile("file", "App/" + appname)
	if err != nil{
		// log.Println(err.Error())
		this.Ctx.ResponseWriter.WriteHeader(423)
		this.Data["json"] = ErrAppSaveErr
		this.ServeJSON()
		return
	}


	app := models.App{
		Name:   app_info.AppName,
		AppName: app_info.Name,
		Version: app_info.Version,
		BuildNum: app_info.MVersion,
		Uuid: appname,
		Oldname: filename,
		Types :app_info.Apptype,
		Describe : "",
	}
	this.Data["json"] = Response{200, "success.", models.AddApps(app)}
	this.ServeJSON()
}


func (this *AppController) Applist() {
	// var app_info AppJson
	models.Applist()
	this.Ctx.ResponseWriter.WriteHeader(423)
	this.Data["json"] = ErrAppTypeErr
	this.ServeJSON()
    return 
}
package controllers

import (
	"fmt"
	"github.com/satori/go.uuid"
	"github.com/xuchengzhi/Library/AppInfo"
	"github.com/xuchengzhi/Library/Encryption"
	"github.com/xuchengzhi/apimonitor/models"
	"path"
	"path/filepath"
	"strings"
	// "github.com/astaxie/beego"
	"log"
)

var (
	ErrAppTypeErr  = ErrResponse{423001, "上传文件非app类型"}
	ErrAppSaveErr  = ErrResponse{423002, "文件保存失败"}
	ErrAppCheckErr = ErrResponse{423003, "解析App失败"}
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
	Img      string
}

func (this *AppController) TestApis() {
	this.Ctx.ResponseWriter.WriteHeader(200)
	this.Data["json"] = Response{200, "success.", "ok"}
	this.ServeJSON()
	return
}

func (this *AppController) FileUp() {
	file, head, err := this.GetFile("file")
	u1 := uuid.Must(uuid.NewV4())
	filename := head.Filename

	apptype := path.Ext(filename)
	log.Println(filename, file, apptype)
	if err != nil {
		log.Println(err.Error())
		this.Ctx.ResponseWriter.WriteHeader(423)
		this.Data["json"] = ErrResponse{403001, "msg"}
		this.ServeJSON()
		return
	}
	defer file.Close()
	var AllowExtMap map[string]bool = map[string]bool{
		".apk": true,
		".ipa": true,
		".IPA": true,
	}
	if _, ok := AllowExtMap[apptype]; !ok {
		this.Ctx.ResponseWriter.WriteHeader(423)
		this.Data["json"] = ErrAppTypeErr
		this.ServeJSON()
		return
	}

	var app_info AppJson
	var appname string

	if apptype == ".apk" {
		appname = fmt.Sprintf("android/%v%v", u1, apptype)
	} else {
		appname = fmt.Sprintf("IOS/%v%v", u1, apptype)
	}

	err = this.SaveToFile("file", "App/"+appname)
	if err != nil {
		// log.Println(err.Error())
		this.Ctx.ResponseWriter.WriteHeader(423)
		this.Data["json"] = ErrAppSaveErr
		this.ServeJSON()
		return
	}
	if apptype == ".apk" {
		stat, tmp := CheckApp.Adr("./App/" + appname)

		if stat {
			app_info.AppName = tmp.Name
			app_info.Version = tmp.Version
			app_info.OldName = filename
			app_info.MVersion = tmp.VCode
			app_info.Apptype = "android"
			app_info.Img = "/img/blob.png"
		}

		if tmp.Name == "com.handwriting.makefont" {
			app_info.Name = "手迹造字"
		} else if tmp.Name == "com.font" {
			app_info.Name = "写字先生"
		}
	}
	if apptype == ".ipa" {
		fullName, _ := filepath.Abs(filepath.Dir("App/" + appname))
		fullName = strings.Replace(fullName, "\\", "/", -1)
		fullName = strings.Replace(fullName, "/IOS", "", -1)
		stat, tmps := CheckApp.IOS(fullName + "/" + appname)
		log.Println(tmps)
		if stat {
			app_info.AppName = tmps.Name
			// app_info.Name = tmps.packageName
			app_info.Version = tmps.Version
			app_info.OldName = filename
			app_info.MVersion = tmps.VCode
			app_info.Apptype = "ios"
		} else {
			this.Ctx.ResponseWriter.WriteHeader(423)
			this.Data["json"] = ErrAppCheckErr
			this.ServeJSON()
			return
		}

		// if tmps.Name == "me.myfont.HandFontMaker" {
		// 	app_info.Name = "手迹造字"
		// } else if tmps.Name == "com.founder.MrWrite" {
		// 	app_info.Name = "写字先生"
		// }
	}
	u2 := fmt.Sprintf("%v", uuid.Must(uuid.NewV4()))
	stat, url := models.CheckUrl(app_info.AppName)
	var PageUrl string
	if stat {
		PageUrl = url
	} else {
		PageUrl = XorEnc.Gmd5(u2)
	}

	app := models.App{
		Name:     app_info.AppName,
		AppName:  app_info.Name,
		Version:  app_info.Version,
		BuildNum: app_info.MVersion,
		Uuid:     appname,
		Oldname:  filename,
		Types:    app_info.Apptype,
		Describe: "",
		Url:      PageUrl,
		Img:      app_info.Img,
	}

	this.Data["json"] = Response{200, "success.", models.AddApps(app)}
	this.ServeJSON()
}

func (this *AppController) Applist() {

	types := this.GetString("type")
	AppName := this.GetString("appname")
	// var app AppJson
	app := models.App{
		Name:    this.GetString("name"),
		AppName: AppName,
		Version: this.GetString("version"),
		Types:   types,
	}
	if types != "" && AppName != "" {
		this.Ctx.ResponseWriter.WriteHeader(200)
		this.Data["json"] = Response{200, "success.", models.AppAny(app)}
		this.ServeJSON()
		return
	} else if types != "" {
		this.Ctx.ResponseWriter.WriteHeader(200)
		this.Data["json"] = Response{200, "success.", models.AppByTypes(types)}
		this.ServeJSON()
		return
	} else if AppName != "" {
		this.Ctx.ResponseWriter.WriteHeader(200)
		this.Data["json"] = Response{200, "success.", models.AppByPackage(AppName)}
		// this.Data["json"] = Response{200, "success.", sss}
		this.ServeJSON()
		return
	} else {
		this.Ctx.ResponseWriter.WriteHeader(200)
		this.Data["json"] = Response{200, "success.", models.AppGroup()}
		this.ServeJSON()
		return
	}
}

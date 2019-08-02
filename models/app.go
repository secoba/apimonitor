package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"log"
	"time"
)

type App struct {
	Id       int64     `json:"id" orm:"column(id);pk;auto;unique"`
	Created  time.Time `json:"create_at" orm:"column(create_at);auto_now_add;type(datetime)"`
	Updated  time.Time `json:"-" orm:"column(update_at);auto_now;type(datetime)"`
	Name     string    `gorm:"Text(191)"`
	AppName  string    `gorm:"not null Text(191)"`
	Version  string    `gorm:"not null Text(191)"`
	BuildNum string    `gorm:"not null Text(191)"`
	Uuid     string    `gorm:"not null Text(191)"`
	Oldname  string    `gorm:"not null Text(191)"`
	Types    string    `gorm:"not null Text(191)"`
	Describe string    `gorm:"Text(191)"`
	Url      string    `gorm:"Text(191)"`
	Img      string    `gorm:"Text(191)"`
}

func (u *App) TableName() string {
	return TableName("apps")
}

func Apps() orm.QuerySeter {
	return orm.NewOrm().QueryTable(new(App))
}

func AddApps(app App) App {

	o := orm.NewOrm()
	st, err := o.Insert(&app)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(st)
	return app
}

type AppListJson struct {
	Count int
	List  interface{}
}

func CheckUrl(AppName string) (bool, string) {
	orm.Debug = true
	o := orm.NewOrm()
	var app App
	var msg bool
	err := o.Raw("SELECT * FROM api_apps WHERE name = ?", AppName).QueryRow(&app)
	if err != nil {
		log.Printf("err:%v", err.Error())
		msg = false
	}

	msg = false
	if app.Url == "" {
		msg = false
	} else {
		msg = true
	}

	return msg, app.Url
}

func AppGroup() AppListJson {
	o := orm.NewOrm()
	var apps AppListJson
	var tmp []App
	_, qs := o.QueryTable("api_apps").GroupBy("Name").All(&tmp)
	log.Println(tmp)
	if qs == nil {

		for _, App := range tmp {

			log.Println(App)
		}
	}
	apps.Count = len(tmp)
	apps.List = tmp
	return apps
}

func AppByPackage(AppName string) AppListJson {
	o := orm.NewOrm()
	var apps AppListJson
	var tmp []App
	_, qs := o.QueryTable("api_apps").Filter("name", AppName).All(&tmp)
	log.Println(tmp)
	if qs == nil {

		for _, App := range tmp {

			log.Println(App)
		}
	}
	apps.Count = len(tmp)
	apps.List = tmp
	return apps
}

func AppByTypes(Types string) AppListJson {
	orm.Debug = true
	o := orm.NewOrm()
	var apps AppListJson
	var tmp []App
	log.Printf("Types=%v", Types)
	_, qs := o.QueryTable("api_apps").Filter("types", Types).All(&tmp)
	if qs == nil {

		for _, App := range tmp {

			log.Println(App)
		}
	}

	apps.Count = len(tmp)
	apps.List = tmp
	return apps
}

func AppAny(app App) AppListJson {
	orm.Debug = true
	o := orm.NewOrm()
	var tmp []App
	var apps AppListJson
	_, qs := o.QueryTable("api_apps").Filter("name", app.AppName).All(&tmp)
	if qs == nil {

		for _, App := range tmp {

			log.Println(App)
		}
	}
	apps.List = tmp
	apps.Count = len(tmp)
	return apps
}

func Applist() AppListJson {
	o := orm.NewOrm()
	var apps AppListJson
	var tmp []App
	_, qs := o.QueryTable("api_apps").All(&tmp)
	if qs == nil {

		for _, App := range tmp {

			log.Println(App)
		}
	}
	apps.Count = len(tmp)
	apps.List = tmp
	return apps
}

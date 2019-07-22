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

func Applist() []App {
	o := orm.NewOrm()
	var apps []App
	_, qs := o.QueryTable("api_apps").All(&apps)

	if qs == nil {

		for _, App := range apps {

			log.Println(App)
		}
	}
	return apps
}

package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"time"
)

type Host struct {
	Id      int64     `json:"id" orm:"column(id);pk;auto;unique"`
	Created time.Time `json:"create_at" orm:"column(create_at);auto_now_add;type(datetime)"`
	Updated time.Time `json:"-" orm:"column(update_at);auto_now;type(datetime)"`
	Name    string
	Host    string
	Ip      string `gorm:"type:varchar(20);not null;"`
}

type Projects struct {
	Id          int64     `json:"id" orm:"column(id);pk;auto;unique"`
	Created     time.Time `json:"create_at" orm:"column(create_at);auto_now_add;type(datetime)"`
	Updated     time.Time `json:"-" orm:"column(update_at);auto_now;type(datetime)"`
	ProjectType string
	Description string
	Status      bool
	Name        string
	User        string
	Image       string
}

type Api struct {
	Id          int64     `json:"id" orm:"column(id);pk;auto;unique"`
	Created     time.Time `json:"create_at" orm:"column(create_at);auto_now_add;type(datetime)"`
	Updated     time.Time `json:"-" orm:"column(update_at);auto_now;type(datetime)"`
	ApiName     string    `json:"apiname" orm:"column(apiname);unique;size(40);unique_index:apiname_idx"`
	ApiPath     string
	ApiType     string
	ApiHost     string
	ApiMethod   string
	ApiTag      string
	ApiBuildNum string
}

type Params struct {
	Id      int64     `json:"id" orm:"column(id);pk;auto;unique"`
	Created time.Time `json:"create_at" orm:"column(create_at);auto_now_add;type(datetime)"`
	Updated time.Time `json:"-" orm:"column(update_at);auto_now;type(datetime)"`
	ApiId   int64
	Keys    string
	Values  string
}

func (u *Api) TableName() string {
	return TableName("apis")
}

func (u *Params) TableName() string {
	return TableName("params")
}

func init() {
	orm.RegisterModel(new(Api))
	orm.RegisterModel(new(Params))
}

func CreateApi(api Api) Api {

	o := orm.NewOrm()
	st, err := o.Insert(&api)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(st)
	return api
}

func (api *Api) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(api, fields...); err != nil {
		return err
	}
	return nil
}

func (api *Api) Delete() error {
	if _, err := orm.NewOrm().Delete(api); err != nil {
		return err
	}
	return nil
}

func ApiGetById(id int) (*Api, error) {
	u := new(Api)

	err := orm.NewOrm().QueryTable(TableName("apis")).Filter("id", id).One(u)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func ApiUpdate(api *Api, fields ...string) error {
	_, err := orm.NewOrm().Update(api, fields...)
	return err
}

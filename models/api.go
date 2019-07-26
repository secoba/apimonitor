package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	// "log"
	"time"
)

type Projects struct {
	Id          int64     `json:"id" orm:"column(id);pk;auto;unique"`
	Created     time.Time `json:"create_at" orm:"column(create_at);auto_now_add;type(datetime)"`
	Updated     time.Time `json:"-" orm:"column(update_at);auto_now;type(datetime)"`
	ProjectType string
	Host        []*Host `orm:"reverse(many)"`
	Api         []*Api  `orm:"reverse(many)"`
	Description string
	Status      bool
	Name        string
	User        *User `orm:"rel(fk)"`
	Image       string
}

type Host struct {
	Id       int64     `json:"id" orm:"column(id);pk;auto;unique"`
	Created  time.Time `json:"create_at" orm:"column(create_at);auto_now_add;type(datetime)"`
	Updated  time.Time `json:"-" orm:"column(update_at);auto_now;type(datetime)"`
	Name     string    `json:"name" orm:"column(name);unique;size(40);"`
	Project  *Projects `orm:"rel(fk)"`
	HostName string
	Ip       string `gorm:"type:varchar(20);not null;"`
}

type Api struct {
	Id          int64     `json:"id" orm:"column(id);pk;auto;unique"`
	Created     time.Time `json:"create_at" orm:"column(create_at);auto_now_add;type(datetime)"`
	Updated     time.Time `json:"-" orm:"column(update_at);auto_now;type(datetime)"`
	ApiName     string    `json:"apiname" orm:"column(apiname);unique;size(40);unique_index:apiname_idx"`
	Project     *Projects `orm:"rel(fk)"`
	ApiPath     string
	ApiType     string
	ApiHost     string
	ApiMethod   string
	ApiTag      string
	ApiBuildNum string
	User        *User `orm:"rel(fk)"`
}

type Params struct {
	Id      int64     `json:"id" orm:"column(id);pk;auto;unique"`
	Created time.Time `json:"create_at" orm:"column(create_at);auto_now_add;type(datetime)"`
	Updated time.Time `json:"-" orm:"column(update_at);auto_now;type(datetime)"`
	ApiId   *Api      `orm:"rel(fk)"`
	Keys    string
	Types   *ParamType `orm:"rel(fk)"`
	Values  string     //*ParamValue `orm:"rel(fk)"`
}

type XOR struct {
	string
}

type ParamValue struct {
	Param []*Params `orm:"reverse(many)"`
	Str   string
	XOR   []string
	TIMES bool
	MD5   []string
}

type ParamType struct {
	Id      int64     `json:"id" orm:"column(id);pk;auto;unique"`
	Created time.Time `json:"create_at" orm:"column(create_at);auto_now_add;type(datetime)"`
	Updated time.Time `json:"-" orm:"column(update_at);auto_now;type(datetime)"`
	Param   []*Params `orm:"reverse(many)"`
	Name    string
	Etype   string
}

func (u *ParamType) TableName() string {
	return TableName("paramtype")
}

func (u *Params) TableName() string {
	return TableName("params")
}

func (u *Projects) TableName() string {
	return TableName("project")
}

func (u *Api) TableName() string {
	return TableName("apis")
}

func init() {
	orm.RegisterModel(new(Host))
	orm.RegisterModel(new(Projects))
	orm.RegisterModel(new(Api))
	orm.RegisterModel(new(ParamType))
	orm.RegisterModel(new(Params))
	// orm.RegisterModel(new(ParamValue))
}

func CreateHost(host Host) Host {

	o := orm.NewOrm()
	st, err := o.Insert(&host)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(st)
	return host
}

func CreateProjects(project Projects) Projects {

	o := orm.NewOrm()
	st, err := o.Insert(&project)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(st)
	return project
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

func CreateParams(p Params) Params {

	o := orm.NewOrm()
	st, err := o.Insert(&p)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(st)
	return p
}

func CreateParamType(py ParamType) ParamType {

	o := orm.NewOrm()
	st, err := o.Insert(&py)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(st)
	return py
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

func ApiInit() {

	var project Projects
	var host Host
	var api Api
	var params Params
	var partype ParamType
	o := orm.NewOrm()
	u := new(User)       //用户
	p := new(Projects)   //项目
	a := new(Api)        //接口
	pt := new(ParamType) //参数类型id
	o.QueryTable(TableName("user")).Filter("id", 1).One(u)
	project.ProjectType = "App"
	project.Description = "手迹造字客户端"
	project.Status = true
	project.User = u
	project.Name = "手迹造字"
	project.Image = ""
	host.Name = "测试"
	host.HostName = "hwdev.xiezixiansheng.com"
	host.Ip = "192.168.248.249"
	CreateProjects(project)
	o.QueryTable(TableName("project")).Filter("id", 1).One(p)
	host.Project = p

	CreateHost(host)
	host.Name = "正式"
	host.HostName = "hw.xiezixiansheng.com"
	host.Ip = "123.1256.589"
	CreateHost(host)

	api.ApiName = "邮箱注册接口"
	api.Project = p
	api.ApiPath = "/mobile.php/Users/email_reg"
	api.ApiType = ""
	api.ApiMethod = "POST"
	api.ApiTag = "D1"
	api.ApiBuildNum = "1.0.0"
	api.User = u

	CreateApi(api)
	o.QueryTable(TableName("apis")).Filter("id", 1).One(a)
	partype.Name = "异或"
	partype.Etype = "XOR"
	CreateParamType(partype)
	partype.Name = "Md5"
	partype.Etype = "MD5"
	CreateParamType(partype)
	partype.Name = "字符串"
	partype.Etype = "String"
	CreateParamType(partype)

	o.QueryTable(TableName("paramtype")).Filter("id", 1).One(pt)
	params.ApiId = a
	params.Keys = "p"
	params.Types = pt
	CreateParams(params)
	o.QueryTable(TableName("paramtype")).Filter("id", 3).One(pt)
	params.ApiId = a
	params.Keys = "sys"
	CreateParams(params)
	o.QueryTable(TableName("paramtype")).Filter("id", 2).One(pt)
	params.ApiId = a
	params.Keys = "t"
	CreateParams(params)
	// params.Types =
}

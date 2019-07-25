package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"time"
)

type User struct {
	Id        int64  `json:"id" orm:"column(id);pk;auto;unique"`
	Phone     string `json:"phone" orm:"column(phone);unique;size(11)"`
	Nickname  string
	Password  string    `json:"-" orm:"column(password);size(40)"`
	Created   time.Time `json:"create_at" orm:"column(create_at);auto_now_add;type(datetime)"`
	Updated   time.Time `json:"-" orm:"column(update_at);auto_now;type(datetime)"`
	UserName  string    `json:"username" orm:"column(username);unique;size(40);"`
	Salt      string
	Email     string
	LastLogin int64
	LastIp    string
	Status    int
	Uimages   string `orm:"default(/static/usr/1.jpg)"`
}

func (u *User) TableName() string {
	return TableName("user")
}

func init() {
	orm.RegisterModel(new(User))
	orm.RegisterModel(new(App))
}

func Users() orm.QuerySeter {
	return orm.NewOrm().QueryTable(new(User))
}

// 检测手机号是否注册
func CheckUserPhone(phone string) bool {
	exist := Users().Filter("phone", phone).Exist()
	return exist
}

// 检测用户昵称是否存在
func CheckUserNickname(username string) bool {
	exist := Users().Filter("username", username).Exist()
	return exist
}

func UserGetByName(UserName string) (*User, error) {
	u := new(User)

	err := orm.NewOrm().QueryTable(TableName("user")).Filter("UserName", UserName).One(u)
	if err != nil {
		return nil, err
	}
	return u, nil
}

//创建用户
func CreateUser(user User) User {

	o := orm.NewOrm()
	st, err := o.Insert(&user)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(st)
	return user
}

//检测手机和昵称是否注册
func CheckUserPhoneOrNickname(phone string, username string) bool {
	cond := orm.NewCondition()
	count, _ := Users().SetCond(cond.And("phone", phone).Or("username", username)).Count()
	if count <= int64(0) {
		return false
	}
	return true
}
func CheckUserAuth(username string, password string) (User, bool) {
	o := orm.NewOrm()
	user := User{
		UserName: username,
		Password: password,
	}
	// fmt.Println(user)
	err := o.Read(&user, "UserName", "Password")
	// fmt.Println(err)
	if err != nil {
		return user, false
	}
	return user, true
}

// User database CRUD methods include Insert, Read, Update and Delete
func (usr *User) Insert() error {
	if _, err := orm.NewOrm().Insert(usr); err != nil {
		return err
	}
	return nil
}

func (usr *User) Read(fields ...string) error {
	if err := orm.NewOrm().Read(usr, fields...); err != nil {
		return err
	}
	return nil
}

func (usr *User) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(usr, fields...); err != nil {
		return err
	}
	return nil
}

func (usr *User) Delete() error {
	if _, err := orm.NewOrm().Delete(usr); err != nil {
		return err
	}
	return nil
}

func UserGetById(id int) (*User, error) {
	u := new(User)

	err := orm.NewOrm().QueryTable(TableName("user")).Filter("id", id).One(u)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func UserUpdate(user *User, fields ...string) error {
	_, err := orm.NewOrm().Update(user, fields...)
	return err
}

func UserInit() {
	var userinfo User
	userinfo.Phone = "13041195556"
	userinfo.UserName = "xuchengzhi"
	userinfo.Nickname = "许成志"
	userinfo.Password = "e10adc3949ba59abbe56e057f20f883e"
	CreateUser(userinfo)
}

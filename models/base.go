package models

import (
	"fmt"
	"github.com/astaxie/beego/config"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
	"strings"
	// "github.com/pelletier/go-toml"
	"flag"
)

var (
	h    bool
	test string
)

func Formats(f *flag.Flag) string {
	res := ""
	if f != nil {
		res = fmt.Sprintf("%v", f.Value)
	} else {
		fmt.Println(nil)
	}
	return res
}

// 数据库连接初始化
func Init() {
	flag.String("runmode", "test", "运行环境")
	flag.BoolVar(&h, "h", false, "this help")
	flag.Usage = usage
	appConf, err := config.NewConfig("ini", "conf/app.conf")
	if err != nil {
		panic(err)
	}
	flag.Parse()
	msg := Formats(flag.Lookup("runmode"))

	if msg == "test" {
		log.Println("连接本地SqlLite数据库")
		// orm.RegisterDriver("sqlite", orm.DR_Sqlite)
		orm.RegisterDataBase("default", "sqlite3", "./datas/test.db")
		orm.RunSyncdb("default", false, true)

	} else {
		log.Printf("连接MySQL数据库%s\n", appConf.String("database::db_host"))
		conn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s",
			appConf.String("database::db_user"),
			appConf.String("database::db_passwd"),
			appConf.String("database::db_host"),
			appConf.String("database::db_port"),
			appConf.String("database::db_name"),
			appConf.String("database::db_charset"))
		orm.RegisterDataBase("default", "mysql", conn)
	}

	//自动建表
	name := "default"
	force := false
	verbose := true
	err = orm.RunSyncdb(name, force, verbose)
	if err != nil {
		fmt.Println(err)
	}
	orm.RunCommand()

}

//返回带前缀的表名
func TableName(str string) string {
	appConf, err := config.NewConfig("ini", "conf/app.conf")
	if err != nil {
		panic(err)
	}
	return appConf.String("database::db_prefix") + str
}

func getAppEnv() string {
	file := "test" //os.Args[0]
	s := strings.Split(file, ".")
	return s[len(s)-1]
}

func usage() {
	fmt.Fprintf(os.Stderr, `App version: Test/1.10.0
Usage: nginx [-hvVtTq] [-s signal] [-c filename] [-p prefix] [-g directives]

Options:
`)
	flag.PrintDefaults()
}

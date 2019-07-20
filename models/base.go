package models

import (
	"fmt"
	"github.com/astaxie/beego/config"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"strings"
	_ "github.com/mattn/go-sqlite3"
	// "github.com/pelletier/go-toml"
)



// 数据库连接初始化
func Init() {

	appConf, err := config.NewConfig("ini", "conf/app.conf")
	if err != nil {
		panic(err)
	}

	
	if getAppEnv() == "test" {
		// orm.RegisterDriver("sqlite", orm.DR_Sqlite)
    	orm.RegisterDataBase("default", "sqlite3", "./datas/test.db")
    	orm.RunSyncdb("default", false, true)

	}else{
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
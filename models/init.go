package models

import (
	"fmt"
	"net/url"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

//Init is used for database init.
func Init() {
	dbhost := beego.AppConfig.String("db.host")
	dbport := beego.AppConfig.String("db.port")
	dbuser := beego.AppConfig.String("db.user")
	dbpassword := beego.AppConfig.String("db.password")
	dbname := beego.AppConfig.String("db.name")
	timezone := beego.AppConfig.String("db.timezone")
	if dbport == "" {
		dbport = "3306"
	}
	dsn := dbuser + ":" + dbpassword + "@tcp(" + dbhost + ":" + dbport + ")/" + dbname + "?charset=utf8"
	fmt.Println(dsn)

	if timezone != "" {
		dsn = dsn + "&loc=" + url.QueryEscape(timezone)
	}

	//必须在初始化时注册一个默认的数据库
	//这个只是在初始化的时候要注册使用的  当你真正需要使用到自己的数据库的时候
	//在model里面 去Using 一个你自己的库  并且定义成全局变量方便在别的文件中使用
	orm.RegisterDataBase("default", "mysql", dsn)

	orm.RegisterModel(new(Auth), new(Role), new(User), new(Device),
		new(Group), new(DeviceKey), new(Versions))

	orm.RunSyncdb("default", false, true)

	if beego.AppConfig.String("runmode") == "dev" {
		orm.Debug = true
	}
}

//TableName ,Base function
func TableName(name string) string {
	return beego.AppConfig.String("db.prefix") + name
}

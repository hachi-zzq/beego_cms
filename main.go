package main

import (
	_ "cms/routers"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego"
)

func init(){
	beego.BConfig.WebConfig.Session.SessionOn = true
	initMysql()
}

func initMysql()  {
	mysqlUser := beego.AppConfig.String("mysqluser")
	mysqlPassword := beego.AppConfig.String("mysqlpass")
	mysqlPort := beego.AppConfig.String("mysqlport")
	mysqlHost := beego.AppConfig.String("mysqlhost")
	mysqlDb := beego.AppConfig.String("mysqldb")

	orm.RegisterDataBase("default", "mysql", mysqlUser+":"+mysqlPassword+"@tcp("+mysqlHost+":"+mysqlPort+")/"+mysqlDb+"?charset=utf8&loc=Asia%2FShanghai", 30)
}

func main() {
	beego.Run()

}


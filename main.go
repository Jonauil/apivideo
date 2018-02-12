package main

import (
	_ "apivideo/routers"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"

	"apivideo/models"
)

func init(){
	//Dsn := beego.AppConfig.String("username")+":"+beego.AppConfig.String("password")+"@tcp("+beego.AppConfig.String("host")+":"+beego.AppConfig.String("port")+")/"+beego.AppConfig.String("dbname")+"?charset=utf8&loc=Asia%2FShanghai"
	dbuser := beego.AppConfig.String("username")
	dbpassword := beego.AppConfig.String("password")
	dbhost := beego.AppConfig.String("host")
	dbport := beego.AppConfig.String("port")
	db := beego.AppConfig.String("dbname")
	//注册mysql Driver
	orm.RegisterDriver("mysql",orm.DRMySQL)
	//构造conn连接
	conn := dbuser + ":" + dbpassword + "@tcp(" +dbhost + ":" + dbport + ")/" + db + "?charset=utf8&loc=Asia%2FShanghai"
	//注册数据库连接
	orm.RegisterDataBase("default","mysql",conn)
	orm.RegisterModel(new(models.Info))
	orm.RunSyncdb("default",false,true)
    //orm.RegisterModel(new(models.Info))
	fmt.Printf("数据库连接成功! %s\n",conn)
}
func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}

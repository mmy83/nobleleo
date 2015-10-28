package main

import (
	"fmt"
	_ "github.com/mmy83/nobleleo/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func init(){
	db_host :=beego.AppConfig.String("db_host")
	db_user :=beego.AppConfig.String("db_user")
	db_pass :=beego.AppConfig.String("db_pass")
	db_name :=beego.AppConfig.String("db_name")
	db_port :=beego.AppConfig.String("db_port")
	db_charset :=beego.AppConfig.String("db_charset")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s",db_user,db_pass,db_host,db_port,db_name,db_charset)
	orm.RegisterDataBase("default","mysql",dsn)
	orm.RunSyncdb("default",false,true)
}

func main() {
	beego.Run()
}


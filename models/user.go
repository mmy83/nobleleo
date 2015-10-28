package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type User struct {
	Id       int64  `orm:"auto"`
	Username string `orm:"size(100)"`
	Password string `orm:size(100)`
}

func init() {
	db_prefix := beego.AppConfig.String("db_prefix")
	orm.RegisterModelWithPrefix(db_prefix,new(User))
}

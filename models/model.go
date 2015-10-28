package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type User struct {
	Id       int64  `orm:"auto"`
	Username string `orm:"size(100);index"`
	Password string `orm:"size(100)"`
}


type Content struct{
	Id int64 `orm:"auto"`
	Title string `orm:"size(255)"`
	Filepath string `orm:"size(500)"`
	Content string `orm:"type(text)"`
}





func init() {
	db_prefix := beego.AppConfig.String("db_prefix")
	orm.RegisterModelWithPrefix(db_prefix,new(User),new(Content))
}
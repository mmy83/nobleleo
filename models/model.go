package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type User struct {
	Id       int64      `orm:"auto;pk"`
	Username string     `orm:"size(100);index;unique"`
	Password string     `orm:"size(100)"`
	Contents []*Content `orm:"reverse(many)"`
}

func (u *User) TableName() string {
	return "users"
}

type Content struct {
	Id       int64     `orm:"auto;pk"`
	Title    string    `orm:"size(255)"`
	Filepath string    `orm:"size(500)"`
	Content  string    `orm:"type(text);null"`
	User     *User     `orm:"rel(fk)"`
	Company  *Company  `orm:"rel(fk)"`
	Category *Category `orm:"rel(fk)"`
}

func (c *Content) TableName() string {
	return "contents"
}

type Company struct {
	Id       int64      `orm:"auto;pk" form:"-"`
	Symbol   string     `orm:"size(6)"  form:"Symbol"   valid:"Required;Length(6)"`
	Sname    string     `orm:"size(20)"  form:"Sname"  varlid:"Required"`
	Contents []*Content `orm:"reverse(many)"  form:"-"`
}

func (company *Company) TableName() string {
	return "companys"
}

type Category struct {
	Id       int64  `orm:"auto;pk"`
	Name     string `orm:"size(100)"`
	Pid      int64
	Contents []*Content `orm:"reverse(many)"`
}

func (category *Category) TableName() string {
	return "categorys"
}

func init() {
	db_prefix := beego.AppConfig.String("db_prefix")
	orm.RegisterModelWithPrefix(db_prefix, new(User), new(Content), new(Company), new(Category))
}

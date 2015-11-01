package admin

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/mmy83/nobleleo/models"
)

type CompanyController struct {
	beego.Controller
}

func (c *CompanyController) Create() {
	c.TplNames = "admin/company/create.tpl"
}

func (this *CompanyController) Store() {
	company := &models.Company{}
	company.Symbol = this.GetString("Symbol")
	company.Sname = this.GetString("Sname")
	o := orm.NewOrm()
	o.Insert(company)
	this.Redirect("/admin/company",302)
}

func (this *CompanyController) Index(){
	this.TplNames="index.tpl"
}

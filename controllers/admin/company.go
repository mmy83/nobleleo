package admin

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/utils/pagination"
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
	this.Redirect("/admin/company", 302)
}

func (this *CompanyController) Index() {
	perPage := 2
	//paginator := pagination.SetPaginator(this.Ctx,perPage,int64(20))
	var companys []models.Company
	o := orm.NewOrm()
	company := new(models.Company)
	qs := o.QueryTable(company)
	//qs.Filter("","")
	nums, _ := qs.Count()
	//nums := int64(10)
	fmt.Print(nums, companys)
	beego.Debug(nums)
	
	paginator := pagination.SetPaginator(this.Ctx,perPage,nums)
	qs = qs.Limit(perPage,paginator.Offset())
	
	beego.Debug(paginator.Offset())
	qs.All(&companys)
	this.Data["companys"] = companys
	this.Data["paginator"] = paginator
	this.TplNames = "admin/company/index.tpl"
}

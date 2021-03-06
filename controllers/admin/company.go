package admin

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/utils/pagination"
	"github.com/astaxie/beego/validation"
	"github.com/mmy83/nobleleo/models"
	"strconv"
)

type CompanyController struct {
	beego.Controller
}

func (c *CompanyController) Create() {
	c.TplNames = "admin/company/create.tpl"
}

func (this *CompanyController) Store() {
	company := &models.Company{}
	//company.Symbol = this.GetString("Symbol")
	//company.Sname = this.GetString("Sname")
	this.ParseForm(company)
	valid := validation.Validation{}
	b, err := valid.Valid(company)
	if err != nil {
		// handle error
		beego.Debug(err)
	}
	if !b {
		// validation does not pass
		// blabla...
		for _, err := range valid.Errors {
			beego.Debug(err.Message)
		}
	}
	o := orm.NewOrm()
	o.Insert(company)
	this.Redirect(beego.UrlFor("CompanyController.Index"), 302)
}

func (this *CompanyController) Index() {
	perPage := 20
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

	paginator := pagination.SetPaginator(this.Ctx, perPage, nums)
	qs = qs.Limit(perPage, paginator.Offset())

	beego.Debug(paginator.Offset())
	qs.All(&companys)
	this.Data["companys"] = companys
	this.Data["paginator"] = paginator
	this.TplNames = "admin/company/index.tpl"
}

func (this *CompanyController) Edit() {
	o := orm.NewOrm()
	id, _ := strconv.Atoi(this.GetString("id"))
	company := models.Company{Id: int64(id)}
	error := o.Read(&company)
	if error != nil {
		fmt.Println(error)
	}
	this.Data["company"] = company
	this.TplNames = "admin/company/edit.tpl"

}

func (this *CompanyController) Update() {

	o := orm.NewOrm()
	id, _ := strconv.Atoi(this.GetString("Id"))
	company := models.Company{Id: int64(id)}
	error := o.Read(&company)
	beego.Debug(error)
	if error == nil {
		beego.Debug("start")
		//company.Sname = this.GetString("Sname")
		//company.Symbol = this.GetString("Symbol")
		this.ParseForm(&company)
		_, error_u := o.Update(&company, "sname", "symbol")
		if error_u != nil {
			beego.Error(error_u)
		} else {
			this.Redirect(beego.UrlFor("CompanyController.Index"), 302)
		}
	} else {
		beego.Error(error)
		this.Redirect(beego.UrlFor("CompanyController.Edit", "id", strconv.Itoa(id)), 302)
	}

}

func (this *CompanyController) Del() {
	id, _ := strconv.Atoi(this.Ctx.Input.Param(":id"))
	beego.Debug("id=", id)
	o := orm.NewOrm()
	if num, err := o.Delete(&models.Company{Id: int64(id)}); err == nil {
		beego.Debug(num)
		this.Redirect(beego.UrlFor("CompanyController.Index"), 302)
	} else {
		beego.Error(err)
	}
}

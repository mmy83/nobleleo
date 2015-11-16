package admin

import (
	//"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/mmy83/nobleleo/models"
	"github.com/astaxie/beego/utils/pagination"
)

type CategoryController struct {
	beego.Controller
}

func (this *CategoryController) Index() {
	perPage := 2
	o := orm.NewOrm()
	var categorys []models.Category
	category := new(models.Category)
	qs := o.QueryTable(category)
	num,_ := qs.Count()
	paginator := pagination.SetPaginator(this.Ctx , perPage ,num)
	qs = qs.Limit(perPage,paginator.Offset())
	qs.All(&categorys)
	this.Data["categorys"] = categorys
	this.Data["paginator"] = paginator
	this.TplNames = "admin/category/index.tpl"
}

func (this *CategoryController) Create(){
	this.TplNames = "admin/category/create.tpl"
}

func (this *CategoryController) Store(){
	
}

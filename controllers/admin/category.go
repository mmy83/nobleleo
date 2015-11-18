package admin

import (
	//"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/mmy83/nobleleo/models"
	"github.com/astaxie/beego/utils/pagination"
	"strconv"
)

type CategoryController struct {
	beego.Controller
}

func (this *CategoryController) Index() {
	perPage := 1000
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
	pid ,_ := strconv.Atoi(this.GetString("pid","0"))
	o := orm.NewOrm()
	category := models.Category{Id:int64(pid)}
	//category.Id = int64(pid)
	err := o.Read(&category)
	if err!=nil{
		beego.Error(err)
	}
	this.Data["category"] = category
	this.TplNames = "admin/category/create.tpl"
}

func (this *CategoryController) Store(){
	pid, _:= strconv.Atoi(this.GetString("Pid"))
	name := this.GetString("Name")
	o := orm.NewOrm()
	category := new(models.Category)
	category.Pid = int64(pid)
	category.Name = name
	//category := &models.Category{Pid:int64(pid),Name:name}
	o.Insert(category)
	this.Redirect(beego.UrlFor("CategoryController.Index"),302)
	
}

func (this * CategoryController)Edit(){
	id,_ :=strconv.Atoi(this.GetString("id"))
	o := orm.NewOrm()
	category := models.Category{Id:int64(id)}
	o.Read(&category)
	this.Data["category"] = category
	this.TplNames = "admin/category/edit.tpl"
}
 
func (this * CategoryController)Update(){
	id,_ :=strconv.Atoi(this.GetString("Id"))
	name := this.GetString("Name")
	o := orm.NewOrm()
	category := models.Category{Id:int64(id)}
	error := o.Read(&category)
	if error!=nil{
		beego.Error("error=",error)
	}
	beego.Debug("category=",category)
	category.Name = name
	beego.Debug("category=",category)
	o.Update(&category,"name")
	this.Redirect(beego.UrlFor("CategoryController.Index"),302)
} 

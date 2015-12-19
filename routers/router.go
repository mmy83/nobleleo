package routers

import (
	"github.com/astaxie/beego"
	"github.com/mmy83/nobleleo/controllers"
	"github.com/mmy83/nobleleo/controllers/admin"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/admin/welcome", &admin.HomeController{}, "get:Welcome")

	beego.Router("/admin/company/create", &admin.CompanyController{}, "get:Create;post:Store")
	beego.Router("/admin/company",&admin.CompanyController{},"get:Index")
	beego.Router("/admin/company/edit",&admin.CompanyController{},"get:Edit;post:Update")
	beego.Router("/admin/company/del/:id",&admin.CompanyController{},"get:Del")
	
	//category
	beego.Router("/admin/category",&admin.CategoryController{},"get:Index")
	beego.Router("/admin/category/create", &admin.CategoryController{}, "get:Create;post:Store")
	beego.Router("/admin/category/edit",&admin.CategoryController{},"get:Edit;post:Update")
	
	//content
	beego.Router("/admin/content",&admin.ContentController{},"get:Index")
	beego.Router("/admin/content/create", &admin.ContentController{}, "get:Create;post:Store")
	beego.Router("/admin/content/edit",&admin.ContentController{},"get:Edit;post:Update")
	
	//home
	beego.Router("/admin/home",&admin.HomeController{},"get:Welcome")
	beego.Router("/admin/home/welcome",&admin.HomeController{},"get:Welcome")
}

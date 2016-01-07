package admin

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/mmy83/nobleleo/models"
	"os"
	"time"
)

type ContentController struct {
	beego.Controller
}

func (this *ContentController) Index() {
	var contents []models.Content
	o := orm.NewOrm()
	qs := o.QueryTable(new(models.Content)).RelatedSel()
	qs.All(&contents)

	beego.Debug(contents)

	this.Data["contents"] = contents
	this.TplNames = "admin/content/index.tpl"
}

func (this *ContentController) Create() {
	o := orm.NewOrm()
	var categorys []models.Category
	category := new(models.Category)
	qs := o.QueryTable(category)
	qs.All(&categorys)
	this.Data["categorys"] = categorys
	var companys []models.Company
	qs = o.QueryTable(new(models.Company))
	qs.All(&companys)
	this.Data["companys"] = companys
	this.TplNames = "admin/content/create.tpl"
}

func (this *ContentController) Store() {
	o := orm.NewOrm()

	f, h, _ := this.GetFile("Filepath")
	defer f.Close()

	newdir := "./static/upload/" + time.Now().Format("20151121")
	beego.Debug(newdir)
	os.MkdirAll(newdir, 0777)
	newfile := newdir + "/" + h.Filename

	this.SaveToFile("Filepath", newfile)

	company_id, _ := this.GetInt64("company_id")
	company := models.Company{Id: company_id}
	o.Read(&company)
	beego.Debug(company)
	category_id, _ := this.GetInt64("category_id")
	category := models.Category{Id: category_id}
	o.Read(&category)
	beego.Debug(category)

	user_id := int64(1)
	user := models.User{Id: user_id}
	o.Read(&user)
	beego.Debug(user)

	title := this.GetString("Title")
	Content := this.GetString("Content")

	content := &models.Content{}
	content.Title = title
	content.Content = Content
	content.Category = &category
	content.Company = &company
	content.Filepath = newfile
	content.User = &user

	beego.Debug(content)

	_, e := o.Insert(content)
	if e != nil {
		beego.Error(e)
	}
	this.Redirect(beego.UrlFor("ContentController.Index"), 302)

}

func (this *ContentController) Edit() {

}
func (this *ContentController) Update() {

}

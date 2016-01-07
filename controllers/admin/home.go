package admin

import (
	"github.com/astaxie/beego"
)

type HomeController struct {
	beego.Controller
}

func (c *HomeController) Welcome() {
	//c.Layout = "admin/layout/main.tpl"
	c.TplNames = "admin/home/index.tpl"
}

func (c *HomeController) Login() {
	c.TplNames = "admin/home/login.tpl"
}

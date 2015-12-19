package admin

import (
	"github.com/astaxie/beego"
)

type HomeController struct {
	beego.Controller
}

func (c *HomeController) Welcome() {
	c.TplNames = "admin/home/index.tpl"
}

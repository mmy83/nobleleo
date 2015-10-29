package admin

import (
	"github.com/astaxie/beego"
)

type ContentController struct{
	beego.Controller
}

func (c *ContentController) Get(){
	c.TplNames = ""
}
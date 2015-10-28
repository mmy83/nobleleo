package routers

import (
	"github.com/astaxie/beego"
	"github.com/mmy83/nobleleo/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
}

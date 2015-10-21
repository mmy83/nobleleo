package routers

import (
	"github.com/mmy83/nobleleo/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}

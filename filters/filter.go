package filters

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

var FilterUser = func(ctx *context.Context) {
	beego.Info("filter print")
	_, ok := ctx.Input.Session("uid").(int)
	if !ok && ctx.Request.RequestURI != "/admin/login" {
		ctx.Redirect(302, "/admin/login")
	}
}

func init() {
	//beego.InsertFilter("/admin/*",beego.BeforeRouter,FilterUser)
}

package routers

import (
	"github.com/astaxie/beego"
	"xx/controllers/user"
)

func init() {
	ns := beego.NewNamespace("/v1", beego.NSAutoRouter(&user.UserController{}))
	beego.AddNamespace(ns)
	// beego.Router("/", &controllers.MainController{})
}

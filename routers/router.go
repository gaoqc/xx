package routers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"xx/controllers"
	"xx/filters"
	"xx/models"
)

func init() {
	ns := beego.NewNamespace("/v1", beego.NSBefore(filters.LoginFilter), beego.NSAutoRouter(&controllers.UserController{}))

	beego.AddNamespace(ns)
	// beego.Router("/", &controllers.MainController{})
	initDb()
	initLog()
	initFilter()
}
func initDb() {

	orm.RegisterDataBase("default", "mysql", "root:gaoqc@123.com@/xx?charset=utf8")
	orm.RegisterModelWithPrefix("t_", new(models.User), new(models.UserAddress))
	orm.Debug = true
	err := orm.RunSyncdb("default", false, true)
	if err != nil {
		fmt.Errorf("err on RunSyncdb:%v", nil)
	}
}
func initLog() {
	beego.SetLogger("file", `{"filename":"/Users/gaoqc/Documents/codes/go/src/xx/logs/xx.log"}`)
	beego.SetLevel(beego.LevelDebug)
}
func initFilter() {

	beego.InsertFilter("/*", beego.BeforeRouter, filters.PrintURL)
}

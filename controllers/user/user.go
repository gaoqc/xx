package user

import (
	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
}

func (c *UserController) Register() {

}
func (c *UserController) Login() {
	loginAcc := c.GetString("loginAcc")
	loginPwd := c.GetString("loginPwd")
	c.Ctx.WriteString("loginAcc:" + loginAcc + ",loginPwd:" + loginPwd)

}

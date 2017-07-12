package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
	"strconv"
	"xx/models"
	"xx/utils"
)

type UserForm struct {
	LoginAcc string `form:"loginAcc" ;valid:"Required"`
	LoginPwd string `form:"loginPwd";valid:"Required"`
	Phone    string `form:"phone";valid:"Mobile"`
	Email    string `form:"email";valid:"Email; MaxSize(100)"`
	TrueName string `form:trueName`
	IdNo     string `form:idNo`
	IdType   int    `form:idType`
}

type UserController struct {
	beego.Controller
}

//入参验证
func (c *UserForm) Valid(v *validation.Validation) {

}

func (c *UserController) Update() {
	idType, _ := c.GetInt("idType")
	num := models.Update(c.GetString("phone"), c.GetString("email"), c.GetString("idNo"), c.GetString("trueName"), c.GetString("loginAcc"), idType)
	if num > 0 {
		c.Ctx.WriteString("update ok!")
	} else {
		c.Ctx.WriteString("update fail,noChange")
	}

}

//注册新用户
func (c *UserController) Register() {

	//校验入参
	var userForm UserForm
	c.ParseForm(&userForm)

	valid := validation.Validation{}
	b, _ := valid.Valid(&userForm)
	var msg = ""
	if !b {
		// validation does not pass
		// blabla...
		for _, err := range valid.Errors {
			msg = msg + err.Key + ":" + err.Message + "\n\r"
		}
		c.Ctx.WriteString(msg)
		return
	}
	//判断是否存在
	exists := models.ExistUser(userForm.LoginAcc, userForm.IdNo, userForm.Phone, userForm.Email)
	if exists {
		msg = " user is areadly exists!"
		c.Ctx.WriteString(msg)
		return
	}

	//新增记录
	id := models.AddUser(&models.User{LoginAcc: userForm.LoginAcc, LoginPwd: userForm.LoginPwd, Phone: userForm.Phone, Email: userForm.Email, TrueName: userForm.TrueName, IdNo: userForm.IdNo, IdType: userForm.IdType})
	msg = "id =" + strconv.FormatInt(id, 10)
	c.Ctx.WriteString(msg)

}
func (c *UserController) Login() {
	loginAcc := c.GetString("loginAcc")
	loginPwd := c.GetString("loginPwd")
	exist, user := models.Login(loginAcc, loginPwd)
	if exist {
		c.Ctx.WriteString(utils.ToJson(user))
	} else {
		c.Ctx.WriteString("user." + loginAcc + " is not exists!")
	}

}

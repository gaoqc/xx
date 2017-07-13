package filters

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"strings"
	"xx/utils"
	"xx/vo"
)

var PrintURL = func(ctx *context.Context) {
	uri := ctx.Input.URI()
	beego.Info("uri:" + uri)
	// for k, v := range ctx.Input.Params() {
	// 	beego.Info(k + "=" + v)
	// }
}

func loginFilterUri() []string {
	return []string{"/v1/user/update"}
}

var LoginFilter = func(ctx *context.Context) {
	uri := ctx.Input.URI()
	ticket := ctx.Input.Session("ticket")
	beego.Debug("uri is" + uri + " ,ticket is :" + utils.ToJson(ticket))
	for _, f := range loginFilterUri() {
		if strings.HasPrefix(uri, f) {
			beego.Debug("uri:" + uri + ",patter:" + f)
			if ticket == nil {
				beego.Warn("uri:" + uri + " not login")
				ctx.Output.JSON(vo.GetRetVO(vo.NoLoginCode, vo.NoLoginMsg, nil), false, false)
			}
			break
		}
	}
}

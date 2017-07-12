package filters

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

var PrintURL = func(ctx *context.Context) {
	uri := ctx.Input.URI()
	beego.Info("uri:" + uri)
	// for k, v := range ctx.Input.Params() {
	// 	beego.Info(k + "=" + v)
	// }
}

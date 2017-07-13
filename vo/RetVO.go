package vo

type RetVO struct {
	Code, Msg string
	Data      interface{}
}

const (
	SuccCode         = "000000"
	SuccMsg          = "success"
	LoginFailCode    = "000100"
	LoginFailMsg     = "登陆失败"
	RegistExistsCode = "000101"
	RegistExistsMsg  = "用户已注册"
	NoLoginCode      = "000102"
	NoLoginMsg       = "未登录"
)

func SuccessVO(d interface{}) RetVO {
	return RetVO{Code: SuccCode, Msg: SuccMsg, Data: d}
}
func GetRetVO(code, msg string, d interface{}) RetVO {
	return RetVO{Code: code, Msg: msg, Data: d}
}

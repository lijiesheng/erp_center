package resp

type ResCode int

const (
	CodeSuccess         ResCode = 10000 + iota // 成功
	CodeInvalidParam                           // 参数不对
	CodeUserExist                              // 用户已经存在
	CodeUserNotExist                           // 用户不存在
	CodeInvalidPassword                        // 密码不正确
	CodeServerBusy                             // 服务繁忙
	CodeNeedLogin                              // 先登录
	CodeInvalidToken                           // token 不对
)

var codeMsgMap = map[ResCode]string{
	CodeSuccess:         "success",
	CodeInvalidParam:    "请求参数错误",
	CodeUserExist:       "用户名已存在",
	CodeUserNotExist:    "用户名不存在",
	CodeInvalidPassword: "用户名或密码错误",
	CodeServerBusy:      "服务繁忙",

	CodeNeedLogin:    "需要登录",
	CodeInvalidToken: "无效的token",
}

func (code ResCode) GetMsg() string {
	msg, ok := codeMsgMap[code]
	if !ok {
		msg = codeMsgMap[CodeServerBusy]
	}
	return msg
}

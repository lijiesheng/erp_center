package request

import (
	"errors"
	"github.com/gin-gonic/gin"
)

const CtxTlAccount = "tlAccount"
const ID = "ID"

const Username = "username"

var ErrorUserNotLogin = errors.New("用户没有登录")

// 获取当前的天狼账号 或者 ID
func GetCurrent(c *gin.Context, param string) (str any, err error) {
	value, exists := c.Get(param)
	if !exists {
		return "", ErrorUserNotLogin
	}
	return value, nil
}

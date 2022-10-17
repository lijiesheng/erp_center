package middlewares

import (
	"erp_center/controller/request"
	"erp_center/controller/resp"
	"erp_center/pkg/jwt"
	"fmt"
	"github.com/gin-gonic/gin"
)

func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		tok := c.Query("tok")
		if tok != "" {
			token, err := jwt.ParseToken(tok)
			if err != nil {
				fmt.Println("err ==>", err)
				resp.ResponseErrorJsonpJSON(c, resp.CodeInvalidToken, resp.CodeInvalidToken.GetMsg(), nil)
				c.Abort() // 退出
				return
			}
			// 将当前的请求的 tl_account 信息保存到请求的上下文 c 上
			c.Set(request.CtxTlAccount, token.TlAccount)
			c.Set(request.ID, token.ID)
			c.Next()
		} else {
			resp.ResponseErrorJsonpJSON(c, resp.CodeNeedLogin, resp.CodeNeedLogin.GetMsg(), nil)
			c.Abort()
			return
		}
	}
}

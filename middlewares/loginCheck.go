package middlewares

import "github.com/gin-gonic/gin"

func SesToken(is_need_login bool) func(c *gin.Context) {
	return func(c *gin.Context) {
		tok := c.Query("tok")
		// 将每个请求都带上 token
		if tok != "" {
			//req_ip := util.RealIP(c)
			// 记录 某个 IP 地址 token 的生成频率

		}
		// 请求过来的SID先记录下来
		//c.Set("_ReqSID")
	}
}

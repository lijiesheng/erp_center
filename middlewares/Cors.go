package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//func Cors(context *gin.Context) {
//	method := context.Request.Method
//	fmt.Println("请求的方式是，", method)
//	// 必须，接受指定域的请求，可以使用*不加以限制，但不安全
//	context.Header("Access-Control-Allow-Origin", "*")
//	context.Header("Access-Control-Allow-Headers", "x-requested-with,content-type")
//	//context.Header("Access-Control-Allow-Origin", context.GetHeader("Origin"))
//	fmt.Println(context.GetHeader("Origin"))
//	// 必须，设置服务器支持的所有跨域请求的方法
//	context.Header("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE, OPTIONS")
//	// 服务器支持的所有头信息字段，不限于浏览器在"预检"中请求的字段
//	context.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Token")
//	// 可选，设置XMLHttpRequest的响应对象能拿到的额外字段
//	context.Header("Access-Control-Expose-Headers", "Access-Control-Allow-Headers, Token")
//	// 可选，是否允许后续请求携带认证信息Cookir，该值只能是true，不需要则不设置
//	context.Header("Access-Control-Allow-Credentials", "true")
//	// 放行所有OPTIONS方法
//	if method == "OPTIONS" {
//		context.AbortWithStatus(http.StatusNoContent)
//		return
//	}
//	context.Next()
//}

//解决跨域问题
//func Cors() gin.HandlerFunc {
//	return func(c *gin.Context) {
//		method := c.Request.Method
//		fmt.Println("请求的方式是，", method)
//		c.Header("Access-Control-Allow-Origin", "*")
//		c.Header("Access-Control-Allow-Credentials", "True")
//
//		// //用于判断request来自ajax还是传统请求
//		c.Header("Access-Control-Allow-Headers", "*")
//
//		//
//		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token,Authorization,Token")
//
//		//允许访问的方式
//		c.Header("Access-Control-Allow-Methods", "POST,GET,OPTIONS,PUT,HEAD")
//		c.Header("Access-Control-Expose-Headers", "Content-Length,Access-Control-Allow-Origin,Access-Control-Allow-Headers,Content-Type")
//
//		//内容类型：如果是post请求必须指定这个属性
//		//当ContentType为application/json，会分两次发送请求。
//		// 第一次，先发送method为options的请求到服务器，这个请求会询问服务器支持哪些请求方法（get，post等），支持哪些请求头等服务器的支持情况。
//		//等到这个请求返回后，如果我们准备发送的请求符合服务器规则，那么才会继续发送第二次请求，否则会在console中报错。
//		//c.Header("Content-Type", "application/json;charset=utf-8")
//
//		//放行索引options
//
//		if method == "OPTIONS" {
//			//c.AbortWithStatus(http.StatusNoContent)
//			c.JSON(200, gin.H{
//				"message": "Hello world!",
//			})
//			c.Abort()
//			return
//		}
//		//处理请求
//		c.Next()
//	}
//}

func CorsHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "*")
		c.Header("Access-Control-Allow-Methods", "POST, GET, PUT, PATCH, OPTIONS")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Expose-Headers", "*")
		if c.Request.Method == "OPTIONS" {
			c.JSON(http.StatusOK, "")
			c.Abort()
			return
		}
		c.Next()
	}

}

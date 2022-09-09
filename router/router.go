package router

import (
	"erp_center/controller/power"
	"erp_center/logger"
	"github.com/gin-gonic/gin"
)

func SetupRiouter(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode) // 设置为发布模式
	}
	r := gin.New()

	// 接收gin框架默认的日志
	// recover掉项目可能出现的panic，并使用zap记录相关日志
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	r.GET("/hello", func(c *gin.Context) {
		// c.JSON：返回JSON格式的数据
		c.JSON(200, gin.H{
			"message": "Hello world!",
		})

	})

	r.GET("/r_login", power.Login)
	r.GET("/st/sys.core/user_info", power.Get_info)

	//r.NoRoute(func(context *gin.Context) { // 没有找到路由
	//	fmt.Println("没有找到路由")
	//	context.HTML(http.StatusNotFound, "views/404.html", nil)
	//})
	return r
}

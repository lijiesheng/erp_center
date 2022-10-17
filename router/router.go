package router

import (
	"erp_center/controller/mypage"
	"erp_center/controller/power"
	"erp_center/logger"
	"erp_center/middlewares"
	"github.com/gin-gonic/gin"
)

func SetupRiouter(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode) // 设置为发布模式
	}
	r := gin.New()
	r.Use(middlewares.CorsHandler())
	// 接收gin框架默认的日志
	// recover掉项目可能出现的panic，并使用zap记录相关日志
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	// 跨域

	// 中间键
	r.GET("/r_login", power.Login)
	r.POST("/r_login", power.PostLogin)
	r.POST("/r_register", power.Register)
	r.Use(middlewares.JWTAuthMiddleware())
	r.GET("/st/sys.core/user_info", power.Get_info)
	r.GET("/st/sys.core/user_pages", power.User_pages)
	r.GET("st/sys.core/home_message_list", power.Home_message_list)

	// 我的 api
	// 1、我的考勤

	//r.GET("/api/api.attendance.my_page/queryMyAttendanceDetail", mypage.QueryMyAttendanceDetail) // 查询我本月考勤明细
	r.GET("/api/api.attendance.my_page/queryMyAttendancInfo", mypage.QueryMyAttendancInfo)   // 查询我的假剩余时间
	r.GET("/api/api.attendance.my_page/queryMyAttendanceList", mypage.QueryMyAttendanceList) // 查询我的考勤信息

	// hr 管理

	r.GET("/hello", func(c *gin.Context) {
		// c.JSON：返回JSON格式的数据
		c.JSON(200, gin.H{
			"message": "Hello world!",
		})

	})

	//r.NoRoute(func(context *gin.Context) { // 没有找到路由
	//	fmt.Println("没有找到路由")
	//	context.HTML(http.StatusNotFound, "views/404.html", nil)
	//})
	return r
}

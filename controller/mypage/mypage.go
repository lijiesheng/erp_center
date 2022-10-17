package mypage

import (
	"github.com/gin-gonic/gin"
)

// 查询我本月考勤明细
//func QueryMyAttendanceDetail(c *gin.Context) {
//	// 获取本月第一天
//	month_begin := util.GetFirstDateOfMonth(time.Now()).Format(util.Yyyy_MM_dd)
//	var id any
//	var err error
//	if id, err = request.GetCurrent(c, request.ID); err != nil {
//
//	}
//	logic.QueryMyAttendanceDetail(month_begin, id.(int), 0)
//}

// 查询我的假剩余时间
func QueryMyAttendancInfo(c *gin.Context) {

}

// 查询我的考勤信息
func QueryMyAttendanceList(c *gin.Context) {

}

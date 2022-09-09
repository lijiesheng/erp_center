package power

import (
	"erp_center/controller/resp"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"net/http"
	"time"
)

type SignUpParam struct {
	TlAccount      string `form:"tl_account" json:"tl_account" binding:"gte=1"`
	HashedPassword string `form:"hashed_password" json:"hashed_password" binding:"gte=1"`
	Jpc            string `form:"jpc" json:"jpc"`
}

// tl_account=crmadmin&hashed_password=7c4a8d09ca3762af61e59520943dc26494f8941
func Login(c *gin.Context) {
	// 1、获取参数
	var p SignUpParam
	if err := c.ShouldBind(&p); err != nil {
		// 请求参数有误，直接返回响应
		zap.L().Error("LoginHandler with invalid params ", zap.Error(err))
		// 判断 err 是不是 validator.ValidationErrors 类型
		_, ok := err.(validator.ValidationErrors) // 验证参数是否传入错误
		if !ok {                                  // 非validator.ValidationErrors类型错误直接返回
			resp.ResponseErrorJsonpJSON(c, resp.CodeInvalidParam, resp.CodeInvalidParam.GetMsg(), err.Error())
			return
		}
		resp.ResponseErrorJsonpJSON(c, resp.CodeInvalidParam, resp.CodeInvalidParam.GetMsg(), err.Error())
		return
	}
	resp.ResponseErrorJsonpJSON(c, resp.CodeInvalidParam, resp.CodeInvalidParam.GetMsg(), nil)
	// 2、参数校验

	// 3、业务处理调用 logic

	// 4、返回数据

}

func Get_info(c *gin.Context) {
	fmt.Println("Get_info 进来")
	data := map[string]any{
		"msg":      "账号不存在,请重新输入!",
		"tg":       "login",
		"_stamp":   time.Now().Unix(),
		"status":   "fai",
		"msg_code": 0,
	}
	fmt.Println("bbb")
	callback := c.DefaultQuery("jpc", "")
	if callback == "" {
		c.Render(http.StatusOK, render.JSON{Data: data})
		return
	}
	//s := `/**/ typeof ` + callback + ` === function && ` + callback
	s := fmt.Sprintf(`/**/ typeof %s  === 'function' %s`, callback, callback)
	//s1, _ := UnicodeEscape(s).MarshalJSON()
	c.Render(http.StatusOK, render.JsonpJSON{Callback: s, Data: data})
	return
}

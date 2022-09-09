package resp

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"
	"net/http"
)

type ResponseData struct {
	Code int `json:"code"`
	Msg  any `json:"msg"`
	Data any `json:"data"`
	Err  any `json:"err"`
}

func ResponseErrorJsonpJSON(c *gin.Context, code ResCode, msg any, err any) {
	jpc := c.DefaultQuery("jpc", "")
	if jpc == "" {
		c.JSON(int(code), ResponseData{Code: int(code), Msg: msg, Data: nil})
		return
	}
	callback := fmt.Sprintf(`/**/ typeof %s  === 'function' && %s`, jpc, jpc)
	//data := map[string]any{
	//	"msg":      "账号不存在,请重新输入!",
	//	"tg":       "login",
	//	"_stamp":   time.Now().Unix(),
	//	"status":   "fai",
	//	"tok":      "t:affXRtX/diFRMZgtYGGWPDGk.dxS3MQ3ISxeQKuetmnNc0Yc9eWuHTza3eiG5ESaXm4",
	//	"msg_code": 0,
	//}
	//c.Render(http.StatusOK, render.JsonpJSON{Callback: callback, Data: data})
	c.Render(http.StatusOK, render.JsonpJSON{Callback: callback, Data: ResponseData{Code: int(code), Msg: msg, Err: err}})
}

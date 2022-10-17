package resp

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"
	"net/http"
	"time"
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
		c.JSON(int(code), ResponseData{Code: int(code), Msg: msg, Data: err})
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

type ResponseSucData struct {
	Msg     string       `json:"msg"`
	Info    []InfoDetail `json:"info"`
	Tok     string       `json:"tok"`
	_Stamp  string       `json:"_stamp"`
	Status  string       `json:"status"`
	V       any          `json:"v"`
	MsgCode int          `json:"msg_code"`
}

type ResponseSucRecords struct {
	Total   int    `json:"total"`
	Status  string `json:"status"`
	Records []any  `json:"records"`
	PgSize  int    `json:"pgSize"`
	Pg      int    `json:"pg"`
	MsgCode int    `json:"msg_code"`
	Msg     string `json:"msg"`
}

type InfoDetail struct {
	Sid string `json:"sid"`
	Tim string `json:"tim"`
	Ip  string `json:"ip"`
}

func ResponseSucJsonpJSON(c *gin.Context) {
	jpc := c.DefaultQuery("jpc", "")
	if jpc == "" {
		//c.JSON(int(code), ResponseData{Code: int(code), Msg: msg, Data: err})
		return
	}
	callback := fmt.Sprintf(`/**/ typeof %s  === 'function' && %s`, jpc, jpc)

	token, exists := c.Get("token")
	if !exists {
		token = ""
	}
	c.Render(http.StatusOK, render.JsonpJSON{Callback: callback, Data: ResponseSucData{
		Msg:     "恭喜您，登录成功。",
		Info:    []InfoDetail{},
		Tok:     token.(string),
		_Stamp:  string(time.Now().UnixNano() / 1e6),
		Status:  "suc",
		MsgCode: 1,
	}})
}

func ResponseSucJsonpJSONWithData(c *gin.Context, data any) {
	jpc := c.DefaultQuery("jpc", "")
	if jpc == "" {
		return
	}
	callback := fmt.Sprintf(`/**/ typeof %s  === 'function' && %s`, jpc, jpc)

	token, exists := c.Get("token")
	if !exists {
		token = ""
	}
	c.Render(http.StatusOK, render.JsonpJSON{Callback: callback, Data: ResponseSucData{
		Msg:     "Suc!",
		Info:    []InfoDetail{},
		Tok:     token.(string),
		_Stamp:  string(time.Now().UnixNano() / 1e6),
		Status:  "suc",
		V:       data,
		MsgCode: 1,
	}})
}

func ResponseSucJsonpJSONWithRecords(c *gin.Context, pg int, pgSize int, records []any, total int) {
	jpc := c.DefaultQuery("jpc", "")
	if jpc == "" {
		return
	}
	callback := fmt.Sprintf(`/**/ typeof %s  === 'function' && %s`, jpc, jpc)

	c.Render(http.StatusOK, render.JsonpJSON{Callback: callback, Data: ResponseSucRecords{
		Msg:     "Suc!",
		MsgCode: 1,
		Pg:      pg,
		PgSize:  pgSize,
		Records: records,
		Status:  "suc",
		Total:   total,
	}})
}

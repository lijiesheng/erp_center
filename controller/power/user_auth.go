package power

import (
	"erp_center/controller/request"
	"erp_center/controller/resp"
	"erp_center/dao/mysql"
	"erp_center/dao/rabbitmq/dead_queue_ttl"
	logic "erp_center/logic/user"
	"erp_center/model"
	"erp_center/pkg/jwt"
	"erp_center/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

// tl_account=crmadmin&hashed_password=7c4a8d09ca3762af61e59520943dc26494f8941
func Login(c *gin.Context) {
	var p mysql.SignUpParam
	if err := c.ShouldBind(&p); err != nil {
		zap.L().Error("LoginHandler with invalid params ", zap.Error(err))
		_, ok := err.(validator.ValidationErrors) // 验证参数是否传入错误
		if !ok {                                  // 非validator.ValidationErrors类型错误直接返回
			resp.ResponseErrorJsonpJSON(c, resp.CodeInvalidParam, resp.CodeInvalidParam.GetMsg(), err.Error())
			return
		}
		resp.ResponseErrorJsonpJSON(c, resp.CodeInvalidParam, resp.CodeInvalidParam.GetMsg(), err.Error())
		return
	}
	var user *model.SysUsers
	var err error
	if user, err = logic.Login(&p); err != nil {
		resp.ResponseErrorJsonpJSON(c, 10000, err.Error(), nil)
		return
	}
	// 生成 token
	token, err := jwt.GenToken(p.TlAccount, p.HashedPassword, int64(user.Id))
	if err != nil {
		fmt.Println("生成 token 错误")
		return
	}
	c.Set("token", token)
	fmt.Println("token==>", token)
	fmt.Println("真实的ip 地址 : ", util.RealIP(c))
	resp.ResponseSucJsonpJSON(c)
}

// 注册完成后，使用 rabbitmq 30分钟推送消息
func PostLogin(c *gin.Context) {
	fmt.Println("进来了")

	c.JSON(200, gin.H{
		"message": "Hello world!",
		"status":  1,
		"success": true,
	})

}

func Register(c *gin.Context) {
	fmt.Println("Register 进来了")
	// 发送一条消息
	dead_queue_ttl.Rabbitmq.PublishRouting("注册账号是" + "hello") // 生成消息

	c.JSON(200, gin.H{
		"message": "Hello world!",
		"status":  1,
		"success": true,
	})
}

// todo
func Get_info(c *gin.Context) {
	var err error
	var account any
	if account, err = request.GetCurrent(c, request.CtxTlAccount); err != nil {
		resp.ResponseErrorJsonpJSON(c, 10000, err.Error(), nil)
	}
	user, err := logic.GetUserByTlAccount(account.(string))
	if err != nil {
		resp.ResponseErrorJsonpJSON(c, 10000, err.Error(), nil)
		return
	}
	resp.ResponseSucJsonpJSONWithData(c, user)
}

// 获取导航
func User_pages(c *gin.Context) {
	typeStr := c.Query("type")
	switch typeStr {
	case "ios":
		nav := logic.AllowPagesParse()
		type V struct {
		}
		resp.ResponseSucJsonpJSONWithData(c, nav)
	case "ebs":
	case "app":
	default:

	}
}

func Home_message_list(c *gin.Context) {
	resp.ResponseSucJsonpJSONWithRecords(c, 1, 50, nil, 0)
}

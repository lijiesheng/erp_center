package power

import (
	"erp_center/controller/request"
	"erp_center/controller/resp"
	"erp_center/dao/mysql"
	"erp_center/dao/rabbitmq/dead_queue_ttl"
	"erp_center/lib"
	logic "erp_center/logic/user"
	"erp_center/model"
	"erp_center/pkg/jwt"
	"erp_center/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"net/http"
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

func PostLogin(c *gin.Context) {
	var err error
	// 1、获取参数
	var loginData mysql.LoginData
	if err = c.ShouldBind(&loginData); err != nil {
		fmt.Printf("%+v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 2、校验参数

	// 3、查询数据库
	var extjs_user *model.ExtjsUser
	extjs_user, err = logic.LoginExtjs(&loginData)
	if err != nil {
		fmt.Printf("%+v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 4、生成 toekn
	token, err := jwt.GenToken(extjs_user.Username, extjs_user.Password, int64(extjs_user.Id))
	if err != nil {
		fmt.Println("生成 token 错误")
		return
	}
	c.Set("token", token)
	// 5、返回前端
	c.JSON(200, gin.H{
		"token":   token,
		"message": "Hello world!",
		"status":  1,
		"success": true,
	})
}

func Register(c *gin.Context) {
	// 1、获取参数
	var registerData mysql.RegisterData
	if err := c.ShouldBind(&registerData); err != nil {
		fmt.Printf("%+v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 2、校验参数
	if registerData.ConPassword != registerData.Password {
		c.JSON(http.StatusBadRequest, gin.H{"error": "两次密码不相同"})
		return
	}
	if !util.VerifyEmailFormat(registerData.Email) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "邮箱错误"})
		return
	}

	// 3、存入数据库
	if err := logic.Register(&registerData); err != nil {
		fmt.Printf("%+v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	// 4、发送一条消息
	//dead_queue_ttl.Rabbitmq.PublishRouting("注册账号是" + "hello") // 生成消息
	dead_queue_ttl.Rabbitmq.PublishRouting(registerData.Username + ";" + registerData.Email) // 生成消息

	// 5、返回
	c.JSON(200, gin.H{
		"message": "Hello world!",
		"status":  1,
		"success": true,
	})
}

// 获取用户名和邮箱，如果正确，修改 extjs_user 的 status = 0
// 5min 之内有效
func ConfirmRegister(c *gin.Context) {
	// 1、获取参数
	var conRegisterData mysql.ConRegisterData
	if err := c.ShouldBind(&conRegisterData); err != nil {
		fmt.Printf("%+v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Printf("%+v\n", conRegisterData)

	// 2、检验数据
	if !util.VerifyEmailFormat(conRegisterData.Email) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "邮箱错误"})
		return
	}

	// 3、判断点击事件是否超过 10min
	//now := time.Now()
	//m, _ := time.ParseDuration("-1m")
	//if !now.Add(10 * m).Before(registerData.RegisterTime) {
	//	c.JSON(http.StatusBadRequest, gin.H{"error": "10min 有效期过了，请重新注册"})
	//	return
	//}

	// 4. 修改数据库
	if err := logic.UpdateUser(&conRegisterData); err != nil {
		fmt.Printf("%+v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	// 5、跳转到登录页面
	c.JSON(200, gin.H{
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

// 获取导航条
func GetNav(c *gin.Context) {
	var err error
	var username any
	if username, err = request.GetCurrent(c, request.CtxTlAccount); err != nil {
		resp.ResponseErrorJsonpJSON(c, 10000, err.Error(), nil)
	}
	fmt.Println("username==>", username)
	//c.JSON(200, gin.H{
	//	"status":  1,
	//	"success": true,
	//	"nav":     lib.NavMy,
	//})
	c.JSON(200, gin.H{
		"expanded": true,
		"children": []lib.Children{
			{
				Text:   "首页",
				Leaf:   true,
				Cls:    "node-link",
				Mod:    "desktop",
				ModURL: "desktop.Desktop",
			},
			{
				Text:     "系统管理",
				Expanded: true,
				Cls:      "node-link",
				Children: []lib.Children{
					{
						Text:   "角色管理",
						Leaf:   true,
						Cls:    "node-link",
						Mod:    "role",
						ModURL: "role.Role",
					},
					{
						Text:   "用户管理",
						Leaf:   true,
						Cls:    "node-link",
						Mod:    "user",
						ModURL: "user.User",
					},
				},
			},
			{
				Text:     "文章管理",
				Expanded: true,
				Children: []lib.Children{
					{
						Text: "文章列表",
						Leaf: true,
					},
					{
						Text: "发布文章",
						Leaf: true,
					},
				},
			},

			{
				Text:     "产品管理",
				Expanded: true,
				Children: []lib.Children{
					{
						Text: "产品列表",
						Leaf: true,
					},
					{
						Text: "新增产品",
						Leaf: true,
					},
					{
						Text: "产品监控",
						Leaf: true,
					},
				},
			},

			{
				Text:     "报表管理",
				Expanded: true,
				Children: []lib.Children{
					{
						Text: "生成报表",
						Leaf: true,
					},
					{
						Text: "报表统计",
						Leaf: true,
					},
					{
						Text: "报表打印",
						Leaf: true,
					},
				},
			},
		},
	})
}

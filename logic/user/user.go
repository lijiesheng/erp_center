package logic

import (
	"database/sql"
	mysql "erp_center/dao/mysql"
	redis "erp_center/dao/redis/user"
	"erp_center/lib"
	"erp_center/model"
	"erp_center/model/constParams"
	"erp_center/model/resData"
)

func Login(p *mysql.SignUpParam) (user *model.SysUsers, err error) {
	if user, err = mysql.Login(p); err != nil {
		if err == mysql.ErrorInvalidPassword {
			if err := redis.ErrorTimeAdd(p.TlAccount); err != nil {
				return nil, err
			}
		}
		return nil, err
	}

	time, err := redis.ErrorTime(p.TlAccount)
	if time > 3 { // todo 通过 读取配置文件得到
		return nil, mysql.UserPasswordErrorGt3
	}

	if err := redis.DelErrorTime(p.TlAccount); err != nil {
		return nil, err
	}
	return user, nil
}

func GetUserByTlAccount(tl_account string) (resDataUser *resData.ResUser, err error) {
	resDataUser = &resData.ResUser{}
	// 查询用户
	user := &model.SysUsers{}
	if user, err = mysql.GetUserByTlAccount(tl_account); err != nil {
		return nil, err
	}
	resDataUser.User = user
	// 查询部门
	dep, err := mysql.GetDepart(int64(user.Id), 0, constParams.GT_DEPART)
	// 查询
	if err != nil && err == sql.ErrNoRows {

	} else if err != nil && err != sql.ErrNoRows {
		return nil, err
	} else if err == nil {
		resDataUser.Dep = dep
		groups, err := mysql.GetGroups(int64(dep.GroupId))
		if err != nil && err != sql.ErrNoRows {
			return nil, err
		}
		if err == nil {
			resDataUser.Title = groups
		}
	}
	resDataUser.Account = tl_account
	resDataUser.DeptName = resDataUser.Title.Name
	resDataUser.Name = resDataUser.User.Name
	resDataUser.TitleName = resDataUser.Title.Name

	return resDataUser, nil
}

func AllowPagesParse() []*lib.Node {
	// 这里先写死
	aaa := &lib.Node{
		Text: "ERP功能导航",
		//ID:      "",
		Leaf:    false,
		IconCls: "folder",
		Children: []lib.Node{
			{
				Text:    "首页",
				ID:      "HomePage",
				Leaf:    true,
				IconCls: "i-homes",
				Qtitle:  "HomePage",
			},
			{
				Text:    "我的",
				ID:      "1",
				Leaf:    false,
				IconCls: "folder",
				Children: []lib.Node{
					{
						Text: "我的考勤",
						//Cls:     "i-shield",
						Leaf: true,
						//Class:   "attendance.my_page",
						ID:      "attendance.my_page",
						IconCls: "i-shield",
						Qtitle:  "attendance.my_page",
					},
					{
						Text:    "密码设置",
						ID:      "sys.myPage",
						Leaf:    true,
						IconCls: "i-key",
						Qtitle:  "sys.myPage",
						//Cls:         "i-key",
						//Class:       "sys.myPage",
						//AllUserShow: true,

					},
					{
						Text:    "场地预约",
						ID:      "hr.venue_booking",
						Leaf:    true,
						IconCls: "i-hand",
						Qtitle:  "hr.venue_booking",
						//Cls:         "i-hand",
						//Class:       "hr.venue_booking",
						//AllUserShow: true,
					},
					{
						Text:    "业务申请",
						ID:      "project.business_apply",
						Leaf:    true,
						IconCls: "i-hand",
						Qtitle:  "project.business_apply",
						//Cls:         "i-hand",
						//Class:       "project.business_apply",
						//AllUserShow: true,
					},
					{
						Text:    "登录日志",
						ID:      "tools.user_logs",
						Leaf:    true,
						IconCls: "i-log",
						Qtitle:  "tools.user_logs",
						//Cls:         "i-log",
						//Class:       "tools.user_logs",
						//AllUserShow: true,
					},
					{
						Text:    "任务管理",
						ID:      "project.task_manager",
						Leaf:    true,
						IconCls: "i-event",
						Qtitle:  "project.task_manager",

						//Cls:     "i-event",
						//Class:   "project.task_manager",
					},
				},
			},
		},
	}

	nodes := []*lib.Node{}
	nodes = append(nodes, aaa)
	return nodes
}

// 注册
func Register(registerData *mysql.RegisterData) error {
	if err := mysql.Register(registerData); err != nil {
		return err
	}
	return nil
}

func UpdateUser(registerData *mysql.ConRegisterData) error {
	if err := mysql.UpdateUser(registerData); err != nil {
		return err
	}
	return nil
}

// 修改数据库
func Update(table string, column []string, data []any) {

}

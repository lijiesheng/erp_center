package mysql

import (
	"database/sql"
	"erp_center/model"
	"fmt"
	"time"
)

type SignUpParam struct {
	TlAccount      string `form:"tl_account" json:"tl_account" binding:"gte=1"`
	HashedPassword string `form:"hashed_password" json:"hashed_password" binding:"gte=1"`
	Jpc            string `form:"jpc" json:"jpc"`
}

// 定义数据
type RegisterData struct {
	Username    string `form:"username" json:"username" binding:"required"`
	Email       string `form:"email" json:"email" binding:"required"`
	Password    string `form:"password" json:"password" binding:"required"`
	ConPassword string `form:"conPassword" json:"conPassword" binding:"required"`
}

// 用户名密码判断
func Login(p *SignUpParam) (user *model.SysUsers, err error) {
	user = new(model.SysUsers)
	sqlstr := `select id,account,name, hashed_password, next_pass_time, status from sys_users where account = ?`
	if err := DBs.DxhjOa.Get(user, sqlstr, p.TlAccount); err != nil {
		if err == sql.ErrNoRows { // 没有找到一条记录
			return nil, ErrorUserNotExist
		}
		return nil, err // sql 报错了
	}
	if user.Status == 1 { // 离职
		return nil, ErrorUserLeave
	}
	if user.HashedPassword != p.HashedPassword { // 密码不正确
		return nil, ErrorInvalidPassword
	}
	if time.Now().After(user.NextPassTime) { // 密码的有效期过了
		return nil, ErrorUserPasswordExpired
	}
	return user, nil
}

// 用户注册
func Register(registerData *RegisterData) error {
	sqlInsert := `insert into extjs_user(username, email, password) values (?,?,?)`
	ret, err := DBs.DxhjOa.Exec(sqlInsert, registerData.Username, registerData.Email, registerData.Password)
	if err != nil {
		return err
	}
	theID, err := ret.LastInsertId() // 新插入数据的id
	if err != nil {
		fmt.Printf("get lastinsert ID failed, err:%v\n", err)
		return err
	}
	fmt.Printf("insert success, the id is %d.\n", theID)
	return nil
}

// 根据指定的字段获取 sys_user
func GetUserByTlAccount(tl_account string) (user *model.SysUsers, err error) {
	user = new(model.SysUsers)
	sqlStr := `select id, g_type, account, cus_id,name, gender, mobile, email, email_trans, employee_no, sub_phone, title_id, depart_id,depart_seq,
       i_card, address, portrait, power, remark, company, next_mail_time, next_pass_time, hashed_password, status, created_at, updated_at
       
       from sys_users where account = ?`
	if err = DBs.DxhjOa.Get(user, sqlStr, tl_account); err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrorUserNotExist
		}
		return nil, err // sql 报错了
	}
	if user.Status == 1 { // 离职
		return nil, ErrorUserLeave
	}
	return user, nil
}

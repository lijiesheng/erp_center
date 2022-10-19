package model

import "time"

// 直接用 http://sql2struct.atotoa.com/ 工具转换
//type SysUsers struct {
//	Id             int       `db:"id" json:"id"`
//	GType          int       `db:"g_type"`
//	Account        string    `db:"account"`
//	ExtLogin       string    `db:"ext_login"`
//	ExtNum         string    `db:"ext_num"`
//	CusId          int       `db:"cus_id"`
//	Name           string    `db:"name"`
//	Gender         string    `db:"gender"`
//	Mobile         string    `db:"mobile"`
//	Email          string    `db:"email"`
//	EmailTrans     string    `db:"email_trans"`
//	EmployeeNo     string    `db:"employee_no"`
//	SubPhone       string    `db:"sub_phone"`
//	TitleId        int       `db:"title_id"`
//	DepartId       int       `db:"depart_id"`
//	DepartSeq      int       `db:"depart_seq"`
//	ICard          string    `db:"i_card"`
//	Address        string    `db:"address"`
//	Portrait       string    `db:"portrait"`
//	Power          string    `db:"power"`
//	Remark         string    `db:"remark"`
//	Company        string    `db:"company"`
//	LastLoginTime  time.Time `db:"last_login_time"`
//	NextMailTime   time.Time `db:"next_mail_time"`
//	NextPassTime   time.Time `db:"next_pass_time"`
//	HashedPassword string    `db:"hashed_password"`
//	Status         int       `db:"status"`
//	CreatedAt      time.Time `db:"created_at"`
//	UpdatedAt      time.Time `db:"updated_at"`
//}

type SysUsers struct {
	Id             int       `db:"id" json:"id"`
	GType          int       `db:"g_type" json:"g_type"`
	Account        string    `db:"account" json:"account"`
	ExtLogin       string    `db:"ext_login" json:"ext_login"`
	ExtNum         string    `db:"ext_num" json:"ext_num"`
	CusId          int       `db:"cus_id" json:"cus_id"`
	Name           string    `db:"name" json:"name"`
	Gender         string    `db:"gender" json:"gender"`
	Mobile         string    `db:"mobile" json:"mobile"`
	Email          string    `db:"email" json:"email"`
	EmailTrans     string    `db:"email_trans" json:"email_trans"`
	EmployeeNo     string    `db:"employee_no" json:"employee_no"`
	SubPhone       string    `db:"sub_phone" json:"sub_phone"`
	TitleId        int       `db:"title_id" json:"title_id"`
	DepartId       int       `db:"depart_id" json:"depart_id"`
	DepartSeq      int       `db:"depart_seq" json:"depart_seq"`
	ICard          string    `db:"i_card" json:"i_card"`
	Address        string    `db:"address" json:"address"`
	Portrait       string    `db:"portrait" json:"portrait"`
	Power          string    `db:"power" json:"power"`
	Remark         string    `db:"remark" json:"remark"`
	Company        string    `db:"company" json:"company"`
	LastLoginTime  time.Time `db:"last_login_time" json:"last_login_time"`
	NextMailTime   time.Time `db:"next_mail_time" json:"next_mail_time"`
	NextPassTime   time.Time `db:"next_pass_time" json:"next_pass_time"`
	HashedPassword string    `db:"hashed_password" json:"hashed_password"`
	Status         int       `db:"status" json:"status"`
	CreatedAt      time.Time `db:"created_at" json:"created_at"`
	UpdatedAt      time.Time `db:"updated_at" json:"updated_at"`
}

type ExtjsUser struct {
	Id        int       `db:"id" json:"id"`
	Username  string    `db:"username" json:"username"`
	Email     string    `db:"email" json:"email"`
	Password  string    `db:"password" json:"password"`
	Token     string    `db:"token" json:"token"`
	Status    int       `db:"status" json:"status"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

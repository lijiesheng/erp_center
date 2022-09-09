package model

import "time"

// 直接用 http://sql2struct.atotoa.com/ 工具转换
type SysUsers struct {
	Id             int       `db:"id"`
	TikId          float64   `db:"tik_id"`
	CreatorId      float64   `db:"creator_id"`
	GType          int       `db:"g_type"`
	Account        string    `db:"account"`
	ExtLogin       string    `db:"ext_login"`
	ExtNum         string    `db:"ext_num"`
	CusId          int       `db:"cus_id"`
	Name           string    `db:"name"`
	Gender         string    `db:"gender"`
	Mobile         string    `db:"mobile"`
	Email          string    `db:"email"`
	EmailTrans     string    `db:"email_trans"`
	EmployeeNo     string    `db:"employee_no"`
	SubPhone       string    `db:"sub_phone"`
	TitleId        int       `db:"title_id"`
	DepartId       int       `db:"depart_id"`
	DepartSeq      int       `db:"depart_seq"`
	ICard          string    `db:"i_card"`
	Address        string    `db:"address"`
	Portrait       string    `db:"portrait"`
	Power          string    `db:"power"`
	Remark         string    `db:"remark"`
	LastLoginTime  time.Time `db:"last_login_time"`
	NextMailTime   time.Time `db:"next_mail_time"`
	NextPassTime   time.Time `db:"next_pass_time"`
	HashedPassword string    `db:"hashed_password"`
	Status         int       `db:"status"`
	CreatedAt      time.Time `db:"created_at"`
	UpdatedAt      time.Time `db:"updated_at"`
}

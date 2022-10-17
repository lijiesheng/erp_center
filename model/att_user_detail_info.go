package model

import "time"

type AttUserDetailInfo struct {
	Id             int       `db:"id" json:"id"`
	UserId         int       `db:"user_id" json:"user_id"`
	UserName       string    `db:"user_name" json:"user_name"`
	DepartId       int       `db:"depart_id" json:"depart_id"`
	CompanyType    string    `db:"company_type" json:"company_type"`
	Type           string    `db:"type" json:"type"`
	BillType       string    `db:"bill_type" json:"bill_type"`
	ApplyDate      time.Time `db:"apply_date" json:"apply_date"`
	StartDate      time.Time `db:"start_date" json:"start_date"`
	EndDate        time.Time `db:"end_date" json:"end_date"`
	Remarks        string    `db:"remarks" json:"remarks"`
	Days           float64   `db:"days" json:"days"`
	Hours          float64   `db:"hours" json:"hours"`
	Isagree        string    `db:"isAgree" json:"isAgree"`   //next_status 为0 并且isAgree 为1 才有效
	AmOrPm         string    `db:"am_or_pm" json:"am_or_pm"` //补签上午或者下午
	CurrStatus     int       `db:"curr_status" json:"curr_status"`
	CurrStatusName string    `db:"curr_status_name" json:"curr_status_name"`
	NextStatus     string    `db:"next_status" json:"next_status"`
	NextStatusName string    `db:"next_status_name" json:"next_status_name"`
	CurrId         int       `db:"curr_id" json:"curr_id"`
	CurrName       string    `db:"curr_name" json:"curr_name"`
	NextId         int       `db:"next_id" json:"next_id"`
	NextName       string    `db:"next_name" json:"next_name"`
	Status         float64   `db:"status" json:"status"`
	CreatedAt      time.Time `db:"created_at" json:"created_at"`
	UpdatedAt      int64     `db:"updated_at" json:"updated_at"`
}

package model

import "time"

type SysGroupUsers struct {
	Id        int       `db:"id" json:"id"`
	TikId     float64   `db:"tik_id" json:"tik_id"`
	CreatorId float64   `db:"creator_id" json:"creator_id"`
	GroupId   float64   `db:"group_id" json:"group_id"`
	TypeId    int       `db:"type_id" json:"type_id"`
	UserId    float64   `db:"user_id" json:"user_id"`
	GroupSeq  float64   `db:"group_seq" json:"group_seq"`
	BeginTime time.Time `db:"begin_time" json:"begin_time"`
	EndTime   time.Time `db:"end_time" json:"end_time"`
	Status    float64   `db:"status" json:"status"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

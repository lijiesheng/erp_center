package model

import "time"

type SysGroups struct {
	Id        int       `db:"id" json:"id"`
	TikId     float64   `db:"tik_id" json:"tik_id"`
	CreatorId float64   `db:"creator_id" json:"creator_id"`
	GType     string    `db:"g_type" json:"g_type"`
	ParentId  float64   `db:"parent_id" json:"parent_id"`
	Name      string    `db:"name" json:"name"`
	Desc      string    `db:"desc" json:"desc"`
	Email     string    `db:"email" json:"email"`
	Status    float64   `db:"status" json:"status"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

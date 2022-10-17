package mysql

import (
	"erp_center/model"
)

func GetDepart(user_id, status, type_id int64) (group_user *model.SysGroupUsers, err error) {
	group_user = new(model.SysGroupUsers)
	sqlStr := `select id, tik_id, creator_id, group_id, type_id, user_id, group_seq, begin_time, end_time, status, created_at, updated_at 
from sys_group_users where user_id = ? and status = ? and type_id = ?`
	if err = DBs.DxhjOa.Get(group_user, sqlStr, user_id, status, type_id); err != nil {
		return nil, err // sql 报错了 或 sql.ErrNoRows
	}
	return group_user, nil
}

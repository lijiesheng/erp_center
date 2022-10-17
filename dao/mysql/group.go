package mysql

import (
	"database/sql"
	"erp_center/model"
)

func GetGroups(id int64) (groups *model.SysGroups, err error) {
	groups = new(model.SysGroups)
	sqlStr := `select id, tik_id, creator_id, g_type , parent_id, name, email, status, created_at, updated_at  from sys_groups where id = ?`
	if err = DBs.DxhjOa.Get(groups, sqlStr, id); err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrorUserNotExist
		}
		return nil, err // sql 报错了
	}
	return groups, nil
}

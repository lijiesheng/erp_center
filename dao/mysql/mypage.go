package mysql

import "erp_center/model"

func QueryMyAttendanceDetail(month_begin string, id, status int) (*model.AttUserDetailInfo, error) {
	sql := `select * from att_user_detail_info where start >= ? and user_id = ? and status = ? and isAgree in (?) and next_status != `
	attUser := &model.AttUserDetailInfo{}
	isAgrees := []string{"1", ""}
	if err := DBs.DxhjOa.Get(attUser, sql, month_begin, id, status, isAgrees, "404"); err != nil {
		return nil, err // sql ζ₯ιδΊ ζ sql.ErrNoRows
	}
	return attUser, nil
}

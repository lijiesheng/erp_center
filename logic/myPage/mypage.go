package logic

type Config struct {
	QJ *ColorAndName
	TX *ColorAndName
	JB *ColorAndName
	BQ *ColorAndName
	WC *ColorAndName
}

type ColorAndName struct {
	Color string `json:"color"`
	Name  string `json:"name"`
}

//func QueryMyAttendanceDetail(month_begin string, id, status int) {
//
//	if data, err := mysql.QueryMyAttendanceDetail(month_begin, id, status); err != nil {
//		if err == sql.ErrNoRows {
//
//		}
//		return
//	}
//	for _, v := range data {
//
//	}
//
//}

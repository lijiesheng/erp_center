package util

import "time"

const Yyyy_MM_dd_hh_mm_ss = "2006-01-02 15:04:05"
const Yyyy_MM_dd = "2006-01-02"

// 获取传入时间所在月份的第一天，即某月第一天的0点。如传入time.Now(), 返回当前月份的第一天0点时间。
// 2022-09-01 00:00:00
func GetFirstDateOfMonth(d time.Time) time.Time {
	d = d.AddDate(0, 0, -d.Day()+1)
	return GetZeroTime(d)
}

//获取传入的时间所在月份的最后一天，即某月最后一天的0点。如传入time.Now(), 返回当前月份的最后一天0点时间。
// 比如 2022-09-30 00:00:00
func GetLastDateOfMonth(d time.Time) time.Time {
	return GetFirstDateOfMonth(d).AddDate(0, 1, -1)
}

// 2022-09-30 23:59:59
func GetLastDateOfMonthLastSecond(d time.Time) time.Time {
	return GetFirstDateOfMonth(d).AddDate(0, 1, -1).Add(23 * time.Hour).Add(59 * time.Minute).Add(59 * time.Second)
}

//获取某一天的0点时间
func GetZeroTime(d time.Time) time.Time {
	return time.Date(d.Year(), d.Month(), d.Day(), 0, 0, 0, 0, d.Location())
}

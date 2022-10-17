package util

import (
	"fmt"
	"testing"
	"time"
)

func TestGetFirstDateOfMonth(t *testing.T) {
	time1 := GetFirstDateOfMonth(time.Now())
	fmt.Println(time1.Format(Yyyy_MM_dd_hh_mm_ss))

}

func TestGetLastDateOfMonth(t *testing.T) {
	time1 := GetLastDateOfMonth(time.Now())
	fmt.Println(time1.Format(Yyyy_MM_dd_hh_mm_ss))
}

func TestGetLastDateOfMonthLastSecond(t *testing.T) {
	time1 := GetLastDateOfMonthLastSecond(time.Now())
	fmt.Println(time1.Format(Yyyy_MM_dd_hh_mm_ss))
}

func TestFormat(t *testing.T) {
	fmt.Println(time.Now().Format(Yyyy_MM_dd_hh_mm_ss)) // 2022-09-20 10:36:33
	fmt.Println(time.Now().Format(Yyyy_MM_dd))          // 2022-09-20
}

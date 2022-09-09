package setting

import (
	"fmt"
	"testing"
)

func TestInit(t *testing.T) {
	Init()
	fmt.Printf("%+v\n", Conf)
	fmt.Printf("%+v\n", Conf.Auth)
	fmt.Printf("**************************\n")
	fmt.Printf("%+v\n", Mysqls.Mysql_tl50_crm_db_opt)
	fmt.Printf("%+v\n", Mysqls.Msyql_dxhj_oa_db_opt)
	fmt.Printf("%+v\n", Mysqls.Mysql_dxhj_fund_db_opt)
	fmt.Printf("%+v\n", Mysqls.Mysql_dxhj_gm_fund_trade)
	fmt.Printf("%+v\n", Mysqls.Mysql_tl50_erp_stock_db_opt)
	fmt.Printf("%+v\n", Mysqls.Mysql_dxhj_sm_zg_db_opt)
	fmt.Printf("%+v\n", Mysqls.Msyql_fund_dx_jydb)

	fmt.Printf("**************************\n")
	fmt.Printf("%+v\n", Redises.RedisSession)
	fmt.Printf("%+v\n", Redises.RedisRes99)
	fmt.Printf("%+v\n", Redises.RedisRes6_1)
	fmt.Printf("%+v\n", Redises.Res6_2)
	fmt.Printf("%+v\n", Redises.Res1)
	fmt.Printf("%+v\n", Redises.Res2)
	fmt.Printf("%+v\n", Redises.Res3)
	fmt.Printf("%+v\n", Redises.Res4)
	fmt.Printf("%+v\n", Redises.Res5)
	fmt.Printf("%+v\n", Redises.Res6)
	fmt.Printf("%+v\n", Redises.Res7)
	fmt.Printf("%+v\n", Redises.Res8)
	fmt.Printf("%+v\n", Redises.Res9)
	fmt.Printf("%+v\n", Redises.Res5_4)
	fmt.Printf("%+v\n", Redises.Res_ac2)
	fmt.Printf("%+v\n", Redises.Res6_5)
	fmt.Printf("%+v\n", Redises.Res6_8)
}

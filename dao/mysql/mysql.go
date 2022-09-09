package mysql

import (
	"erp_center/setting"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var DBs *DBList

type DBList struct {
	DxhjOa       *sqlx.DB
	DxhjFund     *sqlx.DB
	FundTrade    *sqlx.DB
	Tl50ErpStock *sqlx.DB
	DxhjSmZg     *sqlx.DB
	DxJydb       *sqlx.DB
}

func initDB(cfgMysql *setting.Mysql, db **sqlx.DB) (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true&loc=Local",
		cfgMysql.User, cfgMysql.Password, cfgMysql.Host, cfgMysql.Port, cfgMysql.Dbname)
	*db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		fmt.Printf("connect DB failed, err:%v\n", err)
		return err
	}
	(*db).SetMaxOpenConns(20)
	(*db).SetMaxIdleConns(10)
	err = (*db).Ping()
	if err != nil {
		fmt.Printf("没有连接到 msyql")
		return err
	}
	return nil
}

func InitDBS(mysqlsConf *setting.MysqlList) {
	DBs = new(DBList)
	initDB(mysqlsConf.Msyql_dxhj_oa_db_opt, &DBs.DxhjOa)
	initDB(mysqlsConf.Mysql_dxhj_sm_zg_db_opt, &DBs.DxhjSmZg)
	initDB(mysqlsConf.Mysql_tl50_erp_stock_db_opt, &DBs.Tl50ErpStock)
	initDB(mysqlsConf.Mysql_dxhj_fund_db_opt, &DBs.DxhjFund)
	initDB(mysqlsConf.Msyql_fund_dx_jydb, &DBs.DxJydb)
	initDB(mysqlsConf.Mysql_dxhj_gm_fund_trade, &DBs.FundTrade)
}

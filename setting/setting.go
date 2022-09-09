package setting

import (
	"fmt"
	"github.com/spf13/viper"
)

var (
	Conf    *AppConfig
	Mysqls  *MysqlList
	Redises *RedisList
	Logger  *Log
)

// 和 conf/config.yaml 文件匹配
type AppConfig struct {
	Name      string `mapstructure:"name"`
	Mode      string `mapstructure:"mode"`
	Port      int    `mapstructure:"port"`
	Version   string `mapstructure:"version"`
	StartTime string `mapstructure:"start_time"`
	MachineId int    `mapstructure:"machine_id"`
	*Auth
}

type Auth struct {
	JwtExpire string `mapstructure:"jwt_expire"`
}

type Log struct {
	Level      string `mapstructure:"level"`
	Filename   string `mapstructure:"filename"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxAge     int    `mapstructure:"max_age"`
	MaxBackups int    `mapstructure:"  max_backups"`
}

type Mysql struct {
	Host         string `mapstructure:"host"`
	Port         uint32 `mapstructure:"port"`
	User         string `mapstructure:"user"`
	Password     string `mapstructure:"password"`
	Dbname       string `mapstructure:"dbname"`
	MaxOpenConns int32  `mapstructure:"max_open_conns"`
	MaxIdleConns int32  `mapstructure:"max_idle_conns"`
}

type Redis struct {
	Host     string `mapstructure:"host"`
	Port     uint32 `mapstructure:"port"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
	PoolSize int    `mapstructure:"pool_size"`
}

type MysqlList struct {
	Msyql_dxhj_oa_db_opt        *Mysql `mapstructure:"msyql_dxhj_oa_db_opt"`
	Mysql_dxhj_fund_db_opt      *Mysql `mapstructure:"mysql_dxhj_fund_db_opt"`
	Mysql_dxhj_gm_fund_trade    *Mysql `mapstructure:"mysql_dxhj_gm_fund_trade"`
	Mysql_tl50_erp_stock_db_opt *Mysql `mapstructure:"mysql_tl50_erp_stock_db_opt"`
	Mysql_dxhj_sm_zg_db_opt     *Mysql `mapstructure:"mysql_dxhj_sm_zg_db_opt"`
	Msyql_fund_dx_jydb          *Mysql `mapstructure:"msyql_fund_dx_jydb"`
}

type RedisList struct {
	RedisSession *Redis `mapstructure:"redis_session"`
	RedisRes99   *Redis `mapstructure:"redis_res99"`
	RedisRes6_1  *Redis `mapstructure:"redis_res6_1"`
	Res6_2       *Redis `mapstructure:"redis_res6_2"`
	Res1         *Redis `mapstructure:"redis_res1"`
	//Res2         *Redis `mapstructure:"redis_res2"`
	//Res3         *Redis `mapstructure:"redis_res3"`
	//Res4         *Redis `mapstructure:"redis_res4"`
	//Res5         *Redis `mapstructure:"redis_res5"`
	//Res6         *Redis `mapstructure:"redis_res6"`
	//Res7         *Redis `mapstructure:"redis_res7"`
	Res8    *Redis `mapstructure:"redis_res8"`
	Res9    *Redis `mapstructure:"redis_res9"`
	Res5_4  *Redis `mapstructure:"redis_res5_4"`
	Res_ac2 *Redis `mapstructure:"redis_res_ac2"`
	Res6_5  *Redis `mapstructure:"redis_res6_5"`
	Res6_8  *Redis `mapstructure:"redis_res6_8"`
}

func Init() (err error) {
	// 1、读取 config.yaml 配置文件
	viper.SetConfigFile("./conf/config.yaml") // todo , 这里是写死的, 后定看看变不变吧
	err = viper.ReadInConfig()
	if err != nil {
		// 读取配置信息失败
		fmt.Printf("viper.ReadInConfig Conf failed, err:%v\n", err)
		return err
	}
	Conf = &AppConfig{}
	if err := viper.Unmarshal(Conf); err != nil {
		fmt.Printf("viper.Unmarshal Conf failed, err:%v\n", err)
		return err
	}

	// 2、读取 mysql.yaml 配置文件
	viper.SetConfigFile("./conf/mysql.yaml")
	err = viper.ReadInConfig()
	if err != nil {
		// 读取配置信息失败
		fmt.Printf("viper.ReadInConfig Mysqls failed, err:%v\n", err)
		return err
	}
	Mysqls = &MysqlList{}
	if err := viper.Unmarshal(Mysqls); err != nil {
		fmt.Printf("viper.Unmarshal Mysqls failed, err:%v\n", err)
		return err
	}

	// 3、读取 redis.yaml 配置文件
	viper.SetConfigFile("./conf/redis.yaml")
	err = viper.ReadInConfig()
	if err != nil {
		// 读取配置信息失败
		fmt.Printf("viper.ReadInConfig redis failed, err:%v\n", err)
		return err
	}
	Redises = &RedisList{}
	if err := viper.Unmarshal(Redises); err != nil {
		fmt.Printf("viper.Unmarshal redis failed, err:%v\n", err)
		return err
	}

	// 4、读取 logger.yaml 配置文件
	viper.SetConfigFile("./conf/logger.yaml")
	err = viper.ReadInConfig()
	if err != nil {
		// 读取配置信息失败
		fmt.Printf("viper.ReadInConfig logger failed, err:%v\n", err)
		return err
	}
	Logger = new(Log)
	if err := viper.Unmarshal(Logger); err != nil {
		fmt.Printf("viper.Unmarshal logger failed, err:%v\n", err)
		return err
	}
	return nil
}

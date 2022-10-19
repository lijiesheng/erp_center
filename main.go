package main

import (
	"erp_center/dao/mysql"
	"erp_center/dao/rabbitmq/dead_queue_ttl"
	"erp_center/dao/redis"
	"erp_center/logger"
	"erp_center/router"
	"erp_center/setting"
	"fmt"
)

/**
 * @Description
 * @Author lijiesheng
 * @Date 2022/9/6 2:20 PM
 * todo 日志写到文件的问题没有解决
 **/
func main() {
	// 1、读取配置文件
	err := setting.Init()
	if err != nil {
		fmt.Println("setting 读取配置文件失败")
		return
	}

	// 2、初始化 mysql
	mysql.InitDBS(setting.Mysqls)

	// 3、初始化 redis
	redis.Init(setting.Redises)

	// 4、初始化 雪花算法

	// 5、初始化 mongo

	// 6、初始化 logger
	logger.Init(setting.Logger, setting.Conf.Mode)

	// 7、初始化 rabbitmq
	dead_queue_ttl.Init("dead-exchange_1", "dead-key_1")

	// 7.1 消费消息
	go dead_queue_ttl.Rabbitmq.RecieveRoutingDeadQueue()

	// 8、初始化 router  这个放在最后
	r := router.SetupRiouter(setting.Conf.Mode)
	err = r.Run(fmt.Sprintf(":%d", setting.Conf.Port))
	if err != nil {
		fmt.Printf("run server failed, err:%d\n", err)
		return
	}

}

// todo 权限管理

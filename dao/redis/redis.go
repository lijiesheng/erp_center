package redis

import (
	"context"
	"erp_center/setting"
	"fmt"
	"github.com/go-redis/redis/v8" // 注意导入的是新版本
	"time"
)

var RDBS *RedisList

type RedisList struct {
	Redis_session *redis.Client
	Redis_res99   *redis.Client
	Redis_res6_1  *redis.Client
	Redis_res6_2  *redis.Client
	Redis_res1    *redis.Client
	//Redis_res2    *redis.Client
	//Redis_res3    *redis.Client
	//Redis_res4    *redis.Client
	//Redis_res5    *redis.Client
	//Redis_res6    *redis.Client
	//Redis_res7    *redis.Client
	Redis_res8    *redis.Client
	Redis_res9    *redis.Client
	Redis_res5_4  *redis.Client
	Redis_res_ac2 *redis.Client
	Redis_res6_5  *redis.Client
	Redis_res6_8  *redis.Client
}

func initRedis(redisConf *setting.Redis, redisDB **redis.Client) {
	addr := fmt.Sprintf("%s:%d", redisConf.Host, redisConf.Port)
	*redisDB = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: redisConf.Password, // no password set
		DB:       redisConf.DB,       // use default DB
		PoolSize: redisConf.PoolSize, // 连接池大小
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := (*redisDB).Ping(ctx).Result()
	if err != nil {
		fmt.Println("redis init err, ", err)
		return
	} else {
		//fmt.Println(addr, " connect successful")
	}
}

func Init(reidsConfList *setting.RedisList) {
	// 初始化 Redis
	RDBS = new(RedisList)
	initRedis(reidsConfList.RedisSession, &RDBS.Redis_session)
	initRedis(reidsConfList.RedisRes99, &RDBS.Redis_res99)
	initRedis(reidsConfList.RedisRes6_1, &RDBS.Redis_res6_1)
	initRedis(reidsConfList.Res6_2, &RDBS.Redis_res6_2)
	initRedis(reidsConfList.Res1, &RDBS.Redis_res1)
	//initRedis(reidsConfList.Res2, &RDBS.Redis_res2)
	//initRedis(reidsConfList.Res3, &RDBS.Redis_res3)
	//initRedis(reidsConfList.Res4, &RDBS.Redis_res4)
	//initRedis(reidsConfList.Res5, &RDBS.Redis_res5)
	//initRedis(reidsConfList.Res6, &RDBS.Redis_res6)
	//initRedis(reidsConfList.Res7, &RDBS.Redis_res7)
	initRedis(reidsConfList.Res8, &RDBS.Redis_res8)
	initRedis(reidsConfList.Res9, &RDBS.Redis_res9)
	initRedis(reidsConfList.Res5_4, &RDBS.Redis_res5_4)
	initRedis(reidsConfList.Res_ac2, &RDBS.Redis_res_ac2)
	initRedis(reidsConfList.Res6_5, &RDBS.Redis_res6_5)
	initRedis(reidsConfList.Res6_8, &RDBS.Redis_res6_8)
}

package redis

import (
	"context"
	"erp_center/dao/redis"
	"time"
)

// 24 小时，密码是否错误超过 3次
func ErrorTime(tl_account string) (int, error) {
	key := "login_time_" + tl_account //todo 字符串拼接，待改进
	value := redis.RDBS.Redis_res6_8.Get(context.Background(), key)
	if value.Err() != nil {
		if value.Err().Error() == "redis: nil" {
			return 0, nil
		}
		return 0, value.Err()
	}
	return value.Int()
}

// 密码错误加1
func ErrorTimeAdd(tl_account string) (err error) {
	var count int
	key := "login_time_" + tl_account //todo 字符串拼接，待改进
	if count, err = ErrorTime(tl_account); err != nil {
		return err
	}
	if count == 0 {
		redis.RDBS.Redis_res6_8.SetEX(context.Background(), key, 0, 120*time.Second)
	}
	value := redis.RDBS.Redis_res6_8.Incr(context.Background(), key)
	return value.Err()
}

// 删除密码错误次数
func DelErrorTime(tl_account string) (err error) {
	key := "login_time_" + tl_account //todo 字符串拼接，待改进
	value := redis.RDBS.Redis_res6_8.Del(context.Background(), key)
	return value.Err()
}

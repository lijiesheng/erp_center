package util

import (
	"context"
	"erp_center/dao/redis"
	"fmt"
	"github.com/gin-gonic/gin"
	rds "github.com/go-redis/redis/v8" // 注意导入的是新版本
	"net"
	"strconv"
	"strings"
	"time"
)

type Z struct {
	Score  float64
	Member interface{}
}

var (
	session_key = "cus_parent_id"
	secret      = "x34sl42xso3yMtiE"

	session_ttl_new = 180      // 刚创建的Token属于临时，有效期很短
	session_ttl     = 3600 * 4 // 默认Token过期时间 -> 4个小时
	prefix          = "tls:"

	key_ip           = "tip_"
	key_access       = "tac_"
	tok_access_locks = "tok_access_locks"
	safe_ips         = []string{"10.10.", "210.14.133.", "210.14.134.", "218.241.191.82", "221.216.211.238", "220.249.6.174"}

	key_danger_ip    = "dip_"
	danger_ips_locks = "danger_ips_locks"
	danger_ip_ttl    = 86400 * 10

	secs_url_count_ip = 180  // 某个接口单位时间单个IP访问次数
	secs_url_expire   = 1800 // 接口访问统计周期

	project_key       = "def"
	secs_ip_count     = 60       // IP token 生成频率
	secs_access_count = 300      // 每个 token 访问次数
	secs_expire       = 300      // 统计周期  (秒) 5分钟
	secs_lock         = 3600 * 8 // 锁定时长  (秒) 8个小时
)

// 将上一个 Token 删掉，重新创建一个 token
func GenNewToken(c *gin.Context) {
	Drop(c)

}

// 删除 session
func Drop(c *gin.Context) {

}

// 生成新的Token，并且限制IP频率
func GenAnonymousToken(c *gin.Context) {
	// ip := RealIP(c)
	// 记录 某个 IP 地址 token 的生成频率

}

// https://www.qycn.com/xzx/article/9591.html
func RealIP(c *gin.Context) string {
	xForwardedFor := c.Request.Header.Get("X-Forward-For")
	ip := strings.TrimSpace(strings.Split(xForwardedFor, ",")[0])
	if ip != "" {
		return ip
	}
	ip = c.Request.Header.Get("x-real-ip")
	if ip != "" {
		return ip
	}
	if ip, _, err := net.SplitHostPort(strings.TrimSpace(c.Request.RemoteAddr)); err == nil {
		return ip
	}
	return ""
}

func SID(c *gin.Context) {

}

func Gen_anonymous_token(c *gin.Context) {
	req_ip := RealIP(c)
	k := key_ip + project_key + "_" + req_ip // todo 字符串直接拼接
	v := redis.RDBS.Redis_session.Incr(context.Background(), k).String()
	re, err := strconv.Atoi(v)
	if err != nil {
		panic("Gen_anonymous_token strconv.Atoi 错误")
	}
	if re > secs_ip_count {
		redis.RDBS.Redis_session.ZAdd(context.Background(), tok_access_locks, &rds.Z{Score: 2, Member: "two"})
		// 如果不是安全的IP地址。
		if !StartWith(safe_ips, req_ip) {
			redis.RDBS.Redis_session.SetEX(context.Background(), k, secs_lock, time.Duration(re)*time.Second)
			fmt.Println("************")
			fmt.Println("这里可能有问题")
			c.JSON(200, gin.H{
				"message": "Hello world!",
			})
		}
	}
	if re <= 2 {
		redis.RDBS.Redis_session.SetEX(context.Background(), k, secs_expire, time.Duration(re)*time.Second)
	}
	// 可以通过c.Set在请求上下文中设置值，后续的处理函数能够取到该值
	c.Set("_TokIsNew", true)
	c.Set("_NewTokTtl", session_ttl_new)

}

// 构造 Ses 对象
// 有 token 但是 token 错误的情况下，需要报错
func GetSessionData(c *gin.Context) {

}

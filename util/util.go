package util

import (
	"math/rand"
	"strings"
)

type New_token struct {
	Id    string `json:"id"`
	Token string `json:"token"`
}

func StartWith(ips []string, ip string) bool {
	for _, v := range ips {
		if strings.IndexAny(v, ip) >= 0 {
			return true
		}
	}
	return false
}

func sesKey(secret string) *New_token {
	var uid = GetRandomString(24)
	return &New_token{Id: uid, Token: "t:" + sesSign(uid, secret)}
}

//func SesUid(len int) string {
//	//id := uuid.NewV4().String()
//
//}

func GetRandomString(n int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	var result []byte
	for i := 0; i < n; i++ {
		result = append(result, bytes[rand.Intn(len(bytes))])
	}
	return string(result)
}

func sesSign(val, secret string) string {
	return ""
}

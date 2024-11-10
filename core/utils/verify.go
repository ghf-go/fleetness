package utils

import (
	"context"
	"fmt"
	"regexp"
	"time"

	"github.com/redis/go-redis/v9"
)

// 是否是手机号
func IsMobile(name string) bool {
	return regexp.MustCompile(`^1\d{9,}$`).MatchString(name)
}

// 是否是邮箱
func IsEmail(name string) bool {
	return regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`).MatchString(name)
}

// 判断发送的验证码
func VerifyCode(r *redis.Client, code, key string) bool {
	rk := fmt.Sprintf("verify:%s", key)
	return r.Get(context.Background(), rk).String() == code
}

// 保存验证码
func VerifySaveCode(r *redis.Client, code, key string, ttl int) error {
	ts := time.Duration(ttl) * time.Second
	rk := fmt.Sprintf("verify:%s", key)
	if r.TTL(context.Background(), rk).Val()+time.Second*60 > ts {
		return fmt.Errorf("请1分钟重试")
	}
	r.Set(context.Background(), rk, code, ts)
	return nil
}

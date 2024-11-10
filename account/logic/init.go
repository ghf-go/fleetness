package logic

import (
	"github.com/ghf-go/fleetness/core"
	"github.com/ghf-go/fleetness/core/utils"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

const (
	TYPE_MOBILE = 1
	TYPE_EMAIL  = 2
	TYPE_WECHAT = 3
	TYPE_WEIBO  = 4
)

var (
	dbConName    = "default"
	cacheConName = "default"
)

func SetDbConName(name string) {
	dbConName = name
}
func SetCacheConName(name string) {
	cacheConName = name
}

func GetDB(c *core.GContent) *gorm.DB {
	return c.GetDB(dbConName)
}
func GetCahce(c *core.GContent) *redis.Client {
	return c.GetCache(cacheConName)
}

// 密码加密
func Passwd(pass, sign string) string {
	return utils.Md5(pass + "_" + sign)
}

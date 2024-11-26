package signin

import (
	"github.com/ghf-go/fleetness/core"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var (
	dbConName    = "default"
	cacheConName = "default"
	isOnline     = false
	days         = 7
	callHandle   func(uid uint64, sumday, contineday uint)
)

func SetConf(maxday int, handle func(uid uint64, sumday, contineday uint)) {
	days = maxday
	callHandle = handle
}
func SetDbConName(name string) {
	dbConName = name
}
func SetCacheConName(name string) {
	cacheConName = name
}

func getDB(c *core.GContent) *gorm.DB {
	return c.GetDB(dbConName)
}
func getCache(c *core.GContent) *redis.Client {
	return c.GetCache(cacheConName)
}

func Init(api, admin, command *core.WebRouter) {
	isOnline = true
	g := api.Group("signin", nil, core.ApiCheckoutLoginMiddleWare)
	g.Post("info", apiSignInfoAction) //签到信息
	g.Post("sign", apiSignAction)     //签到
}

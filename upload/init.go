package upload

import (
	"github.com/ghf-go/fleetness/core"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var (
	dbConName    = "default"
	cacheConName = "default"
	isOnline     = false
)

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
	g := api.Group("upload", core.ApiCheckoutLoginMiddleWare)
	g.Post("getToken", getTokenAction)
	g.Post("uploadSuccess", uploadSuccessAction)
	g.Post("upload", uploadFileAction)

	ag := admin.Group("upload", core.ApiCheckoutLoginMiddleWare)
	ag.Post("getToken", getTokenAction)
	ag.Post("uploadSuccess", uploadSuccessAction)
	ag.Post("upload", uploadFileAction)
}

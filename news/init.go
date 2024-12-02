package news

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
func getCahce(c *core.GContent) *redis.Client {
	return c.GetCache(cacheConName)
}

func Init(api, admin, command *core.WebRouter) {
	isOnline = true
	ag := api.Group("news", core.ApiCheckoutLoginMiddleWare)
	ag.Post("list", apiNewsListAction)
	ag.Post("detail", apiNewsDetailAction)

	adg := admin.Group("news", core.ApiCheckoutLoginMiddleWare)
	adg.Post("list", adminNewsListAction)
	adg.Post("detail", adminNewsDetailAction)
	adg.Post("del", adminNewsDelAction)
	adg.Post("publish", adminNewsPublicAction)
	adg.Post("save", adminNewsSaveAction)
	adg.Post("wait_pulish", adminNewsWaitPublishListAction)

}

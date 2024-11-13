package praise

import (
	"github.com/ghf-go/fleetness/core"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
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

func getDB(c *core.GContent) *gorm.DB {
	return c.GetDB(dbConName)
}
func getCache(c *core.GContent) *redis.Client {
	return c.GetCache(cacheConName)
}

func Init(api, admin, command *core.WebRouter) {
	g := api.Group("praise", nil)
	g.Post("praise", praiseAction)
	g.Post("unpraise", unPraiseAction)
}

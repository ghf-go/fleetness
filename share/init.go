package share

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

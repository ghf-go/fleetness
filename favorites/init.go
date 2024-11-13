package favorites

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

func GetDB(c *core.GContent) *gorm.DB {
	return c.GetDB(dbConName)
}
func GetCahce(c *core.GContent) *redis.Client {
	return c.GetCache(cacheConName)
}

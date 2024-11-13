package comment

import (
	"github.com/ghf-go/fleetness/core"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var (
	dbConName        = "default"
	cacheConName     = "default"
	isSendAfterAudit = false
)

// 设置是否先发后审核
func SetIsSendAfterAutit(isok bool) {
	isSendAfterAudit = isok
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
func getCahce(c *core.GContent) *redis.Client {
	return c.GetCache(cacheConName)
}

func Init(api, admin, command *core.WebRouter) {
	g := api.Group("comment", nil, core.ApiCheckoutLoginMiddleWare)
	g.Post("list", commentListAction)
	g.Post("comment", commentAction)
	// g.Post("favorite", favoriteAction)
	// g.Post("unfavorite", unFavoriteAction)
}

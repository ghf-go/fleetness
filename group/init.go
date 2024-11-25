package group

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
	g := api.Group("group", nil, core.ApiCheckoutLoginMiddleWare)
	g.Post("save", apiGroupSaveAction)          //保存分组
	g.Post("del", apiGroupDelAction)            //删除分组
	g.Post("list", apiGroupListAction)          //分组列表
	g.Post("item_add", apiGroupAddItemAction)   //添加条目
	g.Post("item_del", apiGroupDelItemAction)   //删除条目
	g.Post("item_move", apiGroupMoveItemAction) //移动条目
}

package blackword

import (
	_ "embed"
	"strings"

	"github.com/ghf-go/fleetness/core"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

//go:embed init.sql
var initSql string
var (
	dbConName    = "default"
	cacheConName = "default"
	isOnline     = false
	isInit       = false
)

func execInitSql(c *core.GContent) {
	if isInit {
		return
	}
	lines := strings.Split(initSql, ";")
	for _, sql := range lines {
		sql = strings.TrimSpace(sql)
		if sql != "" {
			getDB(c).Exec(sql)
		}
	}
	isInit = true
}
func initDB(c *core.GContent) {
	if isInit {
		c.Next()
		return
	}
	lines := strings.Split(initSql, ";")
	for _, sql := range lines {
		sql = strings.TrimSpace(sql)
		if sql != "" {
			getDB(c).Exec(sql)
		}
	}
	isInit = true
	c.Next()
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

	ag := admin.Group("blackword", core.ApiCheckoutLoginMiddleWare, initDB)
	ag.Post("del", adminBlackWordDelAction)
	ag.Post("list", adminBlackWordListAction)
	ag.Post("save", adminBlackWordSaveAction)
}

package comment

import (
	_ "embed"
	"strings"

	"github.com/ghf-go/fleetness/core"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var (
	dbConName        = "default"
	cacheConName     = "default"
	isSendAfterAudit = true
	isOnline         = false
	isInit           = false
)

//go:embed init.sql
var initSql string

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
	isOnline = true
	g := api.Group("comment", initDB)
	g.Post("list", commentListAction)
	g.Post("comment", commentAction)

	adg := admin.Group("comment", initDB)
	adg.Post("list", adminListAction)
	adg.Post("wait_audit", adminWaitAuditListAction)
	adg.Post("audit", adminAuditAction)
}

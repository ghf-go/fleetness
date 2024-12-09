package feed

import (
	_ "embed"
	"strings"

	"github.com/ghf-go/fleetness/core"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var (
	dbConName    = "default"
	cacheConName = "default"
	isOnline     = false
	isInit       = false
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

const (
	FEED_TYPE_BLOG  = 0
	FEED_TYPE_VOTE  = 10
	FEED_TYPE_MVOTE = 20
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
	ag := api.Group("feed", initDB, core.ApiCheckoutLoginMiddleWare)
	ag.Post("create", apiFeedCreateAction)
	ag.Post("list", apiFeedListAction)
	ag.Post("vote", apiFeedVoteAction)

	adg := admin.Group("feed", initDB, core.ApiCheckoutLoginMiddleWare)
	adg.Post("list", adminFeedListAction)
	adg.Post("wait_audit", adminFeedWaitAuditAction)
	adg.Post("audit", adminFeedAuditAction)
}

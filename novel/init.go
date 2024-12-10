package novel

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
	g := api.Group("novel", core.ApiCheckoutLoginMiddleWare, initDB)
	g.Post("author_tx", apiAuthorApplyTxAction)
	g.Post("author_tx_list", apiAuthorTxListAction)
	g.Post("author_income", apiAuthorIncomeAction)
	g.Post("author_novel_info", apiAuthorNovelInfoAction)
	g.Post("author_novel_list", apiAuthorNovelListAction)
	g.Post("author_novel_save", apiAuthorSaveNovelAction)
	g.Post("author_session_save", apiAuthorSectionSaveAction)

	g.Post("user_info", apiInfoAuthorAction)
	g.Post("history", apiNovelHistoryAction)
	g.Post("info", apiNovelInfoAction)
	g.Post("list", apiNovelListAction)
	g.Post("read", apiNovelReadAction)
	g.Post("subscribe", apiSubscribeAction)
	g.Post("unsubscribe", apiUnSubscribeAction)
	g.Post("subscribe_list", apiSubscribeListAction)

	ag := api.Group("novel", core.ApiCheckoutLoginMiddleWare, initDB)
	ag.Post("audit_novel", adminNovelAuditAction)
	ag.Post("wait_novel", adminNovelWaitAuditAction)
	ag.Post("audit_section", adminSectionAduitAction)
	ag.Post("wait_section ", adminSectionWaitAuditAction)
	ag.Post("tx", adminTxListAction)
	ag.Post("pay", adminTxPayAction)

}

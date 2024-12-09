package appbuild

import (
	_ "embed"
	"strings"

	"github.com/ghf-go/fleetness/core"
	"github.com/ghf-go/fleetness/core/utils"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var (
	dbConName    = "default"
	cacheConName = "default"
	isOnline     = false
	isInit       = false
	platforms    = []string{"server", "VueAdmin", "AppH5"}
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
	CASH_SCORE  = "score"
	CASH_AMOUNT = "amount"

	TYPE_MOBILE = 1
	TYPE_EMAIL  = 2
	TYPE_WECHAT = 3
	TYPE_WEIBO  = 4
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

func passwd(pass, sign string) string {
	return utils.Md5(pass + "_" + sign)
}

func Init(api, admin, command *core.WebRouter) {
	isOnline = true
	adg := admin.Group("appbuild", initDB, core.ApiCheckoutLoginMiddleWare)
	adg.Post("modules", adminModuleListAction)
	adg.Post("module_save", adminModuleSaveAction)
	adg.Post("items", adminModuleItemsAction)
	adg.Post("item_detail", adminModuleItemDetailAction)
	adg.Post("item_save", adminModuleItemSaveAction)
	adg.Post("project_conf", adminProjectConfAction)
	adg.Post("project_build", adminProjectAction)
}

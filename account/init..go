package account

import (
	_ "embed"
	"strings"

	"github.com/ghf-go/fleetness/core"
	"github.com/ghf-go/fleetness/core/utils"
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
	api.Post("login", loginByPassAction, initDB)
	g := api.Group("account", initDB, core.ApiCheckoutLoginMiddleWare)
	g.Post("register", registerAction)
	g.Post("changepass", changePassAction)
	g.Post("upinfo", setUserInfoAction)
	g.Post("info", getUserInfoAction)
	g.Post("bind", bindAccountAction)
	g.Post("send_code", sendCodeAction)
	g.Post("cash_log", apiCashLogAction)
	g.Post("addrs", apiUserAddrListAction)
	g.Post("addr_save", apiUserAddrSaveAction)

	admin.Post("login", adminLoginAction, initDB)
	adg := admin.Group("account", initDB, core.ApiCheckoutLoginMiddleWare)
	adg.Post("changepasswd", adminChangeAdminPassAction)
	adg.Post("user_add", adminUserAddAction)
	adg.Post("user_wait_audit", adminUserWaitAuditAction)
	adg.Post("user_audit", adminUserAuditAction)
	adg.Post("user_changepass", adminUserChangePassAction)
	adg.Post("user_list", adminUserListAction) //
	adg.Post("user_stat", adminUserStatAction)

}

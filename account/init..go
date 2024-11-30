package account

import (
	"github.com/ghf-go/fleetness/core"
	"github.com/ghf-go/fleetness/core/utils"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var (
	dbConName    = "default"
	cacheConName = "default"
)

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
	api.Post("login", loginByPassAction)
	g := api.Group("account", core.ApiCheckoutLoginMiddleWare)
	g.Post("register", registerAction)
	g.Post("changepass", changePassAction)
	g.Post("upinfo", setUserInfoAction)
	g.Post("info", getUserInfoAction)
	g.Post("bind", bindAccountAction)
	g.Post("send_code", sendCodeAction)
	g.Post("cash_log", apiCashLogAction)

	admin.Post("login", adminLoginAction)
	adg := admin.Group("account", core.ApiCheckoutLoginMiddleWare)
	adg.Post("changepasswd", adminChangeAdminPassAction)
	adg.Post("user_add", adminUserAddAction)
	adg.Post("user_wait_audit", adminUserWaitAuditAction)
	adg.Post("user_audit", adminUserAuditAction)
	adg.Post("user_changepass", adminUserChangePassAction)
	adg.Post("user_list", adminUserListAction) //
	adg.Post("user_stat", adminUserStatAction)

}

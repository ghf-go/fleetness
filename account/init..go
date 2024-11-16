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
	g := api.Group("account", nil, core.ApiCheckoutLoginMiddleWare)
	g.Post("login", loginByPassAction)
	g.Post("register", registerAction)
	g.Post("changepass", changePassAction)
	g.Post("upinfo", setUserInfoAction)
	g.Post("info", getUserInfoAction)
}

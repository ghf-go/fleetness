package feedback

import (
	_ "embed"
	"strings"

	"github.com/ghf-go/fleetness/core"
	"gorm.io/gorm"
)

var (
	dbConName = "default"
	isOnline  = false
	isInit    = false
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
func SetDbConName(name string) {
	dbConName = name
}
func getDB(c *core.GContent) *gorm.DB {
	return c.GetDB(dbConName)
}
func Init(api, admin, command *core.WebRouter) {
	isOnline = true
	g := api.Group("feedback", initDB, core.ApiCheckoutLoginMiddleWare)
	g.Post("list", apiFeedBackListAction)
	g.Post("send", apiFeedBackSendAction)
	a := admin.Group("feedback", initDB, core.ApiCheckoutLoginMiddleWare)
	a.Post("send", adminFeedBackReplayAction)
	a.Post("list", adminFeedListAction)

}

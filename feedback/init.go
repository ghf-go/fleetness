package feedback

import (
	"github.com/ghf-go/fleetness/core"
	"gorm.io/gorm"
)

var (
	dbConName = "default"
	isOnline  = false
)

func SetDbConName(name string) {
	dbConName = name
}
func getDB(c *core.GContent) *gorm.DB {
	return c.GetDB(dbConName)
}
func Init(api, admin, command *core.WebRouter) {
	isOnline = true
	g := api.Group("feedback", core.ApiCheckoutLoginMiddleWare)
	g.Post("list", apiFeedBackListAction)
	g.Post("send", apiFeedBackSendAction)
	a := admin.Group("feedback", core.ApiCheckoutLoginMiddleWare)
	a.Post("send", adminFeedBackReplayAction)
	a.Post("list", adminFeedListAction)

}

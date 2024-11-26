package order

import (
	"github.com/ghf-go/fleetness/core"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var (
	dbConName    = "default"
	cacheConName = "default"
)

const (
	PAYWAY_WC_H5     = "wx_h5"
	PAYWAY_WC_APP    = "wx_APP"
	PAYWAY_WC_JSAPI  = "wx_jsapi"
	PAYWAY_WC_NATIVE = "wx_native"

	ORDER_STATUS_WAIT_PAY   = 0
	ORDER_STATUS_CLOSED     = 1
	ORDER_STATUS_PAY        = 10
	ORDER_STATUS_REFUND     = 20
	ORDER_STATUS_NOT_EXISTS = -1
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

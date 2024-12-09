package push

import (
	"strings"
	"time"

	_ "embed"

	"github.com/ghf-go/fleetness/core"
	"github.com/ghf-go/fleetness/core/session"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var (
	dbConName    = "default"
	cacheConName = "default"
	isOnline     = false
	isInit       = false
	allSse       = map[string]*core.Sse{}
	userSse      = map[uint64]map[string]*core.Sse{}
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
func SetCacheConName(name string) {
	cacheConName = name
}

func getDB(c *core.GContent) *gorm.DB {
	return c.GetDB(dbConName)
}
func getCahce(c *core.GContent) *redis.Client {
	return c.GetCache(cacheConName)
}

func Init(api, admin, command *core.WebRouter, ge *core.GEngine) {
	isOnline = true
	api.Post("regdevice", apiRegDeviceAction, initDB) //上报用户设备
	ge.RouterAny("sse_notify", func(c *core.GContent) {
		c.Sse(func(s *core.Sse) {
			allSse[s.GetKey()] = s
			uid := s.GetUserId()
			if uid > 0 {
				if r, ok := userSse[uid]; ok {
					r[s.GetKey()] = s
					userSse[uid] = r
				} else {
					userSse[uid] = map[string]*core.Sse{s.GetKey(): s}
				}
				defer func(id uint64, k string) {
					if r, ok := userSse[id]; ok {
						delete(r, k)
						if len(r) > 0 {
							userSse[id] = r
						} else {
							delete(userSse, id)
						}
					}
				}(uid, s.GetKey())
			}
			defer delete(allSse, s.GetKey())
			for {
				if s.IsClose() {
					return
				}
				time.Sleep(5 * time.Second)
			}

		})
	}, session.SessionJwt("1234567890123456", 1800))
}

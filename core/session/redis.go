package session

import (
	"fmt"
	"time"

	"github.com/ghf-go/fleetness/core"
	uuid "github.com/satori/go.uuid"
)

func SessionRedis(redisName string, expire int) core.Handle {
	return func(c *core.GContent) {
		token := c.GetRequest().Header.Get("Token")
		rd := c.GetCache(redisName)
		rk := fmt.Sprintf("sess:%s", uuid.NewV4())
		if token != "" {
			rk = fmt.Sprintf("sess:%s", token)
			if uid, e := rd.Get(c.GetContext(), rk).Result(); e == nil {
				c.SetUserID(uid)
			}
		}
		c.Next()
		if c.IsLogin() {
			rd.Set(c.GetContext(), rk, fmt.Sprintf("%d", c.GetUserID()), time.Second*time.Duration(expire))
			c.GetResponseWriter().Header().Add("Token", rk)
		}
	}
}

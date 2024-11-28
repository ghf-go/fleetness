package main

import (
	_ "embed"
	"fmt"

	"github.com/ghf-go/fleetness/account"
	"github.com/ghf-go/fleetness/blocklist"
	"github.com/ghf-go/fleetness/comment"
	"github.com/ghf-go/fleetness/config"
	"github.com/ghf-go/fleetness/core"
	"github.com/ghf-go/fleetness/core/session"
	"github.com/ghf-go/fleetness/favorites"
	"github.com/ghf-go/fleetness/feed"
	"github.com/ghf-go/fleetness/feedback"
	"github.com/ghf-go/fleetness/follow"
	"github.com/ghf-go/fleetness/group"
	"github.com/ghf-go/fleetness/lottery"
	"github.com/ghf-go/fleetness/message"
	"github.com/ghf-go/fleetness/metrics"
	"github.com/ghf-go/fleetness/praise"
	"github.com/ghf-go/fleetness/push"
	"github.com/ghf-go/fleetness/signin"
)

//go:embed test.yaml
var _confData []byte

func main() {
	ge := core.NewGengine(_confData)
	apigrp := ge.RouterGroup("api", func(c *core.GContent) {
		c.FailJson(404, "接口不存在")
	}, session.SessionJwt("1234567890123456"))

	admingrp := ge.RouterGroup("admin", func(c *core.GContent) {
		c.FailJson(404, "接口不存在")
	}, session.SessionJwt("1234567890123456"))

	praise.Init(apigrp, admingrp, nil)
	favorites.Init(apigrp, admingrp, nil)
	comment.Init(apigrp, admingrp, nil)
	account.Init(apigrp, admingrp, nil)
	feedback.Init(apigrp, admingrp, nil)
	group.Init(apigrp, admingrp, nil)
	follow.Init(apigrp, admingrp, nil)
	signin.SetConf(7, func(uid uint64, sumday, contineday uint) {
		fmt.Printf(" uid: %d sd: %d cd :%d\n", uid, sumday, contineday)
	})
	signin.Init(apigrp, admingrp, nil)
	message.Init(apigrp, admingrp, nil)
	blocklist.Init(apigrp, admingrp, nil)
	lottery.Init(apigrp, admingrp, nil)
	feed.Init(apigrp, admingrp, nil)
	push.Init(apigrp, admingrp, nil, ge)
	config.Init(apigrp, admingrp, nil)
	metrics.Init(apigrp, admingrp, nil)
	// ge.AddAfterJob("测试每个5秒运行一次", 5, func(c *core.GContent) {
	// 	core.AppDebug("测试每个5秒运行一次")
	// })
	// ge.AddAlwaysJob("测试一直运行", func(c *core.GContent) {
	// 	i := 0
	// 	for {
	// 		i++
	// 		core.AppDebug("测试一直运行 %d", i)
	// 		if i > 10 {
	// 			panic("退出一次")
	// 		}
	// 		time.Sleep(time.Second)
	// 	}
	// })
	// ge.AddCronJob("计划任务执行", "* * * * *", func(c *core.GContent) {
	// 	core.AppDebug("计划任务执行")
	// })
	ge.Run()
}

package main

import (
	_ "embed"
	"fmt"

	"github.com/ghf-go/fleetness/account"
	"github.com/ghf-go/fleetness/appbuild"
	"github.com/ghf-go/fleetness/appver"
	"github.com/ghf-go/fleetness/blocklist"
	"github.com/ghf-go/fleetness/category"
	"github.com/ghf-go/fleetness/comment"
	"github.com/ghf-go/fleetness/config"
	"github.com/ghf-go/fleetness/core"
	"github.com/ghf-go/fleetness/core/session"
	"github.com/ghf-go/fleetness/favorites"
	"github.com/ghf-go/fleetness/feed"
	"github.com/ghf-go/fleetness/feedback"
	"github.com/ghf-go/fleetness/follow"
	"github.com/ghf-go/fleetness/friendlinks"
	"github.com/ghf-go/fleetness/group"
	"github.com/ghf-go/fleetness/lottery"
	"github.com/ghf-go/fleetness/message"
	"github.com/ghf-go/fleetness/metrics"
	"github.com/ghf-go/fleetness/news"
	"github.com/ghf-go/fleetness/praise"
	"github.com/ghf-go/fleetness/push"
	"github.com/ghf-go/fleetness/signin"
)

//go:embed test.yaml
var _confData []byte

func main() {
	// fmt.Println(utils.VerifyOtp2Fa("3K3WIFWX7BENFKZO", "569070"))
	// fmt.Println(utils.VerifyOtp2Fa("WWNEB6NOT3MDVGRW", "675674"))
	// return
	// v1 := "1.1.3"
	// v2 := "1.0.20"
	// fmt.Printf("%s -> %s r: %v\n", v1, v2, utils.CheckVersion(v1, v2))
	// return
	ge := core.NewGengine(_confData)
	apigrp := ge.RouterGroup("api", session.SessionJwt("1234567890123456", 8640000))

	admingrp := ge.RouterGroup("admin", session.SessionJwt("1234567890123456", 1800))

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
	appver.Init(apigrp, admingrp, nil)
	appbuild.Init(apigrp, admingrp, nil)
	category.Init(apigrp, admingrp, nil)
	news.Init(apigrp, admingrp, nil)
	friendlinks.Init(apigrp, admingrp, nil)
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

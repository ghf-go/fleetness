package main

import (
	_ "embed"
	"fmt"
	"time"

	"github.com/ghf-go/fleetness/account"
	"github.com/ghf-go/fleetness/blocklist"
	"github.com/ghf-go/fleetness/comment"
	"github.com/ghf-go/fleetness/core"
	"github.com/ghf-go/fleetness/core/session"
	"github.com/ghf-go/fleetness/favorites"
	"github.com/ghf-go/fleetness/feed"
	"github.com/ghf-go/fleetness/feedback"
	"github.com/ghf-go/fleetness/follow"
	"github.com/ghf-go/fleetness/group"
	"github.com/ghf-go/fleetness/lottery"
	"github.com/ghf-go/fleetness/message"
	"github.com/ghf-go/fleetness/praise"
	"github.com/ghf-go/fleetness/signin"
	"github.com/gorilla/websocket"
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

	ge.RouterAny("w", func(c *core.GContent) {
		c.WebSocket(func(con *websocket.Conn) {
			for {
				if e := con.WriteMessage(websocket.TextMessage, []byte("asdfasd")); e != nil {
					fmt.Printf("写入数据失败 %s\n", e.Error())
					return
				}
				time.Sleep(1 * time.Second)
			}
		})

	})
	ge.RouterAny("test", func(c *core.GContent) {
		c.Sse(func(s *core.Sse) {
			for {
				time.Sleep(1 * time.Second)
				if s.Send("test") != nil {
					return
				}
				time.Sleep(1 * time.Second)
				if s.Send("test", "aa") != nil {
					return
				}
			}
		})

	})

	ge.Run()
}

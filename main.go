package main

import (
	_ "embed"
	"fmt"

	"github.com/ghf-go/fleetness/account"
	"github.com/ghf-go/fleetness/comment"
	"github.com/ghf-go/fleetness/core"
	"github.com/ghf-go/fleetness/core/session"
	"github.com/ghf-go/fleetness/favorites"
	"github.com/ghf-go/fleetness/feedback"
	"github.com/ghf-go/fleetness/follow"
	"github.com/ghf-go/fleetness/group"
	"github.com/ghf-go/fleetness/praise"
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

	ge.Run()
}

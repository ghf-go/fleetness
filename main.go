package main

import (
	_ "embed"
	"fmt"

	"github.com/ghf-go/fleetness/core"
	"github.com/ghf-go/fleetness/core/log"
	"github.com/ghf-go/fleetness/core/session"
)

//go:embed test.yaml
var _confData []byte

func main() {
	ge := core.NewGengine(_confData)
	api := ge.RouterGroup("api", func(c *core.GContent) {
		log.Debug(c, "das")
		c.FailJson(404, "接口怒存在")

	}, session.SessionJwt("12312312"),
		func(c *core.GContent) {
			log.Debug(c, "api start")
			c.Next()
			log.Debug(c, "api end")
		})
	api.Get("login", func(c *core.GContent) {
		fmt.Printf("处理的信息login\n")
		c.SetUserID("1234324")
		c.SuccessJson("123")
	})
	api.Get("zz", func(c *core.GContent) {
		fmt.Printf("zz\n")
		c.SuccessJson("123")
	})
	ge.Run()
}

package config

import "github.com/ghf-go/fleetness/core"

type adminActionParam struct {
	Key  string                       `json:"key"`
	Data map[string]map[string]string `json:"data"`
}

// 获取配置参数
func admimGetConfigAction(c *core.GContent) {
	p := &adminActionParam{}
	if e := c.BindJson(p); e != nil || p.Key == "" {
		c.FailJson(403, c.Lang("client_param_error"))
		return
	}
	c.SuccessJson(GetConfigRoot(c, p.Key, p.Data))
}

// 设置配置参数
func adminSetConfigAction(c *core.GContent) {
	p := &adminActionParam{}
	if e := c.BindJson(p); e != nil || p.Key == "" {
		c.FailJson(403, c.Lang("client_param_error"))
		return
	}
	if SetConfigRoot(c, p.Key, p.Data) {
		c.SuccessJson(GetConfigRoot(c, p.Key, p.Data))
		return
	}
	c.FailJson(403, c.Lang("save_fail"))
}

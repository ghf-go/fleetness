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
		c.FailJson(403, "参数错误"+e.Error())
		return
	}
	c.SuccessJson(GetConfigRoot(c, p.Key, p.Data))
}

// 设置配置参数
func adminSetConfigAction(c *core.GContent) {
	p := &adminActionParam{}
	if e := c.BindJson(p); e != nil || p.Key == "" {
		c.FailJson(403, "参数错误")
		return
	}
	if SetConfigRoot(c, p.Key, p.Data) {
		c.SuccessJson(GetConfigRoot(c, p.Key, p.Data))
		return
	}
	c.FailJson(403, "更新失败")
}

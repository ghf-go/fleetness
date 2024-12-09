package appver

import "github.com/ghf-go/fleetness/core"

type apiCheckUpdateActionParam struct {
	Ver string `json:"ver"`
}

// 检查更新
func apiCheckUpdateAction(c *core.GContent) {
	p := &apiCheckUpdateActionParam{}
	if e := c.BindJson(p); e != nil {
		c.FailJson(403, c.Lang("client_param_error"))
		return
	}
	c.SuccessJson(GetLastVer(c, p.Ver))
}

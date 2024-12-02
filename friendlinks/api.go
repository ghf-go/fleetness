package friendlinks

import "github.com/ghf-go/fleetness/core"

type apiFriendLinksListActionParam struct {
	Platform string `json:"platform"`
}

// 友情链接列表
func apiFriendLinksListAction(c *core.GContent) {
	p := &apiFriendLinksListActionParam{}
	if e := c.BindJson(p); e != nil || p.Platform == "" {
		c.FailJson(403, c.Lang("client_param_error"))
		return
	}
	c.SuccessJson(GetFriendLink(c, p.Platform))
}

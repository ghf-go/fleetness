package follow

import "github.com/ghf-go/fleetness/core"

type followParam struct {
	ID uint64 `json:"user_id"`
}

// 关注
func apiFollowAction(c *core.GContent) {}

// 取消关注
func apiUnFollowAction(c *core.GContent) {}

type followListParam struct {
	ID       uint64 `json:"user_id"`
	Page     int    `json:"page"`
	PageSize int    `json:"page_size"`
}

// 关注列表
func apiFollowListAction(c *core.GContent) {}

// 粉丝列表
func apiFollowFanAction(c *core.GContent) {}

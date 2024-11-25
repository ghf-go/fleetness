package follow

import "github.com/ghf-go/fleetness/core"

// 是否已关注
func IsFollow(c *core.GContent, curuid, targetid uint64) bool {
	return false
}

// 是否关注列表
func IsFollows(c *core.GContent) {}

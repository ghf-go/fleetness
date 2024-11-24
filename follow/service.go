package follow

import "github.com/ghf-go/fleetness/core"

// 关注
func Follow(c *core.GContent) {}

// 取消关注
func UnFollow(c *core.GContent) {}

// 关注成员列表
func FollowUids(c *core.GContent) {}

// 是否已关注
func IsFollow(c *core.GContent) {}

// 是否关注列表
func IsFollows(c *core.GContent) {}

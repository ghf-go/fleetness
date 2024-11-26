package feed

import "github.com/ghf-go/fleetness/core"

// 动态列表
func adminFeedListAction(c *core.GContent) {}

// 待审核列表
func adminFeedWaitAuditAction(c *core.GContent) {}

// 审核动态
func adminFeedAuditAction(c *core.GContent) {}

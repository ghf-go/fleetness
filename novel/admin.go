package novel

import "github.com/ghf-go/fleetness/core"

// 待审核小说
func adminNovelWaitAuditAction(c *core.GContent) {}

// 审核小数
func adminNovelAuditAction(c *core.GContent) {}

// 待审核章节
func adminSectionWaitAuditAction(c *core.GContent) {}

// 审核章节
func adminSectionAduitAction(c *core.GContent) {}

// 提现列表
func adminTxListAction(c *core.GContent) {}

// 提现付款
func adminTxPayAction(c *core.GContent) {}

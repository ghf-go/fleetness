package feed

import (
	"github.com/ghf-go/fleetness/core"
	"github.com/ghf-go/fleetness/feed/model"
)

// 动态列表
func adminFeedListAction(c *core.GContent) {
	p := &core.PageParam{}
	if e := c.BindJson(p); e != nil {
		c.FailJson(403, c.Lang("client_param_error"))
		return
	}
	flist := []model.Feed{}
	total := int64(0)
	getDB(c).Where("status!=?", core.STATUS_DEL).Order("id DESC").Offset(p.GetOffset()).Limit(p.GetPageSize()).Find(&flist)
	getDB(c).Model(&model.Feed{}).Where("status!=?", core.STATUS_DEL).Count(&total)
	ret := formatFeedList(c, c.GetUserID(), flist)
	c.SuccessJson(map[string]any{
		"total": total,
		"list":  ret,
	})
}

// 待审核列表
func adminFeedWaitAuditAction(c *core.GContent) {
	p := &core.PageParam{}
	if e := c.BindJson(p); e != nil {
		c.FailJson(403, c.Lang("client_param_error"))
		return
	}
	flist := []model.Feed{}
	total := int64(0)
	getDB(c).Where("status=?", core.STATUS_WAIT_AUDIT).Order("id DESC").Offset(p.GetOffset()).Limit(p.GetPageSize()).Find(&flist)
	getDB(c).Model(&model.Feed{}).Where("status=?", core.STATUS_WAIT_AUDIT).Count(&total)
	ret := formatFeedList(c, c.GetUserID(), flist)
	c.SuccessJson(map[string]any{
		"total": total,
		"list":  ret,
	})
}

type adminFeedAuditActionParam struct {
	ID  uint64 `json:"id"`
	Act string `json:"act"`
}

// 审核动态
func adminFeedAuditAction(c *core.GContent) {
	p := &adminFeedAuditActionParam{}
	if e := c.BindJson(p); e != nil || p.ID == 0 || p.Act == "" {
		c.FailJson(403, c.Lang("client_param_error"))
		return
	}
	status := 0
	switch p.Act {
	case "success":
		status = core.STATUS_SUCCESS
	case "self":
		status = core.STATUS_MY
	case "del":
		status = core.STATUS_DEL
	}
	if getDB(c).Model(&model.Feed{}).Where(p.ID).Update("status", status).RowsAffected == 0 {
		c.FailJson(405, c.Lang("save_fail"))
		return
	}
	c.SuccessJson("success")
}

package comment

import (
	"github.com/ghf-go/fleetness/account"
	"github.com/ghf-go/fleetness/comment/model"
	"github.com/ghf-go/fleetness/core"
	"github.com/ghf-go/fleetness/core/utils"
	"gorm.io/gorm"
)

// 评论列表
func adminListAction(c *core.GContent) {
	p := &core.PageParam{}
	if e := c.BindJson(p); e != nil {
		c.FailJson(403, c.Lang("client_param_error"))
		return
	}
	flist := []model.Comment{}
	total := int64(0)
	getDB(c).Where("status!=?", core.STATUS_DEL).Order("id DESC").Offset(p.GetOffset()).Limit(p.GetPageSize()).Find(&flist)
	getDB(c).Model(&model.Comment{}).Where("status!=?", core.STATUS_DEL).Count(&total)
	ret := account.AppendUserBase(c, utils.ModelList2Map(flist), "user_id", "user_info")
	c.SuccessJson(map[string]any{
		"total": total,
		"list":  ret,
	})
}

// 待审核列表
func adminWaitAuditListAction(c *core.GContent) {
	p := &core.PageParam{}
	if e := c.BindJson(p); e != nil {
		c.FailJson(403, c.Lang("client_param_error"))
		return
	}
	flist := []model.Comment{}
	total := int64(0)
	getDB(c).Where("status=?", core.STATUS_WAIT_AUDIT).Order("id DESC").Offset(p.GetOffset()).Limit(p.GetPageSize()).Find(&flist)
	getDB(c).Model(&model.Comment{}).Where("status=?", core.STATUS_WAIT_AUDIT).Count(&total)
	ret := account.AppendUserBase(c, utils.ModelList2Map(flist), "user_id", "user_info")
	c.SuccessJson(map[string]any{
		"total": total,
		"list":  ret,
	})
}

type adminAuditActionParam struct {
	ID  uint64 `json:"id"`
	Act string `json:"act"`
}

// 审核
func adminAuditAction(c *core.GContent) {
	p := &adminAuditActionParam{}
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
	if getDB(c).Model(&model.Comment{}).Where(p.ID).Update("status", status).RowsAffected == 0 {
		if status == core.STATUS_DEL {
			getDB(c).Model(&model.CommentStat{}).Where(p.ID).Update("target_counts", gorm.Expr("target_counts-1"))
		}
		c.FailJson(405, c.Lang("save_fail"))
		return
	}
	c.SuccessJson("success")
}

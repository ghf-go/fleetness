package blocklist

import (
	"github.com/ghf-go/fleetness/account"
	"github.com/ghf-go/fleetness/blocklist/model"
	"github.com/ghf-go/fleetness/core"
)

type apiParam struct {
	TargetType int    `json:"target_type"`
	TargetId   uint64 `json:"target_id"`
}

// 添加黑名单
func apiAddAction(c *core.GContent) {
	p := &apiParam{}
	if e := c.BindJson(p); e != nil || p.TargetId == 0 {
		c.FailJson(403, c.Lang("client_param_error"))
		return
	}
	m := &model.Blocklist{}
	db := getDB(c)
	uid := c.GetUserID()
	db.First(m, "user_id=? AND target_type=? AND target_id=?", uid, p.TargetType, p.TargetId)
	if m.ID == 0 {
		if db.Save(&model.Blocklist{
			UserID:     uid,
			TargetType: p.TargetType,
			TargetID:   p.TargetId,
			CreateIP:   c.GetIP(),
			UpdateIP:   c.GetIP(),
		}).Error != nil {
			c.FailJson(403, c.Lang("save_fail"))
			return
		}
	}
	c.SuccessJson("success")
}

// 删除黑名单
func apiDelAction(c *core.GContent) {
	p := &apiParam{}
	if e := c.BindJson(p); e != nil || p.TargetId == 0 {
		c.FailJson(403, c.Lang("client_param_error"))
		return
	}
	if getDB(c).Delete(&model.Blocklist{}, "user_id=? AND target_type=? AND target_id=?", c.GetUserID(), p.TargetType, p.TargetId).Error != nil {
		c.FailJson(403, c.Lang("save_fail"))
		return
	}
	c.SuccessJson("success")
}

// 黑名单列表
func apiUserListAction(c *core.GContent) {
	p := &core.PageParam{}
	if e := c.BindJson(p); e != nil {
		c.FailJson(403, c.Lang("client_param_error"))
		return
	}
	blist := []model.Blocklist{}
	getDB(c).Where("user_id=? AND target_type=?", c.GetUserID(), TYPE_USER).Order("id DESC").Offset(p.GetOffset()).Limit(p.GetPageSize()).Find(&blist)
	ret := []map[string]any{}
	for _, item := range blist {
		ret = append(ret, map[string]any{
			"id":        item.ID,
			"create_at": item.CreateAt,
			"user_id":   item.TargetID,
		})
	}
	c.SuccessJson(account.AppendUserBase(c, ret, "user_id", "user_info"))
}

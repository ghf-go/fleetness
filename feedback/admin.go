package feedback

import (
	"time"

	"github.com/ghf-go/fleetness/account"
	"github.com/ghf-go/fleetness/core"
	"github.com/ghf-go/fleetness/feedback/model"
)

// 获取意见反馈列表
func adminFeedListAction(c *core.GContent) {
	p := &core.PageParam{}
	if e := c.BindJson(p); e != nil {
		c.FailJson(403, c.Lang("client_param_error"))
		return
	}
	ret := []model.Feedback{}
	total := int64(0)
	db := getDB(c)
	switch p.TabName {
	case "all":
		db.Model(&model.Feedback{}).Count(&total)
		db.Offset((p.Page - 1) * p.PageSize).Limit(p.PageSize).Order("id desc").Find(&ret)
	case "replyed":
		db.Model(&model.Feedback{}).Where("is_replay=1").Count(&total)
		db.Offset((p.Page-1)*p.PageSize).Limit(p.PageSize).Order("id desc").Find(&ret, "is_replay=1")
	default:
		db.Model(&model.Feedback{}).Where("is_replay=0").Count(&total)
		db.Offset((p.Page-1)*p.PageSize).Limit(p.PageSize).Order("id desc").Find(&ret, "is_replay=0")
	}
	retList := []map[string]any{}
	for _, item := range ret {
		retList = append(retList, map[string]any{
			"id":             item.ID,
			"user_id":        item.UserID,
			"imgs":           item.Imgs,
			"content":        item.Content,
			"replay_content": item.ReplayContent,
			"is_replay":      item.IsReplay,
			"replay_at":      item.ReplayAt,
			"create_at":      item.CreateAt,
			"create_ip":      item.CreateIP,
		})
	}
	c.SuccessJson(map[string]any{
		"total": total,
		"list":  account.AppendUserBase(c, retList, "user_id", "user_info"),
	})
}

type adminFeedBackReplayParam struct {
	ID      uint64 `json:"id"`
	Content string `json:"content"`
}

// 回复意见反馈
func adminFeedBackReplayAction(c *core.GContent) {
	p := &adminFeedBackReplayParam{}
	if e := c.BindJson(p); e != nil {
		c.FailJson(403, c.Lang("client_param_error"))
		return
	}
	if p.ID == 0 || p.Content == "" {
		c.FailJson(403, c.Lang("client_param_error"))
		return
	}
	getDB(c).Model(&model.Feedback{}).Where(p.ID).Updates(map[string]any{
		"is_replay":      1,
		"replay_at":      time.Now(),
		"replay_content": p.Content,
	})
	c.SuccessJson("success")

}

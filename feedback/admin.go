package feedback

import (
	"time"

	"github.com/ghf-go/fleetness/core"
	"github.com/ghf-go/fleetness/feedback/model"
)

type adminFeedListActionParam struct {
	Page     int  `json:"page"`
	PageSize int  `json:"page_size"`
	IsAll    bool `json:"is_all"`
}

// 获取评论列表
func adminFeedListAction(c *core.GContent) {
	p := &adminFeedListActionParam{}
	if e := c.BindJson(p); e != nil {
		c.FailJson(403, e.Error())
		return
	}
	if p.Page <= 0 {
		p.Page = 1
	}
	if p.PageSize < 1 {
		p.PageSize = 10
	}
	ret := []model.Feedback{}
	db := getDB(c).Offset((p.Page - 1) * p.PageSize).Limit(p.PageSize).Order("id desc")
	if !p.IsAll {
		db = db.Where("is_replay=1")
	}
	db.Find(&ret)
	c.SuccessJson(ret)
}

type adminFeedBackReplayParam struct {
	ID      uint64 `json:"id"`
	Content string `json:"content"`
}

// 回复意见反馈
func adminFeedBackReplayAction(c *core.GContent) {
	p := &adminFeedBackReplayParam{}
	if e := c.BindJson(p); e != nil {
		c.FailJson(403, e.Error())
		return
	}
	if p.ID == 0 || p.Content == "" {
		c.FailJson(403, "参数错误")
		return
	}
	getDB(c).Model(&model.Feedback{}).Where(p.ID).Updates(map[string]any{
		"is_replay":      1,
		"replay_at":      time.Now(),
		"replay_content": p.Content,
	})
	c.SuccessJson("OK")

}

package feedback

import (
	"github.com/ghf-go/fleetness/core"
	"github.com/ghf-go/fleetness/feedback/model"
)

// 发布请求
func apiFeedBackSendAction(c *core.GContent) {
	p := &model.Feedback{}
	if e := c.BindJson(p); e != nil {
		c.FailJson(403, e.Error())
		return
	}
	p.CreateIP = c.GetIP()
	p.UpdateIP = c.GetIP()
	p.ReplayAt = nil
	p.UserID = c.GetUserID()
	getDB(c).Save(p)
	c.SuccessJson("success")
}

type apiFeedBackListParam struct {
	Page     int `json:"page"`
	PageSize int `json:"page_size"`
}

// 获取意见返回列表
func apiFeedBackListAction(c *core.GContent) {
	p := &apiFeedBackListParam{}
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
	db := getDB(c).Where("user_id=?", c.GetUserID()).Offset((p.Page - 1) * p.PageSize).Limit(p.PageSize).Order("id desc")

	db.Find(&ret)
	c.SuccessJson(ret)
}

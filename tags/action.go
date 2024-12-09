package tags

import (
	"github.com/ghf-go/fleetness/core"
)

type tagListActionParam struct {
	TargetType uint `json:"target_type"`
}

// 标签列表
func tagListAction(c *core.GContent) {
	p := &tagListActionParam{}
	if e := c.BindJson(p); e != nil || p.TargetType == 0 {
		c.FailJson(403, "参数错误")
		return
	}

	c.SuccessJson(TagList(c, p.TargetType))
}

type addTagActionParam struct {
	TargetId   uint64   `json:"target_id"`
	TagIds     []uint64 `json:"tag_ids"`
	TargetType uint     `json:"target_type"`
}

// 目标条件标签
func addTagAction(c *core.GContent) {
	p := &addTagActionParam{}
	if e := c.BindJson(p); e != nil || p.TargetId == 0 || p.TargetType == 0 {
		c.FailJson(403, "参数错误")
		return
	}
	AddTag(c, p.TargetType, p.TargetId, p.TagIds...)
	c.SuccessJson("success")
}

package blackword

import (
	"github.com/ghf-go/fleetness/blackword/model"
	"github.com/ghf-go/fleetness/core"
)

// 敏感词列表
func adminBlackWordListAction(c *core.GContent) {
	list := []model.BlackWord{}
	getDB(c).Find(&list)
	c.SuccessJson(list)
}

type adminBlackWordSaveActionParam struct {
	Word string `json:"word"`
}

// 保存敏感词
func adminBlackWordSaveAction(c *core.GContent) {
	p := &adminBlackWordSaveActionParam{}
	if e := c.BindJson(p); e != nil || p.Word == "" {
		c.FailJson(403, "参数错误")
		return
	}
	m := &model.BlackWord{
		Word:     p.Word,
		CreateIP: c.GetIP(),
		UpdateIP: c.GetIP(),
	}
	if getDB(c).Save(m).Error != nil {
		c.FailJson(403, "保存失败")
		return
	}
	c.SuccessJson(m)
}

// 删除敏感词
func adminBlackWordDelAction(c *core.GContent) {
	p := &core.PageParam{}
	if e := c.BindJson(p); e != nil || p.ID == 0 {
		c.FailJson(403, "参数错误")
		return
	}
	if getDB(c).Delete(&model.BlackWord{}, p.ID).RowsAffected == 0 {
		c.FailJson(403, "删除失败")
		return
	}
	c.SuccessJson("success")
}

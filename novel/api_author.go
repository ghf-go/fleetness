package novel

import (
	"github.com/ghf-go/fleetness/core"
	"github.com/ghf-go/fleetness/novel/model"
)

// 创建小说
func apiAuthorSaveNovelAction(c *core.GContent) {
	p := &model.NovelInfo{}
	if e := c.BindJson(p); e != nil || p.Name == "" {
		c.FailJson(403, "参数错误")
		return
	}
	uid := c.GetUserID()
	db := getDB(c)
	row := &model.NovelInfo{}
	db.First(row, "name=?", p.Name)
	if row.ID > 0 {
		if row.ID != p.ID {
			c.FailJson(403, "名称已存在")
			return
		}
		if row.UserID != uid {
			c.FailJson(403, "名称已存在")
			return
		}
	}

	if p.ID == 0 {
		p.UserID = uid
		p.UpdateIP = c.GetIP()
		p.CreateIP = c.GetIP()
		if db.Save(p).Error == nil {
			c.SuccessJson("success")
			return
		} else {
			c.FailJson(403, "修改失败")
			return
		}
	} else {
		if db.Model(p).Where(p.ID).Updates(map[string]any{
			"name":          p.Name,
			"logo":          p.Logo,
			"section_price": p.SectionPrice,
			"free_section":  p.FreeSection,
			"is_over":       p.IsOver,
			"is_free":       p.IsFree,
			"is_publish":    p.IsPublish,
			"content":       p.Content,
		}).Error == nil {
			c.SuccessJson("success")
			return
		} else {
			c.FailJson(403, "修改失败")
			return
		}
	}

}

// 发布小说列表
func apiAuthorNovelListAction(c *core.GContent) {
	list := []model.NovelInfo{}
	getDB(c).Find(&list, "user_id=?", c.GetUserID())
	c.SuccessJson(list)
}

// 详情
func apiAuthorNovelInfoAction(c *core.GContent) {}

// 保存发布章节
func apiAuthorSectionSaveAction(c *core.GContent) {}

// 我的收入
func apiAuthorIncomeAction(c *core.GContent) {}

// 申请提现
func apiAuthorApplyTxAction(c *core.GContent) {}

// 提现列表
func apiAuthorTxListAction(c *core.GContent) {}

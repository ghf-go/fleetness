package novel

import (
	"errors"
	"time"

	"github.com/ghf-go/fleetness/core"
	"github.com/ghf-go/fleetness/novel/model"
	"gorm.io/gorm"
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
func apiAuthorNovelInfoAction(c *core.GContent) {
	p := &core.ApiParam{}
	if e := c.BindJson(p); e != nil || p.ID == 0 {
		c.FailJson(403, "参数错误")
		return
	}
	uid := c.GetUserID()
	db := getDB(c)
	info := &model.NovelInfo{}
	db.First(info, p.ID)
	if info.ID == 0 || info.UserID != uid {
		c.FailJson(403, "内容不存在")
		return
	}
	slist := []model.NovelSection{}
	db.Order("id DESC").Find(&slist, "novel_id=?", p.ID)

	c.SuccessJson(map[string]any{
		"info":      info,
		"sections:": slist,
	})

}

type apiAuthorSectionSaveActionParam struct {
	NovelID uint64 `json:"novel_id"`
	Name    string `json:"name"`
	Content string `json:"name"`
	core.ApiParam
}

// 保存发布章节
func apiAuthorSectionSaveAction(c *core.GContent) {
	p := &apiAuthorSectionSaveActionParam{}
	if e := c.BindJson(p); e != nil || p.Name == "" || p.NovelID == 0 {
		c.FailJson(403, "参数错误")
		return
	}
	db := getDB(c)
	uid := c.GetUserID()
	ip := c.GetIP()
	info := &model.NovelInfo{}
	db.First(info, p.ID)
	if info.ID == 0 || info.UserID != uid {
		c.FailJson(403, "内容不存在")
		return
	}

	lastRow := &model.NovelSection{}

	db.Where("novel_id=?", p.NovelID).Order("ID DESC").First(lastRow)
	row := &model.NovelSection{
		CreateIP: ip,
		UpdateIP: ip,
		NovelID:  p.NovelID,
		Name:     p.Name,
		ID:       p.ID,
	}
	if row.ID > 0 {
		row.SectionIndex = lastRow.SectionIndex + 1
	}

	if p.Content == "" { //发布大章节
		if lastRow.ID > 0 && lastRow.Words == 0 {
			c.FailJson(403, "不能联系发布大章节")
			return
		}

		if db.Save(row).Error == nil {
			c.SuccessJson("success")
			return
		}
		c.FailJson(403, "发布失败")
		return
	} else { //发布章节
		row.Words = uint(len([]rune(p.Content)))
		c.Tx(db, func(tx *gorm.DB) (error, any) {
			if tx.Save(row).Error != nil || tx.Save(&model.NovelSectionContent{
				CreateIP:  ip,
				UpdateIP:  ip,
				SectionID: row.ID,
				Content:   p.Content,
			}).Error != nil {
				return errors.New("保存失败"), ""
			}
			return nil, ""
		})
		c.SuccessJson("success")
	}
}

// 我的收入
func apiAuthorIncomeAction(c *core.GContent) {
	p := &core.ApiParam{}
	if e := c.BindJson(p); e != nil {
		c.FailJson(403, "参数错误")
		return
	}
	uid := c.GetUserID()
	db := getDB(c)
	ret := map[string]any{
		"income": 0,
		"list":   []any{},
	}
	buyList := []model.NovelBuyLog{}
	if p.ID == 0 { //总收入
		sum := uint(0)
		ilist := []model.NovelInfo{}
		db.Where("user_id=?", uid).Find(&ilist)
		for _, item := range ilist {
			sum += item.TotalIncome
		}
		ret["income"] = sum
		db.Where("author_id=?", uid).Order("id DESC").Offset(p.GetOffset()).Limit(p.GetPageSize()).Find(&buyList)
	} else { //单本收入
		info := &model.NovelInfo{}
		db.First(info, p.ID)
		if info.ID == 0 || info.UserID != uid {
			c.FailJson(403, "内容不存在")
			return
		}
		ret["income"] = info.TotalIncome
		db.Where("author_id=? AND novel_id=?", uid, p.ID).Order("id DESC").Offset(p.GetOffset()).Limit(p.GetPageSize()).Find(&buyList)
	}
	ret["list"] = buyList
	c.SuccessJson(ret)
}

// 申请提现
func apiAuthorApplyTxAction(c *core.GContent) {
	p := &model.NovelTxLog{}
	if e := c.BindJson(p); e != nil || p.Amount == 0 {
		c.FailJson(403, "参数错误")
		return
	}
	db := getDB(c)
	uid := c.GetUserID()

	p.UserID = uid
	p.CreateIP = c.GetIP()
	p.UpdateIP = c.GetIP()
	p.Date = time.Now()

	lastRow := &model.NovelTxLog{}

	db.Where("user_id=?", uid).Order("id DESC").First(lastRow)
	if lastRow.ID > 0 {
		if p.Date.Sub(lastRow.Date) < 24*time.Hour {
			c.FailJson(403, "24小时内只能提现一次")
			return
		}
	}
	if db.Save(p).Error == nil {
		c.SuccessJson("success")
		return
	}
	c.FailJson(403, "提现失败")
}

// 提现列表
func apiAuthorTxListAction(c *core.GContent) {
	p := &core.ApiParam{}
	if e := c.BindJson(p); e != nil {
		c.FailJson(403, "参数错误")
		return
	}
	list := []model.NovelTxLog{}
	getDB(c).Where("user_id=?", c.GetUserID()).Order("id DESC").Offset(p.GetOffset()).Limit(p.GetPageSize()).Find(&list)
	c.SuccessJson(list)
}

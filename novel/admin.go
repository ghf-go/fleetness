package novel

import (
	"github.com/ghf-go/fleetness/core"
	"github.com/ghf-go/fleetness/novel/model"
)

// 待审核小说
func adminNovelWaitAuditAction(c *core.GContent) {
	p := &core.PageParam{}
	if e := c.BindJson(p); e != nil {
		c.FailJson(403, "参数错误")
		return
	}
	db := getDB(c)
	total := int64(0)
	list := []model.NovelInfo{}
	db.Model(&model.NovelInfo{}).Where("is_audit=0").Count(&total)
	db.Where("is_audit=0").Order("id ASC").Offset(p.GetOffset()).Limit(p.GetPageSize()).Find(&list)
	c.SuccessJson(map[string]any{
		"total": total,
		"list":  list,
	})
}

// 审核小数
func adminNovelAuditAction(c *core.GContent) {
	p := &core.PageParam{}
	if e := c.BindJson(p); e != nil || p.ID == 0 {
		c.FailJson(403, "参数错误")
		return
	}
	if getDB(c).Model(&model.NovelInfo{}).Where(p.ID).Update("is_audit", 1).Error != nil {
		c.FailJson(403, "审核失败")
		return
	}
	c.SuccessJson("success")
}

// 待审核章节
func adminSectionWaitAuditAction(c *core.GContent) {
	p := &core.PageParam{}
	if e := c.BindJson(p); e != nil {
		c.FailJson(403, "参数错误")
		return
	}
	db := getDB(c)
	total := int64(0)
	list := []model.NovelSection{}
	clist := []model.NovelSectionContent{}
	cMap := map[uint64]string{}
	db.Model(&model.NovelSection{}).Where("is_audit=0").Count(&total)
	db.Where("is_audit=0").Order("id ASC").Offset(p.GetOffset()).Limit(p.GetPageSize()).Find(&list)
	cids := []uint64{}
	for _, item := range list {
		cids = append(cids, item.ID)
	}
	if len(cids) > 0 {
		db.Find(&clist, cids)
		for _, item := range clist {
			cMap[item.SectionID] = item.Content
		}
	}
	retList := []map[string]any{}
	for _, item := range list {
		if sc, isOk := cMap[item.ID]; isOk {
			retList = append(retList, map[string]any{
				"content": sc,
				"name":    item.Name,
				"id":      item.ID,
			})
		} else {
			retList = append(retList, map[string]any{
				"content": "",
				"name":    item.Name,
				"id":      item.ID,
			})
		}

	}
	c.SuccessJson(map[string]any{
		"total": total,
		"list":  retList,
	})
}

// 审核章节
func adminSectionAduitAction(c *core.GContent) {
	p := &core.PageParam{}
	if e := c.BindJson(p); e != nil || p.ID == 0 {
		c.FailJson(403, "参数错误")
		return
	}
	if getDB(c).Model(&model.NovelSection{}).Where(p.ID).Update("is_audit", 1).Error != nil {
		c.FailJson(403, "审核失败")
		return
	}

	c.SuccessJson("success")
}

// 提现列表
func adminTxListAction(c *core.GContent) {
	p := &core.PageParam{}
	if e := c.BindJson(p); e != nil {
		c.FailJson(403, "参数错误")
		return
	}
	list := []model.NovelTxLog{}

	db := getDB(c)
	total := int64(0)
	switch p.TabName {
	case "wait":
		db.Model(&model.NovelTxLog{}).Where("status=0").Count(&total)
		db.Where("status=0").Order("id DESC").Offset(p.GetOffset()).Limit(p.GetPageSize()).Find(&list)
	case "audit":
		db.Model(&model.NovelTxLog{}).Where("status=10").Count(&total)
		db.Where("status=10").Order("id DESC").Offset(p.GetOffset()).Limit(p.GetPageSize()).Find(&list)
	case "done":
		db.Model(&model.NovelTxLog{}).Where("status=20").Count(&total)
		db.Where("status=20").Order("id DESC").Offset(p.GetOffset()).Limit(p.GetPageSize()).Find(&list)
	default:
		db.Model(&model.NovelTxLog{}).Count(&total)
		db.Order("id DESC").Offset(p.GetOffset()).Limit(p.GetPageSize()).Find(&list)
	}

	c.SuccessJson(map[string]any{
		"total": total,
		"list":  list,
	})

}

// 提现付款
func adminTxPayAction(c *core.GContent) {}

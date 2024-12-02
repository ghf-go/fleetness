package category

import (
	"github.com/ghf-go/fleetness/category/model"
	"github.com/ghf-go/fleetness/core"
)

type adminCategoryListActionParam struct {
	CType int `json:"category_type"`
}

// 分类列表
func adminCategoryListAction(c *core.GContent) {
	p := &adminCategoryListActionParam{}
	if e := c.BindJson(p); e != nil || p.CType == 0 {
		c.FailJson(403, c.Lang("client_param_error"))
		return
	}
	list := []model.Category{}
	getDB(c).Where("target_type=?", p.CType).Order("parent_id ASC ,sort_index ASC").Find(&list)

	c.SuccessJson(formatCategoryList(list))
}

// 保存分类
func adminCategorySaveAction(c *core.GContent) {
	p := &model.Category{}
	if e := c.BindJson(p); e != nil || p.TargetType == 0 || p.Name == "" {
		c.FailJson(403, c.Lang("client_param_error"))
		return
	}
	p.UpdateIP = c.GetIP()
	if p.ID > 0 {
		getDB(c).Model(p).Where(p.ID).Updates(map[string]any{
			"name":        p.Name,
			"target_type": p.TargetType,
			"parent_id":   p.ParentID,
			"is_show":     p.IsShow,
			"sort_index":  p.SortIndex,
			"update_ip":   p.UpdateIP,
		})
	} else {
		p.CreateIP = c.GetIP()
		getDB(c).Save(p)
	}
	c.SuccessJson("success")
}

// 批量保存或者添加分类
func adminCategorySaveListAction(c *core.GContent) {
	p := []*model.Category{}
	if e := c.BindJson(&p); e != nil || len(p) == 0 {
		c.FailJson(403, c.Lang("client_param_error"))
		return
	}
	db := getDB(c)
	for _, item := range p {
		if item.ID > 0 {
			db.Model(&model.Category{}).Where(item.ID).Updates(map[string]any{
				"name":       item.Name,
				"is_show":    item.IsShow,
				"sort_index": item.SortIndex,
				"update_ip":  c.GetIP(),
			})
		} else {
			item.CreateIP = c.GetIP()
			item.UpdateIP = c.GetIP()
			db.Save(item)
		}

	}
	c.SuccessJson("success")
}

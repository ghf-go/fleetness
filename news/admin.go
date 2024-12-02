package news

import (
	"github.com/ghf-go/fleetness/category"
	"github.com/ghf-go/fleetness/core"
	"github.com/ghf-go/fleetness/news/model"
)

type adminNewsListActionParam struct {
	core.PageParam
	Cid uint64 `json:"category_id"`
}

// 新闻列表
func adminNewsListAction(c *core.GContent) {
	p := &adminNewsListActionParam{}
	if e := c.BindJson(p); e != nil {
		c.FailJson(403, c.Lang("client_param_error"))
		return
	}
	total := int64(0)
	db := getDB(c)
	list := []model.News{}
	if p.Cid == 0 {
		db.Model(&model.News{}).Where("is_del=0").Count(&total)
		db.Where("is_del=0").Order("update_at DESC").Offset(p.GetOffset()).Limit(p.GetPageSize()).Find(&list)
	} else {
		db.Model(&model.News{}).Where("is_del=0 AND category_id=?", p.Cid).Count(&total)
		db.Where("is_del=0 AND category_id=?", p.Cid).Order("update_at DESC").Offset(p.GetOffset()).Limit(p.GetPageSize()).Find(&list)
	}

	c.SuccessJson(map[string]any{
		"total":    total,
		"list":     list,
		"category": category.GetCategoryList(c, core.TARGET_TYPE_NEWS),
	})
}

type adminNewsDetailActionParam struct {
	ID uint64 `json:"id"`
}

// 详情
func adminNewsDetailAction(c *core.GContent) {
	p := &adminNewsDetailActionParam{}
	if e := c.BindJson(p); e != nil || p.ID == 0 {
		c.FailJson(403, c.Lang("client_param_error"))
		return
	}
	row := &model.News{}
	getDB(c).First(row, p.ID)
	if row.ID == 0 {
		c.FailJson(403, c.Lang("client_param_error"))
		return
	}
	c.SuccessJson(map[string]any{
		"detail":   row,
		"category": category.GetCategoryList(c, core.TARGET_TYPE_NEWS),
	})
}

// 保存
func adminNewsSaveAction(c *core.GContent) {
	p := &model.News{}
	if e := c.BindJson(p); e != nil || p.Title == "" || p.CategoryID == 0 || p.Content == "" {
		c.FailJson(403, c.Lang("client_param_error"))
		return
	}
	if p.ID > 0 {
		getDB(c).Model(p).Where(p.ID).Updates(map[string]any{
			"update_ip":   c.GetIP(),
			"title":       p.Title,
			"sub_title":   p.SubTitle,
			"category_id": p.CategoryID,
			"img":         p.Img,
			"content":     p.Content,
			"author":      p.Author,
			"refer":       p.Refer,
		})
	} else {
		p.CreateIP = c.GetIP()
		p.UpdateIP = c.GetIP()
		getDB(c).Save(p)
	}
	c.SuccessJson("success")
}

// 待发布列表
func adminNewsWaitPublishListAction(c *core.GContent) {
	p := &core.PageParam{}
	if e := c.BindJson(p); e != nil {
		c.FailJson(403, c.Lang("client_param_error"))
		return
	}
	total := int64(0)
	db := getDB(c)
	list := []model.News{}
	db.Model(&model.News{}).Where("is_publish=0").Count(&total)
	db.Where("is_publish=0").Order("update_at DESC").Offset(p.GetOffset()).Limit(p.GetPageSize()).Find(&list)

	c.SuccessJson(map[string]any{
		"total": total,
		"list":  list,
	})
}

// 发布
func adminNewsPublicAction(c *core.GContent) {
	p := []uint64{}
	if e := c.BindJson(&p); e != nil || len(p) == 0 {
		c.FailJson(403, c.Lang("client_param_error"))
		return
	}
	getDB(c).Model(&model.News{}).Where(p).Updates(map[string]any{
		"update_ip":  c.GetIP(),
		"is_publish": 1,
	})
	c.SuccessJson("success")
}

// 删除
func adminNewsDelAction(c *core.GContent) {
	p := []uint64{}
	if e := c.BindJson(&p); e != nil || len(p) == 0 {
		c.FailJson(403, c.Lang("client_param_error"))
		return
	}
	getDB(c).Model(&model.News{}).Where(p).Updates(map[string]any{
		"update_ip": c.GetIP(),
		"is_del":    1,
	})
	c.SuccessJson("success")
}

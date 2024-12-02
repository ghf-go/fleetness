package news

import (
	"github.com/ghf-go/fleetness/category"
	"github.com/ghf-go/fleetness/core"
	"github.com/ghf-go/fleetness/news/model"
)

type apiNewsListActionParam struct {
	core.PageParam
	Cid uint64 `json:"category_id"`
}

// 新闻列表
func apiNewsListAction(c *core.GContent) {
	p := &apiNewsListActionParam{}
	if e := c.BindJson(p); e != nil {
		c.FailJson(403, c.Lang("client_param_error"))
		return
	}
	db := getDB(c)
	list := []model.News{}
	if p.Cid == 0 {
		db.Where("is_del=0 AND is_publish=1").Order("update_at DESC").Offset(p.GetOffset()).Limit(p.GetPageSize()).Find(&list)
	} else {
		db.Where("is_del=0  AND is_publish=1 AND category_id=?", p.Cid).Order("update_at DESC").Offset(p.GetOffset()).Limit(p.GetPageSize()).Find(&list)
	}

	c.SuccessJson(map[string]any{
		"list":     formatNews(c, c.GetUserID(), list),
		"category": category.GetCategoryList(c, core.TARGET_TYPE_NEWS),
	})
}

type apiNewsDetailActionParam struct {
	ID uint64 `json:"id"`
}

// 新闻详情
func apiNewsDetailAction(c *core.GContent) {
	p := &apiNewsDetailActionParam{}
	if e := c.BindJson(p); e != nil {
		c.FailJson(403, c.Lang("client_param_error"))
		return
	}
	row := &model.News{}
	getDB(c).First(row, p.ID)
	if row.ID == 0 || row.IsDel == 1 || row.IsPublish == 0 {
		c.FailJson(403, c.Lang("client_param_error"))
		return
	}
	data := formatNews(c, c.GetUserID(), []model.News{*row})
	c.SuccessJson(map[string]any{
		"detail":   data[0],
		"category": category.GetCategoryList(c, core.TARGET_TYPE_NEWS),
	})
}

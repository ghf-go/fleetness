package appver

import (
	"strings"

	"github.com/ghf-go/fleetness/appver/model"
	"github.com/ghf-go/fleetness/core"
)

// 版本列表
func adminVerListAction(c *core.GContent) {
	p := &core.PageParam{}
	if e := c.BindJson(p); e != nil {
		c.FailJson(403, "参数错误")
		return
	}
	db := getDB(c)
	total := int64(0)
	db.Model(&model.AppVer{}).Count(&total)
	list := []model.AppVer{}
	db.Order("app_ver DESC").Offset(p.GetOffset()).Limit(p.GetPageSize()).Find(&list)

	c.SuccessJson(map[string]any{
		"total": total,
		"list":  list,
	})
}

type adminVerPublishActionParam struct {
	Ver string `json:"ver"`
}

// 发布版本
func adminVerPublishAction(c *core.GContent) {
	p := &adminVerPublishActionParam{}
	if e := c.BindJson(p); e != nil {
		c.FailJson(403, "参数错误")
		return
	}
	if getDB(c).Model(&model.AppVer{}).Where("app_ver=?", p.Ver).Updates(map[string]any{
		"is_online": 1,
		"update_ip": c.GetIP(),
	}).Error != nil {
		c.FailJson(403, "操作失败")
		return
	}
	c.SuccessJson("OK")
}

type adminVerSaveActionParam struct {
	Ver     string `json:"ver"`
	ApkUrl  string `json:"apk_url"`
	WgtUrl  string `json:"wgt_url"`
	Content string `json:"content"`
}

// 添加版本
func adminVerSaveAction(c *core.GContent) {
	p := &adminVerSaveActionParam{}
	if e := c.BindJson(p); e != nil || !strings.HasPrefix(p.ApkUrl, "http") || !strings.HasPrefix(p.WgtUrl, "http") {
		c.FailJson(403, "参数错误")
		return
	}
	if getDB(c).Save(&model.AppVer{
		AppVer:     p.Ver,
		ApkUrl:     p.ApkUrl,
		WgtUrl:     p.WgtUrl,
		VerContent: p.Content,
		CreateIP:   c.GetIP(),
		UpdateIP:   c.GetIP(),
	}).Error != nil {
		c.FailJson(403, "操作失败")
		return
	}
	c.SuccessJson("OK")
}

package tags

import (
	"github.com/ghf-go/fleetness/core"
	"github.com/ghf-go/fleetness/tags/model"
)

type adminSaveActionParam struct {
	TargetType uint         `json:"target_type"`
	Data       []model.Tags `json:"list"`
}

// 保存列表
func adminSaveAction(c *core.GContent) {
	p := &adminSaveActionParam{}
	if e := c.BindJson(&p); e != nil || len(p.Data) == 0 || p.TargetType == 0 {
		c.FailJson(403, "参数错误")
		return
	}
	db := getDB(c)
	ip := c.GetIP()
	for _, item := range p.Data {
		item.CreateIP = ip
		item.UpdateIP = ip
		item.TargetType = p.TargetType
		if db.Save(&item).Error != nil && db.Model(&item).Where("target_type=? AND name=?", item.TargetType, item.Name).Updates(map[string]any{
			"name":       item.Name,
			"logo":       item.Logo,
			"bg_color":   item.BgColor,
			"font_color": item.FontColor,
			"update_ip":  item.UpdateIP,
		}).Error != nil {
			c.FailJson(403, "参数错误")
			return
		}
	}
	c.SuccessJson("success")
}

// 删除tag
func adminDelAction(c *core.GContent) {
	p := &core.PageParam{}
	if e := c.BindJson(&p); e != nil || p.ID == 0 {
		c.FailJson(403, "参数错误")
		return
	}
	if getDB(c).Delete(&model.Tags{}, "id=? AND sum_times=0", p.ID).RowsAffected > 0 {
		c.SuccessJson("success")
		return
	}
	c.FailJson(403, "删除失败")
	return
}

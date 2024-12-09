package tags

import (
	"github.com/ghf-go/fleetness/core"
	"github.com/ghf-go/fleetness/tags/model"
)

func formatTag(item model.Tags) map[string]any {
	return map[string]any{
		"id":         item.ID,
		"name":       item.Name,
		"logo":       item.Logo,
		"bg_color":   item.BgColor,
		"font_color": item.FontColor,
		"sum_times":  item.SumTimes,
	}
}

// 标签列表
func TagList(c *core.GContent, targetType uint) []map[string]any {
	list := []model.Tags{}
	getDB(c).Find(&list, "target_type=?", targetType)
	ret := []map[string]any{}
	for _, item := range list {
		ret = append(ret, formatTag(item))
	}
	return ret
}

// 目标添加标签ID
func AddTag(c *core.GContent, targetType uint, targetId uint64, tagids ...uint64) {
	db := getDB(c)
	ip := c.GetIP()
	did := []uint64{}
	nid := []uint64{}
	tlis := []model.TagsIds{}
	if len(tagids) == 0 {
		db.Where("target_id=? AND target_type=?", targetId, targetType).Find(&tlis)
	} else {
		db.Where("target_id=? AND target_type=? AND tag_id NOT IN ?", targetId, targetType, tagids).Find(&tlis)
	}
	for _, id := range tagids {
		if db.Save(&model.TagsIds{
			TagID: id, TargetID: targetId, CreateIP: ip, UpdateIP: ip, TargetType: targetType,
		}).Error == nil {
			nid = append(nid, id)
		}
	}
	if len(tlis) > 0 {
		for _, item := range tlis {
			did = append(did, item.TagID)

		}
		db.Delete(&model.TagsIds{}, "tag_id IN ? AND target_id=? AND target_type=?", did, targetId, targetType)

	}
	did = append(did, nid...)
	for _, id := range did {
		t := int64(0)
		db.Model(&model.TagsIds{}).Where("tag_id=? AND target_type=?", id, targetType).Count(&t)
		db.Model(&model.Tags{}).Where(id).Update("sum_times", t)
	}
}

// 专辑tag信息
func AppendTags(c *core.GContent, targetType uint, data []map[string]any, tkey, outkey string) []map[string]any {
	if !isOnline {
		return data
	}
	tids := []uint64{}
	for _, item := range data {
		if id, ok := item[tkey]; ok {
			tids = append(tids, id.(uint64))
		}
	}
	ret := map[uint64][]map[string]any{}
	if len(tids) > 0 {

		db := getDB(c)
		tagList := []model.Tags{}
		tlist := []model.TagsIds{}
		db.Find(&tagList, "target_type=? ", targetType)
		db.Find(&tlist, "target_type=? AND target_id IN ?", targetType, tids)

		tagMap := map[uint64]map[string]any{}
		for _, item := range tagList {
			tagMap[item.ID] = formatTag(item)
		}
		for _, item := range tlist {
			tlis := []map[string]any{}
			if r, isok := ret[item.TargetID]; isok {
				tlis = r
			}
			if row, ok := tagMap[item.TagID]; ok {
				tlis = append(tlis, row)
			}
			ret[item.TargetID] = tlis
		}

	}
	for _, item := range data {
		if id, ok := item[tkey]; ok {
			if row, ok := ret[id.(uint64)]; ok {
				item[outkey] = row
			} else {
				item[outkey] = []any{}
			}
		}
	}
	return data
}

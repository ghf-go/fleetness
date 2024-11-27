package comment

import (
	"github.com/ghf-go/fleetness/comment/model"
	"github.com/ghf-go/fleetness/core"
)

// 追加是评论数据
func AppendCommentInfo(c *core.GContent, targetType uint, data []map[string]any, tkey, outkey string) []map[string]any {
	if !isOnline {
		return data
	}
	tids := []uint64{}
	for _, item := range data {
		if id, ok := item[tkey]; ok {
			tids = append(tids, id.(uint64))
		}
	}
	ret := map[uint64]uint64{}
	if len(tids) > 0 {

		db := getDB(c)
		slist := []model.CommentStat{}
		db.Find(&slist, "target_type=? AND target_id IN ?", targetType, tids)

		for _, item := range slist {
			if _, ok := ret[item.TargetID]; !ok {
				ret[item.TargetID] = item.TargetCounts
			}
		}

	}
	for _, item := range data {
		if id, ok := item[tkey]; ok {
			if row, ok := ret[id.(uint64)]; ok {
				item[outkey] = row
			} else {
				item[outkey] = 0
			}
		}
	}
	return data
}

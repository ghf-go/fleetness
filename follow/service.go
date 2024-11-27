package follow

import (
	"github.com/ghf-go/fleetness/core"
	"github.com/ghf-go/fleetness/follow/model"
)

// 获取关注列表
func Follows(c *core.GContent, uid uint64) []uint64 {
	if !isOnline {
		return []uint64{}
	}
	ulist := []model.FollowItem{}
	ret := []uint64{}
	getDB(c).Where("user_id=?", uid).Find(&ulist)
	for _, item := range ulist {
		ret = append(ret, item.TargetID)
	}
	return ret
}
func IsFollow(c *core.GContent, uid, targetID uint64) bool {
	if !isOnline {
		return false
	}
	m := &model.FollowItem{}
	getDB(c).First(m, "user_id=? AND target_id=?", uid, targetID)
	return m.ID > 0
}

// 追加是否已经关注
func AppendUserIsFollows(c *core.GContent, data []map[string]any, uidkey, outkey string) []map[string]any {
	if !isOnline {
		return data
	}
	uids := []uint64{}
	for _, item := range data {
		if id, ok := item[uidkey]; ok {
			uids = append(uids, id.(uint64))
		}
	}
	if len(uids) > 0 {
		ulist := []model.FollowItem{}
		getDB(c).Where("user_id=? AND target_id IN ?", c.GetUserID(), uids).Find(&ulist)
		um := map[uint64]bool{}
		for _, ii := range ulist {
			um[ii.TargetID] = true
		}
		for _, item := range data {
			if id, ok := item[uidkey]; ok {
				if _, ok := um[id.(uint64)]; ok {
					item[outkey] = true
				} else {
					item[outkey] = false
				}
			}
		}

	}
	return data
}

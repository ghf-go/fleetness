package blocklist

import (
	"github.com/ghf-go/fleetness/blocklist/model"
	"github.com/ghf-go/fleetness/core"
)

const (
	TYPE_USER = 0
)

// 是否在黑名单中
func InBlockList(c *core.GContent, targetID uint64) bool {
	if !isOnline {
		return false
	}
	m := &model.Blocklist{}
	getDB(c).First(m, "user_id=? AND target_type=? AND target_id=?", targetID, TYPE_USER, c.GetUserID())
	return m.ID > 0
}

// 获取黑名单成员
func BlockList(c *core.GContent, uid uint64) []uint64 {
	ulist := []model.Blocklist{}
	ret := []uint64{}
	getDB(c).Where("user_id=?", uid).Find(&ulist)
	for _, item := range ulist {
		ret = append(ret, item.TargetID)
	}
	return ret
}

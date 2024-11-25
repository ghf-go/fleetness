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

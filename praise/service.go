package praise

import (
	"github.com/ghf-go/fleetness/core"
	"github.com/ghf-go/fleetness/core/utils"
	"github.com/ghf-go/fleetness/praise/model"
)

func GetPraise(c *core.GContent, targetType uint, ids ...uint64) map[uint64]map[string]any {
	db := GetDB(c)
	uid := c.GetUserID()
	slist := []model.PraiseStat{}
	flist := []model.Praise{}
	db.Find(&slist, "target_type=? AND target_id IN (?)", targetType, utils.BuildIntsToString(ids...))
	db.Find(&flist, "target_type=? AND user_id=?  AND target_id IN (?)", targetType, uid, utils.BuildIntsToString(ids...))
	ret := map[uint64]map[string]any{}
	for _, item := range slist {
		if _, ok := ret[item.TargetID]; !ok {
			ret[item.TargetID] = map[string]any{
				"count": item.TargetCounts,
				"is_my": false,
			}
		}
	}
	for _, item := range flist {
		if r, ok := ret[item.TargetID]; ok {
			r["is_my"] = true
			ret[item.TargetID] = r
		} else {
			ret[item.TargetID] = map[string]any{
				"count": 0,
				"is_my": true,
			}
		}
	}
	return ret
}

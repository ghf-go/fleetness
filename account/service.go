package account

import (
	"github.com/ghf-go/fleetness/account/model"
	"github.com/ghf-go/fleetness/core"
)

// 获取账号信息列表
func GetUserInfoList(c *core.GContent, uid ...uint64) map[uint64]map[string]string {
	ret := map[uint64]map[string]string{}

	return ret
}

// 获取账号基本信息
func GetUserBaseList(c *core.GContent, uid ...uint64) map[uint64]any {
	ret := map[uint64]any{}
	list := []model.User{}
	getDB(c).Find(&list, uid)
	for _, item := range list {
		ret[item.ID] = map[string]any{
			"user_id":   item.ID,
			"avatar":    item.Avatar,
			"nick_name": item.NickName,
		}
	}
	return ret
}

// 添加账号信息
func AppendUserBase(c *core.GContent, data []map[string]any, uidkey, outkey string) []map[string]any {

	uids := []uint64{}
	for _, item := range data {
		if id, ok := item[uidkey]; ok {
			uids = append(uids, id.(uint64))
		}
	}
	if len(uids) > 0 {
		umap := GetUserBaseList(c, uids...)
		for _, item := range data {
			if id, ok := item[uidkey]; ok {
				if uinfo, ok := umap[id.(uint64)]; ok {
					item[outkey] = uinfo
				}
			}
		}

	}
	return data

}

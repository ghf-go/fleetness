package account

import (
	"errors"

	"github.com/ghf-go/fleetness/account/model"
	"github.com/ghf-go/fleetness/core"
	"gorm.io/gorm"
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

// 添加用户金额记录
func UserCashLog(c *core.GContent, uid uint64, cashType string, amount int, msg string) error {
	e, _ := c.Tx(getDB(c), func(tx *gorm.DB) (error, any) {
		if tx.Model(&model.UserCash{}).Where("user_id=? AND ukey=?", uid, cashType).Update("val", gorm.Expr("val + ?", amount)).RowsAffected == 0 && tx.Save(&model.UserCash{
			CreateIP: c.GetIP(),
			UpdateIP: c.GetIP(),
			UserID:   uid,
			Ukey:     cashType,
			Val:      amount,
		}).Error != nil {
			return errors.New(c.Lang("save_fail")), ""
		}
		if tx.Save(&model.UserCashLog{
			CreateIP: c.GetIP(),
			UpdateIP: c.GetIP(),
			UserID:   uid,
			Ukey:     cashType,
			Val:      amount,
			Content:  msg,
		}).Error != nil {
			return errors.New(c.Lang("save_fail")), ""
		}
		return nil, ""
	})
	return e
}

// 获取收货地址信息
func GetUserAddr(c *core.GContent, uid uint64, addrids ...uint64) map[string]any {
	aid := uint64(0)
	if len(addrids) > 0 {
		aid = addrids[0]
	}
	row := &model.UserAddr{}
	if aid > 0 {
		getDB(c).First(row, aid)
		if row.ID > 0 && row.UserID == uid {
			return formatUserAddr(row)
		}
	}
	getDB(c).First(row, "user_id=? AND is_default=1", uid)
	if row.ID > 0 {
		return formatUserAddr(row)
	}
	getDB(c).First(row, "user_id=?", uid)
	if row.ID > 0 {
		return formatUserAddr(row)
	}
	return nil
}

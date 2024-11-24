package group

import (
	"errors"

	"github.com/ghf-go/fleetness/core"
	"github.com/ghf-go/fleetness/group/model"
	"gorm.io/gorm"
)

// 保存分组
func GroupSave(c *core.GContent, data *model.Group) error {
	if data.ID == 0 {
		data.UserID = c.GetUserID()
		data.CreateIP = c.GetIP()
		data.UpdateIP = c.GetIP()
		return getDB(c).Save(data).Error
	} else {
		return getDB(c).Model(data).Where("id=? AND user_id=?", data.ID, c.GetUserID()).Updates(map[string]any{
			"update_ip":  c.GetIP(),
			"group_name": data.GroupName,
		}).Error
	}
}

// 删除分组
func GroupDel(c *core.GContent, uid, gid uint64) bool {
	db := getDB(c)
	if db.Delete(&model.Group{}, "id=? AND user_id=?", gid, uid).RowsAffected > 0 {
		db.Delete(&model.GroupItem{}, "group_id=?", gid)
		return true
	}
	return false
}

// 分组列表
func GroupList(c *core.GContent, uid uint64, targettype uint) []map[string]any {
	ret := []model.Group{}

	getDB(c).Order("create_at DESC").Find(&ret, "user_id=? AND target_type=?", uid, targettype)
	rr := []map[string]any{}
	for _, item := range ret {
		rr = append(rr, map[string]any{
			"group_id":   item.ID,
			"group_name": item.GroupName,
			"create_at":  item.CreateAt,
		})
	}
	return rr

}

// 分组添加条目
func GroupItemAdd(c *core.GContent, uid, gid uint64, uids ...uint64) error {
	gm := &model.Group{}
	db := getDB(c)
	db.Where(gid).First(gm)
	if gm.ID == 0 {
		return errors.New("群组不存在")
	}
	if gm.UserID != uid {
		return errors.New("群组不存在")
	}
	tx := db.Begin()
	ulist := []model.GroupItem{}
	tx.Find(&ulist, "group_id=? AND target_id IN ?", gid, uids)
	umap := map[uint64]any{}
	for _, item := range ulist {
		umap[item.TargetID] = 1
	}
	adds := 0
	for _, id := range uids {
		if _, ok := umap[id]; !ok {
			tx.Save(&model.GroupItem{
				GroupID:  gid,
				TargetID: id,
				CreateIP: c.GetIP(),
				UpdateIP: c.GetIP(),
			})
			adds += 1
		}
	}
	if adds > 0 {
		tx.Model(gm).Where(gm.ID).Update("items", gorm.Expr("items+?", adds))
	}
	tx.Commit()
	return nil
}

// 删除分组条目
func GroupItemDel(c *core.GContent, uid, gid uint64, uids ...uint64) error {
	gm := &model.Group{}
	db := getDB(c)
	db.Where(gid).First(gm)
	if gm.ID == 0 {
		return errors.New("群组不存在")
	}
	if gm.UserID != uid {
		return errors.New("群组不存在")
	}
	tx := db.Begin()
	ops := tx.Delete(&model.GroupItem{}, "group_id=? AND target_id IN ?", gid, uids).RowsAffected
	if ops > 0 {
		tx.Model(gm).Where(gm.ID).Update("items", gorm.Expr("items-?", ops))
	}
	tx.Commit()

	return nil
}

// 移动分组条目
func GroupItemMove(c *core.GContent, uid, oldgid, negid uint64, uids ...uint64) {
	GroupItemDel(c, uid, oldgid, uids...)
	GroupItemAdd(c, uid, negid, uids...)
}

// 分组成员列表
func GroupItems(c *core.GContent, uid, gid uint64, page, pageSize int) []map[string]any {
	gm := &model.Group{}
	db := getDB(c)
	db.Where(gid).First(gm)
	fr := []map[string]any{}
	if gm.ID == 0 || gm.UserID != uid {
		return fr
	}
	ret := []model.GroupItem{}
	db.Where("group_id=?", gid).Order("update_at DESC").Offset((page - 1) * pageSize).Limit(pageSize).Find(&ret)

	for _, item := range ret {
		fr = append(fr, map[string]any{
			"id":        item.TargetID,
			"create_at": item.CreateAt,
		})
	}

	return fr
}

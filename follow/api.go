package follow

import (
	"errors"

	"github.com/ghf-go/fleetness/account"
	"github.com/ghf-go/fleetness/core"
	"github.com/ghf-go/fleetness/follow/model"
	"gorm.io/gorm"
)

type followParam struct {
	ID uint64 `json:"user_id"`
}

// 关注
func apiFollowAction(c *core.GContent) {
	p := &followParam{}
	if e := c.BindJson(p); e != nil {
		c.FailJson(403, e.Error())
		return
	}
	fm := &model.FollowItem{}
	db := getDB(c)
	uid := c.GetUserID()
	db.First(fm, "user_id=? AND target_id=?", uid, p.ID)
	if fm.ID > 0 {
		c.SuccessJson("ok")
		return
	}
	ip := c.GetIP()
	e, ret := c.Tx(db, func(tx *gorm.DB) (error, any) {
		err := tx.Save(&model.FollowItem{
			UserID:   uid,
			TargetID: p.ID,
			CreateIP: ip,
			UpdateIP: ip,
		}).Error
		if err != nil {
			return err, ""
		}
		m := &model.Follow{}
		if tx.Model(m).Where("user_id=?", uid).Update("follows", gorm.Expr("follows+1")).RowsAffected <= 0 {
			if ee := tx.Save(&model.Follow{
				UserID:   uid,
				Follows:  1,
				CreateIP: ip,
				UpdateIP: ip,
			}).Error; ee != nil {
				return ee, ""
			}
		}
		if tx.Model(m).Where("user_id=?", p.ID).Update("fans", gorm.Expr("fans+1")).RowsAffected <= 0 {
			if ee := tx.Save(&model.Follow{
				UserID:   p.ID,
				Fans:     1,
				CreateIP: ip,
				UpdateIP: ip,
			}).Error; ee != nil {
				return ee, ""
			}
		}
		return nil, "ok"
	})
	if e != nil {
		c.FailJson(403, "操作失败")
		return
	}
	c.SuccessJson(ret)

}

// 取消关注
func apiUnFollowAction(c *core.GContent) {
	p := &followParam{}
	if e := c.BindJson(p); e != nil {
		c.FailJson(403, e.Error())
		return
	}
	fm := &model.FollowItem{}
	db := getDB(c)
	uid := c.GetUserID()
	db.First(fm, "user_id=? AND target_id=?", uid, p.ID)
	if fm.ID == 0 {
		c.SuccessJson("ok")
		return
	}
	e, ret := c.Tx(db, func(tx *gorm.DB) (error, any) {
		if ee := tx.Delete(&model.FollowItem{}, "user_id=? AND target_id=?", uid, p.ID).Error; ee != nil {
			return ee, ""
		}

		m := &model.Follow{}
		if tx.Model(m).Where("user_id=?", uid).Update("follows", gorm.Expr("follows-1")).RowsAffected <= 0 {
			return errors.New("操作失败"), ""
		}
		if tx.Model(m).Where("user_id=?", p.ID).Update("fans", gorm.Expr("fans-1")).RowsAffected <= 0 {
			return errors.New("操作失败"), ""
		}
		return nil, "ok"
	})
	if e != nil {
		c.FailJson(403, "操作失败")
		return
	}
	c.SuccessJson(ret)
}

type followListParam struct {
	ID uint64 `json:"user_id"`
	core.PageParam
}

// 关注列表
func apiFollowListAction(c *core.GContent) {
	p := &followListParam{}
	if e := c.BindJson(p); e != nil {
		c.FailJson(403, e.Error())
		return
	}
	if p.ID == 0 {
		p.ID = c.GetUserID()
	}
	ulist := []model.FollowItem{}
	getDB(c).Where("user_id=?", p.ID).Offset(p.GetOffset()).Limit(p.GetPageSize()).Order("create_at DESC").Find(&ulist)
	dd := []map[string]any{}
	for _, item := range ulist {
		dd = append(dd, map[string]any{
			"user_id":   item.TargetID,
			"create_at": item.CreateAt,
		})
	}
	c.SuccessJson(account.AppendUserBase(c, dd, "user_id", "user_info"))
}

// 粉丝列表
func apiFollowFanAction(c *core.GContent) {
	p := &followListParam{}
	if e := c.BindJson(p); e != nil {
		c.FailJson(403, e.Error())
		return
	}
	if p.ID == 0 {
		p.ID = c.GetUserID()
	}
	ulist := []model.FollowItem{}
	getDB(c).Where("target_id=?", p.ID).Offset(p.GetOffset()).Limit(p.GetPageSize()).Order("create_at DESC").Find(&ulist)
	dd := []map[string]any{}
	for _, item := range ulist {
		dd = append(dd, map[string]any{
			"user_id":   item.TargetID,
			"create_at": item.CreateAt,
		})
	}
	c.SuccessJson(account.AppendUserBase(c, dd, "user_id", "user_info"))
}

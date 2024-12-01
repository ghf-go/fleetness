package praise

import (
	"github.com/ghf-go/fleetness/core"
	"github.com/ghf-go/fleetness/praise/model"
	"gorm.io/gorm"
)

type praiseParams struct {
	TargetType uint   `json:"target_type"`
	TargetId   uint64 `json:"target_id"`
}

// 赞
func praiseAction(c *core.GContent) {
	req := &praiseParams{}
	if e := c.BindJson(req); e != nil {
		c.FailJson(403, c.Lang("client_param_error"))
		return
	}
	db := getDB(c)
	row := &model.Praise{}
	stat := &model.PraiseStat{}
	db.First(row, "user_id=? AND target_type=? AND target_id=?", c.GetUserID(), req.TargetType, req.TargetId)
	db.First(stat, "target_type=? AND target_id=?", req.TargetType, req.TargetId)
	if row.ID > 0 {
		c.SuccessJson(map[string]any{
			"count": stat.TargetCounts,
			"is_my": true,
		})
	} else {
		tx := db.Begin()
		if tx.Create(&model.Praise{
			UserID:     c.GetUserID(),
			TargetType: req.TargetType,
			TargetID:   req.TargetId,
			CreateIP:   c.GetIP(),
			UpdateIP:   c.GetIP(),
		}).Error != nil {
			tx.Rollback()
			c.FailJson(403, c.Lang("save_fail"))
			return
		}

		if tx.Model(stat).Where("target_type=? AND target_id=?", req.TargetType, req.TargetId).Update("target_counts", gorm.Expr("target_counts+?", 1)).RowsAffected == 0 && tx.Create(&model.PraiseStat{
			TargetType:   req.TargetType,
			TargetID:     req.TargetId,
			TargetCounts: 1,
			CreateIP:     c.GetIP(),
			UpdateIP:     c.GetIP(),
		}).Error != nil {
			tx.Rollback()
			c.FailJson(403, c.Lang("save_fail"))
			return
		}
		tx.Commit()
		db.First(stat, "target_type=? AND target_id=?", req.TargetType, req.TargetId)
		c.SuccessJson(map[string]any{
			"count": stat.TargetCounts,
			"is_my": true,
		})
	}

}

// 取消点赞
func unPraiseAction(c *core.GContent) {
	req := &praiseParams{}
	if e := c.BindJson(req); e != nil {
		c.FailJson(403, c.Lang("client_param_error"))
		return
	}
	db := getDB(c)
	row := &model.Praise{}
	stat := &model.PraiseStat{}
	db.First(row, "user_id=? AND target_type=? AND target_id=?", c.GetUserID(), req.TargetType, req.TargetId)
	db.First(stat, "target_type=? AND target_id=?", req.TargetType, req.TargetId)
	if row.ID == 0 {
		c.SuccessJson(map[string]any{
			"count": stat.TargetCounts,
			"is_my": false,
		})
	} else {
		tx := db.Begin()
		if tx.Delete(row, "user_id=? AND target_type=? AND target_id=?", c.GetUserID(), req.TargetType, req.TargetId).Error != nil {
			tx.Rollback()
			c.FailJson(403, c.Lang("save_fail"))
			return
		}

		if tx.Model(stat).Where("target_type=? AND target_id=?", req.TargetType, req.TargetId).Update("target_counts", gorm.Expr("target_counts-?", 1)).Error != nil {
			tx.Rollback()
			c.FailJson(403, c.Lang("save_fail"))
			return
		}
		tx.Commit()
		db.First(stat, "target_type=? AND target_id=?", req.TargetType, req.TargetId)
		c.SuccessJson(map[string]any{
			"count": stat.TargetCounts,
			"is_my": false,
		})
	}
}

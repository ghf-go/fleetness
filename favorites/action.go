package favorites

import (
	"github.com/ghf-go/fleetness/core"
	"github.com/ghf-go/fleetness/favorites/model"
	"gorm.io/gorm"
)

type favoriteParams struct {
	TargetType uint   `json:"target_type"`
	TargetId   uint64 `json:"target_id"`
}

// 收藏
func favoriteAction(c *core.GContent) {
	req := &favoriteParams{}
	if e := c.BindJson(req); e != nil {
		c.FailJson(403, "参数错误")
		return
	}
	db := getDB(c)
	row := &model.Favorite{}
	stat := &model.FavoriteStat{}
	db.First(row, "user_id=? AND target_type=? AND target_id=?", c.GetUserID(), req.TargetType, req.TargetId)
	db.First(stat, "target_type=? AND target_id=?", req.TargetType, req.TargetId)
	if row.ID > 0 {
		c.SuccessJson(map[string]any{
			"count": stat.TargetCounts,
			"is_my": true,
		})
	} else {
		tx := db.Begin()
		if tx.Create(&model.Favorite{
			UserID:     c.GetUserID(),
			TargetType: req.TargetType,
			TargetID:   req.TargetId,
			CreateIP:   c.GetIP(),
			UpdateIP:   c.GetIP(),
		}).Error != nil {
			tx.Rollback()
			c.FailJson(500, "操作失败")
			return
		}

		if tx.Model(stat).Where("target_type=? AND target_id=?", req.TargetType, req.TargetId).Update("target_counts", gorm.Expr("target_counts+?", 1)).RowsAffected == 0 && tx.Create(&model.FavoriteStat{
			TargetType:   req.TargetType,
			TargetID:     req.TargetId,
			TargetCounts: 1,
			CreateIP:     c.GetIP(),
			UpdateIP:     c.GetIP(),
		}).Error != nil {
			tx.Rollback()
			c.FailJson(500, "操作失败")
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

// 取消收藏
func unFavoriteAction(c *core.GContent) {
	req := &favoriteParams{}
	if e := c.BindJson(req); e != nil {
		c.FailJson(403, "参数错误")
		return
	}
	db := getDB(c)
	row := &model.Favorite{}
	stat := &model.FavoriteStat{}
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
			c.FailJson(500, "操作失败")
			return
		}

		if tx.Model(stat).Where("target_type=? AND target_id=?", req.TargetType, req.TargetId).Update("target_counts", gorm.Expr("target_counts-?", 1)).Error != nil {
			tx.Rollback()
			c.FailJson(500, "操作失败")
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

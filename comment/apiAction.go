package comment

import (
	"github.com/ghf-go/fleetness/comment/model"
	"github.com/ghf-go/fleetness/core"
	"gorm.io/gorm"
)

type commentParams struct {
	TargetType uint   `json:"target_type"`
	TargetId   uint64 `json:"target_id"`
	ReplyID    uint64 `json:"reply_id"`
	Content    string `json:"content"`
}

// 发布评论
func commentAction(c *core.GContent) {
	req := &commentParams{}
	if e := c.BindJson(req); e != nil {
		c.FailJson(403, "参数错误")
		return
	}
	db := getDB(c).Begin()
	if db.Create(&model.Comment{
		UserID:     c.GetUserID(),
		TargetType: req.TargetType,
		TargetID:   req.TargetId,
		CreateIP:   c.GetIP(),
		UpdateIP:   c.GetIP(),
		ReplyID:    req.ReplyID,
		Content:    req.Content,
	}).Error == nil && (db.Model(&model.CommentStat{}).Where("target_type=? AND target_id=?", req.TargetType, req.TargetId).Update("target_counts", gorm.Expr("target_counts+?", 1)).RowsAffected > 0 || db.Create(&model.CommentStat{
		TargetType:   req.TargetType,
		TargetID:     req.TargetId,
		TargetCounts: 1,
		CreateIP:     c.GetIP(),
		UpdateIP:     c.GetIP(),
	}).Error == nil) {
		db.Commit()
		c.SuccessJson("OK")
		return
	}
	db.Rollback()
	c.FailJson(500, "发布失败")

}

type commentListParams struct {
	TargetType uint   `json:"target_type"`
	TargetId   uint64 `json:"target_id"`
	ReplyID    uint64 `json:"reply_id"`
	Page       int    `json:"page"`
	PageSize   int    `json:"page_size"`
}

// 获取评论列表
func commentListAction(c *core.GContent) {
	req := &commentListParams{}
	if e := c.BindJson(req); e != nil {
		c.FailJson(403, "参数错误")
		return
	}
	if req.Page == 0 {
		req.Page = 1
	}
	if req.PageSize == 0 {
		req.PageSize = 20
	}
	clist := []model.Comment{}
	db := getDB(c).Offset((req.Page - 1) * req.PageSize).Limit(req.PageSize).Order("create_at DESC")
	if isSendAfterAudit {
		db.Where("target_type=? AND target_id=? AND reply_id=? AND (status<=20 OR (status<>100 AND user_id=?))", req.TargetType, req.TargetId, req.ReplyID, c.GetUserID()).Find(&clist)
	} else {
		db.Where("target_type=? AND target_id=? AND reply_id=? AND (status=10 OR (status<>100 AND user_id=?))", req.TargetType, req.TargetId, req.ReplyID, c.GetUserID()).Find(&clist)
	}
	uids := []uint64{}
	for _, item := range clist {
		uids = append(uids, item.UserID)
	}
	c.SuccessJson(clist)
}

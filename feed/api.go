package feed

import (
	"errors"
	"strings"

	"github.com/ghf-go/fleetness/blocklist"
	"github.com/ghf-go/fleetness/core"
	"github.com/ghf-go/fleetness/feed/model"
	"github.com/ghf-go/fleetness/follow"
	"gorm.io/gorm"
)

type apiFeedCreateActionParam struct {
	Title    string   `json:"title"`
	Imgs     []string `json:"imgs"`
	Content  string   `json:"content"`
	VoteItem []string `json:"vote_items"`
}

// 发布动态
func apiFeedCreateAction(c *core.GContent) {
	p := &apiFeedCreateActionParam{}
	if e := c.BindJson(p); e != nil {
		c.FailJson(403, "参数错误")
		return
	}
	isVote := false
	if len(p.VoteItem) > 0 {
		isVote = true
		if p.Title == "" {
			c.FailJson(403, "请输入投票标题")
			return
		}
	} else {
		if p.Content == "" && len(p.Imgs) == 0 {
			c.FailJson(403, "请输入内容或者上传照片")
			return
		}
	}
	fm := &model.Feed{
		CreateIP: c.GetIP(),
		UpdateIP: c.GetIP(),
		Title:    p.Title,
		Content:  p.Content,
		UserID:   c.GetUserID(),
		Imgs:     strings.Join(p.Imgs, ","),
	}
	db := getDB(c)
	if isVote {
		fm.FeedType = FEED_TYPE_VOTE
		e, _ := c.Tx(db, func(tx *gorm.DB) (error, any) {
			if tx.Save(fm).Error != nil {
				return errors.New("操作失败"), ""
			}
			for _, item := range p.VoteItem {
				if tx.Save(&model.FeedVote{
					FeedID:   fm.ID,
					Name:     item,
					CreateIP: c.GetIP(),
					UpdateIP: c.GetIP(),
				}).Error != nil {
					return errors.New("操作失败"), ""
				}
			}
			return nil, ""
		})
		if e != nil {
			c.FailJson(403, "操作失败")
			return
		}
	} else {
		fm.FeedType = FEED_TYPE_BLOG
		if db.Save(fm).Error != nil {
			c.FailJson(403, "操作失败")
			return
		}
	}
	c.SuccessJson(formatFeed(c, c.GetUserID(), *fm))
}

type apiFeedListActionParam struct {
	Act string `json:"act"`
	core.PageParam
}

// 动态列表
func apiFeedListAction(c *core.GContent) {
	p := &apiFeedListActionParam{}
	if e := c.BindJson(p); e != nil {
		c.FailJson(403, "参数错误")
		return
	}
	flist := []model.Feed{}
	switch p.Act {
	case "My":
		getDB(c).Where("user_id=?", c.GetUserID()).Order("id DESC").Offset(p.GetOffset()).Limit(p.GetPageSize()).Find(&flist)
	case "Follow":
		getDB(c).Where("user_id IN ? AND status = ?", follow.Follows(c, c.GetUserID()), core.STATUS_SUCCESS).Order("id DESC").Offset(p.GetOffset()).Limit(p.GetPageSize()).Find(&flist)
	default:
		getDB(c).Where("user_id NOT IN ? AND status=?", blocklist.BlockList(c, c.GetUserID()), core.STATUS_SUCCESS).Order("id DESC").Offset(p.GetOffset()).Limit(p.GetPageSize()).Find(&flist)
	}
	ret := formatFeedList(c, c.GetUserID(), flist)
	c.SuccessJson(ret)
}

type apiFeedVoteActionParam struct {
	Id    uint64   `json:"id"`
	Items []uint64 `json:"items"`
}

// 投票
func apiFeedVoteAction(c *core.GContent) {
	p := &apiFeedVoteActionParam{}
	if e := c.BindJson(p); e != nil || p.Id == 0 || len(p.Items) == 0 {
		c.FailJson(403, "参数错误")
		return
	}

	db := getDB(c)
	uid := c.GetUserID()
	vlist := []model.FeedVoteLog{}
	fm := &model.Feed{}

	db.Find(&vlist, "feed_id=? AND user_id=?", p.Id, uid)
	if len(vlist) > 0 {
		c.FailJson(403, "你已经投票")
		return
	}
	db.First(fm, p.Id)
	if fm.ID == 0 || fm.Status != core.STATUS_SUCCESS || (fm.FeedType != FEED_TYPE_VOTE && fm.FeedType != FEED_TYPE_MVOTE) {
		c.FailJson(403, "投票不存在或者已经被删除")
		return
	}
	e, _ := c.Tx(db, func(tx *gorm.DB) (error, any) {
		if tx.Model(&model.FeedVote{}).Where("feed_id=? AND id IN ?", p.Id, p.Items).Update("votes", gorm.Expr("votes+1")).RowsAffected == 0 {
			return errors.New("操作失败"), ""
		}
		for _, id := range p.Items {
			if tx.Save(&model.FeedVoteLog{
				FeedID:   p.Id,
				ItemID:   id,
				UserID:   uid,
				CreateIP: c.GetIP(),
				UpdateIP: c.GetIP(),
			}).Error != nil {
				return errors.New("操作失败"), ""
			}
		}
		return nil, ""
	})
	if e != nil {
		c.FailJson(403, "操作失败")
		return
	}
	c.SuccessJson(formatFeed(c, c.GetUserID(), *fm))
}

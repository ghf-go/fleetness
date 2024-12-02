package message

import (
	"errors"
	"fmt"
	"time"

	"github.com/ghf-go/fleetness/account"
	"github.com/ghf-go/fleetness/blocklist"
	"github.com/ghf-go/fleetness/core"
	"github.com/ghf-go/fleetness/message/model"
	"gorm.io/gorm"
)

type apiSendActionParam struct {
	ToUid   uint64 `json:"user_id"`
	Content string `json:"msg"`
}

// 发送消息
func apiSendAction(c *core.GContent) {
	p := &apiSendActionParam{}
	if e := c.BindJson(p); e != nil || p.ToUid < 1 || p.Content == "" {
		c.FailJson(403, c.Lang("client_param_error"))
		return
	}
	if blocklist.InBlockList(c, p.ToUid) {
		c.FailJson(403, c.Lang("save_fail"))
		return
	}
	uid := c.GetUserID()
	mkey := fmt.Sprintf("%020d-%020d", min(uid, p.ToUid), max(uid, p.ToUid))
	if e, _ := c.Tx(getDB(c), func(tx *gorm.DB) (error, any) {
		if e := tx.Save(&model.MessageContent{
			UpdateIP: c.GetIP(),
			CreateIP: c.GetIP(),
			FromID:   c.GetUserID(),
			RecvID:   p.ToUid,
			Content:  p.Content,
			Mkey:     mkey,
		}).Error; e != nil {
			return e, ""
		}
		cc := []rune(p.Content)
		summary := p.Content
		if len(cc) > 30 {
			summary = string(cc[:30])
		}
		if tx.Model(&model.MessageUser{}).Where("user_id=? AND to_id=?", c.GetUserID(), p.ToUid).Updates(map[string]any{
			"last_time": time.Now().UnixMicro(),
			"last_uid":  c.GetUserID(),
			"last_msg":  summary,
		}).RowsAffected == 0 && tx.Save(&model.MessageUser{
			UserID:   c.GetUserID(),
			UpdateIP: c.GetIP(),
			CreateIP: c.GetIP(),
			ToID:     p.ToUid,
			LastTime: time.Now().UnixMicro(),
			LastUid:  c.GetUserID(),
			LastMsg:  summary,
			Mkey:     mkey,
		}).Error != nil {
			return errors.New(c.Lang("save_fail")), ""
		}
		if tx.Model(&model.MessageUser{}).Where("user_id=? AND to_id=?", p.ToUid, c.GetUserID()).Updates(map[string]any{
			"last_time": time.Now().UnixMicro(),
			"last_uid":  c.GetUserID(),
			"last_msg":  summary,
			"un_reads":  gorm.Expr("un_reads+1"),
		}).RowsAffected == 0 && tx.Save(&model.MessageUser{
			ToID:     c.GetUserID(),
			UpdateIP: c.GetIP(),
			CreateIP: c.GetIP(),
			UserID:   p.ToUid,
			LastTime: time.Now().UnixMicro(),
			LastUid:  c.GetUserID(),
			LastMsg:  summary,
			Mkey:     mkey,
			UnReads:  1,
		}).Error != nil {
			return errors.New(c.Lang("save_fail")), ""
		}
		return nil, ""
	}); e != nil {
		c.FailJson(403, c.Lang("save_fail"))
		return
	}
	c.SuccessJson("success")
}

type apiMessageListActionParam struct {
	LastId uint64 `json:"last_id"`
	HeadId uint64 `json:"head_id"`
	MKey   string `json:"msg_key"`
}

// 消息列表
func apiMessageListAction(c *core.GContent) {
	p := &apiMessageListActionParam{}
	if e := c.BindJson(p); e != nil || p.MKey == "" {
		c.FailJson(403, c.Lang("client_param_error"))
		return
	}
	uid := c.GetUserID()
	mlisg := []model.MessageContent{}

	if p.HeadId > 0 {
		getDB(c).Where("mkey=? AND id<?", p.MKey, p.HeadId).Order("id DESC").Limit(50).Find(&mlisg)
	} else {
		getDB(c).Where("mkey=? AND id>?", p.MKey, p.LastId).Order("id DESC").Limit(50).Find(&mlisg)
		getDB(c).Model(&model.MessageContent{}).Where("mkey=? AND recv_id=?", p.MKey, uid).Update("is_read", 1)
		getDB(c).Model(&model.MessageUser{}).Where("mkey=? AND user_id=?", p.MKey, uid).Update("un_reads", 0)
	}
	tml := []model.MessageContent{}
	for i := len(mlisg) - 1; i >= 0; i-- {
		tml = append(tml, mlisg[i])
	}

	ret := []map[string]any{}
	for _, item := range tml {
		user_id := item.FromID
		my_id := item.RecvID
		if user_id == uid {
			user_id = item.RecvID
			my_id = uid
		}
		ret = append(ret, map[string]any{
			"id":        item.ID,
			"msg":       item.Content,
			"is_read":   item.IsRead,
			"create_at": item.CreateAt,
			"user_id":   user_id,
			"my_id":     my_id,
		})
	}

	ret = account.AppendUserBase(c, ret, "user_id", "user_info")
	ret = account.AppendUserBase(c, ret, "my_id", "my_info")
	c.SuccessJson(ret)
}

type apiChatListActionParam struct {
	LastTime int64 `json:"last_time"`
}

// 回话列表
func apiChatListAction(c *core.GContent) {
	p := &apiChatListActionParam{}
	if e := c.BindJson(p); e != nil {
		c.FailJson(403, c.Lang("client_param_error"))
		return
	}
	mlist := []model.MessageUser{}
	getDB(c).Where("user_id=? AND last_time>?", c.GetUserID(), p.LastTime).Order("last_time DESC").Find(&mlist)
	ret := []map[string]any{}
	for _, item := range mlist {
		ret = append(ret, map[string]any{
			"msg_key":   item.Mkey,
			"un_reads":  item.UnReads,
			"user_id":   item.ToID,
			"last_msg":  item.LastMsg,
			"last_time": item.LastTime,
			"last_uid":  item.LastUid,
		})
	}
	ret = account.AppendUserBase(c, ret, "user_id", "user_info")
	ret = account.AppendUserBase(c, ret, "last_uid", "last_user_info")
	c.SuccessJson(ret)
}

type apiChatDelActionParam struct {
	ToUid uint64 `json:"user_id"`
}

// 删除回话
func apiChatDelAction(c *core.GContent) {
	p := &apiChatDelActionParam{}
	if e := c.BindJson(p); e != nil || p.ToUid == 0 {
		c.FailJson(403, c.Lang("client_param_error"))
		return
	}
	if getDB(c).Delete(&model.MessageUser{}, "user_id=? AND to_id=?", c.GetUserID(), p.ToUid).Error != nil {
		c.FailJson(403, c.Lang("save_fail"))
	}
	c.SuccessJson("success")
}

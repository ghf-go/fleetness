package novel

import (
	"errors"

	"github.com/ghf-go/fleetness/core"
	"github.com/ghf-go/fleetness/novel/model"
	"gorm.io/gorm"
)

// 作者详情
func apiInfoAuthorAction(c *core.GContent) {}

// 小说列表
func apiNovelListAction(c *core.GContent) {}

// 小说详情
func apiNovelInfoAction(c *core.GContent) {}

// 阅读
func apiNovelReadAction(c *core.GContent) {}

// 阅读历史记录
func apiNovelHistoryAction(c *core.GContent) {}

// 订阅
func apiSubscribeAction(c *core.GContent) {
	p := &core.ApiParam{}
	if e := c.BindJson(p); e != nil || p.ID == 0 {
		c.FailJson(403, "参数错误")
		return
	}
	uid := c.GetUserID()
	db := getDB(c)
	row := &model.NovelSubscribe{}
	db.First(row, "user_id=? AND novel_id=?", uid, p.ID)
	if row.ID == 0 {
		row.UserID = uid
		row.NovelID = p.ID
		row.CreateIP = c.GetIP()
		row.UpdateIP = c.GetIP()
		c.Tx(db, func(tx *gorm.DB) (error, any) {
			if tx.Save(row).Error == nil && tx.Model(&model.NovelInfo{}).Where(p.ID).Update("subscribe", gorm.Expr("subscribe+1")).Error == nil {
				return nil, ""
			}
			return errors.New("订阅失败"), ""
		})
	}
	c.SuccessJson("success")

}

// 取消订阅
func apiUnSubscribeAction(c *core.GContent) {
	p := &core.ApiParam{}
	if e := c.BindJson(p); e != nil || p.ID == 0 {
		c.FailJson(403, "参数错误")
		return
	}

	uid := c.GetUserID()
	db := getDB(c)
	row := &model.NovelSubscribe{}
	db.First(row, "user_id=? AND novel_id=?", uid, p.ID)
	if row.ID > 0 {
		c.Tx(db, func(tx *gorm.DB) (error, any) {
			if tx.Delete(row, row.ID).Error == nil && tx.Model(&model.NovelInfo{}).Where(p.ID).Update("subscribe", gorm.Expr("subscribe-1")).Error == nil {
				return nil, ""
			}
			return errors.New("取消订阅失败"), ""
		})
	}
	c.SuccessJson("success")
}

// 订阅列表
func apiSubscribeListAction(c *core.GContent) {
	p := &core.ApiParam{}
	if e := c.BindJson(p); e != nil || p.ID == 0 {
		c.FailJson(403, "参数错误")
		return
	}

	uid := c.GetUserID()
	db := getDB(c)
	list := []model.NovelSubscribe{}
	db.Where("user_id=?", uid).Order("update_at DESC").Offset(p.GetOffset()).Limit(p.GetPageSize()).Find(&list)

	nids := []uint64{}
	for _, item := range list {
		nids = append(nids, item.NovelID)
	}

}

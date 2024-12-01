package lottery

import (
	"math/rand"
	"time"

	"github.com/ghf-go/fleetness/account"
	"github.com/ghf-go/fleetness/core"
	"github.com/ghf-go/fleetness/lottery/model"
)

type apiLotteryParam struct {
	ID uint64 `json:"id"`
	core.PageParam
}

// 获取抽奖配置信息
func apiGetLotteryInfoAction(c *core.GContent) {
	p := &apiLotteryParam{}
	if e := c.BindJson(p); e != nil || p.ID < 1 {
		c.FailJson(403, c.Lang("client_param_error"))
		return
	}
	info := &model.Lottery{}
	db := getDB(c)
	db.First(info, p.ID)
	if info.ID == 0 || info.IsOpen == 0 {
		c.FailJson(403, "活动已经下线")
		return
	}
	itemList := []model.LotteryItem{}
	db.Find(&itemList, "lottery_id=? AND is_online=1", p.ID)

	items := []map[string]any{}
	for _, item := range itemList {
		items = append(items, map[string]any{
			"id":      item.ID,
			"name":    item.Name,
			"content": item.Content,
			"logo":    item.Logo,
		})
	}

	ret := map[string]any{
		"info": map[string]any{
			"id":        info.ID,
			"name":      info.Name,
			"logo":      info.Logo,
			"kind":      info.Kind,
			"content":   info.Content,
			"day_limit": info.DayLimit,
		},
		"items": items,
	}

	c.SuccessJson(ret)

}

// 抽奖
func apiLotteryAction(c *core.GContent) {
	p := &apiLotteryParam{}
	if e := c.BindJson(p); e != nil || p.ID < 1 {
		c.FailJson(403, c.Lang("client_param_error"))
		return
	}
	info := &model.Lottery{}
	db := getDB(c)
	db.First(info, p.ID)
	if info.ID == 0 || info.IsOpen == 0 {
		c.FailJson(403, "活动已经下线")
		return
	}
	uid := c.GetUserID()
	if info.DayLimit > 0 {
		days := int64(0)
		db.Model(&model.LotteryLog{}).Where("lottery_id=? AND user_id=? AND day=?", p.ID, uid, time.Now()).Count(&days)
		if days >= int64(info.DayLimit) {
			c.FailJson(403, "今日已经达到抽奖次数，请明天重试")
			return
		}
	}
	itemList := []model.LotteryItem{}
	db.Find(&itemList, "lottery_id=? AND is_online=1", p.ID)
	sum := 0
	im := map[int]uint64{}
	for _, item := range itemList {
		im[sum] = item.ID
		sum += item.Rate
	}
	rc := rand.Intn(sum)
	okid := uint64(0)
	for k, v := range im {
		if rc >= k {
			okid = v
		}
	}
	if okid == 0 {
		c.FailJson(403, "非常遗憾没有中奖")
		return
	}
	ctime := time.Now()
	db.Save(&model.LotteryLog{
		UserID:    uid,
		LotteryID: p.ID,
		ItemID:    okid,
		Day:       &ctime,
	})
	account.UserCashLog(c, uid, account.CASH_SCORE, 100, "抽奖")
	for _, item := range itemList {
		if item.ID == okid {
			c.SuccessJson(map[string]any{
				"id":      item.ID,
				"name":    item.Name,
				"content": item.Content,
				"logo":    item.Logo,
			})
			return
		}
	}
}

// 我的抽奖记录
func apiLotteryLogAction(c *core.GContent) {
	p := &apiLotteryParam{}
	if e := c.BindJson(p); e != nil || p.ID < 1 {
		c.FailJson(403, c.Lang("client_param_error"))
		return
	}
	db := getDB(c)
	uid := c.GetUserID()
	itemList := []model.LotteryItem{}
	db.Find(&itemList, "lottery_id=? AND is_online=1", p.ID)
	logLost := []model.LotteryLog{}
	db.Where("lottery_id=? AND user_id=?", p.ID, uid).Order("id DESC").Offset(p.GetOffset()).Limit(p.GetPageSize()).Find(&logLost)
	itemMap := map[uint64]map[string]any{}
	for _, item := range itemList {
		itemMap[item.ID] = map[string]any{
			"id":      item.ID,
			"name":    item.Name,
			"content": item.Content,
			"logo":    item.Logo,
		}
	}
	ret := []map[string]any{}
	for _, item := range logLost {
		ret = append(ret, map[string]any{
			"day":       item.Day,
			"create_at": item.CreateAt,
			"item":      itemMap[item.ItemID],
		})
	}

	c.SuccessJson(ret)
}

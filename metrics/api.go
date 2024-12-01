package metrics

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/ghf-go/fleetness/core"
	"github.com/ghf-go/fleetness/core/utils"
	"github.com/ghf-go/fleetness/metrics/model"
	"gorm.io/gorm"
)

type apiTimeSyncActionParam struct {
	Day       int64 `json:"day"`
	UnixMilli int64 `json:"unix_milli"`
}

// 时间同步
func apiTimeSyncAction(c *core.GContent) {
	p := &apiTimeSyncActionParam{}
	c.BindJson(p)
	cday := utils.DayUnixMilli()
	subTime := int64(0)
	if p.Day > 0 {
		subTime = cday - p.Day
	}
	if subTime < 0 { //时间大约服务器时间
		p.UnixMilli = p.Day + subTime
	} else { //小于服务器时间
		p.UnixMilli = p.Day
	}
	c.SuccessJson(map[string]any{
		"day":       cday,
		"last_time": p.UnixMilli,
	})
}

type apiUploadActionParam struct {
	CurDay   int64                        `json:"cur_day"`
	Platform string                       `json:"platform"`
	Data     map[string]map[string]string `json:"data"`
}

// 上报数据
func apiUploadAction(c *core.GContent) {
	p := &apiUploadActionParam{}
	if e := c.BindJson(p); e != nil {
		c.FailJson(403, "参数错误"+e.Error())
		return
	}
	db := getDB(c)
	timeSub := utils.DayUnixMilli() - p.CurDay

	confKeys := []string{}
	for k, _ := range p.Data {
		confKeys = append(confKeys, k)
	}

	clist := []model.MetricsConf{}
	db.Find(&clist, "conf_key IN ?", confKeys)
	kmap := map[string]uint64{}
	for _, item := range clist {
		kmap[item.ConfKey] = item.ID
	}

	saveData := map[string]*model.MetricsStat{}
	for key, item := range p.Data {
		cid, ok := kmap[key]
		if !ok {
			continue
		}
		for t, ev := range item {
			tt, _ := strconv.ParseInt(t, 10, 64)
			dt := utils.UnixMilli2Time(tt + timeSub)
			// day := utils.DayUnixMilli(dt)
			kk := fmt.Sprintf("%d_%d", cid, dt.Format(utils.T_DATE))
			vs := 0
			cs := 0
			ev = strings.ToLower(ev)
			if strings.HasPrefix(ev, "v") {
				vs = 1
			}
			if strings.HasPrefix(ev, "c") {
				cs = 1
			}
			if row, ok := saveData[kk]; ok {
				if vs > 0 {
					row.UserViews = 1
				}
				if cs > 0 {
					row.UserClicks = 1
				}
				row.Clicks += uint(cs)
				row.Views += uint(vs)
				saveData[kk] = row
			} else {
				srow := &model.MetricsStat{
					CreateIP:   c.GetIP(),
					UpdateIP:   c.GetIP(),
					ConfID:     cid,
					Platform:   p.Platform,
					Week:       0,
					Year:       uint(dt.Year()),
					Month:      int(dt.Month()),
					Day:        dt.Day(),
					Date:       dt,
					Views:      uint(vs),
					Clicks:     uint(cs),
					UserViews:  0,
					UserClicks: 0,
				}
				if vs > 0 {
					srow.UserViews = 1
				}
				if cs > 0 {
					srow.UserClicks = 1
				}
				saveData[kk] = srow
			}

		}
	}
	lrow := &model.MetricsStat{}
	db.First(lrow)
	stime := utils.UnixMilli2Time(0)
	if lrow.ID > 0 {
		stime = lrow.Date
	}

	for _, item := range saveData {
		item.Date = utils.UnixMilli2Time(utils.DayUnixMilli(item.Date))
		item.Week = utils.SubWeeks(item.Date, stime)
		if db.Model(&model.MetricsStat{}).Where("conf_id=? AND platform=? AND date=?", item.ConfID, item.Platform, item.Date.Format(utils.T_DATE)).Updates(map[string]any{
			"views":       gorm.Expr("views+?", item.Views),
			"clicks":      gorm.Expr("clicks+?", item.Clicks),
			"user_views":  gorm.Expr("user_views+?", item.UserViews),
			"user_clicks": gorm.Expr("user_clicks+?", item.UserClicks),
		}).RowsAffected == 0 {
			db.Save(item)
		}
	}
	c.SuccessJson("success")

}

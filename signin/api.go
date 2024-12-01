package signin

import (
	"time"

	"github.com/ghf-go/fleetness/core"
	"github.com/ghf-go/fleetness/signin/model"
)

// 签到详情
func apiSignInfoAction(c *core.GContent) {
	cday := time.Now()
	cy, cm, cd := cday.Date()
	row := &model.Signin{}
	db := getDB(c)
	db.Where("user_id=?", c.GetUserID()).Order("id DESC").First(row)
	ret := []map[string]any{}
	unsignIndex := 0
	if row.ID > 0 {
		y, m, d := row.Day.Add(86400 * time.Second).Date()
		ly, lm, ld := row.Day.Date()
		if (cy == y && cm == m && cd == d) || (cy == ly && cm == lm && cd == ld) {
			for i := int(row.Continued); i >= 0; i-- {
				dd := cday.Add(time.Duration(i) * -86400 * time.Second)
				ret = append(ret, map[string]any{
					"day":    dd.Format("2006-01-02"),
					"signed": true,
				})
			}
			unsignIndex = int(row.Continued)
		}
	}
	for i := unsignIndex; i < days; i++ {
		dd := cday.Add(time.Duration(i) * 86400 * time.Second)
		ret = append(ret, map[string]any{
			"day":    dd.Format("2006-01-02"),
			"signed": false,
		})
	}
	c.SuccessJson(ret)
}

// 签到
func apiSignAction(c *core.GContent) {
	row := &model.Signin{}
	db := getDB(c)
	db.Where("user_id=?", c.GetUserID()).Order("id DESC").First(row)
	nrow := &model.Signin{
		CreateIP: c.GetIP(),
		UpdateIP: c.GetIP(),
		UserID:   c.GetUserID(),
		Day:      time.Now(),
	}
	if row.ID > 0 {
		nrow.Times = row.Times + 1
		y, m, d := row.Day.Add(86400 * time.Second).Date()
		ny, nm, nd := nrow.Day.Date()
		if y > ny || m > nm || d > nd {
			c.FailJson(403, "已经签到")
			return
		}
		if y == ny && m == nm && d == nd {
			nrow.Continued = row.Continued + 1
		}
		if nrow.Continued > uint(days) {
			nrow.Continued = 1
		}
	}
	if db.Save(nrow).Error == nil {
		callHandle(c.GetUserID(), nrow.Times, nrow.Continued)
		c.SuccessJson("success")
		return
	}
	c.FailJson(403, c.Lang("save_fail"))
}

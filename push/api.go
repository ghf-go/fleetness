package push

import (
	"github.com/ghf-go/fleetness/core"
	"github.com/ghf-go/fleetness/push/model"
)

type apiRegDeviceActionParam struct {
	Platform string `json:"platform"`
	Board    string `json:"board"`
	Channel  string `json:"channel"`
	Token    string `json:"token"`
	OsVer    string `json:"os_ver"`
	AppVer   string `json:"app_ver"`
	WgtVer   string `json:"wgt_ver"`
}

// 上报用户设备
func apiRegDeviceAction(c *core.GContent) {
	p := &apiRegDeviceActionParam{}
	if e := c.BindJson(p); e != nil || p.Platform == "" || p.Channel == "" || p.Token == "" {
		c.FailJson(403, c.Lang("client_param_error"))
		return
	}
	// PushAllSse("asdf", "asdf")
	dm := &model.AppDevice{}
	getDB(c).First(dm, "channel=? AND token=?", p.Channel, p.Token)
	if dm.ID == 0 {
		getDB(c).Save(&model.AppDevice{
			CreateIP: c.GetIP(),
			UpdateIP: c.GetIP(),
			UserID:   c.GetUserID(),
			Platform: p.Platform,
			Board:    p.Board,
			Channel:  p.Channel,
			Token:    p.Token,
			OsVer:    p.OsVer,
			AppVer:   p.AppVer,
			WgtVer:   p.WgtVer,
		})
	} else {
		getDB(c).Model(dm).Where(dm.ID).Updates(map[string]any{
			"update_ip": c.GetIP(),
			"user_id":   c.GetUserID(),
			// "channel":p.Channel,
			"app_ver": p.AppVer,
			"wgt_ver": p.WgtVer,
			"os_ver":  p.OsVer,
		})
	}
	c.SuccessJson("OK")
}

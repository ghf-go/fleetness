package api

import (
	"github.com/ghf-go/fleetness/account/logic"
	"github.com/ghf-go/fleetness/account/model"
	"github.com/ghf-go/fleetness/core"
	"github.com/ghf-go/fleetness/core/utils"
)

type bindAccountParams struct {
	BindType int    `json:"bind_type"`
	Name     string `json:"name"`
	Code     string `json:"code"`
}

// 绑定账号
func BindAccountAction(c *core.GContent) {
	p := &bindAccountParams{}
	if e := c.BindJson(p); e != nil {
		c.FailJson(403, e.Error())
		return
	}
	if p.BindType == logic.TYPE_EMAIL || p.BindType == logic.TYPE_MOBILE {
		if !utils.VerifyCode(logic.GetCahce(c), p.Code, p.Name) {
			c.FailJson(403, "验证码错误")
			return
		}
	}
	ub := &model.UserBind{}
	db := logic.GetDB(c)
	db.First(ub, "bind_val=?", p.Name)
	if ub.UserID > 0 {
		ub.UserID = c.GetUserID()
		ub.UpdateIP = c.GetIP()
		db.Save(ub)
	} else {
		ub.UserID = c.GetUserID()
		ub.UpdateIP = c.GetIP()
		ub.CreateIP = c.GetIP()
		ub.BindVal = p.Name
		ub.BindType = p.BindType
		db.Create(ub)
	}
	c.SuccessJson("ok")
}

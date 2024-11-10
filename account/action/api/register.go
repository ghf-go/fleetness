package api

import (
	"github.com/ghf-go/fleetness/account/logic"
	"github.com/ghf-go/fleetness/account/model"
	"github.com/ghf-go/fleetness/core"
	"github.com/ghf-go/fleetness/core/utils"
)

type registerParams struct {
	Name string `json:"name"`
	Pass string `json:"pass"`
	Code string `json:"code"`
}

// 注册接口
func RegisterAction(c *core.GContent) {
	parmas := &registerParams{}
	if e := c.BindJson(parmas); e != nil {
		c.FailJson(500, e.Error())
		return
	}
	if parmas.Name == "" || parmas.Pass == "" || parmas.Code == "" {
		c.FailJson(403, "参数错误")
		return
	}

	if !utils.VerifyCode(logic.GetCahce(c), parmas.Code, parmas.Name) {
		c.FailJson(403, "验证码错误")
		return
	}

	ub := &model.UserBind{}
	logic.GetDB(c).First(ub, "bind_val = ?", parmas.Name)
	if ub.UserID > 0 {
		c.FailJson(403, "账号已存在")
		return
	}

	disname := ""
	bindType := 0
	if utils.IsMobile(parmas.Name) {
		bindType = logic.TYPE_MOBILE
		disname = utils.HideMobile(parmas.Name)
	} else if utils.IsEmail(parmas.Name) {
		bindType = logic.TYPE_EMAIL
		disname = utils.HideEmail(parmas.Name)
	} else {
		c.FailJson(403, "参数错误")
		return
	}
	sign := utils.RandStr(16)

	userModel := &model.User{
		NickName: disname,
		PassSign: sign,
		Passwd:   logic.Passwd(parmas.Pass, sign),
		CreateIP: c.GetIP(),
		UpdateIP: c.GetIP(),
	}
	tx := logic.GetDB(c).Begin()
	tx.Create(userModel)
	if userModel.ID <= 0 {
		tx.Rollback()
		c.FailJson(403, "注册失败")
		return
	}
	bindUser := &model.UserBind{
		CreateIP: c.GetIP(),
		UpdateIP: c.GetIP(),
		UserID:   userModel.ID,
		BindVal:  parmas.Name,
		BindType: bindType,
	}
	tx.Create(bindUser)
	if bindUser.ID <= 0 {
		tx.Rollback()
		c.FailJson(403, "注册失败")
		return
	}
	tx.Commit()
	c.SuccessJson("注册成功")
}

package account

import (
	"errors"

	"github.com/ghf-go/fleetness/account/model"
	"github.com/ghf-go/fleetness/core"
	"github.com/ghf-go/fleetness/core/utils"
)

func createUser(c *core.GContent, name, pass string) error {
	ub := &model.UserBind{}
	getDB(c).First(ub, "bind_val = ?", name)
	if ub.UserID > 0 {
		return errors.New("账号已存在")
	}

	disname := ""
	bindType := 0
	if utils.IsMobile(name) {
		bindType = TYPE_MOBILE
		disname = utils.HideMobile(name)
	} else if utils.IsEmail(name) {
		bindType = TYPE_EMAIL
		disname = utils.HideEmail(name)
	} else {
		return errors.New("参数错误")
	}
	sign := utils.RandStr(16)

	userModel := &model.User{
		NickName: disname,
		PassSign: sign,
		Passwd:   passwd(pass, sign),
		CreateIP: c.GetIP(),
		UpdateIP: c.GetIP(),
	}
	tx := getDB(c).Begin()
	tx.Create(userModel)
	if userModel.ID <= 0 {
		tx.Rollback()
		return errors.New("注册失败")
	}
	bindUser := &model.UserBind{
		CreateIP: c.GetIP(),
		UpdateIP: c.GetIP(),
		UserID:   userModel.ID,
		BindVal:  name,
		BindType: bindType,
	}
	tx.Create(bindUser)
	if bindUser.ID <= 0 {
		tx.Rollback()
		return errors.New("账号已存在")
	}
	tx.Commit()
	return nil
}

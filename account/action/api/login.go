package api

import (
	"fmt"

	"github.com/ghf-go/fleetness/account/logic"
	"github.com/ghf-go/fleetness/account/model"
	"github.com/ghf-go/fleetness/core"
)

type loginByPass struct {
	Name string `json:"name"`
	Pass string `json:"pass"`
}

// 密码登录
func LoginByPassAction(c *core.GContent) {
	parmas := &loginByPass{}
	if e := c.BindJson(parmas); e != nil {
		c.FailJson(500, e.Error())
		return
	}
	if parmas.Name == "" || parmas.Pass == "" {
		c.FailJson(403, "参数错误")
		return
	}
	ub := &model.UserBind{}
	logic.GetDB(c).First(ub, "bind_val = ?", parmas.Name)
	if ub.UserID == 0 {
		c.FailJson(403, "账号或者密码错误")
		return
	}
	user := &model.User{}
	logic.GetDB(c).First(user, ub.UserID)
	if user.ID == 0 {
		c.FailJson(403, "账号或者密码错误")
		return
	}
	if logic.Passwd(parmas.Pass, user.PassSign) != user.Passwd {
		c.FailJson(403, "账号或者密码错误")
		return
	}
	c.SetUserID(fmt.Sprintf("%d", user.ID))
	c.SuccessJson("登录成功")

}

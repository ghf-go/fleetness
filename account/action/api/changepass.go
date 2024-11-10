package api

import (
	"github.com/ghf-go/fleetness/account/logic"
	"github.com/ghf-go/fleetness/account/model"
	"github.com/ghf-go/fleetness/core"
	"github.com/ghf-go/fleetness/core/utils"
)

type changePass struct {
	Pass    string `json:"pass"`
	OldPass string `json:"old_pass"`
}

// 修改密码
func ChangePassAction(c *core.GContent) {
	parmas := &changePass{}
	if e := c.BindJson(parmas); e != nil {
		c.FailJson(500, e.Error())
		return
	}
	if parmas.OldPass == "" || parmas.Pass == "" {
		c.FailJson(403, "参数错误")
		return
	}
	if parmas.OldPass == parmas.Pass {
		c.FailJson(403, "新旧密码不能一样")
		return
	}
	user := &model.User{}
	logic.GetDB(c).First(user, c.GetUserID())
	if logic.Passwd(parmas.Pass, user.PassSign) != user.Passwd {
		c.FailJson(403, "旧密码错误")
		return
	}
	sign := utils.RandStr(16)
	if logic.GetDB(c).Model(user).Where("id=?", c.GetUserID()).Updates(map[string]any{
		"pass_sign": sign,
		"passwd":    logic.Passwd(parmas.Pass, sign),
		"update_ip": c.GetIP(),
	}).RowsAffected > 0 {
		c.SuccessJson("ok")
		return
	}
	c.FailJson(403, "修改失败")
}

package account

import (
	"fmt"

	"github.com/ghf-go/fleetness/account/model"
	"github.com/ghf-go/fleetness/core"
	"github.com/ghf-go/fleetness/core/utils"
)

type loginByPass struct {
	Name string `json:"name"`
	Pass string `json:"pass"`
}

// 密码登录
func loginByPassAction(c *core.GContent) {
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
	getDB(c).First(ub, "bind_val = ?", parmas.Name)
	if ub.UserID == 0 {
		c.FailJson(403, "账号或者密码错误")
		return
	}
	user := &model.User{}
	getDB(c).First(user, ub.UserID)
	if user.ID == 0 {
		c.FailJson(403, "账号或者密码错误")
		return
	}
	if passwd(parmas.Pass, user.PassSign) != user.Passwd {
		c.FailJson(403, "账号或者密码错误")
		return
	}
	c.SetUserID(fmt.Sprintf("%d", user.ID))
	c.SuccessJson("登录成功")

}

type registerParams struct {
	Name string `json:"name"`
	Pass string `json:"pass"`
	Code string `json:"code"`
}

// 注册接口
func registerAction(c *core.GContent) {
	parmas := &registerParams{}
	if e := c.BindJson(parmas); e != nil {
		c.FailJson(500, e.Error())
		return
	}
	if parmas.Name == "" || parmas.Pass == "" || parmas.Code == "" {
		c.FailJson(403, "参数错误")
		return
	}

	if !utils.VerifyCode(getCahce(c), parmas.Code, parmas.Name) {
		c.FailJson(403, "验证码错误")
		return
	}

	ub := &model.UserBind{}
	getDB(c).First(ub, "bind_val = ?", parmas.Name)
	if ub.UserID > 0 {
		c.FailJson(403, "账号已存在")
		return
	}

	disname := ""
	bindType := 0
	if utils.IsMobile(parmas.Name) {
		bindType = TYPE_MOBILE
		disname = utils.HideMobile(parmas.Name)
	} else if utils.IsEmail(parmas.Name) {
		bindType = TYPE_EMAIL
		disname = utils.HideEmail(parmas.Name)
	} else {
		c.FailJson(403, "参数错误")
		return
	}
	sign := utils.RandStr(16)

	userModel := &model.User{
		NickName: disname,
		PassSign: sign,
		Passwd:   passwd(parmas.Pass, sign),
		CreateIP: c.GetIP(),
		UpdateIP: c.GetIP(),
	}
	tx := getDB(c).Begin()
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

type changePass struct {
	Pass    string `json:"pass"`
	OldPass string `json:"old_pass"`
}

// 修改密码
func changePassAction(c *core.GContent) {
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
	getDB(c).First(user, c.GetUserID())
	if passwd(parmas.OldPass, user.PassSign) != user.Passwd {
		c.FailJson(403, "旧密码错误")
		return
	}
	sign := utils.RandStr(16)
	if getDB(c).Model(user).Where("id=?", c.GetUserID()).Updates(map[string]any{
		"pass_sign": sign,
		"passwd":    passwd(parmas.Pass, sign),
		"update_ip": c.GetIP(),
	}).RowsAffected > 0 {
		c.SuccessJson("ok")
		return
	}
	c.FailJson(403, "修改失败")
}

type bindAccountParams struct {
	BindType int    `json:"bind_type"`
	Name     string `json:"name"`
	Code     string `json:"code"`
}

// 绑定账号
func bindAccountAction(c *core.GContent) {
	p := &bindAccountParams{}
	if e := c.BindJson(p); e != nil {
		c.FailJson(403, e.Error())
		return
	}
	if p.BindType == TYPE_EMAIL || p.BindType == TYPE_MOBILE {
		if !utils.VerifyCode(getCahce(c), p.Code, p.Name) {
			c.FailJson(403, "验证码错误")
			return
		}
	}
	ub := &model.UserBind{}
	db := getDB(c)
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

// 获取账号信息
func getUserInfoAction(c *core.GContent) {
	list := []model.UserInfo{}
	getDB(c).Find(&list, "user_id=?", c.GetUserID())
	ret := map[string]string{}
	for _, item := range list {
		ret[item.Ukey] = item.Newval
	}
	c.SuccessJson(ret)
}

// 保存账号信息
func setUserInfoAction(c *core.GContent) {
	p := map[string]string{}
	if e := c.BindJson(&p); e != nil {
		c.FailJson(403, e.Error())
		return
	}
	uid := c.GetUserID()
	tx := getDB(c)
	for k, v := range p {
		row := &model.UserInfo{}
		tx.First(row, "user_id=? and ukey=?", uid, k)
		if row.ID > 0 {
			tx.Where("user_id=? and ukey=?", uid, k).Updates(map[string]any{
				"newval":    v,
				"update_ip": c.GetIP(),
			})
		} else {
			tx.Create(&model.UserInfo{
				UserID:   uid,
				Ukey:     k,
				Newval:   v,
				CreateIP: c.GetIP(),
				UpdateIP: c.GetIP(),
			})
		}
	}
	c.SuccessJson("ok")
}

type sendCodeParam struct {
	Name string `json:"name"`
}

// 发送验证码
func sendCodeAction(c *core.GContent) {
	p := &sendCodeParam{}
	if e := c.BindJson(p); e != nil {
		c.FailJson(403, "参数错误")
		return
	}
	if utils.IsEmail(p.Name) {
		code := utils.RandStr(6)
		fmt.Printf("验证码：%s:%s\n", p.Name, code)
		utils.VerifySaveCode(getCahce(c), code, p.Name, 600)
		// if e := c.SendLocalMail("default", p.Name, "验证码", true, []byte("验证码： "+code)); e != nil {
		// 	c.FailJson(500, e.Error())
		// 	return
		// }
		c.SuccessJson("ok")
		return
	}
}
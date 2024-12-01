package account

import (
	"fmt"

	"github.com/ghf-go/fleetness/account/model"
	"github.com/ghf-go/fleetness/core"
	"github.com/ghf-go/fleetness/core/utils"
	"gorm.io/gorm"
)

type adminLoginActionParam struct {
	LoginName string `json:"login_name"`
	Passwd    string `json:"pass"`
	Code      string `json:"code"`
}

// 管理员登录
func adminLoginAction(c *core.GContent) {
	p := &adminLoginActionParam{}
	if e := c.BindJson(p); e != nil || p.Code == "" || p.LoginName == "" || p.Passwd == "" {
		c.FailJson(403, c.Lang("client_param_error"))
		return
	}
	admUser := &model.AdminUser{}
	getDB(c).First(admUser, "login_name=?", p.LoginName)
	if !utils.VerifyOtp2Fa(admUser.TfaKey, p.Code) {
		c.FailJson(403, "验证码错误")
		return
	}
	if admUser.ID == 0 || passwd(p.Passwd, admUser.PassSign) != admUser.Passwd {
		c.FailJson(403, "账号或者密码错误")
		return
	}
	c.SetUserID(fmt.Sprintf("%d", admUser.ID))
	c.SuccessJson("success")

}

type adminChangeAdminPassActionParam struct {
	Passwd string `json:"pass"`
}

// 修改密码
func adminChangeAdminPassAction(c *core.GContent) {
	p := &adminChangeAdminPassActionParam{}
	if e := c.BindJson(p); e != nil || p.Passwd == "" {
		c.FailJson(403, c.Lang("client_param_error"))
		return
	}
	sign := utils.RandStr(10)
	setData := map[string]any{
		"pass_sign": sign,
		"passwd":    passwd(p.Passwd, sign),
		"update_ip": c.GetIP(),
	}
	if getDB(c).Model(&model.AdminUser{}).Where(c.GetUserID()).Updates(setData).RowsAffected == 0 {
		c.FailJson(405, c.Lang("save_fail"))
		return
	}

	c.SuccessJson("success")
}

// 用户列表
func adminUserListAction(c *core.GContent) {
	p := &core.PageParam{}
	if e := c.BindJson(p); e != nil {
		c.FailJson(403, c.Lang("client_param_error"))
		return
	}
	total := int64(0)
	db := getDB(c)
	list := []model.User{}
	db.Model(&model.User{}).Where("is_audit=0").Count(&total)
	db.Where("is_audit=0").Order("update_at DESC").Offset(p.GetOffset()).Limit(p.GetPageSize()).Find(&list)

	c.SuccessJson(map[string]any{
		"total": total,
		"list":  list,
	})
}

type adminUserAddActionParm struct {
	UId    uint64 `json:"uid"`
	Passwd string `json:"pass"`
}

// 添加用户
func adminUserAddAction(c *core.GContent) {}

type adminUserChangePassActionParam struct {
	UId    uint64 `json:"uid"`
	Passwd string `json:"pass"`
}

// 修改用户密码
func adminUserChangePassAction(c *core.GContent) {
	p := &adminUserChangePassActionParam{}
	if e := c.BindJson(p); e != nil || p.Passwd == "" || p.UId == 0 {
		c.FailJson(403, c.Lang("client_param_error"))
		return
	}
	sign := utils.RandStr(10)
	setData := map[string]any{
		"pass_sign": sign,
		"passwd":    passwd(p.Passwd, sign),
		"update_ip": c.GetIP(),
	}
	if getDB(c).Model(&model.User{}).Where(p.UId).Updates(setData).RowsAffected == 0 {
		c.FailJson(405, c.Lang("save_fail"))
		return
	}

	c.SuccessJson("success")
}

// 用户信息统计
func adminUserStatAction(c *core.GContent) {}

type adminUserAuditActionParam struct {
	Uid uint64 `json:"uid"`
	Key string `json:"key"`
	Act string `json:"act"`
}

// 审核用户信息
func adminUserAuditAction(c *core.GContent) {
	p := &adminUserAuditActionParam{}
	if e := c.BindJson(p); e != nil || p.Uid == 0 || p.Key == "" || p.Act == "" {
		c.FailJson(403, c.Lang("client_param_error"))
		return
	}
	if p.Act == "accept" {
		getDB(c).Model(&model.UserInfo{}).Where("user_id=? AND ukey=?", p.Uid, p.Key).Updates(map[string]any{
			"is_audit":  1,
			"uval":      gorm.Expr("newval"),
			"newval":    "",
			"update_ip": c.GetIP(),
		})
	} else {
		getDB(c).Model(&model.UserInfo{}).Where("user_id=? AND ukey=?", p.Uid, p.Key).Updates(map[string]any{
			"is_audit":  1,
			"newval":    "",
			"update_ip": c.GetIP(),
		})
	}
	c.SuccessJson("success")
}

// 账号待审核信息
func adminUserWaitAuditAction(c *core.GContent) {
	p := &core.PageParam{}
	if e := c.BindJson(p); e != nil {
		c.FailJson(403, c.Lang("client_param_error"))
		return
	}
	total := int64(0)
	db := getDB(c)
	list := []model.UserInfo{}
	db.Model(&model.UserInfo{}).Where("is_audit=0").Count(&total)
	db.Where("is_audit=0").Order("update_at DESC").Offset(p.GetOffset()).Limit(p.GetPageSize()).Find(&list)

	c.SuccessJson(map[string]any{
		"total": total,
		"list":  list,
	})

}

package account

import (
	"fmt"
	"time"

	"github.com/ghf-go/fleetness/account/model"
	"github.com/ghf-go/fleetness/core"
	"github.com/ghf-go/fleetness/core/echarts"
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
		c.FailJson(403, "账号或者密码错误"+passwd(p.Passwd, admUser.PassSign))
		return
	}
	c.SetUserID(fmt.Sprintf("%d", admUser.ID))
	c.SuccessJson(map[string]any{
		"nick_name": admUser.NickName,
	})

}

type adminChangeAdminPassActionParam struct {
	Passwd    string `json:"newPass"`
	OldPasswd string `json:"oldPass"`
}

// 修改密码
func adminChangeAdminPassAction(c *core.GContent) {
	p := &adminChangeAdminPassActionParam{}
	if e := c.BindJson(p); e != nil || p.Passwd == "" {
		c.FailJson(403, c.Lang("client_param_error"))
		return
	}
	admUser := &model.AdminUser{}
	getDB(c).First(admUser, c.GetUserID())
	if admUser.ID == 0 || passwd(p.OldPasswd, admUser.PassSign) != admUser.Passwd {
		c.FailJson(403, "账号或者密码错误")
		return
	}
	sign := utils.RandStr(10)
	setData := map[string]any{
		"pass_sign": sign,
		"passwd":    passwd(p.Passwd, sign),
		"update_ip": c.GetIP(),
	}
	if getDB(c).Model(&model.AdminUser{}).Where(c.GetUserID()).Updates(setData).Error != nil {
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
	db.Model(&model.User{}).Count(&total)
	db.Order("update_at DESC").Offset(p.GetOffset()).Limit(p.GetPageSize()).Find(&list)

	c.SuccessJson(map[string]any{
		"total": total,
		"list":  list,
	})
}

type adminUserAddActionParam struct {
	Name   string `json:"name"`
	Passwd string `json:"pass"`
}

// 添加用户
func adminUserAddAction(c *core.GContent) {
	p := &adminUserAddActionParam{}
	if e := c.BindJson(p); e != nil || p.Name == "" || p.Passwd == "" {
		c.FailJson(403, c.Lang("client_param_error"))
		return
	}
	if e := createUser(c, p.Name, p.Passwd); e != nil {
		c.FailJson(403, c.Lang("save_fail"))
	}
	c.SuccessJson("success")
}

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

type adminUserStatActionParam struct {
	StartAt string `json:"start"`
}

// 用户信息统计
func adminUserStatAction(c *core.GContent) {
	p := &adminUserStatActionParam{}
	if e := c.BindJson(p); e != nil {
		c.FailJson(403, c.Lang("client_param_error"))
		return
	}
	if p.StartAt == "" {
		p.StartAt = time.Now().Add(time.Hour * time.Duration((24 * -39))).Format(utils.T_DATE)
	}
	m := &model.User{}
	type row struct {
		D time.Time
		N int64
	}
	rlist := []row{}
	data := map[string]map[string]any{}
	getDB(c).Raw(fmt.Sprintf("SELECT DATE(create_at) AS d,COUNT(1) AS n FROM %s WHERE create_at>=? GROUP BY d", m.TableName()), p.StartAt).Scan(&rlist)
	dayList := map[string]any{}
	for _, item := range rlist {
		dayList[item.D.Format(utils.T_DATE)] = item.N
	}
	data["注册用户数"] = dayList
	c.SuccessJson(echarts.BuildBaseLine("每日新增用户", echarts.FillDateLineData(data)))
}

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

package appbuild

import (
	"github.com/ghf-go/fleetness/appbuild/model"
	"github.com/ghf-go/fleetness/core"
)

// 模块列表
func adminModuleListAction(c *core.GContent) {
	p := &core.PageParam{}
	if e := c.BindJson(p); e != nil {
		c.FailJson(403, "参数错误")
		return
	}
	total := int64(0)
	list := []model.AppBuildModule{}
	getDB(c).Model(&model.AppBuildModule{}).Offset(p.GetOffset()).Limit(p.GetPageSize()).Find(&list)
	c.SuccessJson(map[string]any{
		"total": total,
		"list":  list,
	})
}

// 保存模块
func adminModuleSaveAction(c *core.GContent) {
	p := &model.AppBuildModule{}
	if e := c.BindJson(p); e != nil || p.Name == "" || p.ModuleDesc == "" {
		c.FailJson(403, "参数错误")
		return
	}
	p.UpdateIP = c.GetIP()
	if p.ID > 0 {
		getDB(c).Model(p).Where(p.ID).Updates(map[string]any{
			"name":        p.Name,
			"module_desc": p.ModuleDesc,
			"update_ip":   p.UpdateIP,
		})
	} else {
		p.CreateIP = c.GetIP()
		getDB(c).Save(p)
	}
	c.SuccessJson("OK")
}

type adminModuleItemsActionParam struct {
	ModelId uint64 `json:"module_id"`
}

// 模块文件列表
func adminModuleItemsAction(c *core.GContent) {
	p := &adminModuleItemsActionParam{}
	if e := c.BindJson(p); e != nil || p.ModelId == 0 {
		c.FailJson(403, "参数错误")
		return
	}
	list := []model.AppBuildTemplate{}
	getDB(c).Order("platform ASC").Find(&list)
	c.SuccessJson(map[string]any{
		"platforms": platforms,
		"modules":   list,
	})
}

// 保存文件
func adminModuleItemSaveAction(c *core.GContent) {
	p := &model.AppBuildTemplate{}
	if e := c.BindJson(p); e != nil || p.ModuleID == 0 || p.Platform == "" || p.Fname == "" || p.Fdesc == "" || p.Ftemplate == "" {
		c.FailJson(403, "参数错误")
		return
	}
	p.UpdateIP = c.GetIP()
	if p.ID > 0 {
		getDB(c).Model(p).Where(p.ID).Updates(map[string]any{
			"platform":  p.Platform,
			"module_id": p.ModuleID,
			"fname":     p.Fname,
			"fdesc":     p.Fdesc,
			"ftemplate": p.Ftemplate,
			"update_ip": p.UpdateIP,
		})
	} else {
		p.CreateIP = c.GetIP()
		getDB(c).Save(p)
	}
	c.SuccessJson("OK")
}

type adminModuleItemDetailActionParam struct {
	ItemID uint64 `json:"item_id"`
}

// 模版文件详情
func adminModuleItemDetailAction(c *core.GContent) {
	p := &adminModuleItemDetailActionParam{}
	if e := c.BindJson(p); e != nil || p.ItemID == 0 {
		c.FailJson(403, "参数错误")
		return
	}
	ret := &model.AppBuildTemplate{}
	getDB(c).First(ret, p.ItemID)
	c.SuccessJson(map[string]any{
		"detail":    ret,
		"platforms": platforms,
	})
}

// 获取项目配置
func adminProjectConfAction(c *core.GContent) {
	list := []model.AppBuildModule{}
	getDB(c).Find(&list)
	c.SuccessJson(map[string]any{
		"platforms": platforms,
		"modules":   list,
	})
}

// 生成项目工程模版
func adminProjectAction(c *core.GContent) {}

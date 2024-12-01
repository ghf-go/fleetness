package group

import (
	"github.com/ghf-go/fleetness/core"
	"github.com/ghf-go/fleetness/group/model"
)

type apiGroupListActionParam struct {
	TargetType uint `json:"target_type"`
}

// 分组列表
func apiGroupListAction(c *core.GContent) {
	p := &apiGroupListActionParam{}
	if e := c.BindJson(p); e != nil {
		c.FailJson(403, e.Error())
		return
	}
	c.SuccessJson(GroupList(c, c.GetUserID(), p.TargetType))
}

// 保存分组信息
func apiGroupSaveAction(c *core.GContent) {
	p := &model.Group{}
	if e := c.BindJson(p); e != nil {
		c.FailJson(403, e.Error())
		return
	}
	if e := GroupSave(c, p); e != nil {
		c.FailJson(403, e.Error())
		return
	}
	c.SuccessJson("success")
}

type apiGroupDelActionParam struct {
	ID uint64 `json:"id"`
}

// 删除分组
func apiGroupDelAction(c *core.GContent) {
	p := &apiGroupDelActionParam{}
	if e := c.BindJson(p); e != nil {
		c.FailJson(403, e.Error())
		return
	}
	if GroupDel(c, c.GetUserID(), p.ID) {
		c.SuccessJson("success")
		return
	}
	c.FailJson(405, c.Lang("save_fail"))
}

// 添加分组信息
func apiGroupAddItemAction(c *core.GContent) {
	p := &apiGroupDelItemActionParams{}
	if e := c.BindJson(p); e != nil {
		c.FailJson(403, e.Error())
		return
	}
	if e := GroupItemAdd(c, c.GetUserID(), p.Id, p.TargetIds...); e != nil {
		c.FailJson(403, e.Error())
		return
	}
	c.SuccessJson("success")
}

type apiGroupDelItemActionParams struct {
	Id        uint64   `json:"group_id"`
	TargetIds []uint64 `json:"ids"`
}

// 删除分组信息
func apiGroupDelItemAction(c *core.GContent) {
	p := &apiGroupDelItemActionParams{}
	if e := c.BindJson(p); e != nil {
		c.FailJson(403, e.Error())
		return
	}
	if e := GroupItemDel(c, c.GetUserID(), p.Id, p.TargetIds...); e != nil {
		c.FailJson(403, e.Error())
		return
	}
	c.SuccessJson("success")
}

type apiGroupMoveItemActionParam struct {
	// Id       uint64 `json:"group_id"`
	// Page     int    `json:"page"`
	// PageSize int    `json:"page_size"`
	NewID     uint64   `json:"new_group_id"`
	OldID     uint64   `json:"old_group_id"`
	TargetIds []uint64 `json:"ids"`
}

// 移动分组信息
func apiGroupMoveItemAction(c *core.GContent) {
	p := &apiGroupMoveItemActionParam{}
	if e := c.BindJson(p); e != nil {
		c.FailJson(403, e.Error())
		return
	}
	GroupItemMove(c, c.GetUserID(), p.OldID, p.NewID, p.TargetIds...)
	c.SuccessJson("success")
	// if p.Page < 1 {
	// 	p.Page = 1
	// }
	// if p.PageSize < 1 {
	// 	p.PageSize = 10
	// }
	// if p.Id < 1 {
	// 	c.FailJson(403, c.Lang("client_param_error"))
	// 	return
	// }
	// c.SuccessJson(GroupItems(c, c.GetUserID(), p.Id, p.Page, p.PageSize))
}

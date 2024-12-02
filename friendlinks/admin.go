package friendlinks

import (
	"github.com/ghf-go/fleetness/core"
	"github.com/ghf-go/fleetness/friendlinks/model"
)

// 友情链接列表
func adminFriendLinkListAction(c *core.GContent) {
	list := []model.FriendLinks{}
	getDB(c).Order("link_type ASC,sort_index ASC").Find(&list)
	c.SuccessJson(list)
}

// 保存友情链接
func adminFriendLinkSaveAction(c *core.GContent) {
	list := []model.FriendLinks{}
	if e := c.BindJson(&list); e != nil || len(list) == 0 {
		c.FailJson(403, c.Lang("client_param_error"))
		return
	}
	db := getDB(c)
	for _, item := range list {
		if item.ID > 0 {
			db.Model(&item).Where(item.ID).Updates(map[string]any{
				"update_ip":  c.GetIP(),
				"name":       item.Name,
				"link_type":  item.LinkType,
				"logo":       item.Logo,
				"url":        item.Url,
				"ios":        item.Ios,
				"google_pay": item.GooglePay,
				"bg_img":     item.BgImg,
				"content":    item.Content,
				"is_show":    item.IsShow,
				"sort_index": item.SortIndex,
			})
		} else {
			item.CreateIP = c.GetIP()
			item.UpdateIP = c.GetIP()
			db.Save(&item)
		}
	}
	c.SuccessJson("success")
}

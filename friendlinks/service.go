package friendlinks

import (
	"github.com/ghf-go/fleetness/core"
	"github.com/ghf-go/fleetness/friendlinks/model"
)

// 获取友情链接列表
func GetFriendLink(c *core.GContent, linktype string) []map[string]any {
	ret := []map[string]any{}
	if !isOnline {
		return ret
	}
	list := []model.FriendLinks{}
	getDB(c).Where("is_show=1 AND link_type=?", linktype).Order("sort_index ASC").Find(&list)
	for _, item := range list {
		ret = append(ret, map[string]any{
			"name":        item.Name,
			"logo":        item.Logo,
			"url":         item.Url,
			"ios":         item.Ios,
			"google_play": item.GooglePay,
			"bg_img":      item.BgImg,
			"desc":        item.Content,
		})
	}
	return ret
}

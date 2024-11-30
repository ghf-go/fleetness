package account

import "github.com/ghf-go/fleetness/account/model"

// 格式化收货地址信息
func formatUserAddr(item *model.UserAddr) map[string]any {
	return map[string]any{
		"id":         item.ID,
		"mobile":     item.Mobile,
		"consignee":  item.Consignee,
		"province":   item.Province,
		"city":       item.City,
		"district":   item.District,
		"address":    item.Address,
		"is_default": item.IsDefault,
	}
}

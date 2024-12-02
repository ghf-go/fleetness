package category

import (
	"github.com/ghf-go/fleetness/category/model"
	"github.com/ghf-go/fleetness/core"
)

// 获取分类列表
func GetCategoryList(c *core.GContent, catetype int) []map[string]any {
	if !isOnline {
		return []map[string]any{}
	}
	list := []model.Category{}
	getDB(c).Where("target_type=? AND is_show=1", catetype).Order("parent_id ASC ,sort_index ASC").Find(&list)
	return formatCategoryList(list)
}

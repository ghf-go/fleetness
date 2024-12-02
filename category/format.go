package category

import "github.com/ghf-go/fleetness/category/model"

func formatCategory(data model.Category) map[string]any {
	return map[string]any{
		"id":          data.ID,
		"name":        data.Name,
		"target_type": data.TargetType,
		"parent_id":   data.ParentID,
		"is_show":     data.IsShow,
		"sort_index":  data.SortIndex,
		"subs":        []map[string]any{},
	}
}
func formatCategoryList(data []model.Category) []map[string]any {
	dd := map[uint64]map[string]any{}
	for _, item := range data {
		if item.ParentID > 0 {
			if pl, ok := dd[item.ParentID]; ok {
				if sub, o := pl["subs"]; o {
					a := sub.([]map[string]any)
					a = append(a, formatCategory(item))
					pl["subs"] = a
					dd[item.ParentID] = pl
				}
			}
		} else {
			dd[item.ID] = formatCategory(item)
		}
	}
	ret := []map[string]any{}
	for _, i := range dd {
		ret = append(ret, i)
	}
	return ret
}

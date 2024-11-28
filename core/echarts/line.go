package echarts

import (
	"time"

	"github.com/ghf-go/fleetness/core/utils"
)

// 生产折线图数据
func BuildBaseLine(name string, data map[string]map[string]any) map[string]any {
	itemnames := []string{}
	axisNames := []string{}
	series := []map[string]any{}
	for k, item := range data {
		itemnames = append(itemnames, k)
		if len(axisNames) == 0 {
			for an, _ := range item {
				axisNames = append(axisNames, an)
			}
		}
		ld := []any{}
		for _, v := range item {
			ld = append(ld, v)
		}
		series = append(series, map[string]any{
			"name": k, "type": "line", "stack": "Total",
			"data": ld,
		})
	}
	return map[string]any{
		"title":   map[string]any{"text": name},
		"tooltip": map[string]any{"text": "axis"},
		"legend":  map[string]any{"data": itemnames},
		"grid":    map[string]any{"left": "3%", "right": "4%", "bottom": "3%", "containLabel": true},
		"toolbox": map[string]any{"feature": map[string]any{"saveAsImage": map[string]any{}}},
		"xAxis":   map[string]any{"type": "category", "boundaryGap": false, "data": axisNames},
		"yAxis":   map[string]any{"type": "value"},
		"series":  series,
	}
}

// 填充日期
func FillDateLineData(data map[string]map[string]any) map[string]map[string]any {
	min := time.Now().Format(utils.T_DATE)
	max := ""
	for _, items := range data {
		for d, _ := range items {
			if d < min {
				min = d
			}
			if d > max {
				max = d
			}
		}
	}
	dates := []string{}
	for min <= max {
		dates = append(dates, min)
		t, _ := time.Parse(utils.T_DATE, min)
		min = t.Add(time.Hour * 24).Format(utils.T_DATE)
	}
	ret := map[string]map[string]any{}
	for name, slist := range data {
		item := map[string]any{}
		for _, d := range dates {
			if row, ok := slist[d]; ok {
				item[d] = row
			} else {
				item[d] = 0
			}
		}
		ret[name] = item
	}
	return ret
}

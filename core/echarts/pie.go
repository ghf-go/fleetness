package echarts

// 生产饼图数据
func BuildBasePie(name string, data map[string]any) map[string]any {
	dd := []map[string]any{}
	for k, v := range data {
		dd = append(dd, map[string]any{"value": v, "name": k})
	}
	return map[string]any{
		"title":   map[string]any{"text": name, "left": "center"},
		"tooltip": map[string]any{"trigger": "item"},
		"legend":  map[string]any{"orient": "vertical", "left": "left"},
		"series": []map[string]any{
			{
				"name": name, "type": "pie", "radius": "80%",
				"emphasis": map[string]any{
					"itemStyle": map[string]any{
						"shadowBlur":    10,
						"shadowOffsetX": 0,
						"shadowColor":   "rgba(0, 0, 0, 0.5)",
					},
				},
				"data": dd,
			},
		},
	}
}

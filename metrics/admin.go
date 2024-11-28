package metrics

import (
	"fmt"
	"time"

	"github.com/ghf-go/fleetness/core"
	"github.com/ghf-go/fleetness/core/echarts"
	"github.com/ghf-go/fleetness/core/utils"
	"github.com/ghf-go/fleetness/metrics/model"
)

// 获取打点列表
func adminKeyListAction(c *core.GContent) {
	list := []model.MetricsConf{}
	getDB(c).Find(&list)
	c.SuccessJson(list)
}

type adminStatActionParam struct {
	StarDate string            `json:"start"`
	Keys     map[string]uint64 `json:"keys"`
}

// 获取统计数据
func adminStatAction(c *core.GContent) {
	p := &adminStatActionParam{}
	if e := c.BindJson(p); e != nil {
		c.FailJson(403, "参数错误")
		return
	}
	ids := []uint64{}
	km := map[uint64]string{}
	for k, id := range p.Keys {
		ids = append(ids, id)
		km[id] = k
	}
	type row struct {
		ConfId   uint64
		Platform string
		Views    uint
		Clicls   uint
		UViews   uint
		UClicks  uint
		Date     time.Time
	}
	rlist := []row{}
	ms := &model.MetricsStat{}
	getDB(c).Raw(fmt.Sprintf("SELECT conf_id,platform,views,clicks,user_views,user_clicks,`date` FROM %s WHERE conf_id IN ? AND `date`>= ?", ms.TableName()), ids, p.StarDate).Scan(&rlist)
	data := map[string]map[string]any{}

	for _, item := range rlist {
		day := item.Date.Format(utils.T_DATE)
		if name, ok := km[item.ConfId]; ok {
			{
				{
					nv := name + "-展示"
					if aaa, o := data[nv]; o {
						if dd, o2 := aaa[day]; o2 {
							aaa[day] = dd.(uint) + item.Views
						} else {
							aaa[day] = item.Views
						}
						data[nv] = aaa
					} else {
						data[nv] = map[string]any{day: item.Views}
					}
				}

				{
					nv := name + "-点击"
					if aaa, o := data[nv]; o {
						if dd, o2 := aaa[day]; o2 {
							aaa[day] = dd.(uint) + item.Clicls
						} else {
							aaa[day] = item.Clicls
						}
						data[nv] = aaa
					} else {
						data[nv] = map[string]any{day: item.Clicls}
					}
				}

				{
					nv := name + "-展示用户"
					if aaa, o := data[nv]; o {
						if dd, o2 := aaa[day]; o2 {
							aaa[day] = dd.(uint) + item.UViews
						} else {
							aaa[day] = item.UViews
						}
						data[nv] = aaa
					} else {
						data[nv] = map[string]any{day: item.UViews}
					}
				}

				{
					nv := name + "-点击用户"
					if aaa, o := data[nv]; o {
						if dd, o2 := aaa[day]; o2 {
							aaa[day] = dd.(uint) + item.UClicks
						} else {
							aaa[day] = item.UClicks
						}
						data[nv] = aaa
					} else {
						data[nv] = map[string]any{day: item.UClicks}
					}
				}
			}

			if item.Platform == "ios" {
				p := "ios"
				{
					nv := name + "-" + p + "-展示"
					if aaa, o := data[nv]; o {
						if dd, o2 := aaa[day]; o2 {
							aaa[day] = dd.(uint) + item.Views
						} else {
							aaa[day] = item.Views
						}
						data[nv] = aaa
					} else {
						data[nv] = map[string]any{day: item.Views}
					}
				}

				{
					nv := name + "-" + p + "-点击"
					if aaa, o := data[nv]; o {
						if dd, o2 := aaa[day]; o2 {
							aaa[day] = dd.(uint) + item.Clicls
						} else {
							aaa[day] = item.Clicls
						}
						data[nv] = aaa
					} else {
						data[nv] = map[string]any{day: item.Clicls}
					}
				}

				{
					nv := name + "-" + p + "-展示用户"
					if aaa, o := data[nv]; o {
						if dd, o2 := aaa[day]; o2 {
							aaa[day] = dd.(uint) + item.UViews
						} else {
							aaa[day] = item.UViews
						}
						data[nv] = aaa
					} else {
						data[nv] = map[string]any{day: item.UViews}
					}
				}

				{
					nv := name + "-" + p + "-点击用户"
					if aaa, o := data[nv]; o {
						if dd, o2 := aaa[day]; o2 {
							aaa[day] = dd.(uint) + item.UClicks
						} else {
							aaa[day] = item.UClicks
						}
						data[nv] = aaa
					} else {
						data[nv] = map[string]any{day: item.UClicks}
					}
				}
			}

			if item.Platform == "android" {
				p := "android"
				{
					nv := name + "-" + p + "-展示"
					if aaa, o := data[nv]; o {
						if dd, o2 := aaa[day]; o2 {
							aaa[day] = dd.(uint) + item.Views
						} else {
							aaa[day] = item.Views
						}
						data[nv] = aaa
					} else {
						data[nv] = map[string]any{day: item.Views}
					}
				}

				{
					nv := name + "-" + p + "-点击"
					if aaa, o := data[nv]; o {
						if dd, o2 := aaa[day]; o2 {
							aaa[day] = dd.(uint) + item.Clicls
						} else {
							aaa[day] = item.Clicls
						}
						data[nv] = aaa
					} else {
						data[nv] = map[string]any{day: item.Clicls}
					}
				}

				{
					nv := name + "-" + p + "-展示用户"
					if aaa, o := data[nv]; o {
						if dd, o2 := aaa[day]; o2 {
							aaa[day] = dd.(uint) + item.UViews
						} else {
							aaa[day] = item.UViews
						}
						data[nv] = aaa
					} else {
						data[nv] = map[string]any{day: item.UViews}
					}
				}

				{
					nv := name + "-" + p + "-点击用户"
					if aaa, o := data[nv]; o {
						if dd, o2 := aaa[day]; o2 {
							aaa[day] = dd.(uint) + item.UClicks
						} else {
							aaa[day] = item.UClicks
						}
						data[nv] = aaa
					} else {
						data[nv] = map[string]any{day: item.UClicks}
					}
				}
			}

		}
	}
	c.SuccessJson(echarts.BuildBaseLine("打点统计", echarts.FillDateLineData(data)))
}

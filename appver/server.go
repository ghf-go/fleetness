package appver

import (
	"github.com/ghf-go/fleetness/appver/model"
	"github.com/ghf-go/fleetness/core"
	"github.com/ghf-go/fleetness/core/utils"
)

// 检查最新版本
func GetLastVer(c *core.GContent, ver string) map[string]any {
	md := &model.AppVer{}
	getDB(c).Where("is_online=1").Order("app_ver DESC").First(md)
	if md.ID > 0 && utils.CheckVersion(md.AppVer, ver) {
		return map[string]any{
			"apk_ver": md.AppVer,
			"desc":    md.VerContent,
			"apk_url": md.ApkUrl,
			"wgt_url": md.WgtUrl,
		}
	}
	return map[string]any{}
}

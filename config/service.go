package config

import (
	"errors"

	"github.com/ghf-go/fleetness/config/model"
	"github.com/ghf-go/fleetness/core"
	"gorm.io/gorm"
)

// 获取一级配置信息
func GetConfigRoot(c *core.GContent, confkey string, def map[string]map[string]string) map[string]map[string]string {
	if !isOnline {
		return def
	}
	ret := map[string]map[string]string{}
	vlist := []model.Config{}
	getDB(c).Where("conf_key=?", confkey).Order("group_key ASC").Find(&vlist)
	lgk := ""
	gp := map[string]string{}
	for _, item := range vlist {
		if lgk != item.GroupKey {
			if lgk != "" {
				ret[lgk] = gp
				gp = map[string]string{}
			}
			lgk = item.GroupKey
		}
		gp[item.ItemKey] = item.Val
	}
	if lgk != "" {
		ret[lgk] = gp
	}

	for gk, gv := range def {
		if rg, ok := ret[gk]; ok {
			for k, v := range gv {
				if s, ok := rg[k]; ok {
					if s == "" {
						rg[k] = v
					}
				} else {
					rg[k] = v
				}
			}
			ret[gk] = rg
		} else {
			ret[gk] = gv
		}

	}
	return ret
}

// 获取二级配置
func GetConfigGroup(c *core.GContent, confkey, groupkey string, def map[string]string) map[string]string {
	if !isOnline {
		return def
	}
	vlist := []model.Config{}
	ret := map[string]string{}
	getDB(c).Find(&vlist, "conf_key=? AND group_key=?", confkey, groupkey)
	for _, item := range vlist {
		ret[item.ItemKey] = item.Val
	}
	for k, v := range def {
		if row, ok := ret[k]; ok {
			if row == "" {
				ret[k] = v
			}
		} else {
			ret[k] = v
		}
	}
	return ret
}

// 获取具体配置信息
func GetConfigItem(c *core.GContent, confkey, groupkey, itemkey string, def string) string {
	if !isOnline {
		return def
	}
	row := &model.Config{}
	getDB(c).First(row, "conf_key=? AND group_key=? AND item_key=?", confkey, groupkey, itemkey)
	if row.ID > 0 {
		if row.Val == "" {
			return def
		}
		return row.Val
	}
	return def
}

// 设置一级配置信息
func SetConfigRoot(c *core.GContent, confkey string, val map[string]map[string]string) bool {
	if !isOnline {
		return false
	}
	e, _ := c.Tx(getDB(c), func(tx *gorm.DB) (error, any) {
		for gk, vv := range val {
			for k, v := range vv {
				if tx.Save(&model.Config{
					ConfKey:  confkey,
					GroupKey: gk,
					ItemKey:  k,
					Val:      v,
					CreateIP: c.GetIP(),
					UpdateIP: c.GetIP(),
				}).Error != nil && tx.Model(&model.Config{}).Where("conf_key=? AND group_key=? AND item_key=?", confkey, gk, k).Updates(map[string]any{
					"val":       v,
					"update_ip": c.GetIP(),
				}).Error != nil {
					return errors.New(c.Lang("save_fail")), ""
				}
			}
		}
		return nil, ""
	})
	return e == nil
}

// 设置二级配置
func SetConfigGroup(c *core.GContent, confkey, groupkey string, val map[string]string) bool {
	if !isOnline {
		return false
	}
	e, _ := c.Tx(getDB(c), func(tx *gorm.DB) (error, any) {
		for k, v := range val {
			if tx.Save(&model.Config{
				ConfKey:  confkey,
				GroupKey: groupkey,
				ItemKey:  k,
				Val:      v,
				CreateIP: c.GetIP(),
				UpdateIP: c.GetIP(),
			}).Error != nil && tx.Model(&model.Config{}).Where("conf_key=? AND group_key=? AND item_key=?", confkey, groupkey, k).Updates(map[string]any{
				"val":       v,
				"update_ip": c.GetIP(),
			}).Error != nil {
				return errors.New(c.Lang("save_fail")), ""
			}
		}
		return nil, ""
	})
	return e == nil
}

// 设置具体配置信息
func SetConfigItem(c *core.GContent, confkey, groupkey, itemkey string, val string) bool {
	if !isOnline {
		return false
	}
	if getDB(c).Save(&model.Config{
		ConfKey:  confkey,
		GroupKey: groupkey,
		ItemKey:  itemkey,
		Val:      val,
		CreateIP: c.GetIP(),
		UpdateIP: c.GetIP(),
	}).Error == nil {
		return true
	}
	return getDB(c).Model(&model.Config{}).Where("conf_key=? AND group_key=? AND item_key=?", confkey, groupkey, itemkey).Updates(map[string]any{
		"val":       val,
		"update_ip": c.GetIP(),
	}).Error == nil
}

// Code generated by sql2gorm. DO NOT EDIT.
package model

import (
	"time"
)

// 生成配置表
type AppBuildModule struct {
	ID         uint64    `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"`
	Name       string    `gorm:"column:name;NOT NULL" json:"name"`                                     // 模块名称
	ModuleDesc string    `gorm:"column:module_desc;NOT NULL" json:"module_desc"`                       // 描述
	CreateAt   time.Time `gorm:"->;column:create_at;default:CURRENT_TIMESTAMP;NOT NULL" json:"create_at"` // 创建时间
	CreateIP   string    `gorm:"column:create_ip;NOT NULL" json:"create_ip"`                           // 创建IP
	UpdateAt   time.Time `gorm:"->;column:update_at;default:CURRENT_TIMESTAMP;NOT NULL" json:"update_at"` // 更新时间
	UpdateIP   string    `gorm:"column:update_ip;NOT NULL" json:"update_ip"`                           // 更新IP
}

func (m *AppBuildModule) TableName() string {
	return "t_app_build_module"
}

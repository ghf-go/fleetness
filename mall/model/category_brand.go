// Code generated by sql2gorm. DO NOT EDIT.
package model

import (
	"time"
)

// 商城分类品牌
type MallCategoryBrand struct {
	ID       uint64    `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"`
	CateID   uint64    `gorm:"column:cate_id;default:0;NOT NULL" json:"cate_id"`                     // 分类ID
	BrandID  uint64    `gorm:"column:brand_id;default:0;NOT NULL" json:"brand_id"`                   // 品牌ID
	CreateAt time.Time `gorm:"column:create_at;default:CURRENT_TIMESTAMP;NOT NULL" json:"create_at"` // 创建时间
	CreateIP string    `gorm:"column:create_ip;NOT NULL" json:"create_ip"`                           // 创建IP
	UpdateAt time.Time `gorm:"column:update_at;default:CURRENT_TIMESTAMP;NOT NULL" json:"update_at"` // 更新时间
	UpdateIP string    `gorm:"column:update_ip;NOT NULL" json:"update_ip"`                           // 更新IP
}

func (m *MallCategoryBrand) TableName() string {
	return "t_mall_category_brand"
}


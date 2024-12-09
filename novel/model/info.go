// Code generated by sql2gorm. DO NOT EDIT.
package model

import (
	"time"
)

// 小说列表
type NovelInfo struct {
	ID            uint64    `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"`
	Name          string    `gorm:"column:name;NOT NULL" json:"name"`                                     // 名称
	Logo          string    `gorm:"column:logo;NOT NULL" json:"logo"`                                     // LOGO
	UserID        uint64    `gorm:"column:user_id;default:0;NOT NULL" json:"user_id"`                     // 作者
	Subscribe     uint      `gorm:"column:subscribe;default:0;NOT NULL" json:"subscribe"`                 // 订阅人数
	Words         uint      `gorm:"column:words;default:0;NOT NULL" json:"words"`                         // 字数
	SectionPrice  uint      `gorm:"column:section_price;default:0;NOT NULL" json:"section_price"`         // 章节价格，单位分
	FreeSection   uint      `gorm:"column:free_section;default:0;NOT NULL" json:"free_section"`           // 免费章节
	LastSectionID uint64    `gorm:"column:last_section_id;default:0;NOT NULL" json:"last_section_id"`     // 最新章节
	IsOver        int       `gorm:"column:is_over;default:0;NOT NULL" json:"is_over"`                     // 是否完结
	IsFree        int       `gorm:"column:is_free;default:0;NOT NULL" json:"is_free"`                     // 是否免费
	IsPublish     int       `gorm:"column:is_publish;default:0;NOT NULL" json:"is_publish"`               // 是否上架
	IsAudit       int       `gorm:"column:is_audit;default:0;NOT NULL" json:"is_audit"`                   // 是否已经审核
	TotalIncome   uint      `gorm:"column:total_income;default:0;NOT NULL" json:"total_income"`           // 总收入
	TodayIncome   uint      `gorm:"column:today_income;default:0;NOT NULL" json:"today_income"`           // 今日收入
	Content       string    `gorm:"column:content;NOT NULL" json:"content"`                               // 描述
	CreateAt      time.Time `gorm:"->;column:create_at;default:CURRENT_TIMESTAMP;NOT NULL" json:"create_at"` // 创建时间
	CreateIP      string    `gorm:"column:create_ip;NOT NULL" json:"create_ip"`                           // 创建IP
	UpdateAt      time.Time `gorm:"->;column:update_at;default:CURRENT_TIMESTAMP;NOT NULL" json:"update_at"` // 更新时间
	UpdateIP      string    `gorm:"column:update_ip;NOT NULL" json:"update_ip"`                           // 更新IP
}

func (m *NovelInfo) TableName() string {
	return "t_novel_info"
}


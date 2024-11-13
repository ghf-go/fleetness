package model

import (
	"time"
)

// 收藏统计
type FavoriteStat struct {
	ID           uint64    `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"`
	TargetType   uint      `gorm:"column:target_type;default:0;NOT NULL" json:"target_type"`                // 收藏类型
	TargetID     uint64    `gorm:"column:target_id;default:0;NOT NULL" json:"target_id"`                    // 收藏ID
	TargetCounts uint64    `gorm:"column:target_counts;default:0;NOT NULL" json:"target_counts"`            // 收藏次数
	CreateAt     time.Time `gorm:"->;column:create_at;default:CURRENT_TIMESTAMP;NOT NULL" json:"create_at"` // 创建时间
	CreateIP     string    `gorm:"column:create_ip;NOT NULL" json:"create_ip"`                              // 创建IP
	UpdateAt     time.Time `gorm:"->;column:update_at;default:CURRENT_TIMESTAMP;NOT NULL" json:"update_at"` // 更新时间
	UpdateIP     string    `gorm:"column:update_ip;NOT NULL" json:"update_ip"`                              // 更新IP
}

func (m *FavoriteStat) TableName() string {
	return "t_favorite_stat"
}

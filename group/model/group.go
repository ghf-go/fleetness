package model

import (
	"time"
)

// 分组表
type Group struct {
	ID         uint64    `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"`
	TargetType uint      `gorm:"column:target_type;default:0;NOT NULL" json:"target_type"`                // 分组类型
	UserID     uint64    `gorm:"column:user_id;default:0;NOT NULL" json:"user_id"`                        // 用户id
	GroupName  string    `gorm:"column:group_name;NOT NULL" json:"group_name"`                            // 分组名称
	Items      int       `gorm:"column:items;default:0;NOT NULL" json:"items"`                            // 分组内条数
	CreateAt   time.Time `gorm:"->;column:create_at;default:CURRENT_TIMESTAMP;NOT NULL" json:"create_at"` // 创建时间
	CreateIP   string    `gorm:"column:create_ip;NOT NULL" json:"create_ip"`                              // 创建IP
	UpdateAt   time.Time `gorm:"->;column:update_at;default:CURRENT_TIMESTAMP;NOT NULL" json:"update_at"` // 更新时间
	UpdateIP   string    `gorm:"column:update_ip;NOT NULL" json:"update_ip"`                              // 更新IP
}

func (m *Group) TableName() string {
	return "t_group"
}

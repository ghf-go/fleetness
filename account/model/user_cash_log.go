// Code generated by sql2gorm. DO NOT EDIT.
package model

import (
	"time"
)

// 用户资金日志表
type UserCashLog struct {
	ID       uint64    `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"`
	UserID   uint64    `gorm:"column:user_id;default:0;NOT NULL" json:"user_id"`                     // 用户ID
	Ukey     string    `gorm:"column:ukey;NOT NULL" json:"ukey"`                                     // key
	Val      int       `gorm:"column:val;default:0;NOT NULL" json:"val"`                             // val
	Content  string    `gorm:"column:content;NOT NULL" json:"content"`                               // 描述
	CreateAt time.Time `gorm:"->;column:create_at;default:CURRENT_TIMESTAMP;NOT NULL" json:"create_at"` // 创建时间
	CreateIP string    `gorm:"column:create_ip;NOT NULL" json:"create_ip"`                           // 创建IP
	UpdateAt time.Time `gorm:"->;column:update_at;default:CURRENT_TIMESTAMP;NOT NULL" json:"update_at"` // 更新时间
	UpdateIP string    `gorm:"column:update_ip;NOT NULL" json:"update_ip"`                           // 更新IP
}

func (m *UserCashLog) TableName() string {
	return "t_user_cash_log"
}

package model

import (
	"time"
)

// 用户
type User struct {
	ID       uint64    `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"`
	NickName string    `gorm:"column:nick_name;NOT NULL" json:"nick_name"`                           // 昵称
	Avatar   string    `gorm:"column:avatar;NOT NULL" json:"avatar"`                                 // 头像
	Passwd   string    `gorm:"column:passwd;NOT NULL" json:"passwd"`                                 // 密码
	PassSign string    `gorm:"column:pass_sign;NOT NULL" json:"pass_sign"`                           // 密码加盐
	Status   int       `gorm:"column:status;default:0;NOT NULL" json:"status"`                       // 状态，0正常
	CreateAt time.Time `gorm:"column:create_at;default:CURRENT_TIMESTAMP;NOT NULL" json:"create_at"` // 创建时间
	CreateIP string    `gorm:"column:create_ip;NOT NULL" json:"create_ip"`                           // 创建IP
	UpdateAt time.Time `gorm:"column:update_at;default:CURRENT_TIMESTAMP;NOT NULL" json:"update_at"` // 更新时间
	UpdateIP string    `gorm:"column:update_ip;NOT NULL" json:"update_ip"`                           // 更新IP
}

func (m *User) TableName() string {
	return "t_user"
}

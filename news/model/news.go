// Code generated by sql2gorm. DO NOT EDIT.
package model

import (
	"time"
)

// 新闻表
type News struct {
	ID         uint64    `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"`
	Title      string    `gorm:"column:title;NOT NULL" json:"title"`                                   // 标题
	SubTitle   string    `gorm:"column:sub_title;NOT NULL" json:"sub_title"`                           // 附标题
	CategoryID uint64    `gorm:"column:category_id;default:0;NOT NULL" json:"category_id"`             // 分类
	Img        string    `gorm:"column:img;NOT NULL" json:"img"`                                       // 图片
	Content    string    `gorm:"column:content;NOT NULL" json:"content"`                               // 内容
	Author     string    `gorm:"column:author;NOT NULL" json:"author"`                                 // 作者
	Refer      string    `gorm:"column:refer;NOT NULL" json:"refer"`                                   // 应用地址
	IsPublish  int       `gorm:"column:is_publish;default:0;NOT NULL" json:"is_publish"`               // 是否发布
	IsDel      int       `gorm:"column:is_del;default:0;NOT NULL" json:"is_del"`                       // 是否删除
	CreateAt   time.Time `gorm:"->;column:create_at;default:CURRENT_TIMESTAMP;NOT NULL" json:"create_at"` // 创建时间
	CreateIP   string    `gorm:"column:create_ip;NOT NULL" json:"create_ip"`                           // 创建IP
	UpdateAt   time.Time `gorm:"->;column:update_at;default:CURRENT_TIMESTAMP;NOT NULL" json:"update_at"` // 更新时间
	UpdateIP   string    `gorm:"column:update_ip;NOT NULL" json:"update_ip"`                           // 更新IP
}

func (m *News) TableName() string {
	return "t_news"
}

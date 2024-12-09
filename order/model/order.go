// Code generated by sql2gorm. DO NOT EDIT.
package model

import (
	"time"
)

// 订单表
type Order struct {
	ID          uint64    `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"`
	OrderSn     string    `gorm:"column:order_sn;NOT NULL" json:"order_sn"`                             // 订单编号
	UserID      uint64    `gorm:"column:user_id;default:0;NOT NULL" json:"user_id"`                     // 用户ID
	OrderType   string    `gorm:"column:order_type;NOT NULL" json:"order_type"`                         // 订单类型，支付成功是需要回掉
	OrderSrcID  uint64    `gorm:"column:order_src_id;default:0;NOT NULL" json:"order_src_id"`           // 原订单ID
	OrderName   string    `gorm:"column:order_name;NOT NULL" json:"order_name"`                         // 订单名称
	TotalAmount uint      `gorm:"column:total_amount;default:0;NOT NULL" json:"total_amount"`           // 总金额
	PayAmount   uint      `gorm:"column:pay_amount;default:0;NOT NULL" json:"pay_amount"`               // 支付金额
	PayWay      string    `gorm:"column:pay_way;NOT NULL" json:"pay_way"`                               // 支付方式，ali_h5,ali_app,ali_pc,ali_ercode,wx_h5,wx_app,wx_mini...
	OrderStatus int       `gorm:"column:order_status;default:0;NOT NULL" json:"order_status"`           // 支付状态：0待支付，1已取消，10 已支付，20已退款
	PlatformSn  string    `gorm:"column:platform_sn;NOT NULL" json:"platform_sn"`                       // 三方支付订单号
	PayAt       *time.Time `gorm:"column:pay_at" json:"pay_at"`                                          // 支付成功时间
	RefundAt    *time.Time `gorm:"column:refund_at" json:"refund_at"`                                    // 退款成功时间
	CreateAt    time.Time `gorm:"->;column:create_at;default:CURRENT_TIMESTAMP;NOT NULL" json:"create_at"` // 创建时间
	CreateIP    string    `gorm:"column:create_ip;NOT NULL" json:"create_ip"`                           // 创建IP
	UpdateAt    time.Time `gorm:"->;column:update_at;default:CURRENT_TIMESTAMP;NOT NULL" json:"update_at"` // 更新时间
	UpdateIP    string    `gorm:"column:update_ip;NOT NULL" json:"update_ip"`                           // 更新IP
}

func (m *Order) TableName() string {
	return "t_order"
}


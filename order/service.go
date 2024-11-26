package order

import (
	"fmt"
	"time"

	"github.com/ghf-go/fleetness/core"
	"github.com/ghf-go/fleetness/core/payment"
	"github.com/ghf-go/fleetness/order/model"
)

// 下单
func CreateOrder(c *core.GContent, uid uint64, totalAmount, payAmount uint, orderType string, orderSrcId uint64, payWay, orderDesc string) any {
	om := &model.Order{
		OrderSn:     fmt.Sprintf("%s%d", orderType, time.Now().UnixNano()),
		UserID:      uid,
		OrderType:   orderType,
		OrderName:   orderDesc,
		OrderSrcID:  orderSrcId,
		OrderStatus: 0,
		TotalAmount: totalAmount,
		PayAmount:   payAmount,
		PayWay:      payWay,
	}
	if getDB(c).Save(om).Error != nil {
		return nil
	}
	ret := map[string]any{"payway": payWay}
	switch payWay {
	case PAYWAY_WC_APP:
		ret["data"] = payment.CreateOrderWeChatApp(c, int64(om.PayAmount), om.OrderSn, orderDesc, "")
	case PAYWAY_WC_H5:
		ret["data"] = payment.CreateOrderWeChatH5(c, int64(om.PayAmount), om.OrderSn, orderDesc, "")
	case PAYWAY_WC_JSAPI:
		ret["data"] = payment.CreateOrderWeChatJsapi(c, int64(om.PayAmount), om.OrderSn, orderDesc, "")
	case PAYWAY_WC_NATIVE:
		ret["data"] = payment.CreateOrderWeChatNative(c, int64(om.PayAmount), om.OrderSn, orderDesc, "")
	default:
		return nil
	}
	return ret
}

// 关闭订单
func CloseOrder(c *core.GContent, orderType string, orderSrcId uint64) bool {
	return getDB(c).Model(&model.Order{}).Where("order_type=? AND order_src_id=?", orderType, orderSrcId).Update("order_status", ORDER_STATUS_CLOSED).RowsAffected > 0
}

// 检查订单状态
func CheckOrderStatus(c *core.GContent, orderType string, orderSrcId uint64) int {
	row := &model.Order{}
	getDB(c).First(row, "order_type=? AND order_src_id=?", orderType, orderSrcId)
	if row.ID == 0 {
		return ORDER_STATUS_NOT_EXISTS
	}
	return row.OrderStatus
}

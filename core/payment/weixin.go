package payment

import (
	"context"
	"time"

	"github.com/ghf-go/fleetness/core"
	"github.com/ghf-go/fleetness/core/conf"
	"github.com/ghf-go/fleetness/core/log"
	wepay "github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/core/auth/verifiers"
	"github.com/wechatpay-apiv3/wechatpay-go/core/downloader"
	"github.com/wechatpay-apiv3/wechatpay-go/core/notify"
	"github.com/wechatpay-apiv3/wechatpay-go/core/option"
	"github.com/wechatpay-apiv3/wechatpay-go/services/partnerpayments/jsapi"
	"github.com/wechatpay-apiv3/wechatpay-go/services/partnerpayments/native"
	"github.com/wechatpay-apiv3/wechatpay-go/services/payments/app"
	"github.com/wechatpay-apiv3/wechatpay-go/services/payments/h5"
	"github.com/wechatpay-apiv3/wechatpay-go/services/refunddomestic"
	"github.com/wechatpay-apiv3/wechatpay-go/utils"
)

var (
	wepayClient       *wepay.Client
	wepayClientIsLoad = false
	wepayConf         *conf.WeChatPaymentConfig
	wepayNotify       *notify.Handler
)

// 支付回掉
func NotifyWeChatHandle(c *core.GContent, callback func(*WxNotifyBody) bool) core.Handle {
	return func(c *core.GContent) {
		ret := &WxNotifyBody{}
		_, e := wepayNotify.ParseNotifyRequest(c.GetContext(), c.GetRequest(), ret)
		if e != nil {
			return //失败
		}
		if callback(ret) {
			return
		} else {
			return
		}
	}
}

// 微信 申请退款
func RefundWeChat(c *core.GContent, refoudAmoung int64, orderID, msg, refoundOrderID string) string {
	svc := refunddomestic.RefundsApiService{Client: getWeChatClient(c)}
	resp, _, err := svc.Create(context.Background(),
		refunddomestic.CreateRequest{
			SubMchid:    wepay.String(wepayConf.MchID),
			OutTradeNo:  wepay.String(orderID),
			OutRefundNo: wepay.String(refoundOrderID),
			Reason:      wepay.String(msg),
			NotifyUrl:   wepay.String(wepayConf.RefundNotifyURL),
			Amount: &refunddomestic.AmountReq{
				Currency: wepay.String("CNY"),
				Refund:   wepay.Int64(refoudAmoung),
				Total:    wepay.Int64(refoudAmoung),
			},
		},
	)
	if err != nil {
		return ""
	}
	return *resp.RefundId
}

// 微信 h5下单
func CreateOrderWeChatH5(c *core.GContent, orderAmount int64, orderId, orderDesc, orderAttach string) string {
	svc := h5.H5ApiService{Client: getWeChatClient(c)}
	ctx := context.Background()
	resp, _, err := svc.Prepay(ctx,
		h5.PrepayRequest{
			Appid:         wepay.String(wepayConf.AppID),
			Mchid:         wepay.String(wepayConf.MchID),
			Description:   wepay.String(orderDesc),
			OutTradeNo:    wepay.String(orderId),
			TimeExpire:    wepay.Time(time.Now().Add(time.Minute * 30)),
			Attach:        wepay.String(orderAttach),
			NotifyUrl:     wepay.String(wepayConf.NotifyURL),
			SupportFapiao: wepay.Bool(false),
			Amount: &h5.Amount{
				Currency: wepay.String("CNY"),
				Total:    wepay.Int64(orderAmount),
			},
		},
	)
	if err != nil {
		return ""
	}
	return *resp.H5Url
}

// 微信 app下单
func CreateOrderWeChatApp(c *core.GContent, orderAmount int64, orderId, orderDesc, orderAttach string) string {
	svc := app.AppApiService{Client: getWeChatClient(c)}
	ctx := context.Background()
	resp, _, err := svc.Prepay(ctx,
		app.PrepayRequest{
			Appid:         wepay.String(wepayConf.AppID),
			Mchid:         wepay.String(wepayConf.MchID),
			Description:   wepay.String(orderDesc),
			OutTradeNo:    wepay.String(orderId),
			TimeExpire:    wepay.Time(time.Now().Add(time.Minute * 30)),
			Attach:        wepay.String(orderAttach),
			NotifyUrl:     wepay.String(wepayConf.NotifyURL),
			SupportFapiao: wepay.Bool(false),
			Amount: &app.Amount{
				Currency: wepay.String("CNY"),
				Total:    wepay.Int64(orderAmount),
			},
		},
	)
	if err != nil {
		return ""
	}
	return *resp.PrepayId
}

// 微信 JSAPI下单
func CreateOrderWeChatJsapi(c *core.GContent, orderAmount int64, orderId, orderDesc, orderAttach string) string {
	svc := jsapi.JsapiApiService{Client: getWeChatClient(c)}
	ctx := context.Background()
	resp, _, err := svc.Prepay(ctx,
		jsapi.PrepayRequest{
			SpAppid:       wepay.String(wepayConf.AppID),
			SpMchid:       wepay.String(wepayConf.MchID),
			Description:   wepay.String(orderDesc),
			OutTradeNo:    wepay.String(orderId),
			TimeExpire:    wepay.Time(time.Now().Add(time.Minute * 30)),
			Attach:        wepay.String(orderAttach),
			NotifyUrl:     wepay.String(wepayConf.NotifyURL),
			SupportFapiao: wepay.Bool(false),
			Amount: &jsapi.Amount{
				Currency: wepay.String("CNY"),
				Total:    wepay.Int64(orderAmount),
			},
		},
	)
	if err != nil {
		return ""
	}
	return *resp.PrepayId
}

// 微信 app下单
func CreateOrderWeChatNative(c *core.GContent, orderAmount int64, orderId, orderDesc, orderAttach string) string {
	svc := native.NativeApiService{Client: getWeChatClient(c)}
	ctx := context.Background()
	resp, _, err := svc.Prepay(ctx,
		native.PrepayRequest{
			SpAppid:       wepay.String(wepayConf.AppID),
			SpMchid:       wepay.String(wepayConf.MchID),
			Description:   wepay.String(orderDesc),
			OutTradeNo:    wepay.String(orderId),
			TimeExpire:    wepay.Time(time.Now().Add(time.Minute * 30)),
			Attach:        wepay.String(orderAttach),
			NotifyUrl:     wepay.String(wepayConf.NotifyURL),
			SupportFapiao: wepay.Bool(false),
			Amount: &native.Amount{
				Currency: wepay.String("CNY"),
				Total:    wepay.Int64(orderAmount),
			},
		},
	)
	if err != nil {
		return ""
	}
	return *resp.CodeUrl
}

// 关闭 h5订单
func CloseOrderWeChatH5(c *core.GContent, orderId string) {
	svc := h5.H5ApiService{Client: getWeChatClient(c)}
	ctx := context.Background()
	svc.CloseOrder(ctx,
		h5.CloseOrderRequest{
			Mchid:      wepay.String(wepayConf.MchID),
			OutTradeNo: wepay.String(orderId),
		},
	)
}

// 关闭 app订单
func CloseOrderWeChatApp(c *core.GContent, orderId string) {
	svc := app.AppApiService{Client: getWeChatClient(c)}
	ctx := context.Background()
	svc.CloseOrder(ctx,
		app.CloseOrderRequest{
			Mchid:      wepay.String(wepayConf.MchID),
			OutTradeNo: wepay.String(orderId),
		},
	)
}

// 关闭 Jsapi订单
func CloseOrderWeChatJsapi(c *core.GContent, orderId string) {
	svc := jsapi.JsapiApiService{Client: getWeChatClient(c)}
	ctx := context.Background()
	svc.CloseOrder(ctx,
		jsapi.CloseOrderRequest{
			SpMchid:    wepay.String(wepayConf.MchID),
			OutTradeNo: wepay.String(orderId),
		},
	)
}

// 关闭 Jsapi订单
func CloseOrderWeChatNative(c *core.GContent, orderId string) {
	svc := native.NativeApiService{Client: getWeChatClient(c)}
	ctx := context.Background()
	svc.CloseOrder(ctx,
		native.CloseOrderRequest{
			SpMchid:    wepay.String(wepayConf.MchID),
			OutTradeNo: wepay.String(orderId),
		},
	)
}

// 获取微信支付客户端
func getWeChatClient(c *core.GContent) *wepay.Client {
	if wepayClientIsLoad {
		return wepayClient
	}
	wepayConf = c.GetConf().Payment.WexPay
	mchPrivateKey, err := utils.LoadPrivateKey(wepayConf.PriviteKeyPem)
	if err != nil {
		log.Error(c, "微信支付秘钥错误%s", err.Error())
		return nil
	}
	ret, e := wepay.NewClient(context.Background(), option.WithWechatPayAutoAuthCipher(wepayConf.MchID, wepayConf.MchIdNumber, mchPrivateKey, wepayConf.MchApiV3Key))
	wepayClientIsLoad = true
	if e != nil {
		log.Error(c, "微信支付链接错误%s", e.Error())
		return nil
	}
	downloader.MgrInstance().RegisterDownloaderWithClient(context.Background(), ret, wepayConf.MchID, wepayConf.MchApiV3Key)
	certificateVisitor := downloader.MgrInstance().GetCertificateVisitor(wepayConf.MchID)
	// 3. 使用证书访问器初始化 `notify.Handler`
	wepayNotify = notify.NewNotifyHandler(wepayConf.MchApiV3Key, verifiers.NewSHA256WithRSAVerifier(certificateVisitor))
	wepayClient = ret
	return ret
}

// 微信回掉通知结构
type WxNotifyBody struct {
	Mchid               string             `json:"mchid"`
	TransactionID       string             `json:"transaction_id"`
	OutTradeNo          string             `json:"out_trade_no"`
	RefundID            string             `json:"refund_id"`
	OutRefundNo         string             `json:"out_refund_no"`
	RefundStatus        string             `json:"refund_status"`
	TradeState          string             `json:"trade_state"`
	SuccessTime         time.Time          `json:"success_time"`
	UserReceivedAccount string             `json:"user_received_account"`
	Amount              WxNotifyBodyAmount `json:"amount"`
	Payer               WxNotifyBodyPayer  `json:"payer"`
	AppID               string             `json:"AppID"`
	TradeStateDesc      string             `json:"trade_state_desc"`
	TradeType           string             `json:"trade_type"`
	Attach              string             `json:"attach"`
}
type WxNotifyBodyAmount struct {
	Total       int64 `json:"total"`
	Refund      int64 `json:"refund"`
	PayerTotal  int64 `json:"payer_total"`
	PayerRefund int64 `json:"payer_refund"`
}
type WxNotifyBodyPayer struct {
	Openid string `json:"openid"`
}

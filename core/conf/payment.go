package conf

type PaymentConfig struct {
	WexPay *WeChatPaymentConfig `yaml:"wechat"`
	AliPay *AliPaymentConfig    `yaml:"alipay"`
}

// 微信支付
type WeChatPaymentConfig struct {
	AppID           string `yaml:"app_id"`
	MchID           string `yaml:"mch_id"`
	MchIdNumber     string `yaml:"mch_id_num"`
	MchApiV3Key     string `yaml:"mch_api_v3_key"`
	PriviteKeyPem   string `yaml:"private_key_pem"`
	NotifyURL       string `yaml:"notify_url"`
	RefundNotifyURL string `yaml:"refund_notify_url"`
}

// 支付宝支付
type AliPaymentConfig struct {
	AppID         string `yaml:"app_id"`
	NotifyURL     string `yaml:"notify_url"`
	AppPublicPem  string `yaml:"app_public_pem"`
	AliPublicPem  string `yaml:"ali_public_pem"`
	RootPem       string `yaml:"root_pem"`
	AliGateWay    string `yaml:"gateway"`
	AliPrivateKey string `yaml:"alipay_private_key"`
}

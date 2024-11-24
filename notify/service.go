package notify

import (
	"github.com/sideshow/apns2"
)

var (
	iosclient *apns2.Client
)

// 短信通知
func Sms(mobile string, msg string) {}

// 邮件通知
func Email(title, msg string, emails ...string) {}

// google通知
func GoogleFam() {}

// ios通知
func Anps(notify *apns2.Notification) {
	if iosclient == nil {
		return
	}
	iosclient.Production().Push(notify)
}

// ios 测试环境推送
func AnpsDebug(notify *apns2.Notification) {
	if iosclient == nil {
		return
	}
	iosclient.Development().Push(notify)
}

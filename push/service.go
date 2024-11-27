package push

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

// SSe 全量推送
func PushAllSse(msg, event string) {
	if !isOnline {
		return
	}
	go func() {
		for _, item := range allSse {
			item.Send(msg, event)
		}
	}()
}

// 向用户发送 sse推送
func PushAllSseUser(msg, event string, uids ...uint64) {
	if !isOnline {
		return
	}
	go func() {
		for _, uid := range uids {
			if r, ok := userSse[uid]; ok {
				for _, s := range r {
					s.Send(msg, event)
				}
			}
		}
	}()
}

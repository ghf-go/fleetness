package notify

// 短信通知
func Sms(mobile string, msg string) {}

// 邮件通知
func Email(title, msg string, emails ...string) {}

// google通知
func GoogleFam() {}

// ios通知
func Ios() {}

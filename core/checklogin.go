package core

// 验证是否登录的中间件
func ApiCheckoutLoginMiddleWare(c *GContent) {
	if c.IsLogin() {
		c.Next()
	} else {
		c.FailJson(303, "账号没有登录")
	}
}

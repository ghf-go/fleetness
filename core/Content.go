package core

type GContent struct {
}

// 新建GContent
func newGContent() {

}

// 获取客户端IP
func (c *GContent) GetIP() string {
	return ""
}

func (c *GContent) Abort() {}
func (c *GContent) Next()  {}

// 获取数据库
func (c *GContent) GetDB() {}

// 获取缓存配置
func (c *GContent) GetCache() {}

// 绑定数据
func (c *GContent) BindXml() {}

// 绑定数据
func (c *GContent) BindJson() {}

// 发送队列
func (c *GContent) SendMq() {}

// 接口正常返回
func (c *GContent) SuccessJson() {}

// 接口保存信息
func (c *GContent) FailJson() {}

// 显示模版
func (c *GContent) Display() {}

// 显示模版
func (c *GContent) DisplayLayout() {}

// 发送邮件
func (c *GContent) SendMail() {}

package mall

import "github.com/ghf-go/fleetness/core"

type apiCateListActionParam struct {
	ID uint64 `json:"id"`
}

// 商品详情
func apiGoodsDetailAction(c *core.GContent) {
	p := &apiCateListActionParam{}
	if e := c.BindJson(p); e != nil {
		c.FailJson(403, e.Error())
		return
	}
	if r := GetGoodsDetail(c, p.ID); r != nil {
		c.SuccessJson(r)
		return
	}
	c.FailJson(403, "商品不存在或者已经下架")

}

// 分类列表
func apiCateListAction(c *core.GContent) {
	c.SuccessJson("")
}

// 商品列表
func apiGoodsListAction(c *core.GContent) {
	p := &GetGoodsListQuery{}
	if e := c.BindJson(p); e != nil {
		c.FailJson(403, e.Error())
		return
	}
	c.SuccessJson(GetGoodsList(c, p))
}

package mall

import (
	"github.com/ghf-go/fleetness/core"
	"github.com/ghf-go/fleetness/mall/model"
)

type GetGoodsListQuery struct {
	CateId1  uint64 `json:"cate_id1"`
	CateId2  uint64 `json:"cate_id2"`
	CateId3  uint64 `json:"cate_id3"`
	Name     string `json:"name"`
	SortName string `json:"sort_name"`
	SortVal  string `json:"sort_val"`
	Page     int    `json:"page"`
	PageSize int    `json:"page_size"`
}

// 查询商品列表
func GetGoodsList(c *core.GContent, query *GetGoodsListQuery) []map[string]any {
	return []map[string]any{}
}

// 获取商品详情
func GetGoodsDetail(c *core.GContent, goodsID uint64) map[string]any {
	db := getDB(c)
	goodsDetail := &model.MallGoods{}
	db.First(goodsDetail, goodsID)
	if goodsDetail.ID <= 0 {
		return nil
	}
	return map[string]any{}
}

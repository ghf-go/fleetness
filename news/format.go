package news

import (
	"github.com/ghf-go/fleetness/comment"
	"github.com/ghf-go/fleetness/core"
	"github.com/ghf-go/fleetness/news/model"
	"github.com/ghf-go/fleetness/praise"
)

func formatNews(c *core.GContent, uid uint64, data []model.News) []map[string]any {
	ret := []map[string]any{}
	for _, item := range data {
		ret = append(ret, map[string]any{
			"id":          item.ID,
			"title":       item.Title,
			"sub_title":   item.SubTitle,
			"category_id": item.CategoryID,
			"img":         item.Img,
			"content":     item.Content,
			"author":      item.Author,
			"refer":       item.Refer,
			"create_at":   item.CreateAt,
		})
	}
	ret = comment.AppendCommentInfo(c, core.TARGET_TYPE_NEWS, ret, "id", "comment")
	ret = praise.AppendPraise(c, core.TARGET_TYPE_NEWS, ret, "id", "parise")
	return ret
}

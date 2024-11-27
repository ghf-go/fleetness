package feed

import (
	"strings"

	"github.com/ghf-go/fleetness/account"
	"github.com/ghf-go/fleetness/comment"
	"github.com/ghf-go/fleetness/core"
	"github.com/ghf-go/fleetness/favorites"
	"github.com/ghf-go/fleetness/feed/model"
	"github.com/ghf-go/fleetness/praise"
)

func formatFeedList(c *core.GContent, uid uint64, list []model.Feed) []map[string]any {
	ret := []map[string]any{}

	fids := []uint64{}
	for _, item := range list {
		if item.FeedType == FEED_TYPE_MVOTE || item.FeedType == FEED_TYPE_VOTE {
			fids = append(fids, item.ID)
		}
	}
	vmap := map[uint64]map[string]any{}
	umap := map[uint64]bool{}
	if len(fids) > 0 {
		vlist := []model.FeedVote{}
		vlogs := []model.FeedVoteLog{}
		getDB(c).Find(&vlist, "feed_id IN ?", fids)
		getDB(c).Find(&vlogs, "feed_id IN ? AND user_id=?", fids, uid)
		for _, item := range vlist {
			if row, ok := vmap[item.FeedID]; ok {
				row["total"] = row["total"].(uint) + item.Votes
				row["items"].(map[string]any)[item.Name] = item.Votes
			} else {
				vmap[item.FeedID] = map[string]any{
					"total": item.Votes,
					"items": map[string]any{item.Name: item.Votes},
				}
			}
		}
		for _, item := range vlogs {
			umap[item.FeedID] = true
		}
	}

	for _, item := range list {
		dd := map[string]any{
			"id":        item.ID,
			"title":     item.Title,
			"content":   item.Content,
			"imgs":      strings.Split(item.Imgs, ","),
			"type":      item.FeedType,
			"user_id":   item.UserID,
			"create_at": item.CreateAt,
			"voted":     false,
			"votes":     nil,
		}
		if r, ok := vmap[item.ID]; ok {
			dd["votes"] = r
		}
		if _, ok := umap[item.ID]; ok {
			dd["voted"] = true
		}
		ret = append(ret, dd)
	}
	ret = account.AppendUserBase(c, ret, "user_id", "user_info")
	ret = praise.AppendPraise(c, core.TARGET_TYPE_FEED, ret, "id", "praise")
	ret = comment.AppendCommentInfo(c, core.TARGET_TYPE_FEED, ret, "id", "comments")
	ret = favorites.AppendFavoriteInfo(c, core.TARGET_TYPE_FEED, ret, "id", "favorite")
	return ret
}
func formatFeed(c *core.GContent, uid uint64, data model.Feed) map[string]any {
	ret := formatFeedList(c, uid, []model.Feed{data})
	return ret[0]
}

package blackword

import (
	"strings"

	"github.com/ghf-go/fleetness/blackword/model"
	"github.com/ghf-go/fleetness/core"
)

// 检查是否包敏感词
func CheckBlackWord(c *core.GContent, data string) (bool, []string) {
	isOk := false
	bwbws := []string{}
	if !isOnline {
		return isOk, bwbws
	}
	execInitSql(c)
	list := []model.BlackWord{}
	getDB(c).Find(&list)

	for _, item := range list {
		if strings.Index(data, item.Word) >= 0 {
			isOk = true
			bwbws = append(bwbws, item.Word)
		}
	}
	return isOk, bwbws

}

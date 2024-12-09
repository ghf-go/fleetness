package core

import (
	"fmt"

	"github.com/ghf-go/fleetness/core/utils"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var (
	isAppDebug = true
	dbCon      = map[string]*gorm.DB{}
	cacheCon   = map[string]*redis.Client{}
)

const (
	//审核状态
	STATUS_WAIT_AUDIT = 0
	STATUS_SUCCESS    = 10
	STATUS_MY         = 20
	STATUS_DEL        = 100

	//目标类型
	TARGET_TYPE_USER    = 1
	TARGET_TYPE_FEED    = 2
	TARGET_TYPE_COMMENT = 3
	TARGET_TYPE_NEWS    = 4
	TARGET_TYPE_NOVEL   = 5
	TARGET_TYPE_MALL    = 6
)

// 记录系统日志
func AppDebug(flayout string, arg ...any) {
	if isAppDebug {
		fmt.Printf("[%s] %s\n", utils.FormatDateTime(), fmt.Sprintf(flayout, arg...))
	}

}

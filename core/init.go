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

// 记录系统日志
func AppDebug(flayout string, arg ...any) {
	if isAppDebug {
		fmt.Printf("[%s] %s\n", utils.FormatDateTime(), fmt.Sprintf(flayout, arg...))
	}

}

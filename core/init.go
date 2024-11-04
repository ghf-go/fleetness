package core

import (
	"fmt"

	"github.com/ghf-go/fleetness/core/utils"
	"gorm.io/gorm"
)

var (
	isAppDebug = true
	dbCon      = map[string]*gorm.DB{}
)

// 记录系统日志
func AppDebug(flayout string, arg ...any) {
	if isAppDebug {
		fmt.Printf("[%s] %s\n", utils.FormatDateTime(), fmt.Sprintf(flayout, arg...))
	}

}

package log

import (
	"fmt"
	"runtime"

	"github.com/ghf-go/fleetness/core/utils"
)

type glog interface {
	write(logLevel, format string, arg ...any)
}

// 调试信息
func Debug(flayout string, arg ...any) {
	fmt.Printf("[%s] DEBUG %s\n", utils.FormatDateTime(), fmt.Sprintf(flayout, arg...))
}

// 错误信息
func Error(flayout string, arg ...any) {
	_, f, l, ok := runtime.Caller(1)
	if ok {
		fmt.Printf("[%s] ERROR %s:%d  %s\n", utils.FormatDateTime(), f, l, fmt.Sprintf(flayout, arg...))
	} else {
		fmt.Printf("[%s] ERROR %s\n", utils.FormatDateTime(), fmt.Sprintf(flayout, arg...))
	}
}

// 系统打印信息
func sysDebug(flayout string, arg ...any) {
	fmt.Printf("[%s] DEBUG_SYS %s\n", utils.FormatDateTime(), fmt.Sprintf(flayout, arg...))
}

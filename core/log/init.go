package log

import (
	"context"
	"fmt"
	"runtime"

	"github.com/ghf-go/fleetness/core"
)

var _log glog

type glog interface {
	Write(c *core.GContent, logLevel, format string, arg ...any)
	Close(c context.Context)
}

func getLogIns(c *core.GContent) glog {
	if _log == nil {
		logConf := c.GetConf().Log
		if logConf == nil {
			_log = &logConsole{}
			return _log
		}
		switch logConf.Driver {
		case "file":
			_log = &logFile{
				dirPath: logConf.DirPath,
			}
			return _log
		default:
			_log = &logConsole{}
			return _log
		}

	}
	return _log
}

// 调试信息
func Debug(c *core.GContent, flayout string, arg ...any) {
	getLogIns(c).Write(c, "DEBUG", flayout, arg...)
}

// 调试信息
func Sql(c *core.GContent, flayout string, arg ...any) {
	getLogIns(c).Write(c, "DEBUG", flayout, arg...)
}

// 错误信息
func Error(c *core.GContent, flayout string, arg ...any) {
	_, f, l, ok := runtime.Caller(1)
	if ok {
		getLogIns(c).Write(c, "ERROR", "%s:%d %s", f, l, fmt.Sprintf(flayout, arg...))
	} else {
		getLogIns(c).Write(c, "ERROR", flayout, arg...)
	}
}

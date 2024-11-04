package log

import (
	"context"

	"github.com/ghf-go/fleetness/core"
)

type logFile struct {
	dirPath string
}

func (l *logFile) Write(c *core.GContent, logLevel, format string, arg ...any) {

}

func (l *logFile) Close(c context.Context) {}

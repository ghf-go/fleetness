package log

import (
	"context"
	"fmt"
	"strings"

	"github.com/ghf-go/fleetness/core"
	"github.com/ghf-go/fleetness/core/utils"
)

type logConsole struct {
}

func (l *logConsole) Write(c *core.GContent, logLevel, format string, arg ...any) {
	fmt.Printf("[%s] %s %s %s %s\n", utils.FormatDateTime(), c.ReqID, c.GetIP(), strings.ToUpper(logLevel), fmt.Sprintf(format, arg...))
}
func (l *logConsole) Close(c context.Context) {}

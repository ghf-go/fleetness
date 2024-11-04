package log

import (
	"fmt"
	"strings"

	"github.com/ghf-go/fleetness/core/utils"
)

type LogConsole struct {
}

func (l *LogConsole) write(logLevel, format string, arg ...any) {
	fmt.Printf("[%s] %s %s\n", utils.FormatDateTime(), strings.ToUpper(logLevel), fmt.Sprintf(format, arg...))
}

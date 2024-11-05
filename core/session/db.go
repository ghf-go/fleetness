package session

import (
	"github.com/ghf-go/fleetness/core"
)

func SessionDB(data any) core.Handle {
	return func(c *core.GContent) {
		c.Next()
	}

}

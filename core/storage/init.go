package storage

import (
	"strings"

	"github.com/ghf-go/fleetness/core"
)

var (
	driver Storage = nil
)

type Storage interface {
	BuildToken(fileKey, fileName string) map[string]any
}

func GetStorage(c *core.GContent) Storage {
	if driver == nil {
		con := c.GetConf().Storage
		switch strings.ToLower(con.Driver) {
		case "qiniu":
			driver = newQiniu(con)
		}
	}
	return driver
}

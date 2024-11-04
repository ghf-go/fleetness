package main

import (
	_ "embed"

	"github.com/ghf-go/fleetness/core"
)

//go:embed test.yaml
var _confData []byte

func main() {
	gEngine := core.NewGengine(_confData)
	gEngine.Run()
}

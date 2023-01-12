package autocode

import (
	"log"

	"github.com/Gophercraft/core/vsn"
)

func runMain(build vsn.Build, loc string) {
	g := NewGenerator(build, loc)
	if err := g.Generate(); err != nil {
		log.Fatal(err)
	}
}

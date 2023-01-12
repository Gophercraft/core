package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Gophercraft/core/format/mpq"
	"github.com/Gophercraft/core/vsn"
	"github.com/superp00t/etc/yo"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("gcraft_list_mpq /path/to/my/game/directory")
		return
	}

	var fp = os.Args[1]

	build, err := vsn.DetectGame(fp)
	if err != nil {
		log.Fatal(err)
	}

	yo.Ok(build, "detected")

	s, err := mpq.GetFiles(fp)
	if err != nil {
		log.Fatal(err)
	}

	m, err := mpq.OpenPool(s)
	if err != nil {
		log.Fatal(err)
	}

	for _, v := range m.ListFiles() {
		fmt.Println(v)
	}
}

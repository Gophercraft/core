package main

import (
	"fmt"
	"io"
	"os"

	"github.com/Gophercraft/core/datapack"
	"github.com/Gophercraft/text"
)

func openFile(out string) *os.File {
	fmt.Println("Extracting to", out, "...")
	fl, _ := os.OpenFile(out, os.O_TRUNC|os.O_CREATE|os.O_RDWR, 0700)
	return fl
}

func printTimestamp(out io.Writer) {
	fmt.Fprintf(out, "// DO NOT EDIT: extracted from CMaNGOS database on %s\n", datapack.Timestamp())
}

func openTextWriter(out io.Writer) *text.Encoder {
	j := text.NewEncoder(out)
	return j
}

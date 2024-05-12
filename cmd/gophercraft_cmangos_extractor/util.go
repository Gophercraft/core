package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Gophercraft/core/datapack"
	"github.com/Gophercraft/text"
)

type textFile struct {
	file *os.File
	*text.Encoder
}

func openTextFile(path string) *textFile {
	log.Println("Opening", path)
	tf := new(textFile)
	var err error
	tf.file, err = os.OpenFile(path, os.O_TRUNC|os.O_CREATE|os.O_RDWR, 0700)
	if err != nil {
		panic(err)
	}
	tf.commentf("DO NOT EDIT: extracted from CMaNGOS database on %s", datapack.Timestamp())
	tf.Encoder = text.NewEncoder(tf.file)
	return tf
}

func (tf *textFile) comment(msg string) {
	fmt.Fprintf(tf.file, "// %s\n", msg)
}

func (tf *textFile) commentf(msg string, args ...any) {
	tf.comment(fmt.Sprintf(msg, args...))
}

func (tf *textFile) close() {
	tf.file.Close()
}

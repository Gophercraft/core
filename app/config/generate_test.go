package config

import (
	"os"
	"testing"
)

func TestGenerate(t *testing.T) {
	pth := "I:\\gcauth\\"

	os.MkdirAll(pth, 0777)

	if err := GenerateKeyPair(pth); err != nil {
		panic(err)
	}
}

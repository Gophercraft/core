package config

import (
	"os"
	"testing"
)

func TestGenerate(t *testing.T) {
	pth := "/tmp/gcauth"

	os.RemoveAll(pth)

	if err := GenerateDefaultHome(pth); err != nil {
		t.Fatal(err)
	}
}

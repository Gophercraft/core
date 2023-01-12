package mpq

import (
	"testing"

	"github.com/Gophercraft/log"
)

func TestMPQ(t *testing.T) {
	for _, testExtract := range []struct {
		Volume string
		Files  []string
	}{
		{
			"common-2.MPQ",
			[]string{
				"World\\Maps\\Kalimdor\\Kalimdor_38_39.adt",
			},
		},
	} {
		vol, err := Open(testExtract.Volume)
		if err != nil {
			t.Fatal(err)
		}

		for _, filename := range testExtract.Files {
			file, err := vol.OpenFile(filename)
			if err != nil {
				t.Fatal(err)
			}

			data, err := file.ReadBlock()
			if err != nil {
				t.Fatal(err)
			}

			log.Dump("filename", filename)
			log.Dump("data", data)

			file.Close()
		}

		vol.Close()
	}
}

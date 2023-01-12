package vsn

import (
	"fmt"
	"testing"

	"github.com/Gophercraft/log"
)

type Feature struct {
	Data int
}

var Descriptor1 = map[BuildRange]Feature{
	{12340, 20000}: {500},
	{5875, 8606}:   {69},
}

var Descriptor2 = map[BuildRange]Feature{
	{5875, 8606}:  {69},
	{8605, 10000}: {50},
}

var Descriptor3 = map[BuildRange]map[string]string{
	{5875, 8606}: {
		"Jesse": "Pinkman",
	},
}

func TestQueryDescriptor(t *testing.T) {
	var f *Feature
	err := QueryDescriptors(6000, Descriptor1, &f)
	if err != nil {
		t.Fatal(err)
	}
	if f.Data != 69 {
		t.Fatal(f.Data)
	}
	err = QueryDescriptors(6000, Descriptor2, &f)
	if err == nil {
		t.Fatal("should have failed here")
	} else {
		fmt.Println("This error is supposed to happen: ", err)
	}
	var m map[string]string
	err = QueryDescriptors(6000, Descriptor3, &m)
	if err != nil {
		t.Fatal(err)
	}
	log.Dump("m", m)
}

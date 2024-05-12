package grunt_test

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/Gophercraft/core/grunt"
	"github.com/Gophercraft/core/version"
	"github.com/davecgh/go-spew/spew"
)

type realm_list_test_packet struct {
	build version.Build
	data  string
}

var test_packets = []realm_list_test_packet{
	{
		5875,
		"3d0000000000010100000000434d614e474f5320436c617373696320736572766572003139322e3136382e302e3235323a38303835000ad7a33c0201000200",
	},
}

func TestRealmList(t *testing.T) {
	for _, packet := range test_packets {
		packet_bytes, err := hex.DecodeString(packet.data)
		if err != nil {
			t.Fatal(err)
		}
		packet_reader := bytes.NewReader(packet_bytes)

		fmt.Println(spew.Sdump(packet_bytes))

		var realmlist grunt.RealmList_Server
		if err = grunt.ReadRealmList_Server(packet_reader, packet.build, &realmlist); err != nil {
			t.Fatal(err)
		}

		fmt.Println(spew.Sdump(&realmlist))
	}
}

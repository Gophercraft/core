package ngdp

import (
	"testing"

	"github.com/Gophercraft/log"
)

func TestNGDP(t *testing.T) {
	ag := DefaultAgent()
	ag.HostServer = "http://us.patch.battle.net:1119"

	online, err := ag.OpenOnline("wow")
	if err != nil {
		t.Fatal(err)
	}
	log.Dump("online.BuildConfig", online.BuildConfig)
}

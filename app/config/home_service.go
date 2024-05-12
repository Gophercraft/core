package config

import (
	"fmt"
)

type HomeServiceID uint8

const (
	NotAHomeService HomeServiceID = iota
	CoreService
	WebService
	GruntService
	BNetRPCService
	BNetRESTService
	OldRealmlistService
	MaxHomeService
)

var home_service_text = []string{
	"",
	"core",
	"web",
	"grunt",
	"bnet_rpc",
	"bnet_rest",
	"old_realmlist",
}

func (home_service_ID *HomeServiceID) EncodeWord() (text string, err error) {
	value := *home_service_ID
	if value > MaxHomeService {
		return "", fmt.Errorf("config: cannot encode invalid home service ID %d", value)
	}
	return home_service_text[int(value)], nil
}

func (home_service_ID *HomeServiceID) DecodeWord(value string) (err error) {
	for i, text := range home_service_text {
		if text == value {
			*home_service_ID = HomeServiceID(i)
			return
		}
	}
	err = fmt.Errorf("config: %s is not a home service that exists", value)
	return
}

func (home_service_ID HomeServiceID) String() (s string) {
	s, _ = home_service_ID.EncodeWord()
	return
}

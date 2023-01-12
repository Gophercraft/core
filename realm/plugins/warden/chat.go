package warden

import (
	"fmt"

	"github.com/Gophercraft/core/packet/chat"
)

const plurple = "A635B0FF"

type WarningStage uint8

const (
	WarnOk WarningStage = iota
	WarnInformative
	WarnCritical
	WarnSevere
	WarnFatal
)

var Colors = map[WarningStage]string{
	WarnOk:          "3AEA6CFF",
	WarnInformative: "3A98EAFF",
	WarnCritical:    "F59E42FF",
	WarnSevere:      "CF0000FF",
	WarnFatal:       "330000FF",
}

func (sd *SessionData) Wardenf(ws WarningStage, data string, args ...interface{}) {
	color := Colors[ws]
	sd.Session.Send(&chat.ServerMessage{
		Type: chat.MsgSystem,
		Body: fmt.Sprintf("|c%s[Warden]: |r|c%s|%s|r", plurple, color, fmt.Sprintf(data, args...)),
	})
}

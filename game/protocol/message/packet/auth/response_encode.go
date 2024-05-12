package auth

import (
	"github.com/Gophercraft/core/game/protocol/message"
	"github.com/Gophercraft/core/version"
)

func (response *Response) Encode(build version.Build, to *message.Packet) error {
	to.Type = message.SMSG_AUTH_RESPONSE

	// Use old packet structure
	if build.RemovedIn(NewResponse) {
		return response.encode_0_12340(build, to)
	}

	return response.encode_13164_(build, to)
}

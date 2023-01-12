package packet

import (
	"fmt"

	"github.com/Gophercraft/core/vsn"
)

type ReasonForDisconnect uint32

const (
	InvalidServerHeader = 3
	RSAVerifyFailed     = 4
	ServerCheckFailed   = 24
)

func (dc ReasonForDisconnect) Error() string {
	reason := ""
	switch dc {
	case InvalidServerHeader:
		reason = "because of a protocol error"
	case RSAVerifyFailed:
		reason = "due to an RSA signature failure"
	case ServerCheckFailed:
		reason = "because an HMAC server check failed"
	default:
		reason = "for an unknown reason"
	}

	return fmt.Sprintf("client disconnected %s (error code %d)", reason, dc)
}

type ClientDisconnect struct {
	Reason ReasonForDisconnect
}

func (cd *ClientDisconnect) Decode(build vsn.Build, in *WorldPacket) error {
	cd.Reason = ReasonForDisconnect(in.ReadUint32())
	return nil
}

func (cd *ClientDisconnect) Encode(build vsn.Build, out *WorldPacket) error {
	out.WriteUint32(uint32(cd.Reason))
	return nil
}

func (dc *ClientDisconnect) String() string {
	return dc.Reason.Error()
}

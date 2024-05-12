package auth

import (
	"fmt"

	"github.com/Gophercraft/core/game/protocol/message"
	"github.com/Gophercraft/core/version"
)

type Challenge struct {
	Challenge    []byte // "salt"
	DosChallenge []byte // seeds
	DosZeroBits  uint8
}

func (ac *Challenge) Decode(build version.Build, in *message.Packet) error {
	if in.Type != message.SMSG_AUTH_CHALLENGE {
		return fmt.Errorf("auth: type is not SMSG_AUTH_CHALLENGE")
	}

	switch {
	case version.Range(0, 3368).Contains(build):
		in.ReadBytes(6)
	case version.Range(3494, 3494).Contains(build):
		in.ReadUint8()
	case version.Range(3592, 6005).Contains(build):
		ac.Challenge = in.ReadBytes(4)
		ac.DosChallenge = in.ReadBytes(32)
	case version.Range(6180, 12340).Contains(build):
		ac.DosZeroBits = uint8(in.ReadUint32())
		ac.Challenge = in.ReadBytes(4)
		ac.DosChallenge = in.ReadBytes(32)
	case version.Range(13164, 15595).Contains(build):
		ac.DosChallenge = in.ReadBytes(32)
		ac.Challenge = in.ReadBytes(4)
		ac.DosZeroBits = in.ReadUint8()
	case version.Range(15851, 18414).Contains(build):
		_ = in.ReadUint16()
		ac.DosChallenge = in.ReadBytes(32)
		ac.DosZeroBits = in.ReadUint8()
		ac.Challenge = in.ReadBytes(4)
	case version.Range(19027, version.Max).Contains(build):
		ac.DosChallenge = in.ReadBytes(32)
		ac.Challenge = in.ReadBytes(16)
		ac.DosZeroBits = in.ReadUint8()
	default:
		return fmt.Errorf("auth: unhandled build parsing Challenge: %s", build)
	}

	return nil
}

func (ac *Challenge) Encode(build version.Build, out *message.Packet) error {
	out.Type = message.SMSG_AUTH_CHALLENGE

	var (
		badChallenge4     = fmt.Errorf("auth: invalid Challenge length in Challenge, should be 4")
		badChallenge16    = fmt.Errorf("auth: invalid Challenge length in Challenge, should be 16")
		badDosChallenge32 = fmt.Errorf("auth: bad DosChallenge length in Challenge, should be 32")
	)

	switch {
	case version.Range(0, 3368).Contains(build):
		out.Write(make([]byte, 6))
	case version.Range(3494, 3494).Contains(build):
		out.WriteUint8(0xC)
	case version.Range(3592, 6005).Contains(build):
		if len(ac.Challenge) != 4 {
			return badChallenge4
		}
		if len(ac.DosChallenge) != 32 {
			return badDosChallenge32
		}
		out.Write(ac.Challenge)
		out.Write(ac.DosChallenge)
	case version.Range(6180, 12340).Contains(build):
		if len(ac.Challenge) != 4 {
			return badChallenge4
		}
		if len(ac.DosChallenge) != 32 {
			return badDosChallenge32
		}
		out.WriteUint32(uint32(ac.DosZeroBits))
		out.Write(ac.Challenge[:4])
		out.Write(ac.DosChallenge[:32])
	case version.Range(13164, 15595).Contains(build):
		if len(ac.Challenge) != 4 {
			return badChallenge4
		}
		if len(ac.DosChallenge) != 32 {
			return badDosChallenge32
		}
		out.Write(ac.DosChallenge[:32])
		out.Write(ac.Challenge[:4])
		out.WriteUint8(ac.DosZeroBits)
	case version.Range(15851, 18414).Contains(build):
		out.WriteUint16(0)
		if len(ac.Challenge) != 4 {
			return badChallenge4
		}
		if len(ac.DosChallenge) != 32 {
			return badDosChallenge32
		}
		out.Write(ac.DosChallenge[:32])
		out.WriteUint8(ac.DosZeroBits)
		out.Write(ac.Challenge[:4])
	case version.Range(19027, version.Max).Contains(build):
		if len(ac.Challenge) != 16 {
			return badChallenge16
		}
		if len(ac.DosChallenge) != 32 {
			return badDosChallenge32
		}
		out.Write(ac.DosChallenge[:32])
		out.Write(ac.Challenge[:16])
		out.WriteUint8(ac.DosZeroBits)
	}
	return nil
}

// type AuthResponse struct {
// 	Cmd       uint8
// 	WaitQueue uint32
// }

// func UnmarshalAuthResponse(input []byte) (*AuthResponse, error) {
// 	p := etc.FromBytes(input)
// 	s := &AuthResponse{}
// 	cmd := p.ReadUint8()

// 	if cmd != AUTH_WAIT_QUEUE {
// 		return s, nil
// 	}

// 	s.WaitQueue = p.ReadUint32()
// 	p.ReadUint8()
// 	return s, nil
// }

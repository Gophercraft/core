package auth

import (
	"fmt"

	"github.com/Gophercraft/core/packet"
	"github.com/Gophercraft/core/vsn"
)

type Challenge struct {
	Challenge    []byte // "salt"
	DosChallenge []byte // seeds
	DosZeroBits  uint8
}

func (ac *Challenge) Decode(build vsn.Build, in *packet.WorldPacket) error {
	if in.Type != packet.SMSG_AUTH_CHALLENGE {
		return fmt.Errorf("auth: type is not SMSG_AUTH_CHALLENGE")
	}

	switch {
	case vsn.Range(0, 3368).Contains(build):
		in.ReadBytes(6)
	case vsn.Range(3494, 3494).Contains(build):
		in.ReadByte()
	case vsn.Range(3592, 6005).Contains(build):
		ac.Challenge = in.ReadBytes(4)
		ac.DosChallenge = in.ReadBytes(32)
	case vsn.Range(6180, 12340).Contains(build):
		ac.DosZeroBits = uint8(in.ReadUint32())
		ac.Challenge = in.ReadBytes(4)
		ac.DosChallenge = in.ReadBytes(32)
	case vsn.Range(13164, 15595).Contains(build):
		ac.DosChallenge = in.ReadBytes(32)
		ac.Challenge = in.ReadBytes(4)
		ac.DosZeroBits = in.ReadByte()
	case vsn.Range(15851, 18414).Contains(build):
		_ = in.ReadUint16()
		ac.DosChallenge = in.ReadBytes(32)
		ac.DosZeroBits = in.ReadByte()
		ac.Challenge = in.ReadBytes(4)
	case vsn.Range(19027, vsn.Max).Contains(build):
		ac.DosChallenge = in.ReadBytes(32)
		ac.Challenge = in.ReadBytes(16)
		ac.DosZeroBits = in.ReadByte()
	default:
		return fmt.Errorf("auth: unhandled build parsing Challenge: %s", build)
	}

	return nil
}

func (ac *Challenge) Encode(build vsn.Build, out *packet.WorldPacket) error {
	out.Type = packet.SMSG_AUTH_CHALLENGE

	var (
		badChallenge4     = fmt.Errorf("auth: invalid Challenge length in Challenge, should be 4")
		badChallenge16    = fmt.Errorf("auth: invalid Challenge length in Challenge, should be 16")
		badDosChallenge32 = fmt.Errorf("auth: bad DosChallenge length in Challenge, should be 32")
	)

	switch {
	case vsn.Range(0, 3368).Contains(build):
		out.Write(make([]byte, 6))
	case vsn.Range(3494, 3494).Contains(build):
		out.WriteByte(0xC)
	case vsn.Range(3592, 6005).Contains(build):
		if len(ac.Challenge) != 4 {
			return badChallenge4
		}
		if len(ac.DosChallenge) != 32 {
			return badDosChallenge32
		}
		out.Write(ac.Challenge)
		out.Write(ac.DosChallenge)
	case vsn.Range(6180, 12340).Contains(build):
		if len(ac.Challenge) != 4 {
			return badChallenge4
		}
		if len(ac.DosChallenge) != 32 {
			return badDosChallenge32
		}
		out.WriteUint32(uint32(ac.DosZeroBits))
		out.Write(ac.Challenge[:4])
		out.Write(ac.DosChallenge[:32])
	case vsn.Range(13164, 15595).Contains(build):
		if len(ac.Challenge) != 4 {
			return badChallenge4
		}
		if len(ac.DosChallenge) != 32 {
			return badDosChallenge32
		}
		out.Write(ac.DosChallenge[:32])
		out.Write(ac.Challenge[:4])
		out.WriteByte(ac.DosZeroBits)
	case vsn.Range(15851, 18414).Contains(build):
		out.WriteUint16(0)
		if len(ac.Challenge) != 4 {
			return badChallenge4
		}
		if len(ac.DosChallenge) != 32 {
			return badDosChallenge32
		}
		out.Write(ac.DosChallenge[:32])
		out.WriteByte(ac.DosZeroBits)
		out.Write(ac.Challenge[:4])
	case vsn.Range(19027, vsn.Max).Contains(build):
		if len(ac.Challenge) != 16 {
			return badChallenge16
		}
		if len(ac.DosChallenge) != 32 {
			return badDosChallenge32
		}
		out.Write(ac.DosChallenge[:32])
		out.Write(ac.Challenge[:16])
		out.WriteByte(ac.DosZeroBits)
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
// 	cmd := p.ReadByte()

// 	if cmd != AUTH_WAIT_QUEUE {
// 		return s, nil
// 	}

// 	s.WaitQueue = p.ReadUint32()
// 	p.ReadByte()
// 	return s, nil
// }

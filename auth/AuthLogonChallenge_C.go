package auth

import (
	"encoding/binary"
	"strings"

	"github.com/superp00t/etc"
	"github.com/Gophercraft/core/vsn"
)

// AuthLogonChallenge_C is the first packet sent by a client
// while initiating a connection to an authserver.
type AuthLogonChallenge_C struct {
	Error        ErrorType
	GameName     string // Encode in reverse.
	Version      [3]byte
	Build        uint16
	Platform     string
	OS           string
	Country      string
	TimezoneBias uint32
	IP           uint32
	I            string
}

func (ac *AuthLogonChallenge_C) Type() AuthType {
	return LogonChallenge
}

func (ac *AuthLogonChallenge_C) Recv(build vsn.Build, conn *Connection) error {
	header, err := conn.RecvBuffer(3)
	if err != nil {
		return err
	}

	ac.Error = ErrorType(header.ReadByte())
	var size uint16
	binary.Read(header, binary.LittleEndian, &size)

	in, err := conn.RecvBuffer(int(size))
	if err != nil {
		return err
	}

	ac.GameName = in.ReadInvertedString(4)
	copy(ac.Version[:], in.ReadBytes(3))
	ac.Build = in.ReadUint16()
	ac.Platform = in.ReadInvertedString(4)
	ac.OS = in.ReadInvertedString(4)
	ac.Country = in.ReadInvertedString(4)
	ac.TimezoneBias = in.ReadUint32()
	ac.IP = in.ReadUint32()
	ac.I = string(in.ReadBytes(int(in.ReadByte())))
	return nil
}

func (alcc *AuthLogonChallenge_C) Send(build vsn.Build, out *Connection) error {
	a := etc.NewBuffer()
	a.WriteByte(uint8(alcc.Error))

	b := etc.NewBuffer()
	b.WriteInvertedString(4, alcc.GameName)
	b.Write(alcc.Version[:])
	b.WriteUint16(alcc.Build)
	b.WriteInvertedString(4, alcc.Platform)
	b.WriteInvertedString(4, alcc.OS)
	b.WriteInvertedString(4, alcc.Country)
	b.WriteUint32(alcc.TimezoneBias)
	b.WriteUint32(alcc.IP)
	b.WriteByte(uint8(len(alcc.I)))
	b.Write([]byte(alcc.I))

	a.WriteUint16(uint16(b.Len()))
	a.Write(b.Bytes())

	return out.SendBuffer(a)
}

// LogonChallengePacket_C is a helper function to simplify the client library.
func LogonChallengePacket_C(build vsn.Build, username string) Packet {
	alcc := &AuthLogonChallenge_C{
		Error:        8,
		GameName:     "WoW",
		Version:      Version(build),
		Build:        uint16(build),
		Platform:     "x86",
		OS:           "Win",
		Country:      "enGB",
		TimezoneBias: 0,
		IP:           16777343, //localhost
		I:            strings.ToUpper(username),
	}

	return alcc
}

func Version(build vsn.Build) [3]byte {
	info := build.BuildInfo()
	if build.BuildInfo() == nil {
		return [3]byte{0, 0, 0}
	}

	return [3]byte{
		byte(info.MajorVersion),
		byte(info.MinorVersion),
		byte(info.BugfixVersion),
	}
}

package auth

import (
	"fmt"

	"github.com/Gophercraft/core/game/protocol/message"
	"github.com/Gophercraft/core/version"
)

type EnterEncryptedMode struct {
	Signature []byte
	Enabled   bool
}

func (eem *EnterEncryptedMode) Encode(build version.Build, out *message.Packet) error {
	out.Type = message.SMSG_ENTER_ENCRYPTED_MODE
	if len(eem.Signature) != 256 {
		return fmt.Errorf("packet: invalid Signature length, should be 256")
	}
	out.Write(eem.Signature[:256])
	out.WriteBit(eem.Enabled)
	out.FlushBits()
	return nil
}

func (eem *EnterEncryptedMode) Decode(build version.Build, in *message.Packet) error {
	if in.Len() < 257 {
		return fmt.Errorf("packet: EnterEncryptedMode packet too small to be parsed")
	}
	eem.Signature = in.ReadBytes(256)
	eem.Enabled = in.ReadBit()
	in.FlushBits()
	return nil
}

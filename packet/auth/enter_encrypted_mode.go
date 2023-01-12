package auth

import (
	"fmt"

	"github.com/Gophercraft/core/packet"
	"github.com/Gophercraft/core/vsn"
)

type EnterEncryptedMode struct {
	Signature []byte
	Enabled   bool
}

func (eem *EnterEncryptedMode) Encode(build vsn.Build, out *packet.WorldPacket) error {
	out.Type = packet.SMSG_ENTER_ENCRYPTED_MODE
	if len(eem.Signature) != 256 {
		return fmt.Errorf("packet: invalid Signature length, should be 256")
	}
	out.Write(eem.Signature[:256])
	out.WriteMSBit(eem.Enabled)
	out.FlushBits()
	return nil
}

func (eem *EnterEncryptedMode) Decode(build vsn.Build, in *packet.WorldPacket) error {
	if in.Len() < 257 {
		return fmt.Errorf("packet: EnterEncryptedMode packet too small to be parsed")
	}
	eem.Signature = in.ReadBytes(256)
	eem.Enabled = in.ReadMSBit()
	in.FlushBits()
	return nil
}

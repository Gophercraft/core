package party

import (
	"github.com/Gophercraft/core/game/protocol/message"
	"github.com/Gophercraft/core/version"
)

type InviteRequest struct {
	To string
}

func (ir *InviteRequest) Encode(build version.Build, out *message.Packet) error {
	out.Type = message.CMSG_GROUP_INVITE
	out.WriteCString(ir.To)
	return nil
}

func (ir *InviteRequest) Decode(build version.Build, in *message.Packet) error {
	ir.To = in.ReadCString()
	return nil
}

type Invitation struct {
	From string
}

func (iv *Invitation) Encode(build version.Build, out *message.Packet) error {
	out.Type = message.SMSG_GROUP_INVITE
	out.WriteCString(iv.From)
	return nil
}

func (iv *Invitation) Decode(build version.Build, in *message.Packet) error {
	iv.From = in.ReadCString()
	return nil
}

type Declination struct {
	From string
}

func (d *Declination) Encode(build version.Build, out *message.Packet) error {
	out.Type = message.SMSG_GROUP_DECLINE
	out.WriteCString(d.From)
	return nil
}

func (d *Declination) Decode(build version.Build, in *message.Packet) error {
	d.From = in.ReadCString()
	return nil
}

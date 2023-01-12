package party

import (
	"github.com/Gophercraft/core/packet"
	"github.com/Gophercraft/core/vsn"
)

type InviteRequest struct {
	To string
}

func (ir *InviteRequest) Encode(build vsn.Build, out *packet.WorldPacket) error {
	out.Type = packet.CMSG_GROUP_INVITE
	out.WriteCString(ir.To)
	return nil
}

func (ir *InviteRequest) Decode(build vsn.Build, in *packet.WorldPacket) error {
	ir.To = in.ReadCString()
	return nil
}

type Invitation struct {
	From string
}

func (iv *Invitation) Encode(build vsn.Build, out *packet.WorldPacket) error {
	out.Type = packet.SMSG_GROUP_INVITE
	out.WriteCString(iv.From)
	return nil
}

func (iv *Invitation) Decode(build vsn.Build, in *packet.WorldPacket) error {
	iv.From = in.ReadCString()
	return nil
}

type Declination struct {
	From string
}

func (d *Declination) Encode(build vsn.Build, out *packet.WorldPacket) error {
	out.Type = packet.SMSG_GROUP_DECLINE
	out.WriteCString(d.From)
	return nil
}

func (d *Declination) Decode(build vsn.Build, in *packet.WorldPacket) error {
	d.From = in.ReadCString()
	return nil
}

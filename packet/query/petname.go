package query

import (
	"github.com/Gophercraft/core/guid"
	"github.com/Gophercraft/core/packet"
	"github.com/Gophercraft/core/vsn"
)

type PetName struct {
	PetNumber uint32
	PetGUID   guid.GUID
}

func (p *PetName) Encode(build vsn.Build, out *packet.WorldPacket) (err error) {
	out.WriteUint32(p.PetNumber)
	return p.PetGUID.EncodeUnpacked(build, out)
}

func (p *PetName) Decode(build vsn.Build, in *packet.WorldPacket) (err error) {
	p.PetNumber = in.ReadUint32()
	p.PetGUID, err = guid.DecodeUnpacked(build, in)
	return
}

type PetNameResponse struct {
	PetNumber        uint32
	PetName          string
	PetNameTimestamp uint32
}

func (p *PetNameResponse) Encode(build vsn.Build, out *packet.WorldPacket) (err error) {
	out.WriteUint32(p.PetNumber)
	out.WriteCString(p.PetName)
	out.WriteUint32(p.PetNameTimestamp)
	return
}

func (p *PetNameResponse) Decode(build vsn.Build, in *packet.WorldPacket) (err error) {
	p.PetNumber = in.ReadUint32()
	p.PetName = in.ReadCString()
	p.PetNameTimestamp = in.ReadUint32()
	return
}

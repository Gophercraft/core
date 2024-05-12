package query

import (
	"github.com/Gophercraft/core/game/protocol/message"
	"github.com/Gophercraft/core/guid"
	"github.com/Gophercraft/core/version"
)

type PetName struct {
	PetNumber uint32
	PetGUID   guid.GUID
}

func (p *PetName) Encode(build version.Build, out *message.Packet) (err error) {
	out.WriteUint32(p.PetNumber)
	return p.PetGUID.EncodeUnpacked(build, out)
}

func (p *PetName) Decode(build version.Build, in *message.Packet) (err error) {
	p.PetNumber = in.ReadUint32()
	p.PetGUID, err = guid.DecodeUnpacked(build, in)
	return
}

type PetNameResponse struct {
	PetNumber        uint32
	PetName          string
	PetNameTimestamp uint32
}

func (p *PetNameResponse) Encode(build version.Build, out *message.Packet) (err error) {
	out.WriteUint32(p.PetNumber)
	out.WriteCString(p.PetName)
	out.WriteUint32(p.PetNameTimestamp)
	return
}

func (p *PetNameResponse) Decode(build version.Build, in *message.Packet) (err error) {
	p.PetNumber = in.ReadUint32()
	p.PetName = in.ReadCString()
	p.PetNameTimestamp = in.ReadUint32()
	return
}

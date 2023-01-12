package character

import (
	"fmt"

	"github.com/Gophercraft/core/packet"
	"github.com/Gophercraft/core/realm/wdb/models"
	"github.com/Gophercraft/core/vsn"
)

type Create struct {
	Name       string
	Race       models.Race
	Class      models.Class
	BodyType   uint8
	Skin       uint8
	Face       uint8
	HairStyle  uint8
	HairColor  uint8
	FacialHair uint8
}

func (cc *Create) Decode(build vsn.Build, in *packet.WorldPacket) error {
	cc.Name = in.ReadCString()
	cc.Race = models.Race(in.ReadByte())
	cc.Class = models.Class(in.ReadByte())
	cc.BodyType = in.ReadByte()
	cc.Skin = in.ReadByte()
	cc.Face = in.ReadByte()
	cc.HairStyle = in.ReadByte()
	cc.HairColor = in.ReadByte()
	cc.FacialHair = in.ReadByte()
	return nil
}

func (cc *Create) Encode(build vsn.Build, out *packet.WorldPacket) error {
	out.Type = packet.CMSG_CHAR_CREATE
	out.WriteCString(cc.Name)
	out.WriteByte(uint8(cc.Race))
	out.WriteByte(uint8(cc.Class))
	out.WriteByte(cc.BodyType)
	out.WriteByte(cc.Skin)
	out.WriteByte(cc.Face)
	out.WriteByte(cc.HairStyle)
	out.WriteByte(cc.HairColor)
	out.WriteByte(cc.FacialHair)
	return nil
}

type CreateResult struct {
	Result
}

func decodeResult(build vsn.Build, in *packet.WorldPacket) (Result, error) {
	var resDescript ResultDescriptor
	if err := vsn.QueryDescriptors(build, ResultDescriptors, &resDescript); err != nil {
		return 0, err
	}

	code := in.ReadByte()
	for k, v := range resDescript {
		if v == code {
			return k, nil
		}
	}

	return 0, fmt.Errorf("character: descriptor %s does not have Result for %d", build, code)
}

func encodeResult(res Result, build vsn.Build, out *packet.WorldPacket) error {
	var resDescript ResultDescriptor
	if err := vsn.QueryDescriptors(build, ResultDescriptors, &resDescript); err != nil {
		return err
	}
	code, ok := resDescript[res]
	if !ok {
		return fmt.Errorf("character: encodeResult: descriptor %s does not have %d", build, res)
	}
	out.WriteByte(uint8(code))
	return nil
}

func (ccr *CreateResult) Encode(build vsn.Build, out *packet.WorldPacket) error {
	out.Type = packet.SMSG_CHAR_CREATE
	return encodeResult(ccr.Result, build, out)
}

func (ccr *CreateResult) Decode(build vsn.Build, in *packet.WorldPacket) (err error) {
	ccr.Result, err = decodeResult(build, in)
	return
}

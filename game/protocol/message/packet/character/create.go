package character

import (
	"fmt"

	"github.com/Gophercraft/core/game/protocol/message"
	"github.com/Gophercraft/core/realm/wdb/models"
	"github.com/Gophercraft/core/version"
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

func (cc *Create) Decode(build version.Build, in *message.Packet) error {
	cc.Name = in.ReadCString()
	cc.Race = models.Race(in.ReadUint8())
	cc.Class = models.Class(in.ReadUint8())
	cc.BodyType = in.ReadUint8()
	cc.Skin = in.ReadUint8()
	cc.Face = in.ReadUint8()
	cc.HairStyle = in.ReadUint8()
	cc.HairColor = in.ReadUint8()
	cc.FacialHair = in.ReadUint8()
	return nil
}

func (cc *Create) Encode(build version.Build, out *message.Packet) error {
	out.Type = message.CMSG_CHAR_CREATE
	out.WriteCString(cc.Name)
	out.WriteUint8(uint8(cc.Race))
	out.WriteUint8(uint8(cc.Class))
	out.WriteUint8(cc.BodyType)
	out.WriteUint8(cc.Skin)
	out.WriteUint8(cc.Face)
	out.WriteUint8(cc.HairStyle)
	out.WriteUint8(cc.HairColor)
	out.WriteUint8(cc.FacialHair)
	return nil
}

type CreateResult struct {
	Result
}

func decodeResult(build version.Build, in *message.Packet) (Result, error) {
	var resDescript ResultDescriptor
	if err := version.QueryDescriptors(build, ResultDescriptors, &resDescript); err != nil {
		return 0, err
	}

	code := in.ReadUint8()
	for k, v := range resDescript {
		if v == code {
			return k, nil
		}
	}

	return 0, fmt.Errorf("character: descriptor %s does not have Result for %d", build, code)
}

func encodeResult(res Result, build version.Build, out *message.Packet) error {
	var resDescript ResultDescriptor
	if err := version.QueryDescriptors(build, ResultDescriptors, &resDescript); err != nil {
		return err
	}
	code, ok := resDescript[res]
	if !ok {
		return fmt.Errorf("character: encodeResult: descriptor %s does not have %d", build, res)
	}
	out.WriteUint8(uint8(code))
	return nil
}

func (ccr *CreateResult) Encode(build version.Build, out *message.Packet) error {
	out.Type = message.SMSG_CHAR_CREATE
	return encodeResult(ccr.Result, build, out)
}

func (ccr *CreateResult) Decode(build version.Build, in *message.Packet) (err error) {
	ccr.Result, err = decodeResult(build, in)
	return
}

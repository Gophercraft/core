package character

import (
	"github.com/Gophercraft/core/game/protocol/message"
	"github.com/Gophercraft/core/guid"
	"github.com/Gophercraft/core/version"
)

type Delete struct {
	ID guid.GUID
}

func (dc *Delete) Decode(build version.Build, in *message.Packet) error {
	var err error
	dc.ID, err = guid.DecodeUnpacked(build, in)
	return err
}

func (dc *Delete) Encode(build version.Build, out *message.Packet) error {
	out.Type = message.CMSG_CHAR_DELETE
	dc.ID.EncodeUnpacked(build, out)
	return nil
}

type DeleteResult struct {
	Result
}

func (dcr *DeleteResult) Encode(build version.Build, out *message.Packet) error {
	out.Type = message.SMSG_CHAR_DELETE
	return encodeResult(dcr.Result, build, out)
}

func (dcr *DeleteResult) Decode(build version.Build, in *message.Packet) (err error) {
	dcr.Result, err = decodeResult(build, in)
	return
}

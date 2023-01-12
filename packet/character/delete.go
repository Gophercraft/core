package character

import (
	"github.com/Gophercraft/core/guid"
	"github.com/Gophercraft/core/packet"
	"github.com/Gophercraft/core/vsn"
)

type Delete struct {
	ID guid.GUID
}

func (dc *Delete) Decode(build vsn.Build, in *packet.WorldPacket) error {
	var err error
	dc.ID, err = guid.DecodeUnpacked(build, in)
	return err
}

func (dc *Delete) Encode(build vsn.Build, out *packet.WorldPacket) error {
	out.Type = packet.CMSG_CHAR_DELETE
	dc.ID.EncodeUnpacked(build, out)
	return nil
}

type DeleteResult struct {
	Result
}

func (dcr *DeleteResult) Encode(build vsn.Build, out *packet.WorldPacket) error {
	out.Type = packet.SMSG_CHAR_DELETE
	return encodeResult(dcr.Result, build, out)
}

func (dcr *DeleteResult) Decode(build vsn.Build, in *packet.WorldPacket) (err error) {
	dcr.Result, err = decodeResult(build, in)
	return
}

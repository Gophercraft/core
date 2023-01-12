package account

import (
	"time"

	"github.com/Gophercraft/core/packet"
	"github.com/Gophercraft/core/vsn"
)

// An MD5 hash for a wdb cache
type CacheHash [16]byte

// DataTimes seems to govern what the hashes are for a client's WDB cache
// This is necessary to get logged in correctly + and begin recording cache data
type DataTimes struct {
	Time        time.Time
	Activated   bool
	Mask        Cache
	CacheHashes []CacheHash
	Times       [32]uint32
}

type Cache uint32

const (
	GlobalConfigCache Cache = 1 << iota
	CharacterConfigCache
	GlobalBindingsCache
	CharacterBindingsCache
	GlobalMacrosCache
	CharacterMacrosCache
	CharacterLayoutCache
	CharacterChatCache
)

const (
	Global    = GlobalBindingsCache | GlobalConfigCache | GlobalMacrosCache
	Character = CharacterConfigCache | CharacterBindingsCache | CharacterMacrosCache | CharacterLayoutCache | CharacterChatCache
	All       = Global | Character
)

// 3988-6005
func (d *DataTimes) encodeClassic(build vsn.Build, out *packet.WorldPacket) error {
	const (
		nHash  = 8
		szHash = 16
	)

	times := make([]byte, szHash*nHash)

	n := len(d.CacheHashes)
	if n > nHash {
		n = nHash
	}

	for i := 0; i < n; i++ {
		write := times[i*szHash : (i+1)*szHash]
		copy(write, d.CacheHashes[i][:])
	}

	out.Write(times)
	return nil
}

// 6180-8303
func (d *DataTimes) encodeTBC(build vsn.Build, out *packet.WorldPacket) error {
	for i := 0; i < 31; i++ {
		out.WriteUint32(d.Times[i])
	}
	return nil
}

// 8089-8303
func (d *DataTimes) encodePreWotlk(build vsn.Build, out *packet.WorldPacket) error {
	for i := 0; i < 32; i++ {
		out.WriteUint32(d.Times[i])
	}
	return nil
}

func (d *DataTimes) encodeWotlk(build vsn.Build, out *packet.WorldPacket) error {
	out.WriteInt32(int32(d.Time.Unix()))
	out.WriteBool(d.Activated)
	out.WriteUint32(uint32(d.Mask))
	for i := Cache(0); i < 8; i++ {
		if d.Mask&(1<<i) != 0 {
			out.WriteUint32(d.Times[int(i)])
		}
	}
	return nil
}

func (d *DataTimes) Encode(build vsn.Build, out *packet.WorldPacket) error {
	out.Type = packet.SMSG_ACCOUNT_DATA_TIMES

	if build <= 6005 {
		return d.encodeClassic(build, out)
	}

	// if tbc
	if build <= 8606 {
		return d.encodeTBC(build, out)
	}

	if build < 8303 {
		return d.encodePreWotlk(build, out)
	}

	return d.encodeWotlk(build, out)
}

func (d *DataTimes) decodeClassic(build vsn.Build, in *packet.WorldPacket) error {
	d.CacheHashes = make([]CacheHash, 8)
	for i := 0; i < 8; i++ {
		in.Read(d.CacheHashes[i][:])
	}
	return nil
}

func (d *DataTimes) decodeTBC(build vsn.Build, in *packet.WorldPacket) error {
	for i := 0; i < 31; i++ {
		d.Times[i] = in.ReadUint32()
	}
	return nil
}

func (d *DataTimes) decodePreWotlk(build vsn.Build, in *packet.WorldPacket) error {
	for i := 0; i < 32; i++ {
		d.Times[i] = in.ReadUint32()
	}
	return nil
}

func (d *DataTimes) decodeWotlk(build vsn.Build, in *packet.WorldPacket) error {
	d.Time = time.Unix(int64(in.ReadUint32()), 0)
	d.Activated = in.ReadBool()
	d.Mask = Cache(in.ReadUint32())
	for i := Cache(0); i < 8; i++ {
		if d.Mask&(1<<i) != 0 {
			d.Times[int(i)] = in.ReadUint32()
		}
	}
	return nil
}

func (d *DataTimes) Decode(build vsn.Build, in *packet.WorldPacket) error {
	if build <= 6005 {
		return d.decodeClassic(build, in)
	}

	// if tbc
	if build <= 8606 {
		return d.decodeTBC(build, in)
	}

	if build < vsn.V3_0_2 {
		return d.decodePreWotlk(build, in)
	}

	return d.decodeWotlk(build, in)
}

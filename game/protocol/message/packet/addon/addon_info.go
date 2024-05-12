package addon

import (
	"compress/zlib"
	"fmt"
	"io/ioutil"

	"github.com/Gophercraft/core/game/protocol/message"
	"github.com/Gophercraft/core/version"
	"github.com/superp00t/etc"
)

var PublicKey = []byte{
	0xC3, 0x5B, 0x50, 0x84, 0xB9, 0x3E, 0x32, 0x42, 0x8C, 0xD0, 0xC7, 0x48, 0xFA, 0x0E, 0x5D, 0x54,
	0x5A, 0xA3, 0x0E, 0x14, 0xBA, 0x9E, 0x0D, 0xB9, 0x5D, 0x8B, 0xEE, 0xB6, 0x84, 0x93, 0x45, 0x75,
	0xFF, 0x31, 0xFE, 0x2F, 0x64, 0x3F, 0x3D, 0x6D, 0x07, 0xD9, 0x44, 0x9B, 0x40, 0x85, 0x59, 0x34,
	0x4E, 0x10, 0xE1, 0xE7, 0x43, 0x69, 0xEF, 0x7C, 0x16, 0xFC, 0xB4, 0xED, 0x1B, 0x95, 0x28, 0xA8,
	0x23, 0x76, 0x51, 0x31, 0x57, 0x30, 0x2B, 0x79, 0x08, 0x50, 0x10, 0x1C, 0x4A, 0x1A, 0x2C, 0xC8,
	0x8B, 0x8F, 0x05, 0x2D, 0x22, 0x3D, 0xDB, 0x5A, 0x24, 0x7A, 0x0F, 0x13, 0x50, 0x37, 0x8F, 0x5A,
	0xCC, 0x9E, 0x04, 0x44, 0x0E, 0x87, 0x01, 0xD4, 0xA3, 0x15, 0x94, 0x16, 0x34, 0xC6, 0xC2, 0xC3,
	0xFB, 0x49, 0xFE, 0xE1, 0xF9, 0xDA, 0x8C, 0x50, 0x3C, 0xBE, 0x2C, 0xBB, 0x57, 0xED, 0x46, 0xB9,
	0xAD, 0x8B, 0xC6, 0xDF, 0x0E, 0xD6, 0x0F, 0xBE, 0x80, 0xB3, 0x8B, 0x1E, 0x77, 0xCF, 0xAD, 0x22,
	0xCF, 0xB7, 0x4B, 0xCF, 0xFB, 0xF0, 0x6B, 0x11, 0x45, 0x2D, 0x7A, 0x81, 0x18, 0xF2, 0x92, 0x7E,
	0x98, 0x56, 0x5D, 0x5E, 0x69, 0x72, 0x0A, 0x0D, 0x03, 0x0A, 0x85, 0xA2, 0x85, 0x9C, 0xCB, 0xFB,
	0x56, 0x6E, 0x8F, 0x44, 0xBB, 0x8F, 0x02, 0x22, 0x68, 0x63, 0x97, 0xBC, 0x85, 0xBA, 0xA8, 0xF7,
	0xB5, 0x40, 0x68, 0x3C, 0x77, 0x86, 0x6F, 0x4B, 0xD7, 0x88, 0xCA, 0x8A, 0xD7, 0xCE, 0x36, 0xF0,
	0x45, 0x6E, 0xD5, 0x64, 0x79, 0x0F, 0x17, 0xFC, 0x64, 0xDD, 0x10, 0x6F, 0xF3, 0xF5, 0xE0, 0xA6,
	0xC3, 0xFB, 0x1B, 0x8C, 0x29, 0xEF, 0x8E, 0xE5, 0x34, 0xCB, 0xD1, 0x2A, 0xCE, 0x79, 0xC3, 0x9A,
	0x0D, 0x36, 0xEA, 0x01, 0xE0, 0xAA, 0x91, 0x20, 0x54, 0xF0, 0x72, 0xD8, 0x1E, 0xC7, 0x89, 0xD2,
}

type List struct {
	Addons                   []Info
	LastBannedAddOnTimestamp uint32
}

type Info struct {
	Name         string
	HasKey       bool
	PublicKeyCRC uint32
	UrlCRC       uint32
}

func (ai *Info) Decode(build version.Build, in *message.Packet) error {
	ai.Name = in.ReadCString()
	ai.HasKey = in.ReadBool()
	ai.PublicKeyCRC = in.ReadUint32()
	ai.UrlCRC = in.ReadUint32()
	return nil
}

func (list *List) Encode(build version.Build, out *message.Packet) error {
	return fmt.Errorf("nyi")
}

func (list *List) Decode(build version.Build, in *message.Packet) error {
	if in.Available() < 4 {
		return fmt.Errorf("packet: no addon info")
	}

	size := in.ReadUint32()
	if size == 0 {
		return fmt.Errorf("packet: AddonList too small")
	}

	if size > 0xFFFFF {
		return fmt.Errorf("packet: AddonList too large")
	}

	zReader, err := zlib.NewReader(in)
	if err != nil {
		return err
	}

	rawBytes, err := ioutil.ReadAll(zReader)
	if err != nil {
		return err
	}
	zReader.Close()

	if len(rawBytes) != int(size) {
		return fmt.Errorf("packet: mismatch between uncompressed size at start of packet: %d, and actual decompressed bytes from zlib: %d", int(size), len(rawBytes))
	}

	addonInfo := etc.OfBytes(rawBytes)

	infoCount := int(addonInfo.ReadUint32())
	if infoCount > 64 {
		infoCount = 64
	}

	for i := 0; i < infoCount; i++ {
		ai := Info{}
		err := ai.Decode(build, in)
		if err != nil {
			return err
		}
		list.Addons = append(list.Addons, ai)
	}

	list.LastBannedAddOnTimestamp = in.ReadUint32()

	return nil
}

type ServerState uint8

const (
	StateBanned ServerState = iota
	StateSecureVisible
	StateSecureHidden
)

type ServerAllowed struct {
	State       ServerState
	ProvideInfo bool
	NeedsKey    bool
	Revision    uint32 // Revision (from .toc), can be used by AddonSecureVisible to display "update available" in client addon controls
	URL         string
}

type ServerBanned struct {
	ID         uint32
	NameMD5    [16]byte
	VersionMD5 [16]byte
	Timestamp  uint32
	// Really? (maybe ServerBanned isn't the right name?)
	IsBanned uint32
}

type ServerResponse struct {
	Allowed []ServerAllowed
	Banned  []ServerBanned
}

func (sar *ServerResponse) Decode(build version.Build, out *message.Packet) error {
	return fmt.Errorf("nyi")
}

func (sar *ServerResponse) Encode(build version.Build, out *message.Packet) error {
	out.Type = message.SMSG_ADDON_INFO

	for _, allowed := range sar.Allowed {
		out.WriteUint8(uint8(allowed.State))
		out.WriteBool(allowed.ProvideInfo)

		if allowed.ProvideInfo {
			out.WriteBool(allowed.NeedsKey)
			if allowed.NeedsKey {
				out.Write(PublicKey)
			}
			out.WriteUint32(allowed.Revision)
		}
		out.WriteBool(len(allowed.URL) > 0)
		if len(allowed.URL) > 0 {
			out.WriteCString(allowed.URL)
		}
	}

	out.WriteUint32(uint32(len(sar.Banned)))
	for _, banned := range sar.Banned {
		out.WriteUint32(banned.ID)
		out.Write(banned.NameMD5[:])
		out.Write(banned.VersionMD5[:])
		out.WriteUint32(banned.Timestamp)
		out.WriteUint32(banned.IsBanned)
	}

	return nil
}

// This function is merely to skip addon checking. It's only here to get the client to connect.
// Not a proper implementation of actually banning addons designated as harmful.
// A proper implementation may be added soon™
func SkipServerCheck(build version.Build, info *List) *ServerResponse {
	// var crcCheck uint32 = 0x4c1c776d

	sar := new(ServerResponse)

	for _, addon := range info.Addons {
		allowed := ServerAllowed{}
		allowed.State = StateSecureHidden
		allowed.ProvideInfo = true
		allowed.NeedsKey = !addon.HasKey
		sar.Allowed = append(sar.Allowed, allowed)
	}

	return sar
}

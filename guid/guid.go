// Package guid stores the GUID or (Globally Unique Identifier), a 128-bit data type which can contain a type specifier, a server ID, and a global counter.
// This package uses the 128-bit format, but allows lossy conversion to and from the legacy 64-bit format
package guid

import (
	"encoding/binary"
	"fmt"
	"io"

	"github.com/Gophercraft/core/vsn"
)

var (
	// Nil represents the zero value of a GUID.
	Nil = GUID{0, 0}
	// Oldest known revision, unchanged until NewFormat
	OldFormat vsn.Build = 0
	// Starting in protocol build 19027, GUIDs have 128 bits instead of 64.
	NewFormat vsn.Build = 19027
)

// GUID or Global Unique Identification number

// A GUID cannot refer to more than 1 object.

// Any game object can be referenced by its GUID.
//
// GUIDs can have differently located and sized bitfields, depending on what type they represent.

// When using, keep in mind that a GUID can be any 128-bit number
// There are no hard guarantees about how well-formed it is. It may or may not contain the information you are requesting.

// Gophercraft GUIDs are back-compatible with legacy clients.
// This implementation has a strong bias toward the 6.0.2. GUID version, as it requires twice as much data to be stored.
// Therefore it is trivial to use the new format to store GUIDs on legacy protocol servers.

// The HighType of a GUID is stored in Hi, or the first 64-bit segment of the GUID-128.
//
// Hi can also store a 13-bit Server ID.
// Note that this limited size means that GUID can not be used as a unique character handle
// across Home networks with more than 8192 Worldservers. Gophercraft aims to be able to push realmlists bigger than this in the future.
type GUID struct {
	Hi uint64
	Lo uint64
}

// HighType returns a object type signifier
func (g GUID) HighType() HighType {
	return (HighType(g.Hi>>58) & 0x3F)
}

// HiClassic
func (g GUID) HiClassic() uint32 {
	var data [8]byte
	classic := g.Classic()
	binary.LittleEndian.PutUint64(data[:], classic)
	return binary.LittleEndian.Uint32(data[4:])
}

func (g GUID) LoClassic() uint32 {
	var data [8]byte
	classic := g.Classic()
	binary.LittleEndian.PutUint64(data[:], classic)
	return binary.LittleEndian.Uint32(data[:4])
}

func (g GUID) RealmID() uint32 {
	return uint32(g.Hi>>42) & MaxRealmID
}

func (g GUID) Counter() uint64 {
	switch g.HighType() {
	case Transport:
		return (g.Hi >> 38) & uint64(0xFFFFF)
	default:
		break
	}
	return g.Lo & uint64(0x000000FFFFFFFFFF)
}

func (g GUID) String() string {
	// return fmt.Sprintf("(%s) 0x%016X", g.HighType(), g.Lo)

	if g == Nil {
		return "Nil"
	}

	switch g.HighType() {
	case Player:
		return fmt.Sprintf("Player-%d-%08X", g.RealmID(), g.Counter())
	case Item:
		return fmt.Sprintf("Item-%d-%08X", g.RealmID(), g.Counter())
	default:
		return fmt.Sprintf("%s-%016X-%016X", g.HighType(), g.Hi, g.Lo)
	}
}

func (g GUID) Summary() string {
	return fmt.Sprintf("%s (%016X,%016X)", g.String(), g.Hi, g.Lo)
}

// read byte
func rb(r io.Reader) uint8 {
	b := make([]byte, 1)
	r.Read(b)
	return b[0]
}

func decodeMasked64(mask uint8, r io.Reader) uint64 {
	var res uint64

	for i := uint64(0); i < 8; i++ {
		if (mask & (1 << i)) != 0 {
			res += uint64(rb(r)) << (i * 8)
		}
	}

	return res
}

func DecodePacked64(r io.Reader) uint64 {
	mask := rb(r)
	if mask == 0 {
		return 0
	}

	return decodeMasked64(mask, r)
}

func DecodePacked128(r io.Reader) GUID {
	loMask := rb(r)
	hiMask := rb(r)

	lo := decodeMasked64(loMask, r)
	hi := decodeMasked64(hiMask, r)

	return GUID{hi, lo}
}

func (g GUID) IsUnit() bool {
	return g.HighType() == Creature || g.HighType() == Player
}

func encodeMasked64(value uint64) (uint8, []byte) {
	bitMask := uint8(0)
	packGUID := make([]byte, 8)
	size := 0

	for i := uint64(0); value != 0; i++ {
		// Convert byte
		if (value & 0xFF) > 0 {
			bitMask |= uint8(1 << i)
			packGUID[size] = uint8(value & 0xFF)
			size++
		}

		// Read next byte
		value >>= 8
	}

	return bitMask, packGUID[:size]
}

// EncodePacked GUIDs are encoded in a couple ways. 64-bit GUIDs can be encoded plainly as an 8-byte field,
// or in the "packed" format, a very simple compression mechanism.
// The packed format is a 8-bit mask value, followed by up to 8 bytes.
// If a bit n is true in the mask, it means that there is a byte that follows it at that position.
// If bit n is false, there is no byte at its position and its decoded byte should be zero.
// Example (3 bytes):
//
//	byte( 0 1 0 0 0 0 1 0 ) + byte(31) + byte(36)
//
// Is decoded as (8 bytes):
//
//	[8]byte{0, 31, 0, 0, 0, 0, 36, 0}
//
// The 128-bit format uses the exact same packing scheme, just with a 16-bit mask and up to 16 bytes following it.
func (g GUID) EncodePacked(version vsn.Build, w io.Writer) {
	switch {
	// Packing is not enabled in alpha
	case version.RemovedIn(vsn.V1_12_1):
		g.EncodeUnpacked(version, w)
	case version < NewFormat:
		// Resolve to legacy format
		mask, bytes := encodeMasked64(g.Classic())
		w.Write([]byte{mask})
		if mask > 0 {
			w.Write(bytes)
		}
	default:
		loMask, loBytes := encodeMasked64(g.Lo)
		hiMask, hiBytes := encodeMasked64(g.Hi)
		// 16-bit mask
		w.Write([]byte{loMask, hiMask})
		w.Write(loBytes)
		w.Write(hiBytes)
	}
}

// (2 ** 35 bits) - 1
const MaxCharactersLegacy = 0x7FFFFFFFF

// (2 ** 13 bits) - 1
const MaxRealmID = 0x1FFF

// Classic provides backward compatibility with the legacy GUID format in IO streams.
// It is important to use this function understanding its limitations. the player counter must not be allowed to exceed guid.MaxCharactersLegacy.
func (g GUID) Classic() uint64 {
	highType := uint64(oldHighTypes[g.HighType()])
	// set type as uint16
	u64 := (highType << 48)

	switch g.HighType() {
	case Player, Item, GameObject:
		// This formatting is unique to Gophercraft:
		// type RealmSpecificObjectGUID struct {
		//   Type uint16
		//   RealmID uint13 ; GUID can have one of 8192 (2^13) realms IDs Ideally, I would like to allow more, but Bliz thinks this is all anyone should need.
		//   Counter uint35 ; Amount of characters in the legacy protocol = 0x800000000 or 34,359,738,368 (2^35) This should be more than enough; a conservative estimate of 500 bytes per character would require 2.19 petabytes of storage space to hold all these characters.
		// }
		u64 |= uint64(g.RealmID()) << 35
		// Security risk if counter is allowed to exceed MaxCharactersLegacy
		u64 |= uint64(g.Counter()) & MaxCharactersLegacy
	default:
		// Masked to 48 bits so as to not overwrite the 16-bit type field
		u64 |= g.Lo & 0xFFFFFFFFFFFF
	}

	return u64
}

func (g GUID) EncodeUnpacked(version vsn.Build, w io.Writer) error {
	switch {
	case version < NewFormat:
		e := make([]byte, 8)
		binary.LittleEndian.PutUint64(e, g.Classic())
		_, err := w.Write(e)
		return err
	default:
		panic(fmt.Errorf("update: can't encode GUID in unpacked format in version %d", version))
	}
}

func (g GUID) Cmp(g2 GUID) int {
	if g == g2 {
		return 2 // exact match.
	}

	if g.HighType() == Player {
		if g.Counter() == g2.Counter() {
			return 1 // Same GUID counter, not the same realm ID
		}
	}

	return 0
}

func (g GUID) SetRealmID(realmID uint64) GUID {
	g2 := g
	g2.Hi |= (realmID << 42)
	return g2
}

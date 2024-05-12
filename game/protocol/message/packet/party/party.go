package party

import (
	"encoding/binary"
	"fmt"
	"io"

	"github.com/Gophercraft/core/game/protocol/message"
	"github.com/Gophercraft/core/guid"
	"github.com/Gophercraft/core/version"
)

const (
	GroupNormal = 0
	GroupRaid   = 1
)

type MemberStatus uint8

const (
	MemberOffline = 0x0000
	MemberOnline  = 0x0001 // Lua_UnitIsConnected
	MemberPVP     = 0x0002 // Lua_UnitIsPVP
	MemberDead    = 0x0004 // Lua_UnitIsDead
	MemberGhost   = 0x0008 // Lua_UnitIsGhost
	MemberPVPFFA  = 0x0010 // Lua_UnitIsPVPFreeForAll
	MemberZoneOut = 0x0020 // Lua_GetPlayerMapPosition
	MemberAFK     = 0x0040 // Lua_UnitIsAFK
	MemberDND     = 0x0080 // Lua_UnitIsDND
)

const (
	GroupUpdateNone            = 0x00000000 // nothing
	GroupUpdateStatus          = 0x00000001 // uint16, flags
	GroupUpdateCurrentHealth   = 0x00000002 // uint32
	GroupUpdateMaxHealth       = 0x00000004 // uint32
	GroupUpdatePowerType       = 0x00000008 // uint8
	GroupUpdateCurrentPower    = 0x00000010 // uint16
	GroupUpdateMaxPower        = 0x00000020 // uint16
	GroupUpdateLevel           = 0x00000040 // uint16
	GroupUpdateZone            = 0x00000080 // uint16
	GroupUpdatePosition        = 0x00000100 // uint16, uint16
	GroupUpdateAuras           = 0x00000200 // uint64 mask, for each bit set uint32 spellid + uint8 unk
	GroupUpdatePetGUID         = 0x00000400 // uint64 pet guid
	GroupUpdatePetName         = 0x00000800 // pet name, nullptr terminated string
	GroupUpdatePetModelID      = 0x00001000 // uint16, model id
	GroupUpdatePetCurrentHP    = 0x00002000 // uint32 pet cur health
	GroupUpdatePetMaxHP        = 0x00004000 // uint32 pet max health
	GroupUpdatePetPowerType    = 0x00008000 // uint8 pet power type
	GroupUpdatePetCurrentPower = 0x00010000 // uint16 pet cur power
	GroupUpdatePetMaxPower     = 0x00020000 // uint16 pet max power
	GroupUpdatePetAuras        = 0x00040000 // uint64 mask, for each bit set uint32 spellid + uint8 unk, pet auras...
	GroupUpdateVehicleSeat     = 0x00080000 // uint32 vehicle_seat_id (index from VehicleSeat.dbc)
	GroupUpdatePet             = 0x0007FC00 // all pet flags
	GroupUpdateFull            = 0x0007FFFF // all known flags
)

type Operation uint32

const (
	Invite Operation = 0
	Leave  Operation = 2
	Swap   Operation = 4
)

type Result uint32

const (
	OK Result = iota
	BadPlayerName
	TargetNotInGroup
	TargetNotInInstance
	GroupFull
	AlreadyInGroup
	NotInGroup
	NotLeader
	WrongFaction
	IgnoringYou
	InviteRestricted
	LFGPending
)

type ResultDescriptor map[Result]uint32

var ResultDescriptors = map[version.BuildRange]ResultDescriptor{
	{0, 3368}: {
		OK:                  0,
		BadPlayerName:       1,
		TargetNotInGroup:    2,
		TargetNotInInstance: 3,
		GroupFull:           4,
		AlreadyInGroup:      5,
		NotInGroup:          6,
		NotLeader:           7,
		WrongFaction:        8,
		IgnoringYou:         9,
		InviteRestricted:    13,
	},

	{5875, 6005}: {
		OK:               0,
		BadPlayerName:    1,
		TargetNotInGroup: 2,
		GroupFull:        3,
		AlreadyInGroup:   4,
		NotInGroup:       5,
		NotLeader:        6,
		WrongFaction:     7,
		IgnoringYou:      8,
	},

	{version.V2_0_1, version.Max}: {
		OK:                  0,
		BadPlayerName:       1,
		TargetNotInGroup:    2,
		TargetNotInInstance: 3,
		GroupFull:           4,
		AlreadyInGroup:      5,
		NotInGroup:          6,
		NotLeader:           7,
		WrongFaction:        8,
		IgnoringYou:         9,
		LFGPending:          12,
		InviteRestricted:    13,
	},
}

func EncodeResult(build version.Build, out io.Writer, pr Result) error {
	var desc ResultDescriptor
	if err := version.QueryDescriptors(build, ResultDescriptors, &desc); err != nil {
		return err
	}

	op, ok := desc[pr]
	if !ok {
		return fmt.Errorf("packet: no Result code found for %v in %s", pr, build)
	}

	var u32 [4]byte
	binary.LittleEndian.PutUint32(u32[:], op)

	_, err := out.Write(u32[:])
	return err
}

func DecodeResult(build version.Build, in io.Reader) (Result, error) {
	var desc ResultDescriptor
	if err := version.QueryDescriptors(build, ResultDescriptors, &desc); err != nil {
		return 0, err
	}

	var u32 [4]byte

	if _, err := in.Read(u32[:]); err != nil {
		return 0, err
	}

	code := binary.LittleEndian.Uint32(u32[:])

	for k, vcode := range desc {
		if vcode == code {
			return k, nil
		}
	}

	return 0, fmt.Errorf("packet: no Result code found for %v in %s", code, build)
}

type RequestMemberStats struct {
	Member guid.GUID
}

func (p *RequestMemberStats) Encode(build version.Build, out *message.Packet) error {
	out.Type = message.CMSG_REQUEST_PARTY_MEMBER_STATS
	p.Member.EncodeUnpacked(build, out)
	return nil
}

func (p *RequestMemberStats) Decode(build version.Build, in *message.Packet) error {
	var err error
	p.Member, err = guid.DecodeUnpacked(build, in)
	return err
}

type MemberAuraState struct {
	SpellID     int32
	Flags       uint16
	ActiveFlags uint32
	Points      []float32
}

type MemberStats struct {
	ChangeMask uint32

	Member guid.GUID

	Status        MemberStatus
	CurrentHealth uint32
	MaxHealth     uint32

	PowerType uint8

	Power uint16

	MaxPower uint16

	Level uint16

	ZoneID uint16

	PositionX uint16
	PositionY uint16

	Auras []MemberAuraState

	Pet              guid.GUID
	PetName          string
	PetDisplayID     uint32
	PetCurrentHealth uint32
	PetMaxHealth     uint32
	PetPowerType     uint16
	PetPower         uint16
	PetMaxPower      uint16
	PetAuras         []MemberAuraState

	VehicleSeat uint32
}

func (p *MemberStats) Change(status uint32) bool {
	return p.ChangeMask&status != 0
}

func writeHealth(build version.Build, value uint32, out *message.Packet) {
	if build.AddedIn(version.V2_0_1) {
		out.WriteUint32(value)
		return
	}

	out.WriteUint16(uint16(value))
}

func readHealth(build version.Build, in *message.Packet) uint32 {
	if build.AddedIn(version.V2_0_1) {
		return in.ReadUint32()
	}
	return uint32(in.ReadUint16())
}

func writeAuras(build version.Build, auras []MemberAuraState, out *message.Packet) {
	var mask uint64

	for i, aura := range auras {
		if aura.SpellID > 0 {
			mask |= 1 << uint64(i)
		}
	}

	if build.AddedIn(version.V2_0_1) {
		out.WriteUint64(mask)
	} else {
		out.WriteUint32(uint32(mask))
	}

	for i, aura := range auras {
		if mask&(1<<uint64(i)) != 0 {
			switch {
			case build < version.V2_0_1:
				out.WriteUint16(uint16(aura.SpellID))
			case version.Range(version.V2_0_1, version.V3_3_5a).Contains(build):
				out.WriteInt32(aura.SpellID)
				out.WriteUint8(uint8(aura.Flags))
			default:
				panic("verify behavior in writeAuras for this build")
			}
		}
	}
}

func readAuras(build version.Build, in *message.Packet) []MemberAuraState {
	var mask uint64
	var maskLen int
	if build.AddedIn(version.V2_0_1) {
		mask = in.ReadUint64()
		maskLen = 64
	} else {
		mask = uint64(in.ReadUint32())
		maskLen = 32
	}

	pmas := make([]MemberAuraState, maskLen)

	for i := 0; i < maskLen; i++ {
		if mask&(1<<uint64(i)) != 0 {
			state := &pmas[i]
			switch {
			case build < version.V2_0_1:
				state.SpellID = in.ReadInt32()
			case version.Range(version.V2_0_1, version.V3_3_5a).Contains(build):
				state.SpellID = in.ReadInt32()
				state.Flags = uint16(in.ReadUint8())
			default:
				panic("verify behavior in readAuras for this build")
			}
		}
	}

	return pmas
}

func (p *MemberStats) Encode(build version.Build, out *message.Packet) error {
	out.Type = message.SMSG_PARTY_MEMBER_STATS_FULL
	if build.AddedIn(version.V2_0_1) {
		out.WriteUint8(0)
	}
	p.Member.EncodePacked(build, out)
	out.WriteUint32(p.ChangeMask)
	if p.Change(GroupUpdateStatus) {
		out.WriteUint8(uint8(p.Status))
	}
	if p.Change(GroupUpdateCurrentHealth) {
		writeHealth(build, p.CurrentHealth, out)
	}
	if p.Change(GroupUpdateMaxHealth) {
		writeHealth(build, p.MaxHealth, out)
	}
	if p.Change(GroupUpdatePowerType) {
		out.WriteUint8(p.PowerType)
	}
	if p.Change(GroupUpdateCurrentPower) {
		out.WriteUint16(p.Power)
	}
	if p.Change(GroupUpdateMaxPower) {
		out.WriteUint16(p.MaxPower)
	}
	if p.Change(GroupUpdateLevel) {
		out.WriteUint16(p.Level)
	}
	if p.Change(GroupUpdateZone) {
		out.WriteUint16(p.ZoneID)
	}
	if p.Change(GroupUpdatePosition) {
		out.WriteUint16(p.PositionX)
		out.WriteUint16(p.PositionY)
	}
	if p.Change(GroupUpdateAuras) {
		writeAuras(build, p.Auras, out)
	}
	if p.Change(GroupUpdatePetGUID) {
		p.Pet.EncodeUnpacked(build, out)
	}
	if p.Change(GroupUpdatePetName) {
		out.WriteCString(p.PetName)
	}
	if p.Change(GroupUpdatePetModelID) {
		out.WriteUint16(uint16(p.PetDisplayID))
	}
	if p.Change(GroupUpdatePetCurrentHP) {
		writeHealth(build, p.PetCurrentHealth, out)
	}
	if p.Change(GroupUpdatePetMaxHP) {
		writeHealth(build, p.PetMaxHealth, out)
	}
	if p.Change(GroupUpdatePetPowerType) {
		out.WriteUint8(uint8(p.PetPowerType))
	}
	if p.Change(GroupUpdatePetCurrentPower) {
		out.WriteUint16(p.PetPower)
	}
	if p.Change(GroupUpdatePetMaxPower) {
		out.WriteUint16(p.PetMaxPower)
	}
	if p.Change(GroupUpdatePetAuras) {
		writeAuras(build, p.PetAuras, out)
	}
	if p.Change(GroupUpdateVehicleSeat) {
		out.WriteUint32(uint32(p.VehicleSeat))
	}
	return nil
}

func (p *MemberStats) Decode(build version.Build, in *message.Packet) error {
	var err error
	if build.AddedIn(version.V2_0_1) {
		in.ReadUint8()
	}
	p.Member, err = guid.DecodePacked(build, in)
	if err != nil {
		return err
	}
	p.ChangeMask = in.ReadUint32()
	if p.Change(GroupUpdateStatus) {
		p.Status = MemberStatus(in.ReadUint8())
	}
	if p.Change(GroupUpdateCurrentHealth) {
		p.CurrentHealth = readHealth(build, in)
	}
	if p.Change(GroupUpdateMaxHealth) {
		p.MaxHealth = readHealth(build, in)
	}
	if p.Change(GroupUpdatePowerType) {
		p.PowerType = in.ReadUint8()
	}
	if p.Change(GroupUpdateCurrentPower) {
		p.Power = in.ReadUint16()
	}
	if p.Change(GroupUpdateMaxPower) {
		p.MaxPower = in.ReadUint16()
		p.MaxPower = in.ReadBigUint16()
	}
	if p.Change(GroupUpdateLevel) {
		p.Level = in.ReadUint16()
	}
	if p.Change(GroupUpdateZone) {
		p.ZoneID = in.ReadUint16()
	}
	if p.Change(GroupUpdatePosition) {
		p.PositionX = in.ReadUint16()
		p.PositionY = in.ReadUint16()
	}
	if p.Change(GroupUpdateAuras) {
		p.Auras = readAuras(build, in)
	}
	if p.Change(GroupUpdatePetGUID) {
		p.Pet, err = guid.DecodeUnpacked(build, in)
		if err != nil {
			return err
		}
	}
	if p.Change(GroupUpdatePetName) {
		p.PetName = in.ReadCString()
	}
	if p.Change(GroupUpdatePetModelID) {
		p.PetDisplayID = uint32(in.ReadUint16())
	}
	if p.Change(GroupUpdatePetCurrentHP) {
		p.PetCurrentHealth = readHealth(build, in)
	}
	if p.Change(GroupUpdatePetMaxHP) {
		p.PetMaxHealth = readHealth(build, in)
	}
	if p.Change(GroupUpdatePetPowerType) {
		p.PetPowerType = uint16(in.ReadUint8())
	}
	if p.Change(GroupUpdatePetCurrentPower) {
		p.PetPower = in.ReadUint16()
	}
	if p.Change(GroupUpdatePetMaxPower) {
		p.PetMaxPower = in.ReadUint16()
	}
	if p.Change(GroupUpdatePetAuras) {
		p.PetAuras = readAuras(build, in)
	}
	if p.Change(GroupUpdateVehicleSeat) {
		p.VehicleSeat = in.ReadUint32()
	}
	return nil
}

type CommandResult struct {
	Operation Operation
	Member    string
	Result    Result
}

func (pcr *CommandResult) Encode(build version.Build, out *message.Packet) error {
	out.Type = message.SMSG_PARTY_COMMAND_RESULT
	out.WriteUint32(uint32(pcr.Operation))
	out.WriteCString(pcr.Member)
	return EncodeResult(build, out, pcr.Result)
}

func (pcr *CommandResult) Decode(build version.Build, in *message.Packet) (err error) {
	pcr.Operation = Operation(in.ReadUint32())
	pcr.Member = in.ReadCString()
	pcr.Result, err = DecodeResult(build, in)
	return
}

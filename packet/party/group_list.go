package party

import (
	"github.com/Gophercraft/core/guid"
	"github.com/Gophercraft/core/packet"
	"github.com/Gophercraft/core/realm/wdb/models"
	"github.com/Gophercraft/core/vsn"
)

type GroupType uint8

const (
	Normal              GroupType = 0x00
	Battleground        GroupType = 0x01
	Raid                GroupType = 0x02
	LFGRestricted       GroupType = 0x04
	LookingForDungeon   GroupType = 0x08
	Unk10               GroupType = 0x10
	OnePersonParty      GroupType = 0x20
	IsEveryoneAssistant GroupType = 0x40
)

type InstanceStatus uint

const (
	NotSaved  InstanceStatus = 0
	Saved     InstanceStatus = 1
	Completed InstanceStatus = 2
)

type MapDifficulty uint8

type LootMethod uint8

type LfgRoleFlag uint8

type GroupUpdateFlags uint32

type GroupMemberStatus uint8

type GroupMember struct {
	Name        string
	GUID        guid.GUID
	Status      GroupMemberStatus
	SubGroup    uint8
	UpdateFlags GroupUpdateFlags
	LfgRole     LfgRoleFlag
}

func (member *GroupMember) Encode(build vsn.Build, out *packet.WorldPacket) error {
	out.WriteCString(member.Name)
	member.GUID.EncodeUnpacked(build, out)
	out.WriteByte(uint8(member.Status))
	out.WriteByte(member.SubGroup)
	out.WriteByte(uint8(member.UpdateFlags))
	if build.AddedIn(10958) {
		out.WriteByte(uint8(member.LfgRole))
	}
	return nil
}

func (member *GroupMember) Decode(build vsn.Build, in *packet.WorldPacket) (err error) {
	member.Name = in.ReadCString()
	member.GUID, err = guid.DecodeUnpacked(build, in)
	if err != nil {
		return err
	}
	member.Status = GroupMemberStatus(in.ReadByte())
	member.SubGroup = in.ReadByte()
	member.UpdateFlags = GroupUpdateFlags(in.ReadByte())
	if build.AddedIn(10958) {
		member.LfgRole = LfgRoleFlag(in.ReadByte())
	}
	return nil
}

type GroupList struct {
	GroupType         GroupType
	SubGroup          uint8
	Flags             GroupUpdateFlags
	RolesAssigned     uint8
	GroupTypeStatus   InstanceStatus
	LFGentry          uint32
	GUID              guid.GUID
	Counter           int32
	Members           []GroupMember
	Leader            guid.GUID
	LootMethod        LootMethod
	Looter            guid.GUID
	LootThreshold     models.ItemQuality
	DungeonDifficulty MapDifficulty
	RaidDifficulty    MapDifficulty

	// alpha only:
	LeaderName string
}

func (list *GroupList) Encode(build vsn.Build, out *packet.WorldPacket) (err error) {
	out.Type = packet.SMSG_GROUP_LIST

	out.WriteByte(uint8(list.GroupType))
	out.WriteByte(list.SubGroup)

	out.WriteByte(uint8(list.Flags))
	out.WriteByte(list.RolesAssigned)

	if list.GroupType&LookingForDungeon != 0 {
		out.WriteByte(uint8(list.GroupTypeStatus))
		out.WriteUint32(list.LFGentry)
		if build.AddedIn(14545) {
			out.WriteBool(true)
		}
	}

	err = list.GUID.EncodeUnpacked(build, out)
	if err != nil {
		return
	}

	if build.AddedIn(10958) {
		out.WriteInt32(list.Counter)
	}

	out.WriteInt32(int32(len(list.Members)))

	if build <= vsn.Alpha {
		out.WriteCString(list.LeaderName)
		list.Leader.EncodeUnpacked(build, out)
		out.WriteByte(1)
	}

	for i := 0; i < len(list.Members); i++ {
		if err = list.Members[i].Encode(build, out); err != nil {
			return
		}
	}

	if build > vsn.Alpha {
		err = list.Leader.EncodeUnpacked(build, out)
		if err != nil {
			return
		}
	}

	if len(list.Members) <= 0 {
		return
	}

	out.WriteByte(uint8(list.LootMethod))

	err = list.Looter.EncodeUnpacked(build, out)
	if err != nil {
		return
	}

	out.WriteByte(uint8(list.LootThreshold))

	if build.AddedIn(10192) {
		out.WriteByte(uint8(list.DungeonDifficulty))
	}

	out.WriteByte(uint8(list.RaidDifficulty))

	if build.AddedIn(10958) && build.RemovedIn(13623) {
		out.WriteByte(0x00) // Has something to do with difficulty too
	}

	return nil
}

func (list *GroupList) Decode(build vsn.Build, in *packet.WorldPacket) (err error) {
	list.GroupType = GroupType(in.ReadByte())
	list.SubGroup = in.ReadByte()
	list.Flags = GroupUpdateFlags(in.ReadByte())
	list.RolesAssigned = in.ReadByte()

	if list.GroupType&LookingForDungeon != 0 {
		list.GroupTypeStatus = InstanceStatus(in.ReadByte())
		list.LFGentry = in.ReadUint32()
		if build.AddedIn(14545) {
			in.ReadBool()
		}
	}

	list.GUID, err = guid.DecodeUnpacked(build, in)
	if err != nil {
		return
	}

	if build.AddedIn(10958) {
		list.Counter = in.ReadInt32()
	}

	memberCount := int(in.ReadInt32())
	if memberCount > 200 {
		return nil
	}

	if build <= vsn.Alpha {
		list.LeaderName = in.ReadCString()
		list.Leader, err = guid.DecodeUnpacked(build, in)
		if err != nil {
			return
		}
		in.ReadByte()
	}

	list.Members = make([]GroupMember, memberCount)

	for i := 0; i < memberCount; i++ {
		if err = list.Members[i].Decode(build, in); err != nil {
			return
		}
	}

	if build > vsn.Alpha {
		list.Leader, err = guid.DecodeUnpacked(build, in)
		if err != nil {
			return
		}
	}

	if memberCount <= 0 {
		return
	}

	list.LootMethod = LootMethod(in.ReadByte())

	list.Looter, err = guid.DecodeUnpacked(build, in)
	if err != nil {
		return
	}

	list.LootThreshold = models.ItemQuality(in.ReadByte())

	if build.AddedIn(10192) {
		list.DungeonDifficulty = MapDifficulty(in.ReadByte())
	}

	list.RaidDifficulty = MapDifficulty(in.ReadByte())

	if build.AddedIn(10958) && build.RemovedIn(13623) {
		in.ReadByte() // Has something to do with difficulty too
	}

	return nil
}

package party

import (
	"github.com/Gophercraft/core/game/protocol/message"
	"github.com/Gophercraft/core/guid"
	"github.com/Gophercraft/core/realm/wdb/models"
	"github.com/Gophercraft/core/version"
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

func (member *GroupMember) Encode(build version.Build, out *message.Packet) error {
	out.WriteCString(member.Name)
	member.GUID.EncodeUnpacked(build, out)
	out.WriteUint8(uint8(member.Status))
	out.WriteUint8(member.SubGroup)
	out.WriteUint8(uint8(member.UpdateFlags))
	if build.AddedIn(10958) {
		out.WriteUint8(uint8(member.LfgRole))
	}
	return nil
}

func (member *GroupMember) Decode(build version.Build, in *message.Packet) (err error) {
	member.Name = in.ReadCString()
	member.GUID, err = guid.DecodeUnpacked(build, in)
	if err != nil {
		return err
	}
	member.Status = GroupMemberStatus(in.ReadUint8())
	member.SubGroup = in.ReadUint8()
	member.UpdateFlags = GroupUpdateFlags(in.ReadUint8())
	if build.AddedIn(10958) {
		member.LfgRole = LfgRoleFlag(in.ReadUint8())
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

func (list *GroupList) Encode(build version.Build, out *message.Packet) (err error) {
	out.Type = message.SMSG_GROUP_LIST

	out.WriteUint8(uint8(list.GroupType))
	out.WriteUint8(list.SubGroup)

	out.WriteUint8(uint8(list.Flags))
	out.WriteUint8(list.RolesAssigned)

	if list.GroupType&LookingForDungeon != 0 {
		out.WriteUint8(uint8(list.GroupTypeStatus))
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

	if build <= version.Alpha {
		out.WriteCString(list.LeaderName)
		list.Leader.EncodeUnpacked(build, out)
		out.WriteUint8(1)
	}

	for i := 0; i < len(list.Members); i++ {
		if err = list.Members[i].Encode(build, out); err != nil {
			return
		}
	}

	if build > version.Alpha {
		err = list.Leader.EncodeUnpacked(build, out)
		if err != nil {
			return
		}
	}

	if len(list.Members) <= 0 {
		return
	}

	out.WriteUint8(uint8(list.LootMethod))

	err = list.Looter.EncodeUnpacked(build, out)
	if err != nil {
		return
	}

	out.WriteUint8(uint8(list.LootThreshold))

	if build.AddedIn(10192) {
		out.WriteUint8(uint8(list.DungeonDifficulty))
	}

	out.WriteUint8(uint8(list.RaidDifficulty))

	if build.AddedIn(10958) && build.RemovedIn(13623) {
		out.WriteUint8(0x00) // Has something to do with difficulty too
	}

	return nil
}

func (list *GroupList) Decode(build version.Build, in *message.Packet) (err error) {
	list.GroupType = GroupType(in.ReadUint8())
	list.SubGroup = in.ReadUint8()
	list.Flags = GroupUpdateFlags(in.ReadUint8())
	list.RolesAssigned = in.ReadUint8()

	if list.GroupType&LookingForDungeon != 0 {
		list.GroupTypeStatus = InstanceStatus(in.ReadUint8())
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

	if build <= version.Alpha {
		list.LeaderName = in.ReadCString()
		list.Leader, err = guid.DecodeUnpacked(build, in)
		if err != nil {
			return
		}
		in.ReadUint8()
	}

	list.Members = make([]GroupMember, memberCount)

	for i := 0; i < memberCount; i++ {
		if err = list.Members[i].Decode(build, in); err != nil {
			return
		}
	}

	if build > version.Alpha {
		list.Leader, err = guid.DecodeUnpacked(build, in)
		if err != nil {
			return
		}
	}

	if memberCount <= 0 {
		return
	}

	list.LootMethod = LootMethod(in.ReadUint8())

	list.Looter, err = guid.DecodeUnpacked(build, in)
	if err != nil {
		return
	}

	list.LootThreshold = models.ItemQuality(in.ReadUint8())

	if build.AddedIn(10192) {
		list.DungeonDifficulty = MapDifficulty(in.ReadUint8())
	}

	list.RaidDifficulty = MapDifficulty(in.ReadUint8())

	if build.AddedIn(10958) && build.RemovedIn(13623) {
		in.ReadUint8() // Has something to do with difficulty too
	}

	return nil
}

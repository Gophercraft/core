package quest

import (
	"fmt"

	"github.com/Gophercraft/core/guid"
	"github.com/Gophercraft/core/packet"
	"github.com/Gophercraft/core/vsn"
)

type GiverStatusDescriptor map[GiverStatus]uint32

var GiverStatusDescriptors = map[vsn.BuildRange]GiverStatusDescriptor{
	{0, vsn.Alpha}: {
		// QUEST_GIVER_NONE = 0
		// QUEST_GIVER_TRIVIAL = 1
		// QUEST_GIVER_FUTURE = 2
		// QUEST_GIVER_REWARD = 3
		// QUEST_GIVER_QUEST = 4
		// QUEST_GIVER_NUMITEMS = 5
		None:              0,
		Unavailable:       0,
		LowLevelAvailable: 2,
		Available:         4,
	},
}

type GiverStatusQuery struct {
	ID guid.GUID
}

func (q *GiverStatusQuery) Encode(build vsn.Build, out *packet.WorldPacket) (err error) {
	out.Type = packet.CMSG_QUESTGIVER_STATUS_QUERY
	err = q.ID.EncodeUnpacked(build, out)
	return
}

func (q *GiverStatusQuery) Decode(build vsn.Build, in *packet.WorldPacket) (err error) {
	q.ID, err = guid.DecodeUnpacked(build, in)
	return
}

type GiverStatus uint32

const (
	None GiverStatus = iota
	Unavailable
	LowLevelAvailable
	LowLevelRewardRep
	LowLevelAvailableRep
	Incomplete
	RewardRep
	AvailableRep
	Available
	Reward2
	Reward
)

type GiverStatusResponse struct {
	ID     guid.GUID
	Status GiverStatus
}

func statusSize(build vsn.Build) int {
	switch {
	case build == vsn.Alpha:
		return 4
	default:
		return 1
	}
}

func (r *GiverStatusResponse) Encode(build vsn.Build, out *packet.WorldPacket) (err error) {
	out.Type = packet.SMSG_QUESTGIVER_STATUS
	err = r.ID.EncodeUnpacked(build, out)
	if err != nil {
		return
	}

	// if build.RemovedIn(4000) {
	// 	out.WriteUint32(uint32(r.Status))
	// 	return
	// }

	// if build.RemovedIn(13623) {
	// 	out.WriteByte(uint8(r.Status))
	// }

	var descriptor GiverStatusDescriptor
	if err = vsn.QueryDescriptors(build, GiverStatusDescriptors, &descriptor); err != nil {
		return
	}

	status, ok := descriptor[r.Status]
	if !ok {
		err = fmt.Errorf("no giver status %d", r.Status)
		return
	}

	switch statusSize(build) {
	case 1:
		out.WriteByte(uint8(status))
	case 4:
		out.WriteUint32(uint32(status))
	}

	return
}

func (r *GiverStatusResponse) Decode(build vsn.Build, in *packet.WorldPacket) (err error) {
	r.ID, err = guid.DecodeUnpacked(build, in)
	if err != nil {
		return err
	}

	// if build.RemovedIn(4000) {
	// 	r.Status = GiverStatus(in.ReadUint32())
	// 	return
	// }

	// if build.RemovedIn(13623) {
	// 	r.Status = GiverStatus(in.ReadByte())
	// }

	var descriptor GiverStatusDescriptor
	if err = vsn.QueryDescriptors(build, GiverStatusDescriptors, &descriptor); err != nil {
		return err
	}
	var status uint32

	switch statusSize(build) {
	case 1:
		status = uint32(in.ReadByte())
	case 4:
		status = uint32(in.ReadUint32())
	}

	for k, v := range descriptor {
		if v == status {
			r.Status = k
			break
		}
	}

	return
}

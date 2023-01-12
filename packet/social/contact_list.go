package social

import (
	"fmt"

	"github.com/Gophercraft/core/guid"
	"github.com/Gophercraft/core/packet"
	"github.com/Gophercraft/core/vsn"
)

const DefaultContactFlags uint32 = 0x7

type Contact struct {
	Player guid.GUID
	Flags  uint32
	Note   string
	ContactStatus
}

type ContactList struct {
	Flags    uint32
	Contacts []Contact
}

type ContactStatus struct {
	Status Status
	AreaID uint32
	Level  uint32
	Class  uint32
}

func (c *ContactStatus) Decode(build vsn.Build, in *packet.WorldPacket) (err error) {
	c.Status = Status(in.ReadByte())
	if c.Status != StatusOffline {
		c.AreaID = in.ReadUint32()
		c.Level = in.ReadUint32()
		c.Class = in.ReadUint32()
	}

	return
}

func (c *ContactStatus) Encode(build vsn.Build, out *packet.WorldPacket) (err error) {
	out.WriteByte(uint8(c.Status))

	if c.Status != StatusOffline {
		out.WriteUint32(c.AreaID)
		out.WriteUint32(c.Level)
		out.WriteUint32(c.Class)
	}

	return
}

func (cl *ContactList) Encode(build vsn.Build, out *packet.WorldPacket) error {
	out.Type = packet.SMSG_CONTACT_LIST

	out.WriteUint32(cl.Flags)

	out.WriteUint32(uint32(len(cl.Contacts)))

	for _, c := range cl.Contacts {
		c.Player.EncodeUnpacked(build, out)

		out.WriteUint32(c.Flags)

		out.WriteCString(c.Note)

		if c.Flags&FlagFriend != 0 {
			out.WriteByte(uint8(c.Status))

			if c.Status != StatusOffline {
				out.WriteUint32(c.AreaID)
				out.WriteUint32(c.Level)
				out.WriteUint32(c.Class)
			}
		}
	}

	return nil
}

func (cl *ContactList) Decode(build vsn.Build, in *packet.WorldPacket) error {
	cl.Flags = in.ReadUint32()

	count := in.ReadUint32()

	if count >= 1000 {
		return fmt.Errorf("social: too many contacts, damn it: %d", count)
	}

	cl.Contacts = make([]Contact, count)

	var i uint32 = 0

	var err error

	for ; i < count; i++ {
		c := Contact{}
		c.Player, err = guid.DecodeUnpacked(build, in)
		if err != nil {
			return err
		}

		c.Flags = in.ReadUint32()

		c.Note = in.ReadCString()

		if c.Flags&FlagFriend != 0 {
			if err := c.ContactStatus.Decode(build, in); err != nil {
				return err
			}
		}
	}

	return nil
}

type FriendStatus struct {
	Result FriendResult
	ID     guid.GUID
	Note   string
	ContactStatus
}

func (fs *FriendStatus) Encode(build vsn.Build, out *packet.WorldPacket) error {
	out.Type = packet.SMSG_FRIEND_STATUS

	out.WriteByte(uint8(fs.Result))

	fs.ID.EncodeUnpacked(build, out)

	if build.AddedIn(vsn.V2_0_1) {
		switch fs.Result {
		case FriendAddedOnline:
			out.WriteCString(fs.Note)
		case FriendAddedOffline:
			out.WriteCString(fs.Note)
			return nil
		default:
		}
	}

	if err := fs.ContactStatus.Encode(build, out); err != nil {
		return err
	}

	return nil
}

func (fs *FriendStatus) Decode(build vsn.Build, in *packet.WorldPacket) (err error) {
	fs.Result = FriendResult(in.ReadByte())
	fs.ID, err = guid.DecodeUnpacked(build, in)
	if err != nil {
		return
	}

	if build.AddedIn(vsn.V2_0_1) {
		switch fs.Result {
		case FriendAddedOnline:
			fs.Note = in.ReadCString()
		case FriendAddedOffline:
			fs.Note = in.ReadCString()
			return nil
		default:
		}
	}

	err = fs.ContactStatus.Decode(build, in)
	return
}

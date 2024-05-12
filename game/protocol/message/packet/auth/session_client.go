package auth

import (
	"fmt"

	"github.com/Gophercraft/core/game/protocol/message"
	"github.com/Gophercraft/core/packet"
	"github.com/Gophercraft/core/packet/addon"
	"github.com/Gophercraft/core/version"
	"github.com/Gophercraft/log"
)

type SessionClient struct {
	Build           version.Build
	BuildType       int
	LoginServerID   uint32
	Account         string // 0-terminated string
	LoginServerType uint32
	LocalChallenge  []byte
	RegionID        uint32
	BattlegroupID   uint32
	RealmID         uint32
	DosResponse     uint64
	Digest          []byte
	AddonList       *addon.List
	RealmJoinTicket string
	UseIPv6         bool
}

func (c *SessionClient) Decode(build version.Build, in *message.Packet) error {
	if in.Type != packet.CMSG_AUTH_SESSION {
		return fmt.Errorf("auth: invalid opcode %s", in.Type)
	}

	switch {
	case version.Range(0, 3368).Contains(build):
		c.Build = version.Build(in.ReadUint32())
		c.LoginServerID = in.ReadUint32()
		c.Account = in.ReadCString()

		// Alpha sends the session file (wow.ses) in this packet
		// accountSize := in.ReadUint32()
		// if accountSize >= 0xFFFF {
		// 	return fmt.Errorf("packet: invalid account size")
		// }
		// accountData := etc.OfBytes(in.ReadBytes(int(accountSize))).ReadCString()
		log.Dump("Packet", in.Bytes())
		log.Dump("Account data", c.Account)
		// c.Account = accountData
	case version.Range(5875, 6005).Contains(build):
		c.Build = version.Build(in.ReadUint32())
		c.LoginServerID = in.ReadUint32()
		c.Account = in.ReadCString()
		c.LocalChallenge = in.ReadBytes(4)
		c.Digest = in.ReadBytes(20)
	case version.Range(6180, 12340).Contains(build):
		c.Build = version.Build(in.ReadUint32())
		c.LoginServerID = in.ReadUint32()
		c.Account = in.ReadCString()
		c.LoginServerType = in.ReadUint32()
		c.LocalChallenge = in.ReadBytes(4)
		c.RegionID = in.ReadUint32()
		c.BattlegroupID = in.ReadUint32()
		c.RealmID = in.ReadUint32()
		c.DosResponse = in.ReadUint64()
		c.Digest = in.ReadBytes(20)
		var err error
		c.AddonList = &addon.List{}
		err = c.AddonList.Decode(c.Build, in)
		if err != nil {
			return err
		}
		// This range might be incorrect, check WoWPacketParser definitions if you want to add support for other versions than 4.3.4
	case version.Range(13164, 15595).Contains(build):
		c.Build = version.Build(in.ReadUint32())
		c.LoginServerID = in.ReadUint32()
		c.BattlegroupID = in.ReadUint32()
		c.LoginServerType = in.ReadUint32()
		// What even is the point of this?
		c.Digest = make([]byte, 20)
		c.Digest[10] = in.ReadUint8()
		c.Digest[18] = in.ReadUint8()
		c.Digest[12] = in.ReadUint8()
		c.Digest[5] = in.ReadUint8()
		c.DosResponse = in.ReadUint64()
		c.Digest[15] = in.ReadUint8()
		c.Digest[9] = in.ReadUint8()
		c.Digest[19] = in.ReadUint8()
		c.Digest[4] = in.ReadUint8()
		c.Digest[7] = in.ReadUint8()
		c.Digest[16] = in.ReadUint8()
		c.Digest[3] = in.ReadUint8()
		c.Build = version.Build(in.ReadUint32())
		c.Digest[8] = in.ReadUint8()
		c.RealmID = in.ReadUint32()
		c.BuildType = int(in.ReadInt8())
		c.Digest[17] = in.ReadUint8()
		c.Digest[6] = in.ReadUint8()
		c.Digest[0] = in.ReadUint8()
		c.Digest[1] = in.ReadUint8()
		c.Digest[11] = in.ReadUint8()
		c.LocalChallenge = in.ReadBytes(4)
		c.Digest[2] = in.ReadUint8()
		c.RegionID = in.ReadUint32()
		c.Digest[14] = in.ReadUint8()
		c.Digest[13] = in.ReadUint8()
		var err error
		c.AddonList = &addon.List{}
		err = c.AddonList.Decode(c.Build, in)
		if err != nil {
			return err
		}
		c.UseIPv6 = in.ReadMSBit()
		accountNameLength := int(in.ReadUint32())
		if accountNameLength > in.Available() {
			return fmt.Errorf("packet: invalid account name")
		}
		c.Account = in.ReadFixedString(accountNameLength)
	case version.Range(15851, 18414).Contains(build):
		c.LoginServerID = in.ReadUint32()
		c.LoginServerType = in.ReadUint32()
		c.Digest = make([]byte, 20)
		c.Digest[18] = in.ReadUint8()
		c.Digest[14] = in.ReadUint8()
		c.Digest[3] = in.ReadUint8()
		c.Digest[4] = in.ReadUint8()
		c.Digest[0] = in.ReadUint8()
		//guess
		c.BattlegroupID = in.ReadUint32()
		c.Digest[11] = in.ReadUint8()
		c.LocalChallenge = in.ReadBytes(4)
		c.Digest[19] = in.ReadUint8()
		// what is this
		in.Jump(2)
		c.Digest[2] = in.ReadUint8()
		c.Digest[9] = in.ReadUint8()
		c.Digest[12] = in.ReadUint8()
		c.DosResponse = in.ReadUint64()
		// guess
		c.RegionID = in.ReadUint32()
		c.Digest[16] = in.ReadUint8()
		c.Digest[5] = in.ReadUint8()
		c.Digest[6] = in.ReadUint8()
		c.Digest[8] = in.ReadUint8()
		c.Build = version.Build(in.ReadUint16())
		c.Digest[17] = in.ReadUint8()
		c.Digest[7] = in.ReadUint8()
		c.Digest[13] = in.ReadUint8()
		c.Digest[15] = in.ReadUint8()
		c.Digest[1] = in.ReadUint8()
		c.Digest[10] = in.ReadUint8()

		var err error
		c.AddonList = &addon.List{}
		err = c.AddonList.Decode(c.Build, in)
		if err != nil {
			return err
		}

		c.UseIPv6 = in.ReadMSBit()
		accountNameLength := int(in.ReadUint32())
		if accountNameLength > in.Available() {
			return fmt.Errorf("packet: invalid account name")
		}
		c.Account = in.ReadFixedString(accountNameLength)
	case version.Range(19027, 20886).Contains(build):
		c.Build = version.Build(in.ReadUint16())
		c.BuildType = int(in.ReadInt8())
		c.RegionID = in.ReadUint32()
		c.BattlegroupID = in.ReadUint32()
		c.RealmID = in.ReadUint32()
		c.LocalChallenge = in.ReadBytes(4)
		c.Digest = in.ReadBytes(20)
		c.DosResponse = in.ReadUint64()
		accountNameLength := int(in.ReadUint32())
		if accountNameLength > in.Available() {
			return fmt.Errorf("packet: invalid account name")
		}
		c.Account = in.ReadFixedString(accountNameLength)
		c.UseIPv6 = in.ReadMSBit()
		var err error
		c.AddonList = &addon.List{}
		err = c.AddonList.Decode(c.Build, in)
		if err != nil {
			return err
		}
	case version.Range(21355, 21742).Contains(build):
		c.Build = version.Build(in.ReadUint16())
		c.BuildType = int(in.ReadInt8())
		c.RegionID = in.ReadUint32()
		c.BattlegroupID = in.ReadUint32()
		c.RealmID = in.ReadUint32()
		c.LocalChallenge = in.ReadBytes(16)
		c.Digest = in.ReadBytes(24)
		c.DosResponse = in.ReadUint64()
		var err error
		c.AddonList = &addon.List{}
		err = c.AddonList.Decode(c.Build, in)
		if err != nil {
			return err
		}
		realmJoinTicketLength := int(in.ReadUint32())
		if realmJoinTicketLength > in.Available() {
			return fmt.Errorf("packet: invalid realm join ticket")
		}
		c.RealmJoinTicket = in.ReadFixedString(realmJoinTicketLength)
		c.UseIPv6 = in.ReadMSBit()
	case version.Range(22248, 26972).Contains(build):
		c.DosResponse = in.ReadUint64()
		c.Build = version.Build(in.ReadUint16())
		c.BuildType = int(in.ReadInt8())
		c.RegionID = in.ReadUint32()
		c.BattlegroupID = in.ReadUint32()
		c.RealmID = in.ReadUint32()
		c.LocalChallenge = in.ReadBytes(16)
		c.Digest = in.ReadBytes(24)
		c.UseIPv6 = in.ReadMSBit()
		realmJoinTicketLength := int(in.ReadUint32())
		if realmJoinTicketLength > in.Available() {
			return fmt.Errorf("packet: invalid realm join ticket")
		}
		c.RealmJoinTicket = in.ReadFixedString(realmJoinTicketLength)
	case version.Range(28153, version.Max).Contains(build):
		c.DosResponse = in.ReadUint64()
		c.Build = build // build parameter in this packet was removed
		c.RegionID = in.ReadUint32()
		c.BattlegroupID = in.ReadUint32()
		c.RealmID = in.ReadUint32()
		c.LocalChallenge = in.ReadBytes(16)
		c.Digest = in.ReadBytes(24)
		c.UseIPv6 = in.ReadMSBit()
		realmJoinTicketLength := int(in.ReadUint32())
		if realmJoinTicketLength > in.Available() {
			return fmt.Errorf("packet: invalid realm join ticket")
		}
		c.RealmJoinTicket = in.ReadFixedString(realmJoinTicketLength)
	}

	return nil
}

func (c *SessionClient) Encode(build version.Build, out *message.Packet) error {
	out.Type = message.CMSG_AUTH_SESSION
	out.WriteUint32(uint32(c.Build))
	out.WriteUint32(c.LoginServerID)
	out.WriteCString(c.Account)

	if c.Build.RemovedIn(version.V3_3_5a) {
		out.Write(c.LocalChallenge)
		out.Write(c.Digest)
		// out.Write(c.AddonList.Encode())
	} else {
		out.WriteUint32(c.LoginServerType)
		// out.Write(c.Seed)
		out.WriteUint32(c.RegionID)
		out.WriteUint32(c.BattlegroupID)
		out.WriteUint32(c.RealmID)
		out.WriteUint64(c.DosResponse)
		out.Write(c.Digest)
		// out.Write(c.AddonData)
	}

	// Addon data
	return nil
}

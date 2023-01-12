package auth

import (
	"encoding/binary"

	"github.com/Gophercraft/core/home/config"
	"github.com/Gophercraft/core/home/models"
	"github.com/Gophercraft/core/vsn"
	"github.com/superp00t/etc"
)

type RealmList_C struct {
	Request uint32
}

func (rlc *RealmList_C) Type() AuthType {
	return RealmList
}

func (rlc *RealmList_C) Send(build vsn.Build, out *Connection) error {
	buffer := etc.NewBuffer()
	buffer.WriteUint32(rlc.Request)
	return out.SendBuffer(buffer)
}

func (rlc *RealmList_C) Recv(build vsn.Build, conn *Connection) error {
	var data [4]byte
	conn.Read(data[:])
	rlc.Request = binary.LittleEndian.Uint32(data[:])
	return nil
}

type RealmListing struct {
	Type       config.RealmType //
	Locked     bool
	Flags      uint8
	Name       string
	Address    string
	Population float32
	Characters uint8
	Timezone   uint8
	ID         uint8
}

type RealmList_S struct {
	Realms []RealmListing
	Flags  uint16
}

func (rls *RealmList_S) Type() AuthType {
	return RealmList
}

func MakeRealmlist(rlst []models.Realm) *RealmList_S {
	realmState := &RealmList_S{}

	for _, realm := range rlst {
		realmListing := RealmListing{}
		realmListing.Type = realm.Type
		realmListing.Locked = realm.Locked
		if realm.Offline() {
			realmListing.Flags |= 0x02
		}
		realmListing.Name = realm.Name
		realmListing.Address = realm.Address
		realmListing.Population = float32(realm.ActivePlayers/8000) * 3.0
		realmListing.Characters = 0 // todo: get character count from database map
		realmListing.Timezone = uint8(realm.Timezone)
		realmListing.ID = uint8(realm.ID)
		realmState.Realms = append(realmState.Realms, realmListing)
	}

	return realmState
}

func (rlst *RealmList_S) Send(build vsn.Build, out *Connection) error {
	if rlst.Flags == 0 {
		rlst.Flags |= GruntLegacy

		if build.AddedIn(vsn.V2_0_1) {
			rlst.Flags |= GruntImprovedSystem
		}
	}

	listBuffer := etc.NewBuffer()

	if rlst.Flags&GruntImprovedSystem != 0 {
		listBuffer.WriteUint16(uint16(len(rlst.Realms)))
	} else {
		listBuffer.WriteUint32(0)
		listBuffer.WriteByte(uint8(len(rlst.Realms)))
	}

	for _, v := range rlst.Realms {
		if rlst.Flags&GruntImprovedSystem != 0 {
			listBuffer.WriteByte(uint8(v.Type))
			listBuffer.WriteBool(v.Locked)
		} else {
			listBuffer.WriteUint32(uint32(v.Type))
		}

		listBuffer.WriteByte(v.Flags)
		listBuffer.WriteCString(v.Name)
		listBuffer.WriteCString(v.Address)
		listBuffer.WriteFloat32(v.Population)
		listBuffer.WriteByte(v.Characters) // TODO: character count has to be included
		listBuffer.WriteByte(uint8(v.Timezone))
		listBuffer.WriteByte(uint8(v.ID))
	}

	listBuffer.WriteUint16(rlst.Flags)

	head := etc.NewBuffer()
	head.WriteUint16(uint16(listBuffer.Len()))
	head.Write(listBuffer.Bytes())
	return out.SendBuffer(head)
}

func (rls *RealmList_S) Recv(build vsn.Build, conn *Connection) error {
	// realmlist must be size-prefixed in this way to ensure that all realms can be read
	var size uint16
	err := binary.Read(conn, binary.LittleEndian, &size)
	if err != nil {
		return err
	}

	in, err := conn.RecvBuffer(int(size))
	if err != nil {
		return err
	}

	var numrealms uint16

	if build.RemovedIn(vsn.V2_0_1) {
		in.ReadUint32()
		numrealms = uint16(in.ReadByte())
	} else {
		numrealms = in.ReadUint16()
	}

	for x := uint16(0); x < numrealms; x++ {
		rlst := RealmListing{}
		if build < vsn.V2_0_1 {
			rlst.Type = config.RealmType(in.ReadUint32())
		} else {
			rlst.Type = config.RealmType(in.ReadByte())
			rlst.Locked = in.ReadBool()
		}

		rlst.Flags = in.ReadByte()
		rlst.Name = in.ReadCString()
		rlst.Address = in.ReadCString()
		rlst.Population = in.ReadFloat32()
		rlst.Characters = in.ReadByte()
		rlst.Timezone = in.ReadByte()
		rlst.ID = uint8(in.ReadByte())
		rls.Realms = append(rls.Realms, rlst)
	}

	return nil
}

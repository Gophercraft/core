package grunt

import (
	"bytes"
	"encoding/binary"
	"io"

	"github.com/Gophercraft/core/version"
)

type RealmList_Client struct {
	Request uint32
}

type RealmType uint8

type RealmFlags uint8

const (
	RealmVersionMismatch RealmFlags = 0x01
	RealmOffline         RealmFlags = 0x02
	RealmHasBuildInfo    RealmFlags = 0x04
	RealmRecommended     RealmFlags = 0x20
	RealmNew             RealmFlags = 0x40
	RealmFull            RealmFlags = 0x80
)

type RealmBuildInfo struct {
	Major    uint8
	Minor    uint8
	Patch    uint8
	Revision uint16
}

type Realm struct {
	Type          RealmType //
	Locked        bool
	Flags         RealmFlags
	Name          string
	Address       string
	Population    float32
	NumCharacters uint8
	Category      uint8
	Sort          uint8
	BuildInfo     RealmBuildInfo
}

type RealmList_Server struct {
	Unk1   uint32
	Realms []Realm
	Unk2   uint16
}

func WriteRealmList_Client(writer io.Writer, request *RealmList_Client) (err error) {
	var request_bytes [4]byte
	binary.LittleEndian.PutUint32(request_bytes[:], request.Request)
	_, err = writer.Write(request_bytes[:])
	return
}

func ReadRealmList_Client(reader io.Reader, request *RealmList_Client) (err error) {
	var request_bytes [4]byte
	_, err = io.ReadFull(reader, request_bytes[:])
	request.Request = binary.LittleEndian.Uint32(request_bytes[:])

	return
}

const (
	// after build 6178,
	// the size of the realm count was changed from 8 bits to 16 bits
	// also the size of the realm type was changed from 32 bits to 8 bits
	realm_size_changed      = version.Build(6178)
	realm_locked_byte_added = version.Build(6178)
)



func ReadRealmList_Server(reader io.Reader, build version.Build, realm_list *RealmList_Server) (err error) {
	// read list size
	var list_size int
	var list_size_bytes [2]byte
	_, err = io.ReadFull(reader, list_size_bytes[:])
	if err != nil {
		return
	}
	list_size = int(binary.LittleEndian.Uint16(list_size_bytes[:]))

	// read list bytes
	list_bytes := make([]byte, list_size)
	if _, err = io.ReadFull(reader, list_bytes[:]); err != nil {
		return
	}
	list_reader := bytes.NewReader(list_bytes)

	// read padding
	var realm_list_padding_bytes [4]byte
	if _, err = io.ReadFull(list_reader, realm_list_padding_bytes[:]); err != nil {
		return
	}
	realm_list.Unk1 = binary.LittleEndian.Uint32(realm_list_padding_bytes[:])

	// read realm count
	var realm_count int
	if build.AddedIn(realm_size_changed) {
		// after build 6178 realm count is 16 bits
		var realm_count_bytes [2]byte
		if _, err = io.ReadFull(list_reader, realm_count_bytes[:]); err != nil {
			return err
		}
		realm_count = int(binary.LittleEndian.Uint16(realm_count_bytes[:]))
	} else {
		// before 6178 realm count is 8 bits
		var realm_count_byte [1]byte
		if _, err = io.ReadFull(list_reader, realm_count_byte[:]); err != nil {
			return err
		}
		realm_count = int(realm_count_byte[0])
	}

	// read realms
	realm_list.Realms = make([]Realm, realm_count)
	for i := 0; i < realm_count; i++ {
		if err = ReadRealm(list_reader, build, &realm_list.Realms[i]); err != nil {
			return
		}
	}

	// read final padding
	if _, err = io.ReadFull(list_reader, realm_list_padding_bytes[0:2]); err != nil {
		return
	}
	realm_list.Unk2 = binary.LittleEndian.Uint16(realm_list_padding_bytes[0:2])

	return
}

func WriteRealmList_Server(writer io.Writer, build version.Build, realm_list *RealmList_Server) (err error) {
	list_writer := new(bytes.Buffer)

	// write zero uint32 as padding
	var realm_list_padding_bytes [4]byte
	binary.LittleEndian.PutUint32(realm_list_padding_bytes[:], realm_list.Unk1)
	if _, err = list_writer.Write(realm_list_padding_bytes[:]); err != nil {
		return
	}

	realm_count := len(realm_list.Realms)

	// write realm count
	if build.AddedIn(realm_size_changed) {
		// after build 6178 realm count is 16 bits
		var realm_count_bytes [2]byte
		binary.LittleEndian.PutUint16(realm_count_bytes[:], uint16(realm_count))
		if _, err = list_writer.Write(realm_count_bytes[:]); err != nil {
			return err
		}
	} else {
		// before 6178 realm count is 8 bits
		var realm_count_byte [1]byte
		realm_count_byte[0] = byte(realm_count)
		if _, err = list_writer.Write(realm_count_byte[:]); err != nil {
			return err
		}
	}

	for i := 0; i < realm_count; i++ {
		if err = WriteRealm(list_writer, build, &realm_list.Realms[i]); err != nil {
			return
		}
	}

	// write final padding zero uint16
	binary.LittleEndian.PutUint16(realm_list_padding_bytes[0:2], realm_list.Unk2)
	if _, err = list_writer.Write(realm_list_padding_bytes[0:2]); err != nil {
		return
	}

	// write list size
	list_size := list_writer.Len()
	var list_size_bytes [2]byte
	binary.LittleEndian.PutUint16(list_size_bytes[:], uint16(list_size))
	if _, err = writer.Write(list_size_bytes[:]); err != nil {
		return
	}

	// write list
	if _, err = writer.Write(list_writer.Bytes()); err != nil {
		return
	}

	return
}

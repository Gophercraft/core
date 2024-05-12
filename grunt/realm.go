package grunt

import (
	"encoding/binary"
	"fmt"
	"io"
	"math"

	"github.com/Gophercraft/core/serialization"
	"github.com/Gophercraft/core/version"
)

func ReadRealmBuildInfo(reader io.Reader, build_info *RealmBuildInfo) (err error) {
	var build_info_bytes [5]byte
	if _, err = io.ReadFull(reader, build_info_bytes[:]); err != nil {
		return
	}
	build_info.Major = build_info_bytes[0]
	build_info.Minor = build_info_bytes[1]
	build_info.Patch = build_info_bytes[2]
	build_info.Revision = binary.LittleEndian.Uint16(build_info_bytes[3:5])
	return
}

func WriteRealmBuildInfo(writer io.Writer, build_info *RealmBuildInfo) (err error) {
	var build_info_bytes [5]byte
	build_info_bytes[0] = build_info.Major
	build_info_bytes[1] = build_info.Minor
	build_info_bytes[2] = build_info.Patch
	binary.LittleEndian.PutUint16(build_info_bytes[3:5], build_info.Revision)
	_, err = writer.Write(build_info_bytes[:])
	return
}

func ReadRealm(reader io.Reader, build version.Build, realm *Realm) (err error) {
	// read realm type
	if build.AddedIn(realm_size_changed) {
		var realm_type_byte [1]byte
		_, err = io.ReadFull(reader, realm_type_byte[:])
		if err != nil {
			return
		}
		realm.Type = RealmType(realm_type_byte[0])
	} else {
		var realm_type_bytes [4]byte
		_, err = io.ReadFull(reader, realm_type_bytes[:])
		if err != nil {
			return
		}
		realm.Type = RealmType(binary.LittleEndian.Uint32(realm_type_bytes[:]))
	}

	if build.AddedIn(realm_locked_byte_added) {
		// read if realm is locked
		var locked_byte [1]byte
		_, err = io.ReadFull(reader, locked_byte[:])
		if err != nil {
			return
		}
		if locked_byte[0] != 0 && locked_byte[0] != 1 {
			err = fmt.Errorf("locked byte is not a boolean: %d", locked_byte[0])
			return
		}
		realm.Locked = locked_byte[0] == 1
	}

	// read realm flags
	var realm_flags_byte [1]byte
	_, err = io.ReadFull(reader, realm_flags_byte[:])
	if err != nil {
		return
	}
	realm.Flags = RealmFlags(realm_flags_byte[0])

	// read realm name
	realm.Name, err = serialization.ReadCString(reader, 256)
	if err != nil {
		return
	}

	// read realm address host[:port]
	realm.Address, err = serialization.ReadCString(reader, 256)
	if err != nil {
		return
	}

	// Read population float
	var population_bytes [4]byte
	if _, err = io.ReadFull(reader, population_bytes[:]); err != nil {
		return
	}
	realm.Population = math.Float32frombits(binary.LittleEndian.Uint32(population_bytes[:]))

	// Read number of your characters on realm
	var num_characters_byte [1]byte
	if _, err = io.ReadFull(reader, num_characters_byte[:]); err != nil {
		return
	}
	realm.NumCharacters = num_characters_byte[0]

	// Read realm category
	var realm_category_byte [1]byte
	if _, err = io.ReadFull(reader, realm_category_byte[:]); err != nil {
		return
	}
	realm.Category = realm_category_byte[0]

	// read realm sort
	var realm_sort_byte [1]byte
	if _, err = io.ReadFull(reader, realm_sort_byte[:]); err != nil {
		return
	}
	realm.Sort = realm_sort_byte[0]

	if realm.Flags&RealmHasBuildInfo != 0 {
		// read build info
		if err = ReadRealmBuildInfo(reader, &realm.BuildInfo); err != nil {
			return
		}
	}

	return
}

func WriteRealm(writer io.Writer, build version.Build, realm *Realm) (err error) {
	if build.AddedIn(realm_size_changed) {
		// write realm type
		var realm_type_byte [1]byte
		realm_type_byte[0] = byte(realm.Type)
		_, err = writer.Write(realm_type_byte[:])
		if err != nil {
			return
		}
	} else {
		// write realm type
		var realm_type_bytes [4]byte
		binary.LittleEndian.PutUint32(realm_type_bytes[:], uint32(realm.Type))
		_, err = writer.Write(realm_type_bytes[:])
		if err != nil {
			return
		}
	}

	if build.AddedIn(realm_locked_byte_added) {
		// write if realm is locked
		var locked_byte [1]byte
		if realm.Locked {
			locked_byte[0] = 1
		}
		_, err = writer.Write(locked_byte[:])
		if err != nil {
			return
		}
	}

	// write realm flags
	var realm_flags_byte [1]byte
	realm_flags_byte[0] = byte(realm.Flags)
	_, err = writer.Write(realm_flags_byte[:])
	if err != nil {
		return
	}

	// write realm name
	_, err = writer.Write(append([]byte(realm.Name), 0))
	if err != nil {
		return
	}

	// write realm address host[:port]
	_, err = writer.Write(append([]byte(realm.Address), 0))
	if err != nil {
		return
	}

	// write population float
	var population_bytes [4]byte
	binary.LittleEndian.PutUint32(population_bytes[:], math.Float32bits(realm.Population))
	if _, err = writer.Write(population_bytes[:]); err != nil {
		return
	}

	// write number of your characters on realm
	var num_characters_byte [1]byte
	num_characters_byte[0] = realm.NumCharacters
	if _, err = writer.Write(num_characters_byte[:]); err != nil {
		return
	}

	// write realm category
	var realm_category_byte [1]byte
	realm_category_byte[0] = realm.Category
	if _, err = writer.Write(realm_category_byte[:]); err != nil {
		return
	}

	// write realm sort
	var realm_sort_byte [1]byte
	realm_sort_byte[0] = realm.Sort
	if _, err = writer.Write(realm_sort_byte[:]); err != nil {
		return
	}

	if realm.Flags&RealmHasBuildInfo != 0 {
		// write build info
		if err = WriteRealmBuildInfo(writer, &realm.BuildInfo); err != nil {
			return
		}
	}

	return
}

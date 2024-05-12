package guid

import (
	"fmt"
	"io"

	"github.com/Gophercraft/core/version"
)

//go:generate gcraft_stringer -type=TypeID
type TypeID uint8

const (
	TypeObject TypeID = iota
	TypeItem
	TypeContainer
	TypeAzeriteEmpoweredItem
	TypeAzeriteItem
	TypeUnit
	TypePlayer
	TypeActivePlayer
	TypeGameObject
	TypeDynamicObject
	TypeCorpse
	TypeAreaTrigger
	TypeSceneObject
	TypeConversation

	TypeNPCText TypeID = 100
)

type TypeIDDescriptor map[TypeID]uint8

var (
	TypeIDDescriptors = map[version.Build]TypeIDDescriptor{
		version.Alpha: {
			TypeObject:    0,
			TypeItem:      1,
			TypeContainer: 2,
			TypeUnit:      3,
			TypePlayer:    4,

			TypeGameObject:    5,
			TypeDynamicObject: 6,
			TypeCorpse:        7,
		},
	}
)

func init() {
	TypeIDDescriptors[version.V1_12_1] = TypeIDDescriptors[version.Alpha]
	TypeIDDescriptors[version.V2_4_3] = TypeIDDescriptors[version.Alpha]
	TypeIDDescriptors[version.V3_3_5a] = TypeIDDescriptors[version.Alpha]
}

func DecodeTypeID(version version.Build, in io.Reader) (TypeID, error) {
	desc, ok := TypeIDDescriptors[version]
	if !ok {
		return 0, fmt.Errorf("guid: cannot decode type ID for version %d", version)
	}

	var code [1]byte

	_, err := in.Read(code[:])
	if err != nil {
		return 0, err
	}
	resolved := TypeID(0)
	found := false

	for k, v := range desc {
		if v == code[0] {
			found = true
			resolved = k
			break
		}
	}

	if !found {
		return 0, fmt.Errorf("guid: could not resolve type ID for %d", code)
	}

	return resolved, nil
}

func EncodeTypeID(version version.Build, id TypeID, out io.Writer) error {
	desc, ok := TypeIDDescriptors[version]
	if !ok {
		return fmt.Errorf("guid: cannot encode type ID for version %d", version)
	}

	code, ok := desc[id]
	if !ok {
		return fmt.Errorf("guid: cannot resolve code for typeID: %s", id)
	}

	out.Write([]byte{code})
	return nil
}

package packet

import (
	"fmt"

	"github.com/Gophercraft/core/vsn"
)

type WorldCodes map[WorldType]uint32

type WorldTypeDescriptor struct {
	Codes         WorldCodes
	CodeLookupMap map[uint32]WorldType // Generated at startup
}

var WorldTypeDescriptors = map[vsn.BuildRange]*WorldTypeDescriptor{
	{0, 3368}:      worldTypeDescriptor_Alpha,
	{5875, 6141}:   worldTypeDescriptor_5875,
	{6180, 8606}:   worldTypeDescriptor_8606,
	{9056, 12340}:  worldTypeDescriptor_12340,
	{33369, 33369}: worldTypeDescriptor_33369,
}

func init() {
	// Create fast lookup maps for converting a uint32 to a WorldType.
	// TODO: investigate more efficient method of lookup, perhaps with a binary search of a simple table?
	for _, descriptor := range WorldTypeDescriptors {
		descriptor.CodeLookupMap = make(map[uint32]WorldType, len(descriptor.Codes))
		for wType, code := range descriptor.Codes {
			descriptor.CodeLookupMap[code] = wType
		}
	}
}

func (wtd *WorldTypeDescriptor) Code(wt WorldType) (uint32, error) {
	uint32_, ok := wtd.Codes[wt]
	if !ok {
		return 0, fmt.Errorf("packet: no uint found for %s in descriptor", wt)
	}

	return uint32_, nil
}

func (wtd *WorldTypeDescriptor) LookupCode(u32 uint32) (WorldType, error) {
	wt, ok := wtd.CodeLookupMap[u32]
	if !ok {
		return 0, fmt.Errorf("packet: no WorldType found for 0x%04X in lookup table", u32)
	}
	return wt, nil
}

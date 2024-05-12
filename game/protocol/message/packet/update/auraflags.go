package update

import "fmt"

type AuraFlags uint8

const (
	AuraCancelable AuraFlags = 1 << iota
	AuraUnk2
	AuraUnk3
	AuraUnk4
)

const AuraMaskAll AuraFlags = 0x0F

func SetAuraFlag(flagarray *Value, index int, value AuraFlags) error {
	const (
		max byte = byte(AuraMaskAll)
	)

	var (
		val = byte(value)
	)

	byteIndex := index / 2

	// in 4-bit nibbles
	flagsize := flagarray.Len() * 2

	if index >= flagsize {
		return fmt.Errorf("realm: index %d, couldn't possibly index anything, the size of flags in this descriptor is %d", index, flagsize)
	}

	prev := flagarray.Index(byteIndex).Byte()

	var newflag byte

	shift := index%2 == 1

	if shift {
		newflag = (val & max) << 4
	} else {
		newflag = val & max
	}

	var newvalue byte = newflag | prev
	flag := flagarray.Index(byteIndex)

	if flag.Byte() != newvalue {
		flag.SetByte(newvalue)
	}
	return nil
}

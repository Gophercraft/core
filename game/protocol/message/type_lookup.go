package message

import (
	"fmt"
)

func (typedesc *TypeDescriptor) LookupOpcode(t Type) (op uint32, err error) {
	var ok bool
	op, ok = typedesc.opcodes[t]
	if !ok {
		err = fmt.Errorf("game/protocol/message: LookupOpcode: failed to resolve protocol opcode for %s", t)
	}

	return
}

func (typedesc *TypeDescriptor) LookupType(opcode uint32) (t Type, err error) {
	var ok bool
	t, ok = typedesc.types[opcode]
	if !ok {
		err = fmt.Errorf("game/protocol/message: LookupType: failed to resolve message type for code 0x%04X", opcode)
	}

	return
}

package message

// Type is a symbol for a protocol-level opcode.
//go:generate gcraft_stringer -type=Type -fromString
type Type uint16

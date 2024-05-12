package grunt

import (
	"encoding/binary"
	"io"
)

type AuthReconnectProof_Server struct {
	Error           uint8
	LogonProofFlags LogonProofFlags
}

func ReadAuthReconnectProof_Server(reader io.Reader, proof *AuthReconnectProof_Server) (err error) {
	// read error type
	var error_type_byte [1]byte
	if _, err = io.ReadFull(reader, error_type_byte[:]); err != nil {
		return
	}
	proof.Error = error_type_byte[0]
	if proof.Error != 0 {
		// end if not a success
		return
	}

	// read logon proof flags
	var logon_proof_flags_bytes [2]byte
	if _, err = io.ReadFull(reader, logon_proof_flags_bytes[:]); err != nil {
		return
	}
	proof.LogonProofFlags = LogonProofFlags(binary.LittleEndian.Uint16(logon_proof_flags_bytes[:]))

	return
}

func WriteAuthReconnectProof_Server(writer io.Writer, proof *AuthReconnectProof_Server) (err error) {
	// write error type
	var error_type_byte [1]byte
	error_type_byte[0] = byte(proof.Error)
	if _, err = writer.Write(error_type_byte[:]); err != nil {
		return
	}

	if proof.Error != 0 {
		// no more data if not a success
		return
	}

	var logon_proof_flags_bytes [2]byte
	binary.LittleEndian.PutUint16(logon_proof_flags_bytes[:], uint16(proof.LogonProofFlags))
	if _, err = writer.Write(logon_proof_flags_bytes[:]); err != nil {
		return
	}

	return
}

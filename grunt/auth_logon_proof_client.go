package grunt

import (
	"io"
)

type SecurityFlags uint8

type ClientPINInfo struct {
	Salt  [16]byte
	Proof [20]byte
}

type ClientMatrixInfo struct {
	Proof [16]byte
}

type ClientTokenInfo struct {
	Token string
}

type AuthLogonProof_Client struct {
	// Client public key (A)
	ClientPublicKey [32]byte // 32 long
	// Client proof (M1)
	ClientProof [20]byte // 20 long
	// Version proof (CRC)
	VersionProof [20]byte // 20 long
	NumberOfKeys uint8
	LogonFlags   LogonFlags

	PIN    ClientPINInfo
	Matrix ClientMatrixInfo
	Token  ClientTokenInfo
}

func WriteClientPINInfo(writer io.Writer, pin_info *ClientPINInfo) (err error) {
	if _, err = writer.Write(pin_info.Salt[:]); err != nil {
		return
	}
	if _, err = writer.Write(pin_info.Proof[:]); err != nil {
		return
	}
	return
}

func WriteClientMatrixInfo(writer io.Writer, matrix_info *ClientMatrixInfo) (err error) {
	if _, err = writer.Write(matrix_info.Proof[:]); err != nil {
		return
	}
	return
}

func WriteClientTokenInfo(writer io.Writer, token_info *ClientTokenInfo) (err error) {
	var token_length_byte [1]byte
	token_length_byte[0] = byte(len(token_info.Token))
	if _, err = writer.Write(token_length_byte[:]); err != nil {
		return
	}
	if _, err = writer.Write([]byte(token_info.Token)); err != nil {
		return
	}
	return
}

func WriteAuthLogonProof_Client(writer io.Writer, proof *AuthLogonProof_Client) (err error) {
	// Write client public key
	if _, err = writer.Write(proof.ClientPublicKey[:]); err != nil {
		return
	}

	// Write client proof
	if _, err = writer.Write(proof.ClientProof[:]); err != nil {
		return
	}

	// Write version proof
	if _, err = writer.Write(proof.VersionProof[:]); err != nil {
		return
	}

	// Write number of keys
	var number_of_keys_byte [1]byte
	number_of_keys_byte[0] = proof.NumberOfKeys
	if _, err = writer.Write(number_of_keys_byte[:]); err != nil {
		return
	}

	// Write security flags
	var security_flags_byte [1]byte
	security_flags_byte[0] = byte(proof.LogonFlags)
	if _, err = writer.Write(security_flags_byte[:]); err != nil {
		return
	}

	// Write security PIN
	if proof.LogonFlags&PIN != 0 {
		if err = WriteClientPINInfo(writer, &proof.PIN); err != nil {
			return
		}
	}

	// Write matrix info
	if proof.LogonFlags&Matrix != 0 {
		if err = WriteClientMatrixInfo(writer, &proof.Matrix); err != nil {
			return
		}
	}

	// Write security token
	if proof.LogonFlags&Token != 0 {
		if err = WriteClientTokenInfo(writer, &proof.Token); err != nil {
			return
		}
	}

	return
}

func ReadClientPINInfo(reader io.Reader, pin_info *ClientPINInfo) (err error) {
	if _, err = io.ReadFull(reader, pin_info.Salt[:]); err != nil {
		return
	}
	if _, err = io.ReadFull(reader, pin_info.Proof[:]); err != nil {
		return
	}
	return
}

func ReadClientMatrixInfo(reader io.Reader, matrix_info *ClientMatrixInfo) (err error) {
	if _, err = io.ReadFull(reader, matrix_info.Proof[:]); err != nil {
		return
	}
	return
}

func ReadClientTokenInfo(reader io.Reader, token_info *ClientTokenInfo) (err error) {
	var token_length_byte [1]byte
	if _, err = io.ReadFull(reader, token_length_byte[:]); err != nil {
		return
	}
	token_bytes := make([]byte, token_length_byte[0])
	if _, err = io.ReadFull(reader, token_bytes[:]); err != nil {
		return
	}
	token_info.Token = string(token_bytes)
	return
}

func ReadAuthLogonProof_Client(reader io.Reader, proof *AuthLogonProof_Client) (err error) {
	// Read client public key
	if _, err = io.ReadFull(reader, proof.ClientPublicKey[:]); err != nil {
		return
	}

	// Read client proof
	if _, err = io.ReadFull(reader, proof.ClientProof[:]); err != nil {
		return
	}

	// Read version proof
	if _, err = io.ReadFull(reader, proof.VersionProof[:]); err != nil {
		return
	}

	// Read number of keys
	var number_of_keys_byte [1]byte
	if _, err = io.ReadFull(reader, number_of_keys_byte[:]); err != nil {
		return
	}
	proof.NumberOfKeys = number_of_keys_byte[0]

	// Read security flags
	var security_flags_byte [1]byte
	if _, err = io.ReadFull(reader, security_flags_byte[:]); err != nil {
		return
	}
	proof.LogonFlags = LogonFlags(security_flags_byte[0])

	// Read security PIN
	if proof.LogonFlags&PIN != 0 {
		if err = ReadClientPINInfo(reader, &proof.PIN); err != nil {
			return
		}
	}

	// Read security matrix
	if proof.LogonFlags&Matrix != 0 {
		if err = ReadClientMatrixInfo(reader, &proof.Matrix); err != nil {
			return
		}
	}

	// Read security token
	if proof.LogonFlags&Token != 0 {
		if err = ReadClientTokenInfo(reader, &proof.Token); err != nil {
			return
		}
	}

	return
}

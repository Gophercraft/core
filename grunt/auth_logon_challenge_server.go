package grunt

import (
	"encoding/binary"
	"io"
)

type LogonFlags uint8

const (
	PIN LogonFlags = 1 << iota
	Matrix
	Token
)

type ServerPINInfo struct {
	GridSeed uint32
	Salt     [16]byte
}

type ServerMatrixInfo struct {
	Width          uint8
	Height         uint8
	DigitCount     uint8
	ChallengeCount uint8
	Seed           uint64
}

type ServerTokenInfo struct {
	Required uint8
}

type AuthLogonChallenge_Server struct {
	Error       uint8
	LoginResult LoginResult
	// Server public key (B)
	ServerPublicKey [32]byte // 32 long
	// Generator (G)
	Generator []byte
	// Large safe prime (N)
	LargeSafePrime []byte // 32 long
	// Salt (S)
	Salt             [32]byte // 32 long
	VersionChallenge [16]byte // 16 long
	LogonFlags       LogonFlags
	PIN              ServerPINInfo
	Matrix           ServerMatrixInfo
	Token            ServerTokenInfo
}

func ReadServerPINInfo(reader io.Reader, pin_info *ServerPINInfo) (err error) {
	// read grid seed
	var grid_seed_bytes [4]byte
	if _, err = io.ReadFull(reader, grid_seed_bytes[:]); err != nil {
		return
	}
	pin_info.GridSeed = binary.LittleEndian.Uint32(grid_seed_bytes[:])

	// read salt
	if _, err = io.ReadFull(reader, pin_info.Salt[:]); err != nil {
		return
	}

	return
}

func ReadServerMatrixInfo(reader io.Reader, matrix_info *ServerMatrixInfo) (err error) {
	// read matrix info
	var matrix_info_bytes [12]byte
	if _, err = io.ReadFull(reader, matrix_info_bytes[:]); err != nil {
		return
	}
	matrix_info.Width = matrix_info_bytes[0]
	matrix_info.Height = matrix_info_bytes[1]
	matrix_info.DigitCount = matrix_info_bytes[2]
	matrix_info.ChallengeCount = matrix_info_bytes[3]
	matrix_info.Seed = binary.LittleEndian.Uint64(matrix_info_bytes[4:12])
	return
}

func ReadServerTokenInfo(reader io.Reader, token_info *ServerTokenInfo) (err error) {
	// read token info
	var token_info_bytes [1]byte
	if _, err = io.ReadFull(reader, token_info_bytes[:]); err != nil {
		return
	}
	token_info.Required = token_info_bytes[0]
	return
}

func ReadAuthLogonChallenge_Server(reader io.Reader, challenge *AuthLogonChallenge_Server) (err error) {
	// Read error code
	var error_byte [1]byte
	if _, err = io.ReadFull(reader, error_byte[:]); err != nil {
		return
	}
	challenge.Error = error_byte[0]

	// Read result
	var result_byte [1]byte
	if _, err = io.ReadFull(reader, result_byte[:]); err != nil {
		return
	}
	challenge.LoginResult = LoginResult(error_byte[0])

	if challenge.LoginResult != LoginOk {
		return
	}

	// Read server public key
	if _, err = io.ReadFull(reader, challenge.ServerPublicKey[:]); err != nil {
		return
	}

	// Read generator
	var G_size_byte [1]byte
	if _, err = io.ReadFull(reader, G_size_byte[:]); err != nil {
		return
	}
	G_size := int(G_size_byte[0])
	challenge.Generator = make([]byte, G_size)
	if _, err = io.ReadFull(reader, challenge.Generator); err != nil {
		return
	}

	// Read large safe prime
	var N_size_byte [1]byte
	if _, err = io.ReadFull(reader, N_size_byte[:]); err != nil {
		return
	}
	N_size := int(N_size_byte[0])
	challenge.LargeSafePrime = make([]byte, N_size)
	if _, err = io.ReadFull(reader, challenge.LargeSafePrime); err != nil {
		return
	}

	// Read salt
	if _, err = io.ReadFull(reader, challenge.Salt[:]); err != nil {
		return
	}

	// Read version challenge
	if _, err = io.ReadFull(reader, challenge.VersionChallenge[:]); err != nil {
		return
	}

	// Read logon flags
	var security_flag_byte [1]byte
	if _, err = io.ReadFull(reader, security_flag_byte[:]); err != nil {
		return
	}
	challenge.LogonFlags = LogonFlags(security_flag_byte[0])

	// Read optional info
	if challenge.LogonFlags&PIN != 0 {
		if err = ReadServerPINInfo(reader, &challenge.PIN); err != nil {
			return
		}
	}

	if challenge.LogonFlags&Matrix != 0 {
		if err = ReadServerMatrixInfo(reader, &challenge.Matrix); err != nil {
			return
		}
	}

	if challenge.LogonFlags&Token != 0 {
		if err = ReadServerTokenInfo(reader, &challenge.Token); err != nil {
			return
		}
	}

	return
}

func WriteServerPINInfo(writer io.Writer, pin_info *ServerPINInfo) (err error) {
	// write grid seed
	var grid_seed_bytes [4]byte
	binary.LittleEndian.PutUint32(grid_seed_bytes[:], pin_info.GridSeed)
	if _, err = writer.Write(grid_seed_bytes[:]); err != nil {
		return
	}

	// write salt
	if _, err = writer.Write(pin_info.Salt[:]); err != nil {
		return
	}

	return
}

func WriteServerMatrixInfo(writer io.Writer, matrix_info *ServerMatrixInfo) (err error) {
	// write matrix info
	var matrix_info_bytes [12]byte
	matrix_info_bytes[0] = matrix_info.Width
	matrix_info_bytes[1] = matrix_info.Height
	matrix_info_bytes[2] = matrix_info.DigitCount
	matrix_info_bytes[3] = matrix_info.ChallengeCount
	binary.LittleEndian.PutUint64(matrix_info_bytes[4:12], matrix_info.Seed)

	_, err = writer.Write(matrix_info_bytes[:])
	return
}

func WriteServerTokenInfo(writer io.Writer, token_info *ServerTokenInfo) (err error) {
	// write token info
	var token_info_bytes [1]byte
	token_info_bytes[0] = token_info.Required
	_, err = writer.Write(token_info_bytes[:])
	return
}

func WriteAuthLogonChallenge_Server(writer io.Writer, challenge *AuthLogonChallenge_Server) (err error) {
	// write error
	var error_byte [1]byte
	error_byte[0] = byte(challenge.Error)
	if _, err = writer.Write(error_byte[:]); err != nil {
		return
	}

	// write login result
	var login_result_byte [1]byte
	login_result_byte[0] = byte(challenge.LoginResult)
	if _, err = writer.Write(login_result_byte[:]); err != nil {
		return
	}

	if challenge.LoginResult != LoginOk {
		// no content
		return
	}

	// Write server public key
	if _, err = writer.Write(challenge.ServerPublicKey[:]); err != nil {
		return
	}

	// write generator
	var generator_len_byte [1]byte
	generator_len_byte[0] = byte(len(challenge.Generator))
	if _, err = writer.Write(generator_len_byte[:]); err != nil {
		return
	}
	if _, err = writer.Write(challenge.Generator); err != nil {
		return
	}

	// write large safe prime
	var large_safe_prime_len_byte [1]byte
	large_safe_prime_len_byte[0] = byte(len(challenge.LargeSafePrime))
	if _, err = writer.Write(large_safe_prime_len_byte[:]); err != nil {
		return
	}
	if _, err = writer.Write(challenge.LargeSafePrime); err != nil {
		return
	}

	// write salt
	if _, err = writer.Write(challenge.Salt[:]); err != nil {
		return
	}

	// write version challenge
	if _, err = writer.Write(challenge.VersionChallenge[:]); err != nil {
		return
	}

	// write logon flags
	var logon_flag_byte [1]byte
	logon_flag_byte[0] = byte(challenge.LogonFlags)
	if _, err = writer.Write(logon_flag_byte[:]); err != nil {
		return
	}

	// Write optional info
	if challenge.LogonFlags&PIN != 0 {
		if err = WriteServerPINInfo(writer, &challenge.PIN); err != nil {
			return
		}
	}

	if challenge.LogonFlags&Matrix != 0 {
		if err = WriteServerMatrixInfo(writer, &challenge.Matrix); err != nil {
			return
		}
	}

	if challenge.LogonFlags&Token != 0 {
		if err = WriteServerTokenInfo(writer, &challenge.Token); err != nil {
			return
		}
	}

	return
}

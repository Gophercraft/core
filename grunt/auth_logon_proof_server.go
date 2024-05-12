package grunt

import (
	"encoding/binary"
	"fmt"
	"io"

	"github.com/Gophercraft/core/version"
)

type AccountFlags uint32

type LogonProofFlags uint16

const (
	AccountMessageAvailable LogonProofFlags = 1 << iota
)

type AuthLogonProof_Server struct {
	LoginResult LoginResult
	Unk1        uint16
	// Server proof (M2)
	ServerProof     [20]byte
	AccountFlags    AccountFlags
	SurveyID        uint32
	LogonProofFlags LogonProofFlags
}

const (
	logon_proof_flags_added = version.Build(6178)
	survey_ID_added         = version.Build(8089)
)

func WriteAuthLogonProof_Server(writer io.Writer, build version.Build, proof *AuthLogonProof_Server) (err error) {
	// Write login result
	var login_result_byte [1]byte
	login_result_byte[0] = byte(proof.LoginResult)
	if _, err = writer.Write(login_result_byte[:]); err != nil {
		return
	}

	if proof.LoginResult != LoginOk {
		if proof.LoginResult == LoginUnknownAccount {
			if build.AddedIn(version.V2_0_1) {
				// two bytes
				var unk [2]byte
				binary.LittleEndian.PutUint16(unk[:], proof.Unk1)
				if _, err = writer.Write(unk[:]); err != nil {
					return
				}
			}
		}

		return
	}

	// Write server proof
	if _, err = writer.Write(proof.ServerProof[:]); err != nil {
		return
	}

	// Write account flags
	var account_flags_bytes [4]byte
	binary.LittleEndian.PutUint32(account_flags_bytes[:], uint32(proof.AccountFlags))
	if _, err = writer.Write(account_flags_bytes[:]); err != nil {
		return
	}

	if build.AddedIn(survey_ID_added) {
		// Write survey id
		var survey_ID_bytes [4]byte
		binary.LittleEndian.PutUint32(survey_ID_bytes[:], uint32(proof.SurveyID))
		if _, err = writer.Write(survey_ID_bytes[:]); err != nil {
			return
		}
	}

	if build.AddedIn(logon_proof_flags_added) {
		// Write logon proof flags
		var logon_proof_flags_bytes [2]byte
		binary.LittleEndian.PutUint16(logon_proof_flags_bytes[:], uint16(proof.LogonProofFlags))
		if _, err = writer.Write(logon_proof_flags_bytes[:]); err != nil {
			return
		}
	}

	return
}

func ReadAuthLogonProof_Server(reader io.Reader, build version.Build, proof *AuthLogonProof_Server) (err error) {
	// read login result
	var login_result_byte [1]byte
	if _, err = io.ReadFull(reader, login_result_byte[:]); err != nil {
		fmt.Println("error reading login result", err)
		return
	}
	proof.LoginResult = LoginResult(login_result_byte[0])

	if proof.LoginResult != LoginOk {
		if proof.LoginResult == LoginUnknownAccount {
			if build.AddedIn(version.V2_0_1) {
				// two bytes
				var unk [2]byte
				if _, err = io.ReadFull(reader, unk[:]); err != nil {
					return
				}
				proof.Unk1 = binary.LittleEndian.Uint16(unk[:])
			}
		}

		return
	}

	// read server proof
	if _, err = io.ReadFull(reader, proof.ServerProof[:]); err != nil {
		return
	}

	// read account flags
	var account_flags_bytes [4]byte
	if _, err = io.ReadFull(reader, account_flags_bytes[:]); err != nil {
		return
	}
	proof.AccountFlags = AccountFlags(binary.LittleEndian.Uint32(account_flags_bytes[:]))

	if build.AddedIn(survey_ID_added) {
		// read survey id
		var survey_ID_bytes [4]byte
		if _, err = io.ReadFull(reader, survey_ID_bytes[:]); err != nil {
			return
		}
		proof.SurveyID = binary.LittleEndian.Uint32(survey_ID_bytes[:])
	}

	if build.AddedIn(logon_proof_flags_added) {
		// Write logon proof flags
		var logon_proof_flags_bytes [2]byte
		if _, err = io.ReadFull(reader, logon_proof_flags_bytes[:]); err != nil {
			return
		}
		proof.LogonProofFlags = LogonProofFlags(binary.LittleEndian.Uint16(logon_proof_flags_bytes[:]))
	}

	return
}

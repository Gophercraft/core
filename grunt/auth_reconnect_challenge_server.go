package grunt

import "io"

type AuthReconnectChallenge_Server struct {
	Error            uint8
	ServerProof      [16]byte
	VersionChallenge [16]byte
}

func ReadAuthReconnectChallenge_Server(reader io.Reader, challenge *AuthReconnectChallenge_Server) (err error) {
	// read error type
	var error_type_byte [1]byte
	if _, err = io.ReadFull(reader, error_type_byte[:]); err != nil {
		return
	}
	challenge.Error = error_type_byte[0]
	if challenge.Error != 0 {
		return
	}

	// read server proof
	if _, err = io.ReadFull(reader, challenge.ServerProof[:]); err != nil {
		return
	}

	// read version challenge
	if _, err = io.ReadFull(reader, challenge.VersionChallenge[:]); err != nil {
		return
	}

	return
}

func WriteAuthReconnectChallenge_Server(writer io.Writer, challenge *AuthReconnectChallenge_Server) (err error) {
	// write error type
	var error_type_byte [1]byte
	error_type_byte[0] = byte(challenge.Error)
	if _, err = writer.Write(error_type_byte[:]); err != nil {
		return
	}

	if challenge.Error != 0 {
		return
	}

	// write server proof
	if _, err = writer.Write(challenge.ServerProof[:]); err != nil {
		return
	}

	// write version challenge
	if _, err = writer.Write(challenge.VersionChallenge[:]); err != nil {
		return
	}

	return
}

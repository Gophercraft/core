package grunt

import "io"

type AuthReconnectProof_Client struct {
	R1           [16]byte
	R2           [20]byte
	R3           [20]byte
	NumberOfKeys uint8
}

func ReadAuthReconnectProof_Client(reader io.Reader, proof *AuthReconnectProof_Client) (err error) {
	// read R1
	if _, err = io.ReadFull(reader, proof.R1[:]); err != nil {
		return
	}

	// read R2
	if _, err = io.ReadFull(reader, proof.R2[:]); err != nil {
		return
	}

	// read R3
	if _, err = io.ReadFull(reader, proof.R3[:]); err != nil {
		return
	}

	// Read number of keys
	var number_of_keys_byte [1]byte
	if _, err = io.ReadFull(reader, number_of_keys_byte[:]); err != nil {
		return
	}
	proof.NumberOfKeys = number_of_keys_byte[0]

	return
}

func WriteAuthReconnectProof_Client(writer io.Writer, proof *AuthReconnectProof_Client) (err error) {
	// write R1
	if _, err = writer.Write(proof.R1[:]); err != nil {
		return
	}

	// write R2
	if _, err = writer.Write(proof.R2[:]); err != nil {
		return
	}

	// write R3
	if _, err = writer.Write(proof.R3[:]); err != nil {
		return
	}

	// Write number of keys
	var number_of_keys_byte [1]byte
	number_of_keys_byte[0] = proof.NumberOfKeys
	if _, err = writer.Write(number_of_keys_byte[:]); err != nil {
		return
	}

	return
}

package grunt

import (
	"crypto/rand"
	"crypto/sha1"
	"encoding/binary"
	"io"
)

// Creates a lookup table to be scrambled according to a seed
func GeneratePINKeyMap(seed uint32) (key_map []byte) {
	digits := [10]byte{
		0,
		1,
		2,
		3,
		4,
		5,
		6,
		7,
		8,
		9,
	}

	key_map = make([]byte, 10)

	n := uint32(10)

	i := 0

	grid_seed := seed

	for {
		n_div := grid_seed / n
		n_mod := grid_seed % n

		key_map[i] = digits[n_mod]

		copy_len := (n - n_mod) - 1

		copy(digits[n_mod:n_mod+copy_len], digits[n_mod+1:n_mod+1+copy_len])

		i = i + 1
		n = n - 1

		grid_seed = n_div

		if n == 0 {
			break
		}
	}

	return
}

func ScramblePINNumber(grid_seed uint32, pin_digits []byte) (scrambled_pin_digits []byte) {
	//
	key_map := GeneratePINKeyMap(grid_seed)

	scrambled_pin_digits = make([]byte, len(pin_digits))

	for digit_index, digit := range pin_digits {
		number := digit - '0'
		for key_index, key := range key_map {
			if key == number {
				scrambled_pin_digits[digit_index] = '0' + byte(key_index)
				break
			}
		}
	}
	return
}

func random_pin_grid_seed() uint32 {
	var seed [4]byte
	if _, err := io.ReadFull(rand.Reader, seed[:]); err != nil {
		panic(err)
	}
	return binary.LittleEndian.Uint32(seed[:])
}

// func ScramblePIN(digits string, key_map [10]byte)

// pin digits are in ascii
func GetPINProof(server_salt []byte, pin_digits []byte, client_salt []byte) (client_proof []byte) {
	hash := sha1.New()
	hash.Write(server_salt[:16])
	hash.Write(pin_digits)
	digits_digest := hash.Sum(nil)

	hash = sha1.New()
	hash.Write(client_salt[:16])
	hash.Write(digits_digest[:20])
	client_proof = hash.Sum(nil)

	return
}

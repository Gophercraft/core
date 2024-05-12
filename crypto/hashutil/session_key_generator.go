package hashutil

import (
	"hash"
)

type SessionKeyGenerator struct {
	Hash  func() hash.Hash
	Seeds [3][]byte
}

func NewSessionKeyGenerator(hash_function func() hash.Hash, data []byte) *SessionKeyGenerator {
	size := len(data)
	halfSize := size / 2
	key_generator := &SessionKeyGenerator{
		Hash: hash_function,
	}

	hash1 := H(key_generator.Hash, data[0:halfSize])
	hash0 := make([]byte, len(hash1))
	hash2 := H(key_generator.Hash, data[halfSize:])
	hash0 = H(key_generator.Hash, hash1, hash0, hash2)

	key_generator.Seeds[0] = hash0
	key_generator.Seeds[1] = hash1
	key_generator.Seeds[2] = hash2

	return key_generator
}

func (key_generator *SessionKeyGenerator) Read(data []byte) (int, error) {
	seedIndex := 0

	for i := range data {
		// Progress to next hash
		if seedIndex == len(key_generator.Seeds[0]) {
			key_generator.Seeds[0] = H(key_generator.Hash, key_generator.Seeds[1], key_generator.Seeds[0], key_generator.Seeds[2])
			seedIndex = 0
		}

		data[i] = key_generator.Seeds[0][seedIndex]

		seedIndex++
	}

	return len(data), nil
}

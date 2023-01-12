package warden

import (
	"crypto/sha1"
	"encoding/binary"
	"sync"

	"github.com/Gophercraft/core/crypto"
	"github.com/Gophercraft/core/crypto/arc4"
)

func Checksum(data []byte) uint32 {
	hash := crypto.Hash(sha1.New, data)

	ints := make([]uint32, 5)
	for i := 0; i < 5; i++ {
		ints[i] = binary.LittleEndian.Uint32(hash[i*4 : (i+1)*4])
	}

	var sum uint32
	for i := 0; i < 5; i++ {
		sum = sum ^ ints[i]
	}

	return sum
}

type Session struct {
	sync.Mutex
	Server              bool
	Key                 []byte
	InputKey, OutputKey []byte
	Input, Output       *arc4.ARC4
}

func NewSession(sessionKey []byte, server bool) *Session {
	s := new(Session)
	s.Server = server
	s.Key = sessionKey
	skg := crypto.NewSessionKeyGenerator(sha1.New, sessionKey)
	s.InputKey = make([]byte, 16)
	s.OutputKey = make([]byte, 16)
	if s.Server {
		skg.Read(s.InputKey[:])
		skg.Read(s.OutputKey[:])
	} else {
		skg.Read(s.OutputKey[:])
		skg.Read(s.InputKey[:])
	}

	s.Input = arc4.New(s.InputKey)
	s.Output = arc4.New(s.OutputKey)
	return s
}

func (s *Session) Decrypt(data []byte) {
	s.Input.Decrypt(data)
}

func (s *Session) Encrypt(data []byte) {
	s.Output.Encrypt(data)
}

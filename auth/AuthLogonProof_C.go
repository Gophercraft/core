package auth

import (
	"github.com/Gophercraft/core/crypto"
	"github.com/Gophercraft/core/vsn"
	"github.com/Gophercraft/log"
	"github.com/superp00t/etc"
)

type AuthLogonProof_C struct {
	A            []byte // 32 long
	M1           []byte // 20 long
	CRC          []byte // 20 long
	NumberOfKeys uint8
	SecFlags     uint8
}

func (alpc *AuthLogonProof_C) Type() AuthType {
	return LogonProof
}

// Client sends BE (Big Endian)
// Server reinterpret_casts struct, converting it to LE in C++
// Server converts it back to BE with SetBinary
func (alpc *AuthLogonProof_C) Send(build vsn.Build, conn *Connection) error {
	buf := etc.NewBuffer()
	buf.Write(alpc.A)
	buf.Write(alpc.M1)
	if len(alpc.CRC) == 0 {
		buf.Write(crypto.RandBytes(20))
	} else {
		if len(alpc.CRC) != 20 {
			panic("invalid CRC length")
		}

		buf.Write(alpc.CRC)
	}
	buf.WriteByte(alpc.NumberOfKeys)
	buf.WriteByte(alpc.SecFlags)
	return conn.SendBuffer(buf)
}

func (alpc *AuthLogonProof_C) Recv(build vsn.Build, conn *Connection) error {
	in, err := conn.RecvBuffer(74)
	if err != nil {
		return nil
	}
	input := in.Bytes()
	alpc.A = input[0:32]
	alpc.M1 = input[32:52]
	alpc.CRC = input[52:72]
	alpc.NumberOfKeys = input[72]
	alpc.SecFlags = input[73]
	log.Dump("alpc", alpc)
	return nil
}

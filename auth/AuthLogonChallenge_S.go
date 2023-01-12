package auth

import (
	"github.com/superp00t/etc"
	"github.com/Gophercraft/core/vsn"
)

type AuthLogonChallenge_S struct {
	Error            ErrorType
	B                []byte // 32 long
	G                []byte
	N                []byte // 32 long
	S                []byte // 32 long
	VersionChallenge []byte // 16 long
	SecurityFlags    uint8
	// Unk4  uint8
}

func (acls *AuthLogonChallenge_S) Type() AuthType {
	return LogonChallenge
}

func (acls *AuthLogonChallenge_S) Send(build vsn.Build, conn *Connection) error {
	buf := etc.NewBuffer()
	buf.WriteByte(0x00)
	buf.WriteByte(uint8(acls.Error))

	if acls.Error != GruntSuccess {
		return conn.SendBuffer(buf)
	}

	buf.Write(acls.B)
	// G
	buf.WriteByte(uint8(len(acls.G)))
	buf.Write(acls.G)

	// N
	buf.WriteByte(uint8(len(acls.N)))
	buf.Write(acls.N)
	buf.Write(acls.S)
	buf.Write(acls.VersionChallenge)
	buf.WriteByte(acls.SecurityFlags)
	return conn.SendBuffer(buf)
}

func (acls *AuthLogonChallenge_S) Recv(build vsn.Build, conn *Connection) error {
	alcs := &AuthLogonChallenge_S{}
	in, err := conn.RecvBuffer(1 + 1 + 32 + 1)
	if err != nil {
		return err
	}
	in.ReadByte() // Always zero
	alcs.Error = ErrorType(in.ReadByte())
	alcs.B = in.ReadBytes(32)
	gLen := in.ReadByte()
	gDat, err := conn.RecvBuffer(int(gLen))
	if err != nil {
		return err
	}
	alcs.G = gDat.Bytes()
	var nLen [1]byte
	conn.Read(nLen[:])
	nDat, err := conn.RecvBuffer(int(nLen[0]))
	if err != nil {
		return err
	}
	alcs.N = nDat.Bytes()
	in, err = conn.RecvBuffer(32 + 16 + 1)
	if err != nil {
		return err
	}
	alcs.S = in.ReadBytes(32)
	alcs.VersionChallenge = in.ReadBytes(16)
	alcs.SecurityFlags = in.ReadByte()
	return nil
}

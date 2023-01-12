package auth

import (
	"github.com/Gophercraft/core/vsn"
	"github.com/superp00t/etc"
)

type AuthLogonProof_S struct {
	Error        ErrorType
	M2           []byte
	AccountFlags uint32
	SurveyID     uint32
	Unk3         uint16
}

func (proof *AuthLogonProof_S) Type() AuthType {
	return LogonProof
}

func (proof *AuthLogonProof_S) Send(build vsn.Build, conn *Connection) error {
	buf := etc.NewBuffer()
	buf.WriteByte(uint8(proof.Error))

	if proof.Error != GruntSuccess {
		if build >= vsn.V2_0_1 {
			buf.WriteUint16(0)
		}
		return conn.SendBuffer(buf)
	}

	buf.Write(proof.M2)
	buf.WriteUint32(proof.AccountFlags)

	if build.AddedIn(vsn.V2_0_1) {
		buf.WriteUint32(proof.SurveyID)
		buf.WriteUint16(proof.Unk3)
	}

	return conn.SendBuffer(buf)
}

func (proof *AuthLogonProof_S) Recv(build vsn.Build, conn *Connection) error {
	sizeIn := 1 + 20 + 4
	if build >= vsn.V2_0_1 {
		sizeIn += 6
	}

	in, err := conn.RecvBuffer(sizeIn)
	if err != nil {
		return err
	}

	proof.Error = ErrorType(in.ReadByte())
	proof.M2 = in.ReadBytes(20)
	proof.AccountFlags = in.ReadUint32()

	if build >= vsn.V2_0_1 {
		proof.SurveyID = in.ReadUint32()
		proof.Unk3 = in.ReadUint16()
	}

	return nil
}

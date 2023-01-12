package auth

import (
	"github.com/Gophercraft/core/vsn"
	"github.com/superp00t/etc"
)

type ReconnectChallenge_C struct {
	Result LoginResult

	ChallengeData []byte
}

func (rc *ReconnectChallenge_C) Type() AuthType {
	return ReconnectChallenge
}

func (rc *ReconnectChallenge_C) Send(build vsn.Build, out *Connection) error {
	msg := etc.NewBuffer()
	msg.WriteByte(uint8(rc.Result))
	var challengeData [16]byte
	copy(challengeData[:], rc.ChallengeData[:16])
	msg.Write(challengeData[:])
	var checksumSalt [16]byte
	if build <= vsn.V1_12_2 {
		msg.Write(checksumSalt[:])
	}
	return out.SendBuffer(msg)
}

func (rc *ReconnectChallenge_C) Recv(build vsn.Build, in *Connection) error {
	msgSize := 1 + 16
	if build <= vsn.V1_12_2 {
		msgSize += 16
	}

	msg, err := in.RecvBuffer(msgSize)
	if err != nil {
		return err
	}

	rc.Result = LoginResult(msg.ReadByte())
	rc.ChallengeData = make([]byte, 16)
	msg.Read(rc.ChallengeData)

	return nil
}

package grunt

import "io"

//go:generate stringer -type=MessageType
type MessageType uint8

const (
	LogonChallenge     MessageType = 0x00
	LogonProof         MessageType = 0x01
	ReconnectChallenge MessageType = 0x02
	ReconnectProof     MessageType = 0x03
	RealmList          MessageType = 0x10
	XferInitiate       MessageType = 0x30
	XferData           MessageType = 0x31
	XferAccept         MessageType = 0x32
	XferResume         MessageType = 0x33
	XferCancel         MessageType = 0x34
)

const MaxMessageType = 0x35

func WriteMessageType(writer io.Writer, message_type MessageType) (err error) {
	var mt [1]byte
	mt[0] = byte(message_type)
	_, err = writer.Write(mt[:])
	return
}

func ReadMessageType(reader io.Reader) (message_type MessageType, err error) {
	var mt [1]byte
	_, err = io.ReadFull(reader, mt[:])
	message_type = MessageType(mt[0])
	return
}

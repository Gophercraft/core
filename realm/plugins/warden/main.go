package warden

import "github.com/Gophercraft/core/realm"

func WardenMain(s *realm.Session, done <-chan bool) {
	sdp := s.FindProcess(wardenMain)
	if sdp == nil {
		panic("wardenMain called")
		return
	}

	sd := sdp.Data.(*SessionData)

	for {
		select {
		case <-done:
			return
		case crypto := <-sd.NewCrypto:
			sd.Crypto = crypto
		case wardenData := <-sd.WardenPacketQueue:
			sd.Crypto.Output.Encrypt(wardenData.Buffer.Bytes())
			sd.Session.SendPacket(wardenData)
		}
	}
}

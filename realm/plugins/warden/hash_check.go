package warden

import (
	"bytes"

	"github.com/Gophercraft/core/crypto/arc4"
	packetwarden "github.com/Gophercraft/core/packet/warden"
)

func (sd *SessionData) HandleHashResult(hr *packetwarden.ClientHashResult) {
	if !bytes.Equal(hr.Response, sd.Module.ClientKeySeedHash) {
		sd.Fail(FailHash, "bad hash result")
		return
	}

	sd.Wardenf(WarnInformative, "Correct Warden module verified")
	sd.Crypto.Lock()
	sd.Crypto.Input = arc4.New(sd.Module.ClientKeySeed)
	sd.Crypto.Output = arc4.New(sd.Module.ServerKeySeed)
	sd.Crypto.Unlock()
}

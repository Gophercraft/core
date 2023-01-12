package warden

import (
	"fmt"
	"log"

	packetwarden "github.com/Gophercraft/core/packet/warden"
)

type Module struct {
	Hash              []byte
	Module            []byte
	ModuleKey         []byte
	Seed              []byte
	ServerKeySeed     []byte
	ClientKeySeed     []byte
	ClientKeySeedHash []byte
}

func (sd *SessionData) TransferModule() error {
	const chunkSize = 500

	log.Println("Beginning transfer of Warden module.")
	pack := new(packetwarden.ServerModuleTransfer)
	pack.Data = make([]byte, chunkSize)
	sizeLeft := len(sd.Module.Module)
	pos := 0
	burstSize := 0

	for {
		if sizeLeft == 0 {
			break
		}

		if sizeLeft < chunkSize {
			burstSize = sizeLeft
		} else {
			burstSize = chunkSize
		}

		sizeLeft -= burstSize
		pos += burstSize

		sd.SendWardenData(&packetwarden.ServerData{
			Requests: []packetwarden.ServerRequest{
				&packetwarden.ServerModuleTransfer{
					sd.Module.Module[pos : pos+burstSize],
				},
			},
		})
		fmt.Printf("Warden transfer %d %d, %d/%d\n", sizeLeft, burstSize, pos, len(sd.Module.Module))
	}

	return nil
}

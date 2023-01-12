package warden

import (
	"time"

	"github.com/Gophercraft/core/crypto/warden"
	"github.com/Gophercraft/core/packet"
	packetwarden "github.com/Gophercraft/core/packet/warden"
	"github.com/Gophercraft/core/realm"
)

const sessionProcessName = "gophercraft.realm.plugins.warden.WardenMain"

type SessionData struct {
	Session           *realm.Session
	Crypto            warden.Session
	newCrypto         chan *warden.Session
	Module            *Module
	CurrentChecks     *packetwarden.ServerRequestCheatChecks
	WardenPacketQueue chan *packet.WorldPacket
	ServerTicks       uint32
}

// Called after user passes wait list
func InitWardenSession(s *realm.Session) {
	w := &SessionData{}
	w.Session = s
	w.Crypto = *warden.NewSession(s.SessionKey, true)
	// TODO: research: Explore alternative modules
	w.Module = Module_79C0768D657977D697E10BAD956CCED1
	w.WardenPacketQueue = make(chan *packet.WorldPacket, 8)

	mu := &packetwarden.ServerModuleUse{}
	copy(mu.ModuleID[:], w.Module.Hash)
	copy(mu.ModuleKey[:], w.Module.ModuleKey)
	mu.Size = uint32(len(w.Module.Module))

	req := packetwarden.ServerData{
		Requests: []packetwarden.ServerRequest{
			mu,
		},
	}

	w.SendWardenData(&req)

	proc := s.CreateProcess(sessionProcessName, WardenMain)
	proc.Data = w
}

func (sd *SessionData) SetCurrentChecks(req *packetwarden.ServerRequestCheatChecks) {
	sd.CurrentChecks = req
}

func (sd *SessionData) HandleTimingCheck(result *packetwarden.ClientCheckResult) {
	ms := uint32(result.Time.UnixNano() / int64(time.Millisecond))
	ourTicks := result.NewClientTicks + (ms - sd.ServerTicks)
	sd.Wardenf("client took %d", ourTicks)
}

func (sd *SessionData) SendWardenData(req *packetwarden.ServerData) {
	wardenData := packet.NewWorldPacket(packet.SMSG_WARDEN_DATA)
	if err := req.Encode(sd.Session.Build(), &packetwarden.Writer{
		packetwarden.CryptoData{
			ClientKey: sd.Crypto.InputKey,
			ServerKey: sd.Crypto.OutputKey,
		},
		wardenData,
	}); err != nil {
		panic(err)
	}
	sd.WardenPacketQueue <- wardenData
}

func HandleWardenData(s *realm.Session, wPacket *packet.WorldPacket) {
	wardenProcess := s.FindProcess(wardenMain)
	if wardenProcess == nil {
		return
	}
	sd := wardenProcess.Data.(*SessionData)
	sd.Crypto.Input.Decrypt(wPacket.Buffer.Bytes())

	clientData := &packetwarden.ClientData{}
	clientData.Decode(sd.Session.Build(), &packetwarden.Reader{
		CryptoData: packetwarden.CryptoData{
			ClientKey: sd.Crypto.InputKey,
			ServerKey: sd.Crypto.OutputKey,
		},
		Reader: wPacket,
	})

	switch clientData.Result.Command() {
	case packetwarden.CClientModuleMissing:
		sd.TransferModule()
	case packetwarden.CClientHashResult:
		sd.HandleHashResult(hr)
	case packetwarden.CClientCheatChecksResult:
		cc := clientData.Result.(*packetwarden.ClientCheatChecksResult)
		sd.HandleClientCheatChecksResult(cc)
	}
}

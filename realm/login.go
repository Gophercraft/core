package realm

import (
	"time"

	"github.com/Gophercraft/core/packet"
	"github.com/Gophercraft/core/packet/character"
	"github.com/Gophercraft/core/packet/login"
	"github.com/Gophercraft/core/realm/wdb/models"
	"github.com/Gophercraft/core/tempest"
	"github.com/Gophercraft/core/vsn"
	"github.com/Gophercraft/log"
)

func (s *Session) SendLoginFailure(failure character.LoginResult) {
	s.Send(&character.LoginStatus{
		Type:   packet.SMSG_CHARACTER_LOGIN_FAILED,
		Result: failure,
	})
}

func (s *Session) HandlePlayerLogin(join *login.Player) {
	if s.HasState(InWorld) {
		return
	}

	// todo: handle player already in world

	log.Println("Player join requested", join.Character)

	if sess, _ := s.Server.GetSessionByGUID(join.Character); sess != nil {
		s.SendLoginFailure(character.LoginDuplicateCharacter)
		return
	}

	var chr models.Character

	found, err := s.DB().Where("game_account = ?", s.GameAccount).Where("id = ?", join.Character.Counter()).Get(&chr)
	if err != nil {
		panic(err)
	}

	if found {
		s.InitPlayerSession(&chr)
		log.Println("GUID found for character", chr.Name, join.Character)
		s.SetupOnLogin()
		return
	}

	// Todo handle unknown GUID
	s.SendLoginFailure(character.LoginNoCharacter)
}

func (s *Session) HandleLogoutRequest() {
	if !s.HasState(InWorld) {
		return
	}

	// TODO: deny if in combat
	// TODO: Impose timeout if configured

	if s.Server.LogoutTime() != 0 {
		var resp login.LogoutResponse
		resp.Status = login.LogoutOK
		resp.Instant = false
		s.Send(&resp)

		s.CreateProcess(LogoutRequestPtag, func(s *Session, cancel <-chan bool) {
			for {
				select {
				case <-time.After(s.Server.LogoutTime()):
					go s.Logout()
				case <-cancel:
					return
				}
			}
		})
		return
	} else {
		if s.Build().AddedIn(vsn.V1_12_1) {
			var resp login.LogoutResponse
			resp.Status = login.LogoutOK
			resp.Instant = true
			s.Send(&resp)
		}

		s.Logout()
	}

	// // s.SendObjectDelete(s.GUID())

	// s.CleanupPlayer()2

	// s.SetState(CharacterSelectMenu)

	// resp := packet.NewWorldPacket(packet.SMSG_LOGOUT_COMPLETE)
	// s.SendPacket(resp)
}

func (s *Session) Logout() {
	if !s.HasState(InWorld) {
		return
	}

	resp := packet.NewWorldPacket(packet.SMSG_LOGOUT_COMPLETE)
	s.SendPacket(resp)

	log.Println("Logging out")

	s.CleanupPlayer()

	s.SetState(CharacterSelectMenu)

	// s.SendObjectDelete(s.GUID())

	log.Println("player cleaned up")
}

func (s *Server) LogoutTime() time.Duration {
	// return 5 * time.Second
	return 0
}

const LogoutRequestPtag = "Gophercraft/core/realm/login#logging out"

func (s *Session) HandleLogoutCancel() {
	// response := packet.NewWorldPacket(packet.SMSG_LOGOUT_CANCEL_ACK)
	// s.SendPacket(response)

	// time.Sleep(5 * time.Second)

	s.KillProcessesWithTag(LogoutRequestPtag)

	// s.CreateProcess(LogoutRequestPtag, func(s *Session, cancel <-chan bool) {
	// 	select {
	// 	case <-time.After(s.Server.LogoutTime()):
	// 	case <-cancel:
	// 		return
	// 	}
	// })
}

func (s *Session) SendVerifyLoginPacket() {
	s.Send(&login.VerifyWorld{
		MapID: s.Char.Map,
		Position: tempest.C4Vector{
			s.Char.X,
			s.Char.Y,
			s.Char.Z,
			s.Char.O,
		},
	})

	log.Println("Sent verify login packet")
}

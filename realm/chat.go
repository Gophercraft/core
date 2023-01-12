package realm

import (
	"fmt"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/Gophercraft/core/guid"
	"github.com/Gophercraft/core/packet"
	"github.com/Gophercraft/core/realm/wdb/models"

	"github.com/Gophercraft/core/home/rpcnet"

	"github.com/Gophercraft/core/packet/chat"
	"github.com/Gophercraft/log"
)

func (s *Session) SystemChat(data string) {
	lines := strings.Split(data, "\n")

	for _, ln := range lines {
		s.sysChatLine(ln)
	}
}

func (s *Session) sysChatLine(ln string) {
	s.Send(&chat.ServerMessage{
		Type:     chat.MsgSystem,
		Language: chat.LangUniversal,
		Body:     ln,
	})
}

func (s *Session) Sysf(data string, args ...interface{}) {
	s.SystemChat(fmt.Sprintf(data, args...))
}

// prints something in green text with [SERVER] prepended
// for use in global announcements
func (s *Session) Annf(data string, args ...interface{}) {
	s.SystemChat("|cFF00FF00[SERVER] " + fmt.Sprintf(data, args...) + "|r")
}

func (s *Session) ColorPrintf(color string, data string, args ...interface{}) {
	printed := fmt.Sprintf(data, args...)
	if len(printed) < 255 {
		// We can send this as one packet.
		s.sysChatLine("|c" + color + printed + "|r")
		return
	}

	lines := strings.Split(printed, "\n")

	for _, ln := range lines {
		s.sysChatLine("|c" + color + ln + "|r")
	}
}

func (s *Session) Warnf(data string, args ...interface{}) {
	s.ColorPrintf("FFFFFF00", data, args...)
}

func (s *Session) Kv(name, format string, args ...any) {
	const (
		keyColor = "FF00FF00"
		valColor = DemoColor
	)

	line := fmt.Sprintf(format, args...)

	s.Sysf("|c%s%s:|r |c%s%s|r",
		keyColor,
		name,

		valColor,
		line,
	)
}

func (s *Session) printfObjMgr(data string, args ...interface{}) {
	msg := fmt.Sprintf(data, args...)
	s.Sysf("|cFFD97438[Object Manager]|r %s", msg)
	log.DefaultLogger.LogLine(&log.Line{
		Time:     time.Now(),
		Category: "object",
		Text:     msg,
	})
}

func (s *Session) NoSuchPlayer(playerName string) {
	s.Warnf("The player '%s' could not be found.", playerName)
}

func (s *Session) PlayerName() string {
	return s.Char.Name
}

func (s *Session) Tag() chat.Tag {
	if s.PlayerSession != nil {
		switch s.GameMode() {
		case models.GameMode_God:
			return chat.TagGM
		}
	}

	return chat.TagNone
}

func (s *Session) IsGM() bool {
	return s.Tier >= rpcnet.Tier_GameMaster
}

func (s *Session) IsAdmin() bool {
	return s.Tier >= rpcnet.Tier_Admin
}

func (s *Session) HandleChat(cm *chat.ClientMessage) {
	if len(cm.Body) > 255 {
		return
	}

	if cm.Body == "" {
		return
	}

	if !utf8.ValidString(cm.Body) {
		return
	}

	if strings.HasPrefix(cm.Body, ".") && len(cm.Body) > 1 {
		s.HandleCommand(cm.Body)
		return
	}

	ss := s.Server.Call(ChatEvent, cm.Type, cm)
	if ss {
		return
	}

	switch cm.Type {
	case chat.MsgWhisper:
		s.HandleWhisper(cm)
	case chat.MsgSay:
		s.HandleSay(cm)
	case chat.MsgParty:
		s.HandlePartyMessage(cm)
	}
}

func (s *Session) IsEnemy(wo WorldObject) bool {
	return false
}

func (s *Session) IsIgnoring(player guid.GUID) bool {
	return false
}

func (s *Session) HandleWhisper(whisper *chat.ClientMessage) {
	target := whisper.To
	targetSession, err := s.Server.GetSessionByPlayerName(target)
	if err != nil {
		s.Send(&chat.Error{
			Type: packet.SMSG_CHAT_PLAYER_NOT_FOUND,
			Name: target,
		})
		return
	}

	if s.Server.BoolVar("Chat.LanguageBarrier") {
		if s.IsEnemy(targetSession) {
			s.Send(&chat.Error{
				Type: packet.SMSG_CHAT_WRONG_FACTION,
				Name: target,
			})
			return
		}
	}

	if targetSession.IsIgnoring(s.GUID()) {
		s.Send(&chat.Error{
			Type: packet.SMSG_CHAT_IGNORED_ACCOUNT_MUTED,
			Name: target,
		})
		return
	}

	s.Send(&chat.ServerMessage{
		Type:       chat.MsgWhisperInform,
		SenderGUID: s.GUID(),
		TargetGUID: targetSession.GUID(),
		Body:       whisper.Body,
	})

	targetSession.Send(&chat.ServerMessage{
		Type:       chat.MsgWhisper,
		SenderGUID: s.GUID(),
		Body:       whisper.Body,
	})
}

type ChatHistory struct {
	RecentMessage []*chat.ClientMessage
	// Simili
	SimilitaryToPrevious float32 // must be in range 0.0 - 1.0
}

func (s *Session) String() string {
	sessionIdent := "unknown"
	playerInfo := ""

	switch {
	case s.HasState(Authed):
		sessionIdent = fmt.Sprintf("Account: %d, GameAccount: %d", s.Account, s.GameAccount)
	case s.State() == Handshaking:
		sessionIdent = "Unconnected"
	}

	if s.HasState(InWorld) {
		playerInfo = fmt.Sprintf(" currently in world, currently playing as character %s \"%s\"", s.GUID(), s.PlayerName())
	}

	return fmt.Sprintf("<Session %s%s>", sessionIdent, playerInfo)
}

func (s *Session) IsSpam(cm *chat.ClientMessage) bool {
	if len(cm.Body) >= 255 {
		return true
	}

	if cm.Language == chat.LangAddon {
		return false
	}

	return false
}

func (s *Session) HandleSay(say *chat.ClientMessage) {
	// if s.IsSpam(say) {
	// 	s.CommitMessageLog()
	// 	return
	// }

	var lang chat.Language = chat.LangUniversal
	if !s.Server.BoolVar("Chat.LanguageBarrier") {
		// TODO: use the existing structure

		pck := &chat.ServerMessage{
			Type:       chat.MsgSay,
			Language:   lang,
			Name:       s.PlayerName(),
			SenderGUID: s.GUID(),
			Body:       say.Body,
			Tag:        s.Tag(),
		}

		s.SendArea(pck)
		return
	}

	// if s.Team().Relation()
}

package auth

import (
	"sync/atomic"

	packet_auth "github.com/Gophercraft/core/game/protocol/message/packet/auth"
	pb_auth "github.com/Gophercraft/core/home/protocol/pb/auth"
	"github.com/Gophercraft/core/i18n"
)

type SessionState uint8

const (
	Ended SessionState = 1 << iota
	Authed
	PassedWaitQueue
)

type SessionContext struct {
	challenge    *packet_auth.Challenge
	state        atomic.Uint32
	tier         pb_auth.AccountTier
	locale       i18n.Locale
	account      uint64
	game_account uint64
	session_key  []byte
}

func (context *SessionContext) HasState(ss SessionState) bool {
	return context.state.Load()&uint32(ss) != 0
}

func (context *SessionContext) SetState(ss SessionState) {
	context.state.Store(context.state.Load() | uint32(ss))
}

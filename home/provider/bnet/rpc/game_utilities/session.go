package game_utilities

import (
	"context"

	"github.com/Gophercraft/core/version"
)

type context_key uint8

const (
	session_context_key context_key = iota
)

type Session struct {
	Build        version.Build
	GameAccount  uint64
	ClientSecret []byte
}

func GetSession(ctx context.Context) (session *Session) {
	value := ctx.Value(session_context_key)
	if value == nil {
		return
	}
	session = value.(*Session)
	return
}

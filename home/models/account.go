package models

import (
	"github.com/Gophercraft/core/home/rpcnet"
	"github.com/Gophercraft/core/i18n"
)

// Account represents the authentication account
// This account tracks what privileges you are allowed, and what your password is.
type Account struct {
	ID       uint64      `xorm:"'id' pk autoincr"`
	Tier     rpcnet.Tier `xorm:"'tier'"`
	Locale   i18n.Locale `xorm:"'locale'"`
	Platform string      `xorm:"'platform'"`
	Username string      `xorm:"'username'"`
	// SRP Identity hash
	IdentityHash []byte `xorm:"'identity_hash'"`
	// Bcrypt hash - should be used when SRP IdentityHash is not required
	WebLoginHash []byte `xorm:"'web_login_hash'"`
	Locked       bool   `xorm:"'locked'"`
}

// GameAccount represents the the game data account.
// GameAccounts track what character lists are tied to your account.
// (Legacy servers) By changing your account to Active, you can select a different character list upon logging in
type GameAccount struct {
	ID     uint64 `xorm:"'id' pk autoincr"`
	Name   string `xorm:"'name'"`
	Active bool   `xorm:"'active'"`
	Owner  uint64 `xorm:"'owner'"`
}

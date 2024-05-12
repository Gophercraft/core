package models

import "time"

// GameAccount represents the the game data account.
// GameAccounts track what character lists are tied to your account.
// (Legacy servers) By changing your account to Active, you can select a different character list upon logging in
type GameAccount struct {
	ID     uint64 `database:"1:auto_increment,index,exclusive"`
	Name   string
	Active bool
	Owner  uint64
	// Suspended = temporarily banned
	Suspended bool
	// Banned = permanently banned
	Banned      bool
	UnsuspendAt time.Time
}

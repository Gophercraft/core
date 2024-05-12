package models

import "time"

type CharacterCount struct {
	GameAccountID uint64 `database:"1"`
	RealmID       uint64 `database:"2"`
	Count         uint32 `database:"3"`
}

type LastCharacterLoggedIn struct {
	GameAccountID  uint64    `database:"1"`
	RealmID        uint64    `database:"2"`
	CharacterID    uint64    `database:"3"`
	CharacterName  string    `database:"4"`
	LastPlayedTime time.Time `database:"5"`
}

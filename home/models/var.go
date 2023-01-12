package models

type CVar struct {
	RealmID uint64 `xorm:"'server_id'"`
	Key     string `xorm:"'key'"`
	Value   string
}

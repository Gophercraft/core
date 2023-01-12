package models

type Map struct {
	ID           uint32
	Directory    string
	InstanceType uint32
	MapType      uint32
	Name         string
	MinLevel     uint32
	MaxLevel     uint32
	MaxPlayers   uint32
	Descriptions []string
}

package main

type PlayerCreateActionButton struct {
	Race   uint8
	Class  uint8
	Button uint8
	Action uint32
	Type   uint8
}

func (PlayerCreateActionButton) TableName() string {
	return "Playercreateinfo_action"
}

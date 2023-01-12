package social

type Status uint8

const (
	StatusOffline Status = 0x00
	StatusOnline  Status = 0x01
	StatusAFK     Status = 0x02
	StatusDND     Status = 0x04
	StatusRAF     Status = 0x08
)

const (
	FlagFriend         = 0x01
	FlagIgnored        = 0x02
	FlagMuted          = 0x04
	FlagRecruitAFriend = 0x08
)

type FriendResult uint8

const (
	FriendDBError         = 0x00
	FriendListFull        = 0x01
	FriendOnline          = 0x02
	FriendOffline         = 0x03
	FriendNotFound        = 0x04
	FriendRemoved         = 0x05
	FriendAddedOnline     = 0x06
	FriendAddedOffline    = 0x07
	FriendAlready         = 0x08
	FriendSelf            = 0x09
	FriendEnemy           = 0x0A
	FriendIgnoreFull      = 0x0B
	FriendIgnoreSelf      = 0x0C
	FriendIgnoreNotFound  = 0x0D
	FriendIgnoreAlready   = 0x0E
	FriendIgnoreAdded     = 0x0F
	FriendIgnoreRemoved   = 0x10
	FriendIgnoreAmbiguous = 0x11 // That name is ambiguous, type more of the player's server name
	FriendMuteFull        = 0x12
	FriendMuteSelf        = 0x13
	FriendMuteNotFound    = 0x14
	FriendMuteAlready     = 0x15
	FriendMuteAdded       = 0x16
	FriendMuteRemoved     = 0x17
	FriendMuteAmbiguous   = 0x18 // That name is ambiguous, type more of the player's server name
	FriendUnk1            = 0x19 // no message at client
	FriendUnk2            = 0x1A
	FriendUnk3            = 0x1B
	FriendUnknown         = 0x1C // Unknown friend response from server
)

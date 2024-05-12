package game_utilities

import "fmt"

type RealmAddress uint32

func NewRealmAddress(region, site uint8, ID uint16) (addr RealmAddress) {
	addr |= RealmAddress(region) << 24
	addr |= RealmAddress(site) << 16
	addr |= RealmAddress(ID) & 0xFFFF
	return
}

func (addr RealmAddress) String() string {
	return fmt.Sprintf("%d-%d-%d", addr.Region(), addr.Site(), addr.ID())
}

func (addr RealmAddress) ID() uint16 {
	return uint16(addr & 0xFFFF)
}

func (addr RealmAddress) Region() uint8 {
	return uint8((addr >> 24) & 0xFF)
}

func (addr RealmAddress) Site() uint8 {
	return uint8((addr >> 16) & 0xFF)
}

func set_realm_address(field **uint32, address RealmAddress) {
	var addr = uint32(address)
	*field = &addr
}

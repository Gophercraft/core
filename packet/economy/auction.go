package economy

import (
	"fmt"

	"github.com/Gophercraft/core/guid"
	"github.com/Gophercraft/core/packet"
	"github.com/Gophercraft/core/realm/wdb/models"
	"github.com/Gophercraft/core/vsn"
)

type AuctionListItemsRequest struct {
	Auctioneer          guid.GUID
	ListFrom            uint32
	SearchedName        string
	LevelMin            uint8
	LevelMax            uint8
	AuctionSlotID       uint32
	AuctionMainCategory uint32
	AuctionSubCategory  uint32
	Quality             uint32
	Usable              uint8
}

type AuctionListing struct {
	ID                   uint32
	Entry                uint32
	EnchantmentID        uint32
	ItemRandomPropertyID uint32
	ItemSuffixFactor     uint32
	Owner                guid.GUID
	StartBid             models.Money
	OutBid               models.Money
	ExpireTime           uint32
	CurrentBidder        guid.GUID
	CurrentBid           models.Money
}

func (al *AuctionListing) Encode(version vsn.Build, e *packet.WorldPacket) error {
	e.WriteUint32(al.ID)
	e.WriteUint32(al.Entry)
	e.WriteUint32(al.EnchantmentID)
	e.WriteUint32(al.ItemRandomPropertyID)
	e.WriteUint32(al.ItemSuffixFactor)
	al.Owner.EncodeUnpacked(version, e)
	e.WriteInt32(int32(al.StartBid))
	e.WriteInt32(int32(al.OutBid))
	e.WriteUint32(al.ExpireTime)
	al.CurrentBidder.EncodeUnpacked(version, e)
	e.WriteInt32(al.CurrentBid.Int32())
	return nil
}

func (al *AuctionListing) Decode(build vsn.Build, e *packet.WorldPacket) error {
	al.ID = e.ReadUint32()
	al.Entry = e.ReadUint32()
	al.EnchantmentID = e.ReadUint32()
	al.ItemRandomPropertyID = e.ReadUint32()
	al.ItemSuffixFactor = e.ReadUint32()
	var err error
	al.Owner, err = guid.DecodeUnpacked(build, e)
	if err != nil {
		return err
	}
	al.StartBid = models.Money(e.ReadInt32())
	al.OutBid = models.Money(e.ReadInt32())
	al.ExpireTime = e.ReadUint32()
	al.CurrentBidder, err = guid.DecodeUnpacked(build, e)
	if err != nil {
		return err
	}

	al.CurrentBid = models.Money(e.ReadInt32())
	return nil
}

type AuctionPage struct {
	Listings   []AuctionListing
	TotalCount uint32
}

func (alir *AuctionPage) Decode(build vsn.Build, in *packet.WorldPacket) error {
	const maxPage = 100

	count := in.ReadUint32()

	if count > maxPage {
		return fmt.Errorf("server sent more auction house listings than expected")
	}

	alir.Listings = make([]AuctionListing, count)

	for l := uint32(0); l < count; l++ {
		listing := &alir.Listings[l]

		if err := listing.Decode(build, in); err != nil {
			return err
		}
	}

	alir.TotalCount = in.ReadUint32()
	return nil
}

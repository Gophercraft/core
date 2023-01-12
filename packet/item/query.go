package item

import (
	"fmt"

	"github.com/Gophercraft/core/guid"
	"github.com/Gophercraft/core/i18n"
	"github.com/Gophercraft/core/packet"
	"github.com/Gophercraft/core/packet/query"
	"github.com/Gophercraft/core/packet/update"
	"github.com/Gophercraft/core/realm/wdb/models"
	"github.com/Gophercraft/core/vsn"
)

type QuerySingle struct {
	ID   uint32
	GUID guid.GUID
}

func (iq *QuerySingle) Encode(build vsn.Build, out *packet.WorldPacket) error {
	out.Type = packet.CMSG_ITEM_QUERY_SINGLE
	out.WriteUint32(iq.ID)
	iq.GUID.EncodeUnpacked(build, out)
	return nil
}

func (iq *QuerySingle) Decode(build vsn.Build, in *packet.WorldPacket) error {
	iq.ID = in.ReadUint32()
	if iq.ID >= query.EntryNotFound {
		return fmt.Errorf("packet: client sent item query ID larger than query.EntryNotFound flag (%d, hex 0x%016x). Unless the client is requesting gibberish, this should not happen without a ludicrous amount of item templates installed on the server", iq.ID, iq.ID)
	}
	var err error
	iq.GUID, err = guid.DecodeUnpacked(build, in)
	return err
}

func DmgCount(build vsn.Build) int {
	if build.AddedIn(9767) {
		return 2
	}
	return 5
}

type ResponseSingle struct {
	QueryID uint32
	// Locale should be set by client and server. If not, English will be used as the default!
	Locale i18n.Locale
	Item   *models.ItemTemplate
}

func (ir *ResponseSingle) Encode(build vsn.Build, out *packet.WorldPacket) error {
	out.Type = packet.SMSG_ITEM_QUERY_SINGLE_RESPONSE

	qID := ir.QueryID
	if qID >= query.EntryNotFound {
		return fmt.Errorf("packet: Item query response ID is too large. this should not happen ever")
	}

	if ir.Item == nil {
		qID |= query.EntryNotFound
	}

	out.WriteUint32(qID)

	if qID&query.EntryNotFound != 0 {
		//lol
		return nil
	}

	// out.WriteUint32(ir.Item.Entry)
	out.WriteUint32(ir.Item.Class)
	out.WriteUint32(ir.Item.Subclass)
	if build.AddedIn(vsn.V2_0_1) {
		out.WriteUint32(ir.Item.SoundOverrideSubclass)
	}
	out.WriteCString(ir.Item.Name.GetLocalized(ir.Locale))
	out.WriteCString("")
	out.WriteCString("")
	out.WriteCString("")
	out.WriteUint32(ir.Item.DisplayID)
	out.WriteUint32(uint32(ir.Item.Quality))

	err := ir.Item.Flags.Encode(build, out)
	if err != nil {
		panic(err)
	}

	out.WriteUint32(uint32(ir.Item.BuyPrice))
	out.WriteUint32(uint32(ir.Item.SellPrice))
	out.WriteUint32(uint32(ir.Item.InventoryType))
	out.WriteInt32(int32(ir.Item.AllowableClass))
	out.WriteInt32(int32(ir.Item.AllowableRace))
	out.WriteUint32(ir.Item.ItemLevel)
	out.WriteUint32(uint32(ir.Item.RequiredLevel))
	out.WriteUint32(ir.Item.RequiredSkill) // id from SkillLine.dbc
	out.WriteUint32(ir.Item.RequiredSkillRank)

	if build.AddedIn(vsn.V1_12_1) {
		out.WriteUint32(ir.Item.RequiredSpell) // id from Spell.dbc
		out.WriteUint32(ir.Item.RequiredHonorRank)
		out.WriteUint32(ir.Item.RequiredCityRank)
		out.WriteUint32(ir.Item.RequiredReputationFaction) // id from Faction.dbc
		out.WriteUint32(ir.Item.RequiredReputationRank)
	}

	out.WriteUint32(ir.Item.MaxCount)
	out.WriteUint32(ir.Item.Stackable)
	out.WriteUint32(uint32(ir.Item.ContainerSlots))

	for x := 0; x < 10; x++ {
		if x >= len(ir.Item.Stats) {
			out.WriteUint32(0)
			out.WriteInt32(0)
		} else {
			code, err := ir.Item.Stats[x].Type.Uint32(build)
			if err != nil {
				return err
			}
			out.WriteUint32(uint32(code))
			out.WriteInt32(ir.Item.Stats[x].Value)
		}
	}

	for x := 0; x < DmgCount(build); x++ {
		if x >= len(ir.Item.Damage) {
			out.WriteFloat32(0)
			out.WriteFloat32(0)
			out.WriteUint32(0)
		} else {
			out.WriteFloat32(ir.Item.Damage[x].Min)
			out.WriteFloat32(ir.Item.Damage[x].Max)
			out.WriteUint32(uint32(ir.Item.Damage[x].Type))
		}
	}

	out.WriteUint32(ir.Item.Armor)
	out.WriteUint32(ir.Item.HolyRes)
	out.WriteUint32(ir.Item.FireRes)
	out.WriteUint32(ir.Item.NatureRes)
	out.WriteUint32(ir.Item.FrostRes)
	out.WriteUint32(ir.Item.ShadowRes)

	if build.AddedIn(vsn.V1_12_1) {
		out.WriteUint32(ir.Item.ArcaneRes)
	}

	out.WriteUint32(ir.Item.Delay)
	out.WriteUint32(ir.Item.AmmoType)
	out.WriteFloat32(ir.Item.RangedModRange)

	for x := 0; x < 5; x++ {
		if x >= len(ir.Item.Spells) {
			out.WriteUint32(0)
			out.WriteUint32(0)
			out.WriteUint32(0)
			out.WriteInt32(-1)
			out.WriteUint32(0)
			out.WriteInt32(-1)
		} else {
			out.WriteUint32(ir.Item.Spells[x].ID)
			out.WriteUint32(ir.Item.Spells[x].Trigger)
			out.WriteInt32(ir.Item.Spells[x].Charges)
			out.WriteInt32(int32(ir.Item.Spells[x].Cooldown))
			out.WriteUint32(ir.Item.Spells[x].Category)
			out.WriteInt32(int32(ir.Item.Spells[x].CategoryCooldown))
		}
	}

	out.WriteUint32(uint32(ir.Item.Bonding))
	out.WriteCString(ir.Item.Description.GetLocalized(ir.Locale))
	out.WriteUint32(ir.Item.PageText)
	out.WriteUint32(ir.Item.LanguageID)
	out.WriteUint32(ir.Item.PageMaterial)
	out.WriteUint32(ir.Item.StartQuest)
	out.WriteUint32(ir.Item.LockID)
	out.WriteInt32(ir.Item.Material)
	out.WriteUint32(ir.Item.Sheath)

	if build.AddedIn(vsn.V1_12_1) {
		out.WriteUint32(ir.Item.RandomProperty)
		if build.AddedIn(vsn.V2_0_1) {
			out.WriteUint32(ir.Item.RandomSuffix)
		}
		out.WriteUint32(ir.Item.Block)
		out.WriteUint32(ir.Item.Itemset)
		out.WriteUint32(ir.Item.MaxDurability)
		out.WriteUint32(ir.Item.Area)
		out.WriteInt32(ir.Item.Map)
		out.WriteInt32(ir.Item.BagFamily)

		if build.AddedIn(vsn.V2_0_1) {
			out.WriteInt32(ir.Item.TotemCategory)
			for socket := 0; socket < 3; socket++ {
				if socket >= len(ir.Item.Socket) {
					out.WriteInt32(-1)
					out.WriteInt32(-1)
				} else {
					out.WriteInt32(int32(ir.Item.Socket[socket].Color))
					out.WriteInt32(ir.Item.Socket[socket].Content)
				}
			}

			out.WriteUint32(ir.Item.SocketBonus)
			out.WriteInt32(ir.Item.GemProperties)
			out.WriteInt32(ir.Item.RequiredDisenchantSkill)
			out.WriteFloat32(ir.Item.ArmorDamageModifier)

			if build.AddedIn(8209) {
				out.WriteInt32(ir.Item.Duration)
			}

			if build.AddedIn(9056) {
				out.WriteUint32(uint32(ir.Item.ItemLimitCategory))
			}

			if build.AddedIn(9767) {
				out.WriteUint32(uint32(ir.Item.HolidayID))
			}
		}
	}

	return nil
}

func (ir *ResponseSingle) Decode(build vsn.Build, in *packet.WorldPacket) error {
	ir.QueryID = in.ReadUint32()

	if ir.QueryID&query.EntryNotFound != 0 {
		return nil
	}

	ir.Item.Entry = in.ReadUint32()
	ir.Item.Class = in.ReadUint32()
	ir.Item.Subclass = in.ReadUint32()
	if build.AddedIn(vsn.V2_0_1) {
		ir.Item.SoundOverrideSubclass = in.ReadUint32()
	}
	ir.Item.Name = i18n.Text{
		ir.Locale: in.ReadCString(),
	}
	in.ReadCString()
	in.ReadCString()
	in.ReadCString()
	ir.Item.DisplayID = in.ReadUint32()
	ir.Item.Quality = models.ItemQuality(in.ReadInt32())

	var err error
	ir.Item.Flags, err = update.DecodeItemFlags(build, in)
	if err != nil {
		panic(err)
	}

	ir.Item.BuyPrice = models.Money(in.ReadUint32())
	ir.Item.SellPrice = models.Money(in.ReadUint32())
	ir.Item.InventoryType = models.InventoryType(in.ReadUint32())
	ir.Item.AllowableClass = models.ClassMask(in.ReadInt32())
	ir.Item.AllowableRace = models.RaceMask(in.ReadInt32())
	ir.Item.ItemLevel = in.ReadUint32()
	ir.Item.RequiredLevel = uint8(in.ReadUint32())
	ir.Item.RequiredSkill = in.ReadUint32() // id from SkillLine.dbc
	ir.Item.RequiredSkillRank = in.ReadUint32()

	if build.AddedIn(vsn.V1_12_1) {
		ir.Item.RequiredSpell = in.ReadUint32() // id from Spell.dbc
		ir.Item.RequiredHonorRank = in.ReadUint32()
		ir.Item.RequiredCityRank = in.ReadUint32()
		ir.Item.RequiredReputationFaction = in.ReadUint32() // id from Faction.dbc
		ir.Item.RequiredReputationRank = in.ReadUint32()
	}

	ir.Item.MaxCount = in.ReadUint32()
	ir.Item.Stackable = in.ReadUint32()
	ir.Item.ContainerSlots = uint8(in.ReadUint32())

	ir.Item.Stats = make([]models.ItemStat, 0, 10)

	for x := 0; x < 10; x++ {
		statType := in.ReadUint32()
		statValue := in.ReadInt32()

		if statType == 0 && statValue <= 0 {
			break
		}

		var stat models.ItemStat
		if err := stat.Type.Resolve(build, statType); err != nil {
			return err
		}

		stat.Value = statValue

		ir.Item.Stats = append(ir.Item.Stats, stat)
	}

	for x := 0; x < DmgCount(build); x++ {
		dmgMin := in.ReadFloat32()
		dmgMax := in.ReadFloat32()
		dmgType := in.ReadUint32()

		if dmgMin == 0 && dmgMax == 0 {
			continue
		}

		ir.Item.Damage = append(ir.Item.Damage, models.ItemDamage{
			Min:  dmgMin,
			Max:  dmgMax,
			Type: uint8(dmgType),
		})
	}

	ir.Item.Armor = in.ReadUint32()
	ir.Item.HolyRes = in.ReadUint32()
	ir.Item.FireRes = in.ReadUint32()
	ir.Item.NatureRes = in.ReadUint32()
	ir.Item.FrostRes = in.ReadUint32()
	ir.Item.ShadowRes = in.ReadUint32()

	if build.AddedIn(vsn.V1_12_1) {
		ir.Item.ArcaneRes = in.ReadUint32()
	}

	ir.Item.Delay = in.ReadUint32()
	ir.Item.AmmoType = in.ReadUint32()
	ir.Item.RangedModRange = in.ReadFloat32()

	for x := 0; x < 5; x++ {
		var sp models.ItemSpell
		sp.ID = in.ReadUint32()
		sp.Trigger = in.ReadUint32()
		sp.Charges = in.ReadInt32()
		sp.Cooldown = int64(in.ReadUint32())
		sp.Category = in.ReadUint32()
		sp.CategoryCooldown = int64(in.ReadUint32())
		if sp.ID != 0 {
			ir.Item.Spells = append(ir.Item.Spells, sp)
		}
	}

	ir.Item.Bonding = models.ItemBind(in.ReadUint32())
	ir.Item.Description = i18n.Text{
		ir.Locale: in.ReadCString(),
	}

	ir.Item.PageText = in.ReadUint32()
	ir.Item.LanguageID = in.ReadUint32()
	ir.Item.PageMaterial = in.ReadUint32()
	ir.Item.StartQuest = in.ReadUint32()
	ir.Item.LockID = in.ReadUint32()
	ir.Item.Material = in.ReadInt32()
	ir.Item.Sheath = in.ReadUint32()

	if build > 3368 {
		ir.Item.RandomProperty = in.ReadUint32()
		if build.AddedIn(vsn.V2_0_1) {
			ir.Item.RandomSuffix = in.ReadUint32()
		}
		ir.Item.Block = in.ReadUint32()
		ir.Item.Itemset = in.ReadUint32()
		ir.Item.MaxDurability = in.ReadUint32()
		ir.Item.Area = in.ReadUint32()
		ir.Item.Map = in.ReadInt32()
		ir.Item.BagFamily = in.ReadInt32()

		if build.AddedIn(vsn.V2_0_1) {
			ir.Item.TotemCategory = in.ReadInt32()
			for socket := 0; socket < 3; socket++ {
				var sck models.ItemSocket
				sck.Color = models.SocketColor(in.ReadInt32())
				sck.Content = in.ReadInt32()
				if sck.Color < 0 && sck.Content < 0 {
					continue
				}

				ir.Item.Socket = append(ir.Item.Socket, sck)
			}

			ir.Item.SocketBonus = in.ReadUint32()
			ir.Item.GemProperties = in.ReadInt32()
			ir.Item.RequiredDisenchantSkill = in.ReadInt32()
			ir.Item.ArmorDamageModifier = in.ReadFloat32()

			if build.AddedIn(8209) {
				ir.Item.Duration = in.ReadInt32()
			}

			if build.AddedIn(9056) {
				ir.Item.ItemLimitCategory = in.ReadUint32()
			}

			if build.AddedIn(9767) {
				ir.Item.HolidayID = in.ReadUint32()
			}
		}
	}
	return nil
}

type Push struct {
	Recipient          guid.GUID
	Received           bool
	Created            bool
	ShowInChat         bool
	Bag                models.ItemSlot
	Slot               models.ItemSlot
	Entry              uint32
	PropertySeed       uint32
	RandomPropertiesID uint32
	Count              uint32
	GetItemCount       uint32
}

func (ip *Push) Encode(build vsn.Build, out *packet.WorldPacket) error {
	out.Type = packet.SMSG_ITEM_PUSH_RESULT
	ip.Recipient.EncodeUnpacked(build, out)
	out.WriteBool32(ip.Received)
	out.WriteBool32(ip.Created)
	out.WriteBool32(ip.ShowInChat)
	out.WriteInt8(int8(ip.Bag))
	out.WriteInt32(int32(ip.Slot))
	out.WriteUint32(ip.Entry)
	if build.AddedIn(vsn.V1_12_1) {
		out.WriteUint32(ip.PropertySeed)
		out.WriteUint32(ip.RandomPropertiesID)
	}
	out.WriteUint32(ip.Count)
	out.WriteUint32(ip.GetItemCount)
	return nil
}

func (ip *Push) Decode(build vsn.Build, in *packet.WorldPacket) error {
	var err error
	ip.Recipient, err = guid.DecodeUnpacked(build, in)
	if err != nil {
		return err
	}
	ip.Received = in.ReadBool32()
	ip.Created = in.ReadBool32()
	ip.ShowInChat = in.ReadBool32()
	ip.Bag = models.ItemSlot(in.ReadInt8())
	ip.Slot = models.ItemSlot(in.ReadInt32())
	ip.Entry = in.ReadUint32()
	if build.AddedIn(vsn.V1_12_1) {
		ip.PropertySeed = in.ReadUint32()
		ip.RandomPropertiesID = in.ReadUint32()
	}
	ip.Count = in.ReadUint32()
	ip.GetItemCount = in.ReadUint32()
	return nil
}

type Destroy struct {
	Bag, Slot models.ItemSlot
	Count     uint32
}

func (id *Destroy) Encode(build vsn.Build, out *packet.WorldPacket) error {
	out.WriteInt8(int8(id.Bag))
	out.WriteInt8(int8(id.Slot))
	out.WriteByte(uint8(id.Count))
	return nil
}

func (id *Destroy) Decode(build vsn.Build, in *packet.WorldPacket) error {
	id.Bag = models.ItemSlot(in.ReadInt8())
	id.Slot = models.ItemSlot(in.ReadInt8())
	id.Count = uint32(in.ReadByte())
	return nil
}

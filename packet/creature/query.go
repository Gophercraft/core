package creature

import (
	"github.com/Gophercraft/core/guid"
	"github.com/Gophercraft/core/i18n"
	"github.com/Gophercraft/core/packet"
	"github.com/Gophercraft/core/packet/query"
	"github.com/Gophercraft/core/realm/wdb/models"
	"github.com/Gophercraft/core/vsn"
)

type Query struct {
	Entry uint32
	GUID  guid.GUID
}

func (q *Query) Decode(build vsn.Build, in *packet.WorldPacket) (err error) {
	q.Entry = in.ReadUint32()
	q.GUID, err = guid.DecodeUnpacked(build, in)
	return
}

func (q *Query) Encode(build vsn.Build, out *packet.WorldPacket) (err error) {
	out.Type = packet.CMSG_CREATURE_QUERY
	out.WriteUint32(q.Entry)
	q.GUID.EncodeUnpacked(build, out)
	return
}

type QueryResponse struct {
	Locale           i18n.Locale
	ID               uint32
	DisplayID        uint32
	CreatureFamilyID uint32
	Creature         *models.CreatureTemplate
}

func (q *QueryResponse) Encode(build vsn.Build, out *packet.WorldPacket) error {
	out.Type = packet.SMSG_CREATURE_QUERY_RESPONSE

	crt := q.Creature

	if crt == nil {
		out.WriteUint32(q.ID | query.EntryNotFound)
	} else {
		var creatureTypeFlags uint32
		if crt.Tameable {
			creatureTypeFlags |= 1
		} // Makes the mob tameable (must also be a beast and have family set)
		if crt.VisibleToGhosts {
			creatureTypeFlags |= 2
		} // Sets Creatures that can ALSO be seen when player is a ghost. Used in CanInteract function by client, can’t be attacked
		if crt.BossLevel {
			creatureTypeFlags |= 4
		}
		if crt.DontPlayWoundParryAnim {
			creatureTypeFlags |= 8
		}
		if crt.HideFactionTooltip {
			creatureTypeFlags |= 16
		} // Controls something in client tooltip related to creature faction
		if crt.SpellAttackable {
			creatureTypeFlags |= 64
		}
		if crt.DeadInteract {
			creatureTypeFlags |= 128
		}
		if crt.HerbLoot {
			creatureTypeFlags |= 256
		} // Uses Skinning Loot Field
		if crt.MiningLoot {
			creatureTypeFlags |= 512
		} // Makes Mob Corpse Mineable – Uses Skinning Loot Field
		if crt.DontLogDeath {
			creatureTypeFlags |= 1024
		}
		if crt.MountedCombat {
			creatureTypeFlags |= 2048
		}
		if crt.CanAssist {
			creatureTypeFlags |= 4096
		} //	Can aid any player or group in combat. Typically seen for escorting NPC’s
		if crt.PetHasActionBar {
			creatureTypeFlags |= 8192
		} // 	checked from calls in Lua_PetHasActionBar
		if crt.MaskUID {
			creatureTypeFlags |= 16384
		}
		if crt.EngineerLoot {
			creatureTypeFlags |= 32768
		} //	Makes Mob Corpse Engineer Lootable – Uses Skinning Loot Field
		if crt.ExoticPet {
			creatureTypeFlags |= 65536
		} // Tamable as an exotic pet. Normal tamable flag must also be set.
		if crt.UseDefaultCollisionBox {
			creatureTypeFlags |= 131072
		}
		if crt.IsSiegeWeapon {
			creatureTypeFlags |= 262144
		}
		if crt.ProjectileCollision {
			creatureTypeFlags |= 524288
		}
		if crt.HideNameplate {
			creatureTypeFlags |= 1048576
		}
		if crt.DontPlayMountedAnim {
			creatureTypeFlags |= 2097152
		}
		if crt.IsLinkAll {
			creatureTypeFlags |= 4194304
		}
		if crt.InteractOnlyWithCreator {
			creatureTypeFlags |= 8388608
		}
		if crt.ForceGossip {
			creatureTypeFlags |= 134217728
		}

		out.WriteUint32(q.ID)
		out.WriteCString(crt.Name.GetLocalized(q.Locale))
		out.WriteCString("")
		out.WriteCString("")
		out.WriteCString("")
		out.WriteCString(crt.SubName.GetLocalized(q.Locale))
		out.WriteUint32(creatureTypeFlags)
		out.WriteUint32(crt.CreatureType)

		out.WriteUint32(q.CreatureFamilyID)
		out.WriteUint32(crt.Rank)
		out.WriteUint32(0)
		out.WriteUint32(crt.PetSpellDataId)

		var displayID uint32

		if q.DisplayID != 0 {
			displayID = q.DisplayID
		} else {
			if len(crt.DisplayIDs) > 0 {
				displayID = crt.DisplayIDs[0]
			}
		}

		out.WriteUint32(displayID)
		civilian := uint16(0)

		if crt.Civilian {
			civilian = 1
		}

		out.WriteUint16(civilian)
	}
	return nil
}

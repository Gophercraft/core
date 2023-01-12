package realm

import (
	"fmt"
	"log"

	"github.com/Gophercraft/core/home/rpcnet"
	"github.com/Gophercraft/core/packet/character"
	"github.com/Gophercraft/core/packet/spell"

	"github.com/Gophercraft/core/format/dbc/dbdefs"
	"github.com/Gophercraft/core/realm/wdb"
	"github.com/Gophercraft/core/realm/wdb/models"

	"github.com/Gophercraft/core/guid"
)

func (s *Session) HandleDeleteCharacter(del *character.Delete) {
	var char models.Character

	found, _ := s.DB().Where("id = ?", del.ID.Counter()).Get(&char)

	if !found {
		s.Send(&character.DeleteResult{
			Result: character.CharDeleteFailed,
		})
		return
	}

	if char.GameAccount != s.GameAccount {
		return
	}

	s.Server.ScrubCharacter(del.ID)

	s.Send(&character.DeleteResult{
		Result: character.CharDeleteSuccess,
	})
}

// ScrubCharacter deletes a character PERMANENTLY from a server.
func (s *Server) ScrubCharacter(chr guid.GUID) {
	s.DB.Where("id = ?", chr.Counter()).Delete(new(models.Character))
	s.DB.Where("owner = ?", chr.Counter()).Delete(new(models.Item))
}

func (s *Session) findCharacterEquipment(char uint64) []character.Equipment {
	equipLen := character.EquipLen(s.Build())
	itemList := make([]character.Equipment, equipLen)
	var inventory []models.Inventory
	err := s.DB().Where("player = ?", char).Where("bag = ?", models.Backpack).Where("slot < 19").Find(&inventory)
	if err != nil {
		log.Fatal(err)
	}

	invList := make([]*models.Inventory, 24)

	for i := range inventory {
		inv := &inventory[i]

		if inv.Slot < models.PaperDoll_Head || inv.Slot > models.PaperDoll_Bag4 {
			continue
		}

		invList[int(inv.Slot)] = inv
	}

	x := 0

	for i := models.PaperDoll_Head; i < models.ItemSlot(len(itemList)); i++ {
		invRef := invList[int(i)]
		if invRef == nil {
			x++
			continue
		}

		var item models.Item
		found, err := s.DB().Where("id = ?", invRef.ItemID).Get(&item)
		if !found {
			panic(err)
		}

		pi := character.Equipment{
			Model: item.DisplayID,
			Type:  uint8(item.ItemType),
		}

		// No transmog
		if item.DisplayID == 0 {
			var itt *models.ItemTemplate
			s.DB().Lookup(wdb.BucketKeyStringID, item.ItemID, &itt)
			if itt != nil {
				pi.Model = itt.DisplayID
			}
		}

		if len(item.Enchantments) > 0 {
			pi.Enchantment = item.Enchantments[0]
		}

		itemList[x] = pi
		x++
	}

	return itemList
}

func (s *Session) ValidName(name string) bool {
	if name == "System" {
		return false
	}

	if s.Tier >= rpcnet.Tier_Admin {
		return true
	}

	return true
}

func (s *Session) HandleRequestCharacterList() {
	log.Println("Character list requested")

	// Load characters for this GameAccount
	// from DB
	var chars []models.Character
	err := s.Server.DB.Where("game_account = ?", s.GameAccount).Find(&chars)
	if err != nil {
		panic(err)
	}

	// Start building packet structure
	list := &character.List{}
	list.Chars = make([]character.Char, len(chars))

	for i, char := range chars {
		characterGUID := guid.RealmSpecific(guid.Player, s.Server.RealmID(), char.ID)

		var flags character.Flags

		log.Println("Getting session")

		sess, err := s.Server.GetSessionByGUID(characterGUID)
		if err != nil {
			log.Println(err)
		}

		log.Println("Done getting session")

		if sess != nil {
			// Player will be warned if their character is already in-world
			flags |= character.FlagLockedForTransfer
		}

		// Load options
		if char.HideHelm {
			flags |= character.FlagHideHelm
		}

		if char.HideCloak {
			flags |= character.FlagHideCloak
		}

		if char.Ghost {
			flags |= character.FlagGhost
		}

		level := char.Level
		if level > 255 {
			level = 255
		}

		log.Println("Adding list")

		list.Chars[i] = character.Char{
			GUID:       characterGUID,
			Name:       char.Name,
			Race:       character.Race(char.Race),
			Class:      character.Class(char.Class),
			BodyType:   char.BodyType,
			Skin:       char.Skin,
			Face:       char.Face,
			HairStyle:  char.HairStyle,
			HairColor:  char.HairColor,
			FacialHair: char.FacialHair,
			FirstLogin: char.FirstLogin,
			Level:      uint8(char.Level),
			Flags:      flags,
			Zone:       char.Zone, // Goldshire. Once the login test is complete,
			Map:        char.Map,  // and players can move around within Goldshire without error
			X:          char.X,    // we can replace this with database data.
			Y:          char.Y,
			Z:          char.Z,
			Equipment:  s.findCharacterEquipment(char.ID),
		}
	}

	log.Println("Sending list")
	s.Send(list)
}

func (s *Server) CreateCharacter(playerCharacter *models.Character) character.Result {
	_, err := s.GetGUIDByPlayerName(playerCharacter.Name)
	if err == nil {
		return character.CharCreateNameInUse
	}

	// TODO: validate race and class combo

	var race *dbdefs.Ent_ChrRaces
	s.DB.Lookup(wdb.BucketKeyUint32ID, uint32(playerCharacter.Race), &race)

	var class *dbdefs.Ent_ChrClasses
	s.DB.Lookup(wdb.BucketKeyUint32ID, uint32(playerCharacter.Class), &class)

	if race == nil {
		fmt.Println("race not found", playerCharacter.Race)
		return character.CharCreateRestrictedRaceclass
	}

	if class == nil {
		return character.CharCreateRestrictedRaceclass
	}

	defaultLevel := uint32(s.UintVar("Char.StartLevel"))
	classLevel := class.StartingLevel
	if classLevel == 0 {
		classLevel = 1
	}

	playerCharacter.Level = defaultLevel

	if defaultLevel > uint32(classLevel) {
		playerCharacter.Level = defaultLevel
	} else {
		playerCharacter.Level = uint32(classLevel)
	}

	startPos := s.PosVar("Char.StartPosition")

	if startPos == nil {
		playerCharacter.FirstLogin = true

		var plci *models.PlayerCreateInfo

		s.DB.Range(func(pci *models.PlayerCreateInfo) bool {
			if pci.Class == playerCharacter.Class && pci.Race == playerCharacter.Race {
				plci = pci
				return false
			}
			return true
		})

		if plci != nil {
			playerCharacter.Zone = plci.Zone
			playerCharacter.Map = plci.Map
			playerCharacter.X = plci.Position.X
			playerCharacter.Y = plci.Position.Y
			playerCharacter.Z = plci.Position.Z
			playerCharacter.O = plci.Position.W
		}
	} else {
		playerCharacter.Zone = 0
		playerCharacter.Map = startPos.MapID
		playerCharacter.X = startPos.X
		playerCharacter.Y = startPos.Y
		playerCharacter.Z = startPos.Z
		playerCharacter.O = startPos.W
	}

	if _, err := s.DB.Insert(playerCharacter); err != nil {
		log.Fatal(err)
	}

	var (
		// map to check if a paperdoll already has an item in this slot
		checkSlot = map[models.ItemSlot]bool{}
		// The bag where inventory items are inserted
		bag models.ItemSlot = models.Backpack
		// This doesn't need to be the actual maximum.
		// This is just the point at which items start being placed in different bags
		maxBackpackSlots = 16
		// Slot begins at the end of the equip slots (i.e. the beginning of the backpack slots)
		slot models.ItemSlot = models.GetStartItemSlot(s.Build())
		// List of inventory objects. To be inserted at the end of player creation
		inventory []models.Inventory
		// The index where weapons can be safely stored. After a weapon is placed in the main hand,
		// wpn will become PaperDoll_OffHand
		// And if another weapon comes after that, it is discarded
		wpn = models.PaperDoll_MainHand
		// The index where container items can be equipped in the bottom right screen
		bagslot = models.PaperDoll_Bag1
		// The index where rings can be equipped
		finger = models.PaperDoll_Finger1
		// The index where trinkets can be equipped
		trinket = models.PaperDoll_Trinket1
		// The index into the number of slots
		bagIndex int
		bagSlots [4]uint8

		createItems []models.PlayerCreateItem
	)

	s.DB.Range(func(csi *models.PlayerCreateItem) {
		// Check if class and race masks include our class and race.
		if csi.Class.Has(playerCharacter.Class) && csi.Race.Has(playerCharacter.Race) {
			createItems = append(createItems, *csi)
		}
	})

	// First, handle all bag items.
	// Handling these separately allows us to properly distribute create items into other bags should the need arise.
	for _, csi := range createItems {
		switch csi.Equip {
		case models.EquipContainer:
			var itt *models.ItemTemplate
			s.DB.Lookup(wdb.BucketKeyStringID, csi.Item, &itt)
			if itt == nil {
				panic(fmt.Errorf("no item template found in PlayerCreateItem: %s", csi.Item))
			}

			num := csi.Amount
			if num == 0 {
				num = 1
			}

			// Max 4 bags
			if num > 4 {
				num = 4
			}

			for i := uint32(0); i < num; i++ {
				// Create the equip item
				var item models.Item
				item.ItemID = csi.Item
				// Always 1 for bag items.
				item.StackCount = 1
				item.DisplayID = itt.DisplayID
				item.ItemType = itt.InventoryType
				s.DB.Insert(&item)
				// Create the item's inventory. (Which player owns the item and where it's located)
				var inv models.Inventory
				inv.Player = playerCharacter.ID
				inv.Bag = models.Backpack
				inv.Slot = bagslot
				inv.ItemID = item.ID

				inventory = append(inventory, inv)

				bagSlots[bagIndex] = itt.ContainerSlots
				bagslot++
				bagIndex++
				if bagIndex == 4 {
					break
				}
			}
		}
	}

addCreateItems:
	for _, csi := range createItems {
		var itt *models.ItemTemplate
		s.DB.Lookup(wdb.BucketKeyStringID, csi.Item, &itt)
		if itt == nil {
			panic(fmt.Errorf("no item template found in PlayerCreateItem: %s", csi.Item))
		}

		// Infer the item slot based on Equip type, as well as the InventoryType of the item template.
		// If more than one candidates for a slot exist, the item will not be added.
		// If this is not desirable, put the item under EquipInventory
		switch csi.Equip {
		case models.EquipContainer:
			// already handled
		case models.EquipPaperDoll:
			// Create the equip item
			var item models.Item
			item.ItemID = csi.Item
			item.StackCount = csi.Amount
			if item.StackCount < 1 {
				item.StackCount = 1
			}
			if itt.Stackable > 0 {
				if item.StackCount > itt.Stackable {
					// Cannot exist having more than is stackable
					item.StackCount = itt.Stackable
				}
			}
			item.DisplayID = itt.DisplayID
			item.ItemType = itt.InventoryType
			// Create the item's inventory. (Which player owns the item and where it's located)
			var inv models.Inventory
			inv.Player = playerCharacter.ID
			inv.Bag = models.Backpack
			switch itt.InventoryType {
			case models.IT_Weapon, models.IT_MainHand, models.IT_TwoHand:
				// Bruh unless this race is born with 3 arms you're out of luck
				if wpn == models.PaperDoll_OffHand+1 {
					continue addCreateItems
				}
				inv.Slot = wpn
				wpn++
			case models.IT_Finger:
				// No more fingers
				if finger == models.PaperDoll_Finger2+1 {
					continue addCreateItems
				}
				inv.Slot = finger
				finger++
			case models.IT_Trinket:
				// No more trinket slots
				if trinket == models.PaperDoll_Trinket2+1 {
					continue addCreateItems
				}
				inv.Slot = trinket
				trinket++
			default:
				// this should work in most cases
				var err error
				inv.Slot, err = itt.InventoryType.PaperDollSlot()
				if err != nil {
					panic(err)
				}
			}

			// PaperDoll already has an item of this slot.
			if checkSlot[inv.Slot] {
				continue addCreateItems
			}

			checkSlot[inv.Slot] = true

			// Register item and get ID for inventory
			s.DB.Insert(&item)
			inv.ItemID = item.ID
			inventory = append(inventory, inv)
		case models.EquipInventory:
			// Here's where things like consumables and Hearthstones are put
			maxAmountPerItem := itt.Stackable
			if maxAmountPerItem == 0 {
				maxAmountPerItem = 1
			}

			amountRemaining := csi.Amount

			if amountRemaining == 0 {
				amountRemaining = 1
			}

			for amountRemaining > 0 {
				if bag == models.Bag4+1 {
					panic("bag limit reached")
				}

				var itemStack models.Item
				itemStack.ItemID = csi.Item
				itemStack.StackCount = amountRemaining
				if itemStack.StackCount > maxAmountPerItem {
					itemStack.StackCount = maxAmountPerItem
				}

				itemStack.DisplayID = itt.DisplayID
				itemStack.ItemType = itt.InventoryType

				s.DB.Insert(&itemStack)

				var invStack models.Inventory
				invStack.ItemID = itemStack.ID
				invStack.Player = playerCharacter.ID
				invStack.Bag = bag
				invStack.Slot = slot

				inventory = append(inventory, invStack)

				amountRemaining -= itemStack.StackCount

				slot++
				if bag == models.Backpack {
					if slot == models.ItemSlot(maxBackpackSlots) {
						slot = 0
						bag = models.Bag1
					}
				} else {
					maxSlots := bagSlots[int(bag)]
					if uint8(uint32(slot)) == maxSlots {
						slot = 0
						bag++

						if bag == models.Bag4+1 {
							break
						}
					}
				}
			}
		}

		// s.DB.Insert(&item)
		// inventory = append(inventory, inv)
		// slot++
	}

	if len(inventory) > 0 {
		_, err := s.DB.Insert(&inventory)
		if err != nil {
			panic(err)
		}
	}

	var learnedAbilities []models.LearnedAbility

	s.DB.Range(func(v *models.PlayerCreateAbility) {
		if v.Race.Has(playerCharacter.Race) && v.Class.Has(playerCharacter.Class) {
			learnedAbilities = append(learnedAbilities, models.LearnedAbility{
				Player: playerCharacter.ID,
				Spell:  v.Spell,
			})
		}
	})

	var actionButtons []models.ActionButton

	s.DB.Range(func(v *models.PlayerCreateActionButton) {
		if (v.Race == playerCharacter.Race) && (v.Class == playerCharacter.Class) {
			actionButtons = append(actionButtons, models.ActionButton{
				Player: playerCharacter.ID,
				Button: v.Button,
				Action: v.Action,
				Type:   v.Type,
				Misc:   v.Misc,
			})
		}
	})

	if len(actionButtons) > 0 {
		s.DB.Insert(&actionButtons)
	}

	if len(learnedAbilities) > 0 {
		s.DB.Insert(&learnedAbilities)
	}

	return character.CharCreateSuccess
}

func (s *Session) HandleCreateCharacter(cc *character.Create) {
	if cc.Name == "" {
		s.Send(&character.CreateResult{
			Result: character.CharNameNoName,
		})
		return
	}

	log.Println("Registered name: ", cc.Name)

	playerCharacter := models.Character{}
	playerCharacter.ID = 0 // will be overwritten by insert
	playerCharacter.GameAccount = s.GameAccount
	playerCharacter.RealmID = s.Server.RealmID()
	playerCharacter.Name = cc.Name
	playerCharacter.Race = cc.Race
	playerCharacter.Class = cc.Class
	playerCharacter.BodyType = cc.BodyType
	playerCharacter.Skin = cc.Skin
	playerCharacter.Face = cc.Face
	playerCharacter.HairStyle = cc.HairStyle
	playerCharacter.HairColor = cc.HairColor
	playerCharacter.FacialHair = cc.FacialHair

	result := s.Server.CreateCharacter(&playerCharacter)

	s.Send(&character.CreateResult{
		Result: result,
	})
}

// Get the native display ID for a race and body type
func (s *Server) GetNative(race models.Race, bodyType uint8) uint32 {
	var races *dbdefs.Ent_ChrRaces
	s.DB.Lookup(wdb.BucketKeyUint32ID, uint32(race), &races)
	// Just a gopher
	if races == nil {
		panic(race)
		return 2838
	}

	// gender is a fuck
	switch bodyType {
	case 1:
		return uint32(races.FemaleDisplayID)
	default:
		return uint32(races.MaleDisplayID)
	}
}

func (s *Session) GetCharacterRace() models.Race {
	return s.Char.Race
}

func (s *Session) GetCharacterClass() models.Class {
	return s.Char.Class
}

func (s *Session) GetCharacterLevel() uint32 {
	return s.Char.Level
}

func (s *Session) GetCharacterHealth() uint32 {
	return s.Char.Health
}

func (s *Session) GetCharacterPower() spell.PowerType {
	return s.Server.ClassPowerType(s.GetCharacterClass())
}

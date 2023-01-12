package realm

import (
	"fmt"
	"sort"

	"github.com/Gophercraft/core/guid"
	"github.com/Gophercraft/core/packet/character"
	"github.com/Gophercraft/core/packet/item"
	"github.com/Gophercraft/core/packet/update"
	"github.com/Gophercraft/core/realm/wdb"
	"github.com/Gophercraft/core/realm/wdb/models"
	"github.com/Gophercraft/core/vsn"
	"github.com/Gophercraft/log"
	"github.com/superp00t/etc"
)

type Item struct {
	ItemID string
	*update.ValuesBlock
}

func (i *Item) Entry() uint32 {
	return i.Get("Entry").Uint32()
}

func (i *Item) GUID() guid.GUID {
	if i == nil {
		return guid.Nil
	}

	return i.Get("GUID").GUID()
}

func (i *Item) PropertySeed() uint32 {
	value := i.Get("PropertySeed")
	if value == nil {
		return 0
	}
	return value.Uint32()
}

func (i *Item) RandomPropertiesID() uint32 {
	value := i.Get("RandomPropertiesID")
	if value == nil {
		return 0
	}
	return value.Uint32()
}

func (i *Item) StackCount() uint32 {
	sc := i.Get("StackCount").Uint32()
	if sc == 0 {
		return 1
	}

	return sc
}

func (i *Item) ContainerNumSlots() uint32 {
	return i.Get("NumSlots").Uint32()
}

func (i *Item) BagEmpty() bool {
	gArray := i.Get("Slots")

	for x := 0; x < gArray.Len(); x++ {
		g := gArray.Index(x).GUID()
		if g != guid.Nil {
			return false
		}
	}

	return true
}

func (i *Item) IsBag() bool {
	return i.ValuesBlock.TypeMask&guid.TypeMaskContainer != 0
}

func (i *Item) ID() uint64 {
	return i.GUID().Counter()
}

func (i *Item) TypeID() guid.TypeID {
	if i.IsBag() {
		return guid.TypeContainer
	}

	return guid.TypeItem
}

func (i *Item) Values() *update.ValuesBlock {
	return i.ValuesBlock
}

func (s *Session) GetItemTemplate(it models.Item) *models.ItemTemplate {
	itmp, err := s.DB().GetItemTemplate(it.ItemID)
	if err != nil {
		panic(err)
	}
	return itmp
}

func (s *Session) GetItemTemplateByEntry(entry uint32) *models.ItemTemplate {
	itmp, err := s.DB().GetItemTemplateByEntry(entry)
	if err != nil {
		panic(err)
	}
	return itmp
}

func (s *Session) NewItem(it models.Item) *Item {
	template := s.GetItemTemplate(it)
	mask := guid.TypeMaskObject | guid.TypeMaskItem

	if template.ContainerSlots > 0 {
		mask |= guid.TypeMaskContainer
	}

	values, err := update.NewValuesBlock(s.Build(), mask)
	if err != nil {
		panic(err)
	}

	i := &Item{it.ItemID, values}
	g := guid.RealmSpecific(guid.Item, s.Server.RealmID(), it.ID)
	i.SetGUID("GUID", g)
	i.SetUint32("Entry", template.Entry)
	// code, err := template.Flags.Resolve(s.Build())
	// if err != nil {
	// 	panic(err)
	// }

	if mask&guid.TypeMaskContainer != 0 {
		i.SetUint32("NumSlots", uint32(template.ContainerSlots))
	}

	i.SetFloat32("ScaleX", 1.0)

	i.SetGUID("Owner", s.GUID())
	i.SetGUID("Contained", s.GUID())
	if it.Creator != 0 {
		i.SetGUID("Creator", guid.RealmSpecific(guid.Player, s.Server.RealmID(), it.Creator))
	}

	if s.Build().AddedIn(vsn.V1_12_1) {
		i.SetUint32("Durability", template.MaxDurability)
		i.SetUint32("MaxDurability", template.MaxDurability)
	}

	if template.Stackable != 0 {
		i.SetUint32("StackCount", it.StackCount)
	}

	i.SetUint32("Duration", uint32(template.Duration))

	//	todo: source charges from item struct
	for x := 0; x < len(template.Spells); x++ {
		i.Get("SpellCharges").Index(x).SetInt32(template.Spells[x].Charges)
	}

	// log.Warn(template.Flags, flg, fmt.Sprintf("0x%08X\n", uint32(code)))

	// field flags are not item flags
	// i.SetUint32("Flags", uint32(code))

	return i
}

func (s *Session) PlayerID() uint64 {
	return s.GUID().Counter()
}

func (s *Session) SetVisibleItemEntry(index int, entry uint32) {
	if s.Build() > 3368 {
		visibleItems := s.ValuesBlock.Get("VisibleItems")

		if visibleItems.Len() <= index {
			return
		}

		visibleItems.Index(index).Field("Entry").SetUint32(entry)
	}
}

func (s *Session) InitInventoryManager() {
	if s.Build() <= 3368 {
		s.SetUint32("NumInvSlots", 0x89)
	}

	log.Println("Set numInvSlots")

	var inv []models.Inventory
	s.DB().Where("player = ?", s.PlayerID()).Find(&inv)

	log.Println("found inv")

	displaySlots := map[models.ItemSlot]uint64{}

	for _, v := range inv {
		if v.Bag == models.Backpack && v.Slot <= models.ItemSlot(character.EquipLen(s.Build())) {
			displaySlots[v.Slot] = v.ItemID
		}
	}

	log.Println("Setting visible items")

	for i, itemID := range displaySlots {
		log.Println("Looking up", itemID)
		var instance models.Item
		found, err := s.DB().Where("id = ?", itemID).Get(&instance)
		if !found {
			panic(err)
		}

		var itemTemplate *models.ItemTemplate
		s.DB().Lookup(wdb.BucketKeyStringID, instance.ItemID, &itemTemplate)
		if itemTemplate == nil {
			panic("could not find item template " + instance.ItemID)
		}

		// s.ValuesBlock.SetArrayValue(update.PlayerVisibleItems, int(i), "Creator", guid.RealmSpecific(guid.Player, s.Server.RealmID(), item.Creator))
		s.SetVisibleItemEntry(int(i), itemTemplate.Entry)
	}

	log.Println("Set visible items")

	s.Items = make(map[guid.GUID]*Item)

	var createObjects []Object

	log.Println("getting item slots")

	inventorySlots := s.Get("InventorySlots")

	log.Println("Building item creates", len(inv))

	for _, v := range inv {
		log.Println(v.Bag, v.Slot)
		if v.Bag == models.Backpack {
			var it models.Item
			found, err := s.DB().Where("id = ?", v.ItemID).Get(&it)
			if !found {
				panic(err)
			}

			itemObject := s.NewItem(it)

			inventorySlots.Index(int(v.Slot)).SetGUID(itemObject.GUID())

			if itemObject.IsBag() {
				bagSlots := itemObject.Get("Slots")

				for _, bagContent := range inv {
					log.Println("bag content", bagContent.Bag, bagContent.Slot)
					if bagContent.Bag == v.Slot {
						var bagContentItem models.Item
						found, err := s.DB().Where("id = ?", bagContent.ItemID).Get(&bagContentItem)
						if !found {
							panic(err)
						}

						bagContentObject := s.NewItem(bagContentItem)
						bagContentObject.SetGUID("Contained", itemObject.GUID())
						s.Items[bagContentObject.GUID()] = bagContentObject

						bagSlots.Index(int(bagContent.Slot)).SetGUID(bagContentObject.GUID())

						createObjects = append(createObjects, bagContentObject)
					}
				}
			}

			createObjects = append(createObjects, itemObject)

			s.Items[itemObject.GUID()] = itemObject
		}
	}

	log.Println("Sending", len(createObjects), "item object creates")
	s.SendObjectCreate(createObjects...)
}

func (s *Session) HandleItemQuerySingle(siq *item.QuerySingle) {
	fmt.Println("player queried item...", siq.ID)

	var it *models.ItemTemplate
	s.DB().Lookup(wdb.BucketKeyEntry, siq.ID, &it)
	fmt.Println("Queried", siq.ID, it != nil)

	s.Send(&item.ResponseSingle{
		QueryID: siq.ID,
		Locale:  s.Locale,
		Item:    it,
	})
}

func (s *Session) HandleSwapInventoryItem(swp *item.SwapBackpackRequest) {
	s.SwapItem(models.Backpack, swp.SrcSlot, models.Backpack, swp.DstSlot)
}

func (s *Session) HandleSwapItem(swp *item.SwapRequest) {
	s.SwapItem(swp.SrcBag, swp.SrcSlot, swp.DstBag, swp.DstSlot)
}

func (s *Session) GetItemByPos(bag, slot models.ItemSlot) (*models.Inventory, *Item) {
	var target guid.GUID

	// TODO: Check with different versions
	maxSlot := models.ItemSlot(24)

	if bag == models.Backpack {
		target = s.Get("InventorySlots").Index(int(slot)).GUID()
	} else {
		if bag > maxSlot {
			return nil, nil
		}

		bagGUID := s.Get("InventorySlots").Index(int(bag)).GUID()
		if bagGUID == guid.Nil {
			fmt.Println("bag", bag, "doesnt exist")
			return nil, nil
		}

		bagIt, ok := s.Items[bagGUID]
		if !ok {
			return nil, nil
		}

		target = bagIt.Get("Slots").Index(int(slot)).GUID()
	}

	if target == guid.Nil {
		fmt.Println("target does not exist in", bag, slot)
		return nil, nil
	}

	it, ok := s.Items[target]
	if !ok {
		panic("Item referenced in inventory but does not exist in inventory manager: " + target.String())
	}

	return &models.Inventory{
		ItemID: it.ID(),
		Player: s.PlayerID(),
		Bag:    bag,
		Slot:   slot,
	}, it

	// var inv models.Inventory
	// found, _ := s.DB().Where("player = ?", s.PlayerID()).Where("bag = ?", bag).Where("slot = ?", slot).Get(&inv)
	// if !found {
	// 	return nil, nil
	// }

	// trg := guid.RealmSpecific(guid.Item, s.Server.RealmID(), inv.ItemID)

	// return &inv, it
}

func (s *Session) IsEquipmentPos(bag, slot models.ItemSlot) bool {
	if bag != models.Backpack {
		return false
	}

	return slot < models.PaperDoll_Bag1
}

func (s *Session) IsValidPos(bag, slot models.ItemSlot) bool {
	// main backpack
	if bag == models.Backpack {
		return slot < 39
	}

	if bag > models.PaperDoll_Bag4 {
		return false
	}

	// check bag slots
	bagGUID := s.ValuesBlock.Get("InventorySlots").Index(int(bag)).GUID()

	if bagGUID == guid.Nil {
		return false
	}

	bagItem := s.Items[bagGUID]
	if bagItem == nil {
		return false
	}

	if bagItem.IsBag() == false {
		return false
	}

	return true
}

func (s *Session) HasItem(entry string) bool {
	for _, i := range s.Items {
		if i.ItemID == entry {
			return true
		}
	}
	return false
}

func (s *Session) EquippableIn(it *Item, slot models.ItemSlot) bool {
	itt := s.GetItemTemplateByEntry(it.Entry())

	if itt.InventoryType == models.IT_Weapon {
		// TODO: figure out why it's like this
		if slot == (models.PaperDoll_MainHand) || slot == (models.PaperDoll_OffHand) {
			return true
		}
	}

	iType, ok := models.ItemDisplaySlots[models.InventoryType(itt.InventoryType)]
	if !ok {
		return false
	}

	return models.ItemSlot(iType) == slot
}

// UNSAFE! will cause database/game corruption or exploits if called incorrectly, or with invalid parameters.
// Just transfers an item internally, irrespective of restrictions.
// If something is in dstInv (which it should not be), it will be lost forever.
func (s *Session) transferItemUnsafe(srcInv *models.Inventory, deleteSrc bool, dstBag, dstSlot models.ItemSlot) {
	if srcInv.Bag == models.Backpack {
		if deleteSrc {
			s.ValuesBlock.Get("InventorySlots").Index(int(srcInv.Slot)).SetGUID(guid.Nil)

			if srcInv.Slot < models.PaperDoll_Bag1 {
				// remove armor and show change
				s.SetVisibleItemEntry(int(srcInv.Slot), 0)
			}
		}
	} else {
		if deleteSrc {
			bgItem := s.GetBagItem(srcInv.Bag)
			bgItem.Get("Slots").Index(int(srcInv.Slot)).SetGUID(guid.Nil)
		}
	}

	log.Println("transferring item unsafe: from bag ", srcInv.Bag, "slot", srcInv.Slot, "to", dstBag, dstSlot)

	srcInv.Bag = dstBag
	srcInv.Slot = dstSlot

	s.DB().AsyncTx(func(tx *wdb.Tx) {
		tx.Where("item_id = ?", srcInv.ItemID).Cols("bag", "slot").Update(srcInv)
	})

	if dstBag == models.Backpack {
		s.ValuesBlock.Get("InventorySlots").Index(int(dstSlot)).SetGUID(guid.RealmSpecific(guid.Item, s.Server.RealmID(), srcInv.ItemID))

		if dstSlot < models.PaperDoll_Bag1 {
			// show armor change to other players
			var it models.Item
			found, err := s.DB().Where("id = ?", srcInv.ItemID).Get(&it)
			if !found {
				panic(err)
			}

			tpl := s.GetItemTemplate(it)

			s.SetVisibleItemEntry(int(dstSlot), tpl.Entry)
		}
	} else {
		bgItem := s.GetBagItem(dstBag)
		bgItem.Get("Slots").Index(int(dstSlot)).SetGUID(guid.RealmSpecific(guid.Item, s.Server.RealmID(), srcInv.ItemID))
	}
}

func (s *Session) SwapItem(srcBag, srcSlot, dstBag, dstSlot models.ItemSlot) {
	log.Println("swapping from source bag", srcBag, "slot", srcSlot, "-> bag", dstBag, "slot", dstSlot)

	if !s.IsAlive() {
		s.Send(&item.InvChangeError{
			Result: item.InvYouAreDead,
		})
		return
	}

	// Start position validation

	if !s.IsValidPos(srcBag, srcSlot) {
		log.Warn("Invalid pos src ", srcBag, srcSlot)
		s.Send(&item.InvChangeError{
			Result: item.InvItemNotFound,
		})
		return
	}

	if !s.IsValidPos(dstBag, dstSlot) {
		log.Warn("Invalid pos dst ", dstBag, dstSlot)
		s.Send(&item.InvChangeError{
			Result: item.InvItemsCantBeSwapped,
		})
		return
	}

	srcInv, src := s.GetItemByPos(srcBag, srcSlot)
	dstInv, dst := s.GetItemByPos(dstBag, dstSlot)

	if src == nil {
		s.Send(&item.InvChangeError{
			Result: item.InvItemNotFound,
		})
		return
	}

	// cannot put bag in itself.
	if src.IsBag() && dstBag == srcSlot {
		s.Send(&item.InvChangeError{
			Result:  item.InvNonEmptyBagOverOtherBag,
			SrcItem: src.GUID(),
			DstItem: dst.GUID(),
		})
		return
	}

	if dst != nil {
		if src.IsBag() && dst.IsBag() {
			// todo: seamless swap
			if !src.BagEmpty() || !dst.BagEmpty() {
				s.Send(&item.InvChangeError{
					Result:  item.InvNonEmptyBagOverOtherBag,
					SrcItem: src.GUID(),
					DstItem: dst.GUID(),
				})
				return
			}
		}
	} else {
		if src.IsBag() {
			if !src.BagEmpty() {
				s.Send(&item.InvChangeError{
					Result:  item.InvNonEmptyBagOverOtherBag,
					SrcItem: src.GUID(),
					DstItem: dst.GUID(),
				})
				return
			}
		}
	}

	if dstBag == models.Backpack {
		if srcBag == models.Backpack {
			// What the hell? You can't put your shirt in your pants slot, come on...
			// TODO: Actually you can swap rings and trinkets: work on this.
			if srcSlot < models.PaperDoll_Bag1 && dstSlot < models.PaperDoll_Bag1 {
				s.Send(&item.InvChangeError{
					Result:  item.InvNonEmptyBagOverOtherBag,
					SrcItem: src.GUID(),
					DstItem: dst.GUID(),
				})
				return
			}
		}

		// is the target slot an equipment slot?
		if dstSlot < models.PaperDoll_Bag1 {
			if !s.EquippableIn(src, dstSlot) {
				log.Warn(src.Entry(), "not equippable in", dstSlot)
				s.Send(&item.InvChangeError{
					Result:  item.InvItemsCantBeSwapped,
					SrcItem: src.GUID(),
					DstItem: dst.GUID(),
				})
				return
			}
		}
	}

	// Target is not empty. We have to transfer the target back to src to complete the slot.
	if dst != nil && srcBag == models.Backpack && srcSlot < 19 {
		if !s.EquippableIn(dst, srcSlot) {
			s.Send(&item.InvChangeError{
				Result:  item.InvItemCantBeEquipped,
				SrcItem: src.GUID(),
				DstItem: dst.GUID(),
			})
			return
		}
	}

	// TODO: implement bag checks
	// TODO: swap filled bags

	// merge stacks
	if dst != nil && src != nil {
		// same type
		if src.Entry() == dst.Entry() {
			tpl := s.GetItemTemplateByEntry(src.Entry())
			if tpl.Stackable != 0 {
				availableSpace := tpl.Stackable - dst.StackCount()
				if availableSpace == 0 {
					fmt.Println("no space", dst)
				} else {
					srcHas := src.StackCount()
					if availableSpace < srcHas {
						fmt.Println("Destination can't hold all of src's stack.", srcHas, availableSpace)
						// destination can't hold all of src's stack
						if err := s.modifyStackCount(dst.GUID(), tpl.Stackable); err != nil {
							panic(err)
						}

						if err := s.modifyStackCount(src.GUID(), srcHas-availableSpace); err != nil {
							panic(err)
						}
					} else {
						// destination can hold src's stack
						if err := s.modifyStackCount(dst.GUID(), dst.StackCount()+srcHas); err != nil {
							panic(err)
						}

						if _, err := s.removeItemByGUID(src.GUID()); err != nil {
							panic(err)
						}
					}

					return
				}
			}
		}
	}

	s.transferItemUnsafe(srcInv, true, dstBag, dstSlot)

	if dst != nil {
		s.transferItemUnsafe(dstInv, false, srcBag, srcSlot)
	}

	s.SendBagUpdate(srcBag)

	if dstBag == models.Backpack {
		src.SetGUID("Contained", s.GUID())
		s.SendItemUpdate(src)
	} else {
		src.SetGUID("Contained", s.GetBagItem(dstBag).GUID())
		s.SendItemUpdate(src)
	}

	if dstBag != srcBag {
		s.SendBagUpdate(dstBag)
	}
}

func (s *Session) getEquippableInventorySlot(ty models.InventoryType) (models.ItemSlot, item.InvResult) {
	// todo: check for dual wield capability
	if ty == models.IT_Weapon {
		return models.PaperDoll_MainHand, item.InvOk
	}

	if ty == models.IT_Bag {
		for x := models.PaperDoll_Bag1; x <= models.PaperDoll_Bag4; x++ {
			if s.Get("InventorySlots").Index(int(x)).GUID() == guid.Nil {
				return x, item.InvOk
			}
		}
	}

	u, ok := models.ItemDisplaySlots[ty]
	if !ok {
		return 0, item.InvItemCantBeEquipped
	}

	return models.ItemSlot(u), item.InvOk
}

func (s *Session) HandleAutoEquipItem(ae *item.AutoEquipItemRequest) {
	srcInv, src := s.GetItemByPos(ae.SrcBag, ae.SrcSlot)

	if srcInv == nil {
		return
	}

	template := s.GetItemTemplateByEntry(src.Entry())
	dstSlot, err := s.getEquippableInventorySlot(models.InventoryType(template.InventoryType))
	if err != item.InvOk {
		s.Send(&item.InvChangeError{
			Result:  item.InvItemCantBeEquipped,
			SrcItem: src.GUID(),
		})
		return
	}

	s.SwapItem(ae.SrcBag, ae.SrcSlot, models.Backpack, dstSlot)
}

func (s *Session) PaperDollEnd() models.ItemSlot {
	return models.ItemSlot(character.EquipLen(s.Build()))
}

func (s *Session) StartInv() models.ItemSlot {
	switch {
	case vsn.Range(0, 8606).Contains(s.Build()):
		return 23
	}
	return 23
}

func (s *Session) HandleAutoStoreBagItem(ab *item.AutoStoreBag) {
	srcInv, src := s.GetItemByPos(ab.SrcBag, ab.SrcSlot)
	if srcInv == nil {
		return
	}

	if ab.DstBag == models.Backpack {
		slots := s.ValuesBlock.Get("InventorySlots")
		startInv := s.StartInv()
		endInv := models.ItemSlot(slots.Len())
		for i := startInv; i < endInv; i++ {
			slot := slots.Index(int(i))
			if slot.GUID() == guid.Nil {
				s.SwapItem(ab.SrcBag, ab.SrcSlot, models.Backpack, i)
				return
			}
		}

		s.Send(&item.InvChangeError{
			Result:  item.InvInventoryFull,
			SrcItem: src.GUID(),
		})
		return
	}

	bag := s.GetBagItem(ab.DstBag)

	slots := bag.Get("Slots")
	startInv := models.ItemSlot(0)
	endInv := models.ItemSlot(bag.ContainerNumSlots())

	for i := startInv; i < endInv; i++ {
		slot := slots.Index(int(i))
		if slot.GUID() == guid.Nil {
			s.SwapItem(ab.SrcBag, ab.SrcSlot, ab.DstBag, i)
			return
		}
	}

	s.Send(&item.InvChangeError{
		Result:  item.InvInventoryFull,
		SrcItem: src.GUID(),
	})
}

// only call if you have lock
func (s *Session) removeItemByGUID(g guid.GUID) (uint32, error) {
	it, ok := s.Items[g]
	if !ok {
		return 0, fmt.Errorf("no such item: %s", g)
	}

	var inv models.Inventory
	found, err := s.DB().Where("item_id = ?", it.ID()).Get(&inv)
	if err != nil {
		return 0, err
	}

	if !found {
		return 0, fmt.Errorf("could not find inventory slot for item %s", g)
	}

	if inv.Bag == models.Backpack {
		s.ValuesBlock.Get("InventorySlots").Index(int(inv.Slot)).SetGUID(guid.Nil)
		if inv.Slot < 19 {
			s.SetVisibleItemEntry(int(inv.Slot), 0)
			s.UpdatePlayer()
		} else {
			s.UpdateSelf()
		}
	} else {
		bagItem := s.GetBagItem(inv.Bag)
		bagItem.Get("Slots").Index(int(inv.Slot)).SetGUID(guid.Nil)
		s.SendObjectChanges(update.Owner, bagItem)
	}

	stackCount := it.StackCount()

	delete(s.Items, g)
	s.SendObjectDelete(g)

	s.DB().AsyncTx(func(tx *wdb.Tx) {
		tx.Where("item_id = ?", it.ID()).Delete(new(models.Inventory))
		tx.Where("id = ?", it.ID()).Delete(new(models.Item))
	})

	return stackCount, nil
}

func (s *Session) SendItemUpdate(it *Item) {
	packet := etc.NewBuffer()

	enc, err := update.NewEncoder(s.Build(), packet, 1)
	if err != nil {
		panic(err)
	}

	if err = enc.AddBlock(it.GUID(), it.ValuesBlock, update.Owner); err != nil {
		panic(err)
	}

	s.SendRawUpdateObjectData(packet.Bytes(), 0)
	// it.ValuesBlock.ClearChanges()
	// s.ValuesBlock.ClearChanges()
}

func (s *Session) modifyStackCount(item guid.GUID, count uint32) error {
	it, ok := s.Items[item]
	if !ok {
		return fmt.Errorf("no item %s", item)
	}

	var itemData models.Item
	found, err := s.DB().Where("id = ?", it.ID()).Get(&itemData)
	if err != nil {
		return err
	}

	if !found {
		return fmt.Errorf("could not find item in database for %s", it.GUID())
	}

	itemData.StackCount = count
	s.DB().AsyncTx(func(tx *wdb.Tx) {
		tx.Where("id = ?", it.ID()).Cols("stack_count").Update(&itemData)
	})

	it.SetUint32("StackCount", count)
	s.SendItemUpdate(it)

	return nil
}

func (s *Session) RefreshInventory() {
	// var inv []models.Inventory
	// s.DB().Where("player = ?", s.PlayerID()).Find(&inv)

	// // var nheap []models.Inventory

	// // // de-select for equipped items.
	// // for _, v := range inv {
	// // 	if v.Bag == models.Backpack {
	// // 		if v.Slot > 23 {
	// // 			nheap = append(nheap, v)
	// // 		}
	// // 	} else {
	// // 		nheap = append(nheap, v)
	// // 	}
	// // }

	// s.SortInventory()

	return
}

// May fail. run s.VerifyAvailableSpaceFor(itemID) before executing
func (s *Session) AddItem(itemID string, count int, received, created bool) error {
	if count == 0 {
		count = 1
	}

	// get entry from itemID
	var template *models.ItemTemplate
	s.DB().Lookup(wdb.BucketKeyStringID, itemID, &template)

	if template == nil {
		return fmt.Errorf("no such item: %s", itemID)
	}

	s.RefreshInventory()

	if count < 0 {
		// negative count means subtract items
		countRemaining := uint32(-count)
		for _, inventory := range s.Inventory {
			if countRemaining == 0 {
				return nil
			}

			itemGUID := guid.RealmSpecific(guid.Item, s.Server.RealmID(), inventory.ItemID)
			item, ok := s.Items[itemGUID]
			if !ok {
				return fmt.Errorf("no inventory for %s", itemGUID)
			}

			if item.Entry() == template.Entry {
				if template.Stackable == 0 {
					i, err := s.removeItemByGUID(item.GUID())
					if err != nil {
						return err
					}

					countRemaining -= i
				} else {
					stackCount := item.StackCount()

					if stackCount <= countRemaining {
						// This slot has less then the remaining count, so we can just remove it entirely.
						removed, err := s.removeItemByGUID(item.GUID())
						if err != nil {
							return err
						}

						countRemaining -= removed
					} else {
						removed := stackCount - countRemaining
						// this slot has more than the remaining count of items to be removed, so let's remove the remaining count of items to be destroyed from the item object
						if err := s.modifyStackCount(item.GUID(), removed); err != nil {
							return err
						}

						countRemaining -= removed
					}
				}
			}
		}

		if countRemaining > 0 {
			return fmt.Errorf("could not remove %d items", countRemaining)
		}

		return nil
	}

	sentItem := false

	countRemaining := uint32(count)

	if template.Stackable != 0 {
		// See if we have other items of this kind, and if we can merge.
		for _, mergeItem := range s.Items {
			if countRemaining == 0 {
				return nil
			}

			if mergeItem.Entry() == template.Entry {
				stackCount := mergeItem.StackCount()
				if stackCount < template.Stackable {
					// we have a mergeable item slot!
					availableStackCount := template.Stackable - stackCount

					var inv models.Inventory
					fnd, err := s.DB().Where("item_id = ?", mergeItem.ID()).Get(&inv)
					if !fnd {
						panic(err)
					}

					// we can add the remaining items and stop.
					if countRemaining <= availableStackCount {
						if !sentItem {
							s.SendNewItem(mergeItem, received, created, true, inv.Bag, inv.Slot, uint32(count))
							sentItem = true
						}
						s.modifyStackCount(mergeItem.GUID(), stackCount+countRemaining)
						return nil
					}

					s.SendNewItem(mergeItem, received, created, true, inv.Bag, inv.Slot, template.Stackable)
					// we can stack, but it will overflow.
					s.modifyStackCount(mergeItem.GUID(), template.Stackable)

					countRemaining -= availableStackCount
				}
			}
		}
	}

	// Transfer to empty slots.
	if countRemaining > 0 {
		type bagReceiverPos struct {
			Bag  models.ItemSlot
			Slot models.ItemSlot
		}

		var freeSlots []bagReceiverPos

		slots := s.ValuesBlock.Get("InventorySlots")

		for x := s.StartInv(); x < 39; x++ {
			gp := slots.Index(int(x)).GUID()

			if gp == guid.Nil {
				freeSlots = append(freeSlots, bagReceiverPos{
					Bag:  models.Backpack,
					Slot: x,
				})
			}
		}

		for x := models.ItemSlot(0); x < 4; x++ {
			if s.IsValidPos(x, 0) {
				bgItem := s.GetBagItem(x)
				bgSlots := bgItem.Get("Slots")
				for bagSlot := uint32(0); bagSlot < bgItem.ContainerNumSlots(); bagSlot++ {
					gp := bgSlots.Index(int(bagSlot)).GUID()
					if gp == guid.Nil {
						freeSlots = append(freeSlots, bagReceiverPos{
							Bag:  x,
							Slot: models.ItemSlot(bagSlot),
						})
					}
				}
			}
		}

		for _, pos := range freeSlots {
			if countRemaining == 0 {
				break
			}

			newItem := models.Item{
				ItemType:  template.InventoryType,
				ItemID:    itemID,
				DisplayID: 0,
			}

			if template.Stackable != 0 {
				if countRemaining >= template.Stackable {
					newItem.StackCount = template.Stackable
					if newItem.StackCount == 0 {
						newItem.StackCount = 1
					}
				} else {
					newItem.StackCount = countRemaining
				}
			} else {
				newItem.StackCount = 1
			}

			if _, err := s.DB().Insert(&newItem); err != nil {
				panic(err)
			}

			invObject := models.Inventory{
				ItemID: newItem.ID,
				Player: s.PlayerID(),
				Bag:    pos.Bag,
				Slot:   pos.Slot,
			}

			s.DB().AsyncTx(func(tx *wdb.Tx) {
				tx.Insert(&invObject)
			})

			it := s.NewItem(newItem)
			s.Items[it.GUID()] = it

			s.SetBagGUIDSlot(pos.Bag, pos.Slot, it.GUID())
			s.UpdateSelf()

			if !sentItem {
				s.SendNewItem(it, received, created, true, pos.Bag, pos.Slot, uint32(count))
				sentItem = true
			}

			s.SendObjectCreate(it)

			countRemaining -= newItem.StackCount
		}
	}

	// TODO: place in additional bags

	if countRemaining > 0 {
		return fmt.Errorf("could not add %d items", countRemaining)
	}

	return nil
}

// func (s *Session) GetItemCount(entry uint32) uint32 {
// 	tpl := s.GetItemTemplateByEntry(entry)
// 	i64, err := s.DB().Where("player = ?", s.SumInt()
// }

func (s *Session) SendNewItem(newitem *Item, received, created, showInChat bool, bag, slot models.ItemSlot, count uint32) {
	s.Send(&item.Push{
		Recipient:          s.GUID(),
		Received:           received,
		Created:            created,
		ShowInChat:         showInChat,
		Bag:                bag,
		Slot:               slot,
		Count:              count,
		Entry:              newitem.Entry(),
		PropertySeed:       newitem.PropertySeed(),
		RandomPropertiesID: newitem.RandomPropertiesID(),
	})

	// TODO: share with group
}

func (s *Session) HandleDestroyItem(id *item.Destroy) {
	if !s.IsAlive() {
		s.Send(&item.InvChangeError{
			Result: item.InvYouAreDead,
		})
		return
	}

	_, itemObject := s.GetItemByPos(id.Bag, id.Slot)

	if itemObject.IsBag() && itemObject.BagEmpty() == false {
		s.Send(&item.InvChangeError{
			Result: item.InvCanOnlyDoWithEmptyBags,
		})

		return
	}

	if id.Count != 0 {
		s.modifyStackCount(itemObject.GUID(), itemObject.Get("StackCount").Uint32()-uint32(id.Count))
	} else {
		_, err := s.removeItemByGUID(itemObject.GUID())
		if err != nil {
			panic(err)
		}
	}
}

func (s *Session) SetBagGUIDSlot(bag, slot models.ItemSlot, g guid.GUID) {
	if bag == models.Backpack {
		s.ValuesBlock.Get("InventorySlots").Index(int(slot)).SetGUID(g)
		return
	}

	bagItem := s.GetBagItem(bag)
	bagItem.ValuesBlock.Get("Slots").Index(int(slot)).SetGUID(g)
}

func (s *Session) GetBagItem(bag models.ItemSlot) *Item {
	if bag > models.PaperDoll_Bag4 {
		panic("invalid bag")
	}

	bagGUID := s.ValuesBlock.Get("InventorySlots").Index(int(bag)).GUID()

	if bagGUID == guid.Nil {
		panic("failed bag check, call IsValidPos before calling this function")
	}

	bagItem := s.Items[bagGUID]
	if bagItem == nil {
		panic(bagGUID.String() + " refers to non-existent item")
	}

	if bagItem.IsBag() == false {
		panic(bagGUID.String() + " is not a bag")
	}

	return bagItem
}

func (s *Session) SendBagUpdate(bag models.ItemSlot) {
	if bag == models.Backpack {
		s.UpdateSelf()
		// s.Map().PropagateChanges(s.GUID())
		return
	}

	bagItem := s.GetBagItem(bag)
	s.SendItemUpdate(bagItem)
}

func (s *Session) SortInventory() {
	sort.Slice(s.Inventory, func(i, j int) bool {
		_i := s.Inventory[i]
		_j := s.Inventory[j]

		// Backpack always takes precedence to exterior bags
		if _i.Bag == models.Backpack && _j.Bag != models.Backpack {
			return true
		}

		if _j.Bag == models.Backpack && _i.Bag != models.Backpack {
			return false
		}

		if _i.Bag == models.Backpack && _j.Bag == models.Backpack {
			return _i.Slot < _j.Slot
		}

		if _i.Bag != _j.Bag {
			return _i.Bag < _j.Bag
		}

		if _i.Bag == _j.Bag {
			return _i.Slot < _j.Slot
		}

		panic("should never be reached")
		return false
	})
}

func (s *Session) HandleSplitItem(isr *item.SplitRequest) {
	fmt.Println("Splitting item", isr.SrcBag, "(", isr.SrcSlot, ") ->", isr.DstBag, isr.DstSlot)

	if !s.IsValidPos(isr.SrcBag, isr.SrcSlot) {
		s.Warnf("Invalid src pos: %d %d", isr.SrcBag, isr.SrcSlot)
		return
	}

	if !s.IsValidPos(isr.DstBag, isr.DstSlot) {
		s.Warnf("Invalid dst pos: %d %d", isr.DstBag, isr.DstSlot)
		return
	}

	_, src := s.GetItemByPos(isr.SrcBag, isr.SrcSlot)
	if src == nil {
		s.Warnf("Could not find source for that item.")
		return
	}

	// cannot split to equipment
	if s.IsEquipmentPos(isr.DstBag, isr.DstSlot) {
		s.Warnf("Destination is equipment position.")
		return
	}

	tpl := s.GetItemTemplateByEntry(src.Entry())

	if tpl.Stackable == 0 {
		s.Warnf("Template is unstackable.")
		return
	}

	if isr.Count >= src.StackCount() {
		s.Warnf("Attempted to split %d, more than you have in source: %d", isr.Count, src.StackCount())
		return
	}

	if isr.Count >= tpl.Stackable {
		s.Warnf("Attempted to split %d, more than is stackable: %d", isr.Count, tpl.Stackable)
		s.Send(&item.InvChangeError{
			Result:  item.InvTriedToSplitMoreThanCount,
			SrcItem: src.GUID(),
		})
		return
	}

	_, dst := s.GetItemByPos(isr.DstBag, isr.DstSlot)
	if dst != nil {
		if dst.Entry() != src.Entry() {
			s.Send(&item.InvChangeError{
				Result:  item.InvItemCantStack,
				SrcItem: src.GUID(),
				DstItem: dst.GUID(),
			})
			return
		}

		availableSpace := tpl.Stackable - dst.StackCount()
		if isr.Count > availableSpace {
			s.Send(&item.InvChangeError{
				Result:  item.InvItemCantStack,
				SrcItem: src.GUID(),
				DstItem: dst.GUID(),
			})
			return
		}

		if err := s.modifyStackCount(dst.GUID(), dst.StackCount()+isr.Count); err != nil {
			panic(err)
		}
	} else {
		s.modifyStackCount(src.GUID(), src.StackCount()-isr.Count)

		// create new item
		var newItem models.Item
		found, err := s.DB().Where("id = ?", src.ID()).Get(&newItem)
		if !found {
			panic(err)
		}

		newItem.ID = 0
		newItem.StackCount = isr.Count

		if _, err := s.DB().Insert(&newItem); err != nil {
			panic(err)
		}

		newInv := models.Inventory{
			ItemID: newItem.ID,
			Player: s.PlayerID(),
			Bag:    isr.DstBag,
			Slot:   isr.DstSlot,
		}

		newItemObject := s.NewItem(newItem)

		s.DB().Insert(&newInv)

		s.Items[newItemObject.GUID()] = newItemObject

		s.SendObjectCreate(newItemObject)

		s.SetBagGUIDSlot(isr.DstBag, isr.DstSlot, newItemObject.GUID())
		s.SendBagUpdate(isr.DstBag)
	}
}

package commands

import (
	"os"
	"runtime"
	"strings"

	"github.com/Gophercraft/core/guid"
	"github.com/Gophercraft/core/realm"
	"github.com/Gophercraft/core/realm/wdb/models"
)

func cmdDebugInv(s *realm.Session) {
	s.Warnf("Player:")

	for i := 0; i < 39; i++ {
		gid := s.Get("InventorySlots").Index(i).GUID()
		if gid != guid.Nil {
			s.Warnf(" %d: %s", i, gid)
		}
	}

	for i := 19; i < 23; i++ {
		g := s.Get("InventorySlots").Index(i).GUID()
		if g != guid.Nil {
			s.Warnf("Bag %d:", i)

			gArray := s.GetBagItem(models.ItemSlot(i)).Get("Slots")

			for idx := 0; idx < gArray.Len(); idx++ {
				it := gArray.Index(idx).GUID()
				if it != guid.Nil {
					s.Warnf(" %d: %s", idx, s.DebugGUID(it))
				}
			}
		}
	}
}

// func getObjectArgument(c *C) (guid.GUID, error) {

// }

// func x_ForceUpdate(c *C) {

// }

func cmdGoroutines(s *realm.Session) {
	// pprof.Lookup("goroutine").WriteTo(os.Stdout, 1)
	buf := make([]byte, 1<<20)
	stacklen := runtime.Stack(buf, true)
	os.Stdout.Write(buf[:stacklen])
}

func cmdShowSQL(s *realm.Session, on bool) {
	s.DB().ShowSQL(on)
}

func cmdTrackedGUIDs(s *realm.Session) {
	s.GuardTrackedGUIDs.Lock()
	defer s.GuardTrackedGUIDs.Unlock()

	s.Warnf("%d tracked GUIDs:", len(s.TrackedGUIDs))

	for _, id := range s.TrackedGUIDs {
		s.Warnf("%s", s.DebugGUID(id))
	}
}

func cmdListProps(s *realm.Session) {
	var list []string
	for _, prop := range s.CharacterProps {
		list = append(list, string(prop.Value))
	}
	s.Warnf("Props: %s", strings.Join(list, ", "))
}

func cmdToggleBoolProp(s *realm.Session, propId string) {
	var old bool
	var new bool
	s.GetProp(models.PropID(propId), &old)
	new = !old

	s.Warnf("Setting %t to %t", old, new)

	s.SetBoolProp(models.PropID(propId), new)
}

func cmdRemoveProp(s *realm.Session, propId string) {
	if len(propId) > 8 {
		s.Warnf("too many characters for a prop ID")
		return
	}

	s.RemoveProp(models.PropID(propId))
}

func cmdRefreshPlayer(s *realm.Session) {
	invSlots := s.Get("InventorySlots")
	for i := 0; i < invSlots.Len(); i++ {
		slot := invSlots.Index(i)
		slot.SetGUID(slot.GUID())
	}
	s.UpdateSelf()
}

func cmdDebugMapInfo(s *realm.Session) {
	m := s.Map()

	if m != nil {
		m.DebugMapInfo(s)
	}
}

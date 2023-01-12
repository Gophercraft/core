package realm

import (
	"fmt"
	"io"

	"github.com/arl/math32"
	"github.com/superp00t/etc"

	"github.com/Gophercraft/core/home/config"
	"github.com/Gophercraft/core/i18n"
	"github.com/Gophercraft/core/packet/update"
	"github.com/Gophercraft/core/tempest"
	"github.com/Gophercraft/core/vsn"
	"github.com/Gophercraft/log"

	"github.com/Gophercraft/core/guid"
	"github.com/Gophercraft/core/packet"
	_ "github.com/Gophercraft/core/packet/update/descriptorsupport"
	"github.com/Gophercraft/core/realm/wdb"
	"github.com/Gophercraft/core/realm/wdb/models"
)

type SessionSet []*Session

type InstanceType uint32

type Instance struct {
	InstanceType

	Name i18n.Text
}

type Object interface {
	GUID() guid.GUID
	TypeID() guid.TypeID
	// Values() must return a pointer to a values block that can be modified by the server
	Values() *update.ValuesBlock
}

// Objects that have a presence in the world in a specific location (players, creatures)
type WorldObject interface {
	Object
	// movement data
	Movement() *update.MovementBlock
}

// Used to interface with players
type Seer interface {
	Object

	GetSight() *Sight
}

const (
	UseCompressionSmartly int = iota
	ForceCompressionOff
	ForceCompressionOn
)

const updatePacketLengthThreshold = 100

func (s *Session) SendRawUpdateObjectData(encoded []byte, forceCompression int) {
	// initialize uncompressed packet, sending this is sometimes more efficient than enabling zlib compression
	sPacket := packet.NewWorldPacket(packet.SMSG_UPDATE_OBJECT)

	var compressionEnabled = false

	// detect if compression has been forcefully disabled/enabled
	switch forceCompression {
	case UseCompressionSmartly:
		// compression is disabled if there is no benefit
		compressionEnabled = len(encoded) > updatePacketLengthThreshold
	case ForceCompressionOff:
		compressionEnabled = false
	case ForceCompressionOn:
		compressionEnabled = true
	}

	uncompressedLength := uint32(len(encoded))

	if compressionEnabled {
		sPacket = packet.NewWorldPacket(packet.SMSG_COMPRESSED_UPDATE_OBJECT)
		compressedData := packet.Compress(encoded)
		sPacket.WriteUint32(uncompressedLength)
		sPacket.Write(compressedData)

		if s.BoolProp(ObjectDebug) {
			compressionRatio := float64(len(compressedData)) / float64(uncompressedLength)

			s.printfObjMgr("Sending compressed update. %d => %d bytes compression ratio: %f", uncompressedLength, sPacket.Len(), compressionRatio)
		}

	} else {
		if s.BoolProp(ObjectDebug) {
			s.printfObjMgr("Sending uncompressed update. %d bytes", uncompressedLength)
		}
		sPacket.Write(encoded)
	}

	s.SendPacket(sPacket)
}

func (s *Session) Notify(notif UnitNotification, wo WorldObject, args ...any) {
	switch notif {
	case UnitNotifyCreate:
		s.SendObjectCreate(wo)
	case UnitNotifyDelete, UnitNotifyOutOfRange:
		s.SendObjectDelete(wo.GUID())
	case UnitNotifyMovement:
		s.NotifyMovement(args[0].(packet.WorldType), wo)
	}
}

func (s *Session) NotifyMovement(wt packet.WorldType, wo WorldObject) {
	var mp update.MovementPacket
	mp.Type = wt
	mp.Server = true
	mp.GUID = wo.GUID()
	mp.Info = wo.Movement().Info
	s.Send(&mp)
}

func (m *Map) AllObjects() WorldObjectSet {
	objects := m.GetObjectsInRange(tempest.CAaSphere{
		Radius: math32.Inf(1),
	})
	return objects
}

type WorldObjectSet []WorldObject

type FilterFunc func(wo WorldObject) bool

func WithoutGUID(g guid.GUID) FilterFunc {
	return func(wo WorldObject) bool {
		if wo.GUID() == g {
			return false
		}
		return true
	}
}

func IsUnit(wo WorldObject) bool {
	if _, ok := wo.(Unit); ok {
		return true
	}

	return false
}

func (wobjs WorldObjectSet) Filter(fn FilterFunc) WorldObjectSet {
	nw := make(WorldObjectSet, 0, len(wobjs))
	for _, obj := range wobjs {
		if fn(obj) {
			nw = append(nw, obj)
		}
	}
	return nw
}

func (wobjs WorldObjectSet) WithoutGUID(g guid.GUID) WorldObjectSet {
	return wobjs.Filter(WithoutGUID(g))
}

func (wobjs WorldObjectSet) OfTypeID(id guid.TypeID) WorldObjectSet {
	return wobjs.Filter(func(wo WorldObject) bool {
		return wo.TypeID() == id
	})
}

func (m *Map) Vars() config.WorldVars {
	return m.Config().WorldVars
}

func (m *Map) VisibilityDistance() float32 {
	if m.ShowEverywhere {
		return tempest.Infinity
	}
	return m.Phase.Server.FloatVar("Sync.VisibilityRange")
}

// Refers to the objects that nearTo can see.
func (m *Map) VisibleObjects(viewer WorldObject) WorldObjectSet {
	return m.VisibleObjectsInRange(viewer, m.VisibilityDistance())
}

func (m *Map) VisibleObjectsInRange(viewer WorldObject, radius float32) WorldObjectSet {
	rangeSphere := tempest.CAaSphere{
		Position: viewer.Movement().Position.C3(),
		Radius:   radius,
	}

	if seer, ok := viewer.(Seer); ok {
		rangeSphere.Position = seer.GetSight().Position.C3()
	}

	return m.GetObjectsInRange(rangeSphere)
}

// func (m *Map) ObjectsNearPosition(nearPos tempest.C3Vector, limit float32) WorldObjectSet {
// 	wo := m.GetObjectsInRange(tempest.CAaSphere{})

// 	m.Lock()
// 	var wo []WorldObject
// 	for _, obj := range m.Objects {
// 		if nearPos.Distance(obj.Movement().Position.C3()) <= limit {
// 			wo = append(wo, obj)
// 		}
// 	}
// 	m.Unlock()
// 	return wo
// }

// The inverse of VisibleObjects. Refers to WorldObjects that can see Object
func (m *Map) ViewersOf(object WorldObject) WorldObjectSet {
	sphere := tempest.CAaSphere{
		Position: object.Movement().Position.C3(),
	}

	sphere.Radius = m.VisibilityDistance()

	if sphere.Radius != tempest.Infinity {
		sphere.Radius *= 2
	}

	objects := m.GetObjectsInRange(sphere)

	vobjects := objects.Filter(func(viewer WorldObject) bool {
		return m.CanSee(viewer, object, sphere.Radius)
	})

	return vobjects
}

func (wos WorldObjectSet) Iter(iterFunc func(WorldObject)) {
	for _, wo := range wos {
		iterFunc(wo)
	}
}

func (wos WorldObjectSet) Sessions() SessionSet {
	var ss SessionSet

	for _, v := range wos {
		if s, ok := v.(*Session); ok {
			ss = append(ss, s)
		}
	}

	return ss
}

func (s *Session) SendObjectChanges(viewMask update.VisibilityFlags, object Object) {
	packet := etc.NewBuffer()

	enc, err := update.NewEncoder(s.Build(), packet, 1)
	if err != nil {
		panic(err)
	}

	if s.BoolProp(ObjectDebug) {
		s.printfObjMgr("Updating %s: %s", object.GUID(), object.Values().ChangeMask)
	}

	if err = enc.AddBlock(object.GUID(), object.Values(), viewMask); err != nil {
		panic(err)
	}

	s.SendRawUpdateObjectData(packet.Bytes(), 0)
}

func setNumBlocks(buffer *etc.Buffer, num int) {
	wpos := int64(buffer.Len())
	if _, err := buffer.Seek(0, io.SeekStart); err != nil {
		panic(err)
	}
	buffer.WriteUint32(uint32(num))
	buffer.Seek(wpos, io.SeekStart)
}

// Use this function to send manipulated update field changes, without polluting the main store
func (s *Session) ApplyUpdateMods(viewedObject Object, fields *update.ValuesBlock) {

}

// Send spawn packets for 1+ objects. This function is optimized to fit as many Create/Spawn blocks into one compressed packet as possible.
func (s *Session) SendObjectCreate(objects ...Object) {
	if len(objects) == 0 {
		return
	}

	if s.BoolProp(ObjectDebug) {
		s.printfObjMgr("Spawning %d objects", len(objects))
	}

	uPacket := etc.NewBuffer()

	const compression = UseCompressionSmartly

	// The number of blocks will be written afterward with setNumBlocks, use zero as placeholder value.
	enc, err := update.NewEncoder(s.Build(), uPacket, 0)
	if err != nil {
		panic(err)
	}

	var numBlocks int

	for i, wo := range objects {
		log.Println("Serializing", wo.GUID())

		// Which fields should be included?
		viewMask := s.Server.queryRelationshipMask(wo.GUID(), s.GUID())

		movementBlock := &update.MovementBlock{}
		if wobj, ok := wo.(WorldObject); ok {
			movementBlock = wobj.Movement()
			if wo.GUID() == s.GUID() {
				movementBlock.UpdateFlags |= update.UpdateFlagSelf
			}
			if wo.TypeID() == guid.TypePlayer {
				movementBlock.Player = true
			}
		}

		blockType := update.CreateObject

		if wo.TypeID() == guid.TypeUnit || wo.TypeID() == guid.TypePlayer {
			blockType = update.SpawnObject
		}

		objectSize := int(wo.Values().StorageDescriptor.Elem().Type().Size())

		// Guesstimate whether adding this block will overflow the maximum packet length (doesn't need to be perfect)
		sendOffUpdates := uPacket.Len()+objectSize > packet.MaxLength

		if s.Build() == vsn.Alpha && i > 0 {
			sendOffUpdates = true
		}

		if sendOffUpdates {
			setNumBlocks(uPacket, numBlocks)

			if s.BoolProp(ObjectDebug) {
				s.printfObjMgr("Sending object update (uncompressed: %d bytes, %d blocks)", uPacket.Len(), numBlocks)
			}

			// If so, send the packet
			s.SendRawUpdateObjectData(uPacket.Bytes(), compression)
			uPacket = etc.NewBuffer()

			numBlocks = 0

			enc, err = update.NewEncoder(s.Build(), uPacket, 0)
			if err != nil {
				panic(err)
			}
		}

		if s.BoolProp(ObjectDebug) {
			s.printfObjMgr("Creating %s", wo.GUID())
		}

		// Serialize create data to buffer
		if err = enc.AddBlock(wo.GUID(), &update.CreateBlock{
			BlockType:     blockType,
			ObjectType:    wo.TypeID(),
			MovementBlock: movementBlock,
			ValuesBlock:   wo.Values(),
		}, viewMask); err != nil {
			panic(err)
		}

		numBlocks++

		if uPacket.Len() > packet.MaxLength {
			panic("maximum packet length for object creates exceeded")
		}
	}

	if uPacket.Len() > 0 {
		if s.BoolProp(ObjectDebug) {
			s.printfObjMgr("Sending object update (uncompressed: %d bytes, %d blocks)", uPacket.Len(), numBlocks)
		}
		setNumBlocks(uPacket, numBlocks)
		s.SendRawUpdateObjectData(uPacket.Bytes(), compression)
	}
}

func (s *Session) SendObjectDelete(g guid.GUID) {
	if s.BoolProp(ObjectDebug) {
		s.printfObjMgr("Deleting %s", g)
	}
	s.Send(&update.Destroy{
		Object: g,
	})
}

func (s *Session) SendArea(f packet.Codec) {
	wPacket := &packet.WorldPacket{Buffer: etc.NewBuffer()}
	if err := f.Encode(s.Build(), wPacket); err != nil {
		panic(err)
	}

	s.SendAreaPacket(wPacket)
}

func (s *Session) SendAreaPacket(p *packet.WorldPacket) {
	s.SendPacket(p)

	// broadcast
	s.Map().
		ViewersOf(s).
		WithoutGUID(s.GUID()).
		Sessions().
		SendPacket(p)
}

func (s SessionSet) SendPacket(p *packet.WorldPacket) {
	for _, v := range s {
		v.SendPacket(p)
	}
}

func (s SessionSet) Iter(iterFunc func(*Session)) {
	for _, sess := range s {
		iterFunc(sess)
	}
}

func (ws *Server) queryRelationshipMask(src, target guid.GUID) update.VisibilityFlags {
	if src == target {
		return update.Owner
	}

	// todo: determine party relationship

	return 0
}

func (m *Map) PropagateChanges(id guid.GUID) {
	o := m.GetObject(id)

	m.PropagateObjectChanges(o)
}

func (m *Map) PropagateObjectChanges(o WorldObject) {
	valuesStore := o.Values()

	// if the Object is a player session, send them their own changes.
	// if session, ok := o.(*Session); ok {
	// 	session.SendObjectChanges(m.Phase.Server.queryRelationshipMask(o.GUID(), o.GUID()), session)
	// }

	// transmit appropriate changes.
	for _, v := range m.VisibleObjects(o).Sessions() {
		v.SendObjectChanges(m.Phase.Server.queryRelationshipMask(o.GUID(), v.GUID()), o)
	}

	valuesStore.ClearChanges()
}

func (m *Map) Config() *config.World {
	return m.Phase.Server.Config
}

// Send updated fields directly to client. Use for setting private fields.
func (s *Session) UpdateSelf() {
	s.SendObjectChanges(update.Owner, s)
	// s.ClearChanges()
}

// Broadcast changes to nearby players
func (s *Session) UpdatePlayer() {
	s.Map().PropagateObjectChanges(s)
}

func (ws *Server) AllSessions() SessionSet {
	ss := make(SessionSet, len(ws.PlayerList))
	index := 0
	for _, v := range ws.PlayerList {
		ss[index] = v
		index++
	}
	return ss
}

// Increments the dynamic conter for typeID and returns the result.
func (ws *Server) NextDynamicCounter(typeID guid.TypeID) uint64 {
	ws.GuardCounters.Lock()
	next := ws.DynamicCounters[typeID] + 1
	ws.DynamicCounters[typeID] = next
	ws.GuardCounters.Unlock()
	return next
}

func (s *Session) AddTrackedGUID(g guid.GUID) {
	s.GuardTrackedGUIDs.Lock()
	defer s.GuardTrackedGUIDs.Unlock()

	if s.IsTrackedGUID(g) {
		return
	}

	s.TrackedGUIDs = append(s.TrackedGUIDs, g)
}

func (s *Session) RemoveTrackedGUID(g guid.GUID) {
	s.GuardTrackedGUIDs.Lock()
	defer s.GuardTrackedGUIDs.Unlock()

	idx := -1

	for i, v := range s.TrackedGUIDs {
		if v == g {
			idx = i
			break
		}
	}

	if idx >= 0 {
		s.TrackedGUIDs = append(s.TrackedGUIDs[:idx], s.TrackedGUIDs[idx+1:]...)
	}
}

func (ph *Phase) SpawnUnit(id string, mapID uint32, pos tempest.C4Vector) {
	var cr *models.CreatureTemplate
	ph.Server.DB.Lookup(wdb.BucketKeyStringID, id, &cr)
	if cr == nil {
		panic(fmt.Errorf("No CreatureTemplate could be found with the ID %s", id))
		return
	}

	creature := ph.Server.NewCreature(cr, pos)
	ph.Map(mapID).AddObject(creature)
}

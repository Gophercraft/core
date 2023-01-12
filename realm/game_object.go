package realm

import (
	"fmt"
	"math"

	"github.com/Gophercraft/core/guid"
	"github.com/Gophercraft/core/packet/gameobject"
	"github.com/Gophercraft/core/packet/update"
	"github.com/Gophercraft/core/realm/wdb"
	"github.com/Gophercraft/core/realm/wdb/models"
	"github.com/Gophercraft/core/tempest"
	"github.com/Gophercraft/core/vsn"
)

type GameObject struct {
	ID       string
	Map      *Map
	Group    *SpawnGroupInstance
	Position tempest.C4Vector
	*update.ValuesBlock
}

func (g *GameObject) String() string {
	return fmt.Sprintf("<GameObject %s (%s)>", g.ID, g.GUID())
}

func (g *GameObject) GUID() guid.GUID {
	return g.Get("GUID").GUID()
}

func (g *GameObject) TypeID() guid.TypeID {
	return guid.TypeGameObject
}

func (g *GameObject) Values() *update.ValuesBlock {
	return g.ValuesBlock
}

func (g *GameObject) Living() bool {
	return false
}

func (g *GameObject) SetPosition(build vsn.Build, pos tempest.C4Vector) {
	if build.RemovedIn(vsn.V3_0_2) {
		g.SetFloat32("PosX", pos.X)
		g.SetFloat32("PosY", pos.Y)
		g.SetFloat32("PosZ", pos.Z)
		g.SetFloat32("Facing", pos.W)
	}

	g.Position = pos
}

func (g *GameObject) Speeds() update.Speeds {
	return nil
}

func (g *GameObject) SetOrientation(orientation float32) {
	g.Position.W = orientation
	rot := tempest.C4Quaternion{}
	rot.Z = float32(math.Sin(float64(orientation) / 2))
	rot.W = float32(math.Cos(float64(orientation) / 2))
	g.SetRotation(rot)
}

func (g *GameObject) SetRotation(rot tempest.C4Quaternion) {
	array := g.Get("Rotation")
	array.Index(0).SetFloat32(rot.X)
	array.Index(1).SetFloat32(rot.Y)
	array.Index(2).SetFloat32(rot.Z)
	array.Index(3).SetFloat32(rot.W)
}

func (ws *Server) NextGameObjectGUID() guid.GUID {
	g := guid.RealmSpecific(guid.GameObject, ws.RealmID(), ws.NextDynamicCounter(guid.TypeGameObject))
	return g
}

func (ws *Server) NewGameObject(tpl *models.GameObjectTemplate, pos tempest.C4Vector) *GameObject {
	valuesBlock, err := update.NewValuesBlock(ws.Build(), guid.TypeMaskObject|guid.TypeMaskGameObject)
	if err != nil {
		panic(err)
	}
	gobj := &GameObject{
		tpl.ID,
		nil,
		nil,
		pos,
		valuesBlock,
	}

	gobj.SetGUID("GUID", ws.NextGameObjectGUID())

	gobj.SetUint32("Entry", tpl.Entry)
	gobj.SetFloat32("ScaleX", tpl.Size)

	gobj.SetUint32("DisplayID", tpl.DisplayID)
	gobj.SetUint32("Faction", tpl.Faction)

	typeID := gobj.Get("TypeID")

	if typeID != nil {
		typeID.SetUint32(tpl.Type)
	}

	// gobj.SetUint32("TypeID", tpl.Type)

	gobj.SetUint32("Flags", uint32(tpl.Flags))

	gobj.SetPosition(ws.Build(), pos)
	gobj.SetOrientation(pos.W)

	animProgress := gobj.Get("AnimProgress")
	if animProgress != nil {
		animProgress.SetUint32(100)
	}

	gobj.SetUint32("State", 0x01)
	return gobj
}

func (m *Map) SpawnGameObject(gobjID string, pos tempest.C4Vector) error {
	ws := m.Phase.Server

	tpl, err := ws.DB.GetGameObjectTemplate(gobjID)
	if err != nil {
		return err
	}

	gobj := ws.NewGameObject(tpl, pos)

	return m.AddObject(gobj)
}

func (s *Session) HandleGameObjectQuery(goq *gameobject.Query) {
	tpl, _ := s.DB().GetGameObjectTemplateByEntry(goq.ID)

	s.Send(&gameobject.QueryResponse{
		ID:         goq.ID,
		Locale:     s.Locale,
		GameObject: tpl,
	})
}

func (gobj *GameObject) GameObjectType() uint32 {
	server := gobj.Map.Phase.Server
	if server.Build() <= vsn.Alpha {
		var gameobjectInfo *models.GameObjectTemplate
		server.DB.Lookup(wdb.BucketKeyStringID, gobj.ID, &gameobjectInfo)
		if gameobjectInfo == nil {
			return 0
		}
		return gameobjectInfo.Type
	}
	return gobj.Get("TypeID").Uint32()
}

func (gobj *GameObject) Entry() uint32 {
	return gobj.Get("Entry").Uint32()
}

func (s *Session) HandleGameObjectUse(use *gameobject.Use) {
	if use.ID.HighType() != guid.GameObject {
		return
	}

	wo := s.Map().GetObject(use.ID)
	if wo == nil {
		return
	}

	gobj := wo.(*GameObject)

	switch gobj.GameObjectType() {
	case gameobject.TypeChair:
		s.SitChair(gobj)
	}
}

func (s *Session) GetGameObjectTemplateByEntry(entry uint32) *models.GameObjectTemplate {
	var gobjTemplate *models.GameObjectTemplate
	s.DB().Lookup(wdb.BucketKeyEntry, entry, gobjTemplate)
	return gobjTemplate
}

func (gobj *GameObject) Movement() *update.MovementBlock {
	return &update.MovementBlock{
		UpdateFlags: update.UpdateFlagHasPosition,
		Position:    gobj.Position,
		Info: &update.MovementInfo{
			Position: gobj.Position,
		},
	}
}

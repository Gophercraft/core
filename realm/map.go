package realm

import (
	"github.com/Gophercraft/core/guid"
	"github.com/Gophercraft/core/tempest"
	"github.com/Gophercraft/log"
)

func (m *Map) GetObject(id guid.GUID) WorldObject {
	var or objectrequest
	or.Op = o_get | o_byid
	or.Value = id
	or.Confirm = make(chan error)

	if !m.sendrequest(&or) {
		return nil
	}

	err := <-or.Confirm
	if err != nil {
		log.Warn("Problem trying to get object", id, err)
		return nil
	}

	object := or.Value.(WorldObject)
	return object
}

func (m *Map) AddObjects(objects WorldObjectSet) error {
	var or objectrequest
	or.Op = o_modify | o_add
	// or.ID = newObject.GUID()
	or.Confirm = make(chan error)
	or.Value = objects

	if !m.sendrequest(&or) {
		return nil
	}

	return <-or.Confirm
}

func (m *Map) AddObject(newObject ...WorldObject) error {
	return m.AddObjects(newObject)
}

func (m *Map) GetObjectsInRange(vrange tempest.CAaSphere) WorldObjectSet {
	var or objectrequest
	or.Op = o_get | o_byrange
	or.Confirm = make(chan error)
	or.Value = &rangeObjectRequest{
		Range: vrange,
	}
	if !m.sendrequest(&or) {
		return nil
	}
	<-or.Confirm
	return or.Value.(WorldObjectSet)
}

func (m *Map) GetObjectsNearPosition(pos tempest.C3Vector) WorldObjectSet {
	return m.GetObjectsInRange(tempest.CAaSphere{
		Position: pos,
		Radius:   m.VisibilityDistance(),
	})
}

// Update an object's position within the Map. This has bearing on how the object is synchronized in relation to others, but does not modify the object's internal position.
func (m *Map) UpdateMapPosition(object guid.GUID) error {
	var or objectrequest
	or.Op = o_modify | o_position
	or.Confirm = make(chan error)
	or.Value = object

	if !m.sendrequest(&or) {
		return nil
	}

	err := <-or.Confirm
	return err
}

func (m *Map) RemoveObjects(objects []guid.GUID) error {
	var or objectrequest
	or.Op = o_modify
	or.Confirm = make(chan error)
	or.Value = objects

	if !m.sendrequest(&or) {
		return nil
	}

	return <-or.Confirm
}

func (m *Map) RemoveObject(object ...guid.GUID) error {
	return m.RemoveObjects(object)
}

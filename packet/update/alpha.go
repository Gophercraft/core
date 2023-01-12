package update

import (
	"io"

	"github.com/Gophercraft/core/guid"
	"github.com/Gophercraft/core/vsn"
)

func encodeMovementInfoAlpha(version vsn.Build, out io.Writer, mi *MovementInfo) error {
	mi.TransportGUID.EncodeUnpacked(version, out)
	mi.TransportPosition.Encode(out)
	mi.Position.Encode(out)
	writeFloat32(out, mi.Pitch)
	return encodeMoveFlags(version, out, mi.Flags)
}

func (mb *MovementBlock) writeDataAlpha(e *Encoder) error {
	if mb.Info == nil {
		panic(mb.ID.String())
	}

	if err := encodeMovementInfoAlpha(e.Build, e, mb.Info); err != nil {
		return err
	}

	writeUint32(e, mb.Info.FallTime) // fall time?

	var slist []SpeedType
	if err := vsn.QueryDescriptors(e.Build, SpeedLists, &slist); err != nil {
		return err
	}

	for _, sType := range slist {
		writeFloat32(e, mb.Speeds[sType])
	}

	if mb.Info.Flags&MoveFlagSplineEnabled != 0 {
		encodeSplineFlags(e.Build, e, mb.Spline.Flags)

		if mb.Spline.Flags&SplineFinalPoint != 0 {
			mb.Spline.Endpoint.Encode(e)
		}

		if mb.Spline.Flags&SplineFinalTarget != 0 {
			mb.Spline.FacingTarget.EncodeUnpacked(e.Build, e)
		}

		if mb.Spline.Flags&SplineFinalAngle != 0 {
			writeFloat32(e, mb.Spline.FacingAngle)
		}

		writeInt32(e, mb.Spline.TimePassed)
		writeInt32(e, mb.Spline.Duration)
		writeUint32(e, uint32(len(mb.Spline.Spline)))

		for _, pt := range mb.Spline.Spline {
			pt.Encode(e)
		}
	}

	return nil
}

func (cb *CreateBlock) writeDataAlpha(e *Encoder, mask VisibilityFlags, create bool) error {
	err := guid.EncodeTypeID(e.Build, cb.ObjectType, e)
	if err != nil {
		return err
	}
	if err = cb.MovementBlock.WriteData(e, mask, true); err != nil {
		return err
	}

	mb := cb.MovementBlock
	if mb.UpdateFlags&UpdateFlagSelf != 0 {
		writeUint32(e, 1)
	} else {
		writeUint32(e, 0)
	}

	if mb.Player {
		writeUint32(e, 1)
	} else {
		writeUint32(e, 0)
	}

	writeUint32(e, 0) // Timer ID?
	// writeUint32(e, 0)

	mb.Victim.EncodeUnpacked(e.Build, e)

	return cb.ValuesBlock.WriteData(e, mask, true)
}

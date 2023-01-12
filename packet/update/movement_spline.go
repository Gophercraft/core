package update

import (
	"fmt"
	"io"

	"github.com/Gophercraft/core/guid"
	"github.com/Gophercraft/core/tempest"
	"github.com/Gophercraft/core/vsn"
)

type SplineFlags uint32
type SplineFlagDescriptor map[SplineFlags]uint32

const (
	SplineNone SplineFlags = 1 << iota
	// x00-xFF(first byte) used as animation Ids storage in pair with Animation flag
	SplineDone
	SplineFalling // Affects elevation computation, can't be combined with Parabolic flag
	SplineNoSpline
	SplineParabolic // Affects elevation computation, can't be combined with Falling flag
	SplineCanSwim
	SplineFlying           // Smooth movement(Catmullrom interpolation mode), flying animation
	SplineOrientationFixed // Model ori    = 0x00008000,
	SplineFinalTarget
	SplineFinalPoint
	SplineFinalAngle
	SplineCatmullrom // Used Catmullrom interpolation mode
	SplineCyclic     // Movement by cycled spline
	SplineEnterCycle // Everytimes appears with cyclic flag in monster move packet, erases first spline vertex after first cycle done
	SplineAnimation  // Plays animation after some time passed
	SplineFrozen     // Will never arrive
	SplineTransportEnter
	SplineTransportExit
	SplineBackward
	SplineWalkmode
	SplineBoardVehicle
	SplineExitVehicle
	SplineOrientationInverted
)

var (
	SplineFlagDescriptors = map[vsn.BuildRange]SplineFlagDescriptor{
		{0, 5875}: {
			SplineDone:        0x00000001,
			SplineFalling:     0x00000002,
			SplineFlying:      0x00000200,
			SplineNoSpline:    0x00000400,
			SplineFinalPoint:  0x00010000,
			SplineFinalTarget: 0x00020000,
			SplineFinalAngle:  0x00040000,
			SplineCyclic:      0x00100000,
			SplineEnterCycle:  0x00200000,
			SplineFrozen:      0x00400000,
		},

		{vsn.V2_0_1, vsn.V3_3_5a}: {
			// x00-xFF(first byte) used as animation Ids storage in pair with Animation flag
			SplineDone:                0x00000100,
			SplineFalling:             0x00000200, // Affects elevation computation, can't be combined with Parabolic flag
			SplineNoSpline:            0x00000400,
			SplineParabolic:           0x00000800, // Affects elevation computation, can't be combined with Falling flag
			SplineWalkmode:            0x00001000,
			SplineFlying:              0x00002000, // Smooth movement(Catmullrom interpolation mode), flying animation
			SplineOrientationFixed:    0x00004000, // Model orientation fixed
			SplineFinalPoint:          0x00008000,
			SplineFinalTarget:         0x00010000,
			SplineFinalAngle:          0x00020000,
			SplineCatmullrom:          0x00040000, // Used Catmullrom interpolation mode
			SplineCyclic:              0x00080000, // Movement by cycled spline
			SplineEnterCycle:          0x00100000, // Everytimes appears with cyclic flag in monster move packet, erases first spline vertex after first cycle done
			SplineAnimation:           0x00200000, // Plays animation after some time passed
			SplineFrozen:              0x00400000, // Will never arrive
			SplineBoardVehicle:        0x00800000,
			SplineExitVehicle:         0x01000000,
			SplineOrientationInverted: 0x08000000,
		},
	}
)

type MoveSpline struct {
	Flags                SplineFlags
	Facing               tempest.C3Vector
	FacingTarget         guid.GUID
	FacingAngle          float32
	TimePassed           int32
	Duration             int32
	ID                   uint32
	DurationMod          float32
	DurationModNext      float32
	VerticalAcceleration float32
	EffectStartTime      int32
	Spline               []tempest.C3Vector
	SplineMode           uint8
	Endpoint             tempest.C3Vector
}

func decodeSplineFlags(version vsn.Build, in io.Reader) (SplineFlags, error) {
	var descriptor SplineFlagDescriptor
	err := vsn.QueryDescriptors(version, SplineFlagDescriptors, &descriptor)
	if err != nil {
		return 0, fmt.Errorf("update: getting spline descriptor: %s", err)
	}

	sf, err := readUint32(in)
	if err != nil {
		return 0, err
	}

	out := SplineFlags(0)
	// translate packet bits to virtual Gophercraft bits
	for k, v := range descriptor {
		if sf&v != 0 {
			out |= k
		}
	}

	return out, nil
}

func encodeSplineFlags(version vsn.Build, out io.Writer, sf SplineFlags) error {
	var descriptor SplineFlagDescriptor
	err := vsn.QueryDescriptors(version, SplineFlagDescriptors, &descriptor)
	if err != nil {
		return fmt.Errorf("update: getting spline descriptor: %s", err)
	}

	u32 := uint32(0)

	for k, v := range descriptor {
		if sf&k != 0 {
			u32 |= v
		}
	}

	return writeUint32(out, u32)
}

func DecodeMoveSpline(version vsn.Build, in io.Reader) (*MoveSpline, error) {
	ms := new(MoveSpline)
	var err error
	ms.Flags, err = decodeSplineFlags(version, in)
	if err != nil {
		return nil, err
	}

	// Flag order reversed
	if version.AddedIn(vsn.V2_4_3) {
		if ms.Flags&SplineFinalAngle != 0 {
			ms.FacingAngle, err = readFloat32(in)
			if err != nil {
				return nil, err
			}

		} else if ms.Flags&SplineFinalTarget != 0 {
			ms.FacingTarget, err = guid.DecodeUnpacked(version, in)
			if err != nil {
				return nil, err
			}
		} else if ms.Flags&SplineFinalPoint != 0 {
			ms.Facing, err = tempest.DecodeC3Vector(in)
			if err != nil {
				return nil, err
			}
		}
	} else {
		if ms.Flags&SplineFinalPoint != 0 {
			ms.Facing, err = tempest.DecodeC3Vector(in)
			if err != nil {
				return nil, err
			}
		} else if ms.Flags&SplineFinalTarget != 0 {
			ms.FacingTarget, err = guid.DecodeUnpacked(version, in)
			if err != nil {
				return nil, err
			}
		} else if ms.Flags&SplineFinalAngle != 0 {
			ms.FacingAngle, err = readFloat32(in)
			if err != nil {
				return nil, err
			}
		}
	}

	ms.TimePassed, err = readInt32(in)
	if err != nil {
		return nil, err
	}
	ms.Duration, err = readInt32(in)
	if err != nil {
		return nil, err
	}
	ms.ID, err = readUint32(in)
	if err != nil {
		return nil, err
	}

	if version.AddedIn(vsn.V3_3_5a) {
		ms.DurationMod, err = readFloat32(in)
		if err != nil {
			return nil, err
		}
		ms.DurationModNext, err = readFloat32(in)
		if err != nil {
			return nil, err
		}
		ms.VerticalAcceleration, err = readFloat32(in)
		if err != nil {
			return nil, err
		}
		ms.EffectStartTime, err = readInt32(in)
		if err != nil {
			return nil, err
		}
	}

	nodeLength, err := readInt32(in)
	if err != nil {
		return nil, err
	}

	if nodeLength > 0xFFFF {
		return nil, fmt.Errorf("spline overread")
	}

	for i := int32(0); i < nodeLength; i++ {
		p3, err := tempest.DecodeC3Vector(in)
		if err != nil {
			return nil, err
		}
		ms.Spline = append(ms.Spline, p3)
	}

	if version.AddedIn(vsn.V3_3_5a) {
		ms.SplineMode, err = readUint8(in)
		if err != nil {
			return nil, err
		}
	}

	ms.Endpoint, err = tempest.DecodeC3Vector(in)
	return ms, err
}

func EncodeMoveSpline(version vsn.Build, out io.Writer, ms *MoveSpline) error {
	if err := encodeSplineFlags(version, out, ms.Flags); err != nil {
		return err
	}

	if ms.Flags&SplineFinalPoint != 0 {
		if err := ms.Facing.Encode(out); err != nil {
			return err
		}
	} else if ms.Flags&SplineFinalTarget != 0 {
		ms.FacingTarget.EncodeUnpacked(version, out)
	} else if ms.Flags&SplineFinalAngle != 0 {
		writeFloat32(out, ms.FacingAngle)
	}

	writeInt32(out, ms.TimePassed)
	writeInt32(out, ms.Duration)
	writeUint32(out, ms.ID)

	if version.AddedIn(vsn.V3_3_5a) {
		writeFloat32(out, ms.DurationMod)
		writeFloat32(out, ms.DurationModNext)
		writeFloat32(out, ms.VerticalAcceleration)
		writeInt32(out, ms.EffectStartTime)
	}

	writeUint32(out, uint32(len(ms.Spline)))

	for _, p3 := range ms.Spline {
		if err := p3.Encode(out); err != nil {
			return err
		}
	}

	if version.AddedIn(vsn.V3_3_5a) {
		writeUint8(out, ms.SplineMode)
	}

	err := ms.Endpoint.Encode(out)
	return err
}

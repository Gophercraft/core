package login

import (
	"time"

	"github.com/Gophercraft/core/game/protocol/message"
	"github.com/Gophercraft/core/version"
)

const (
	minMask      = 0x3f
	hourMask     = 0x1f
	weekdayMask  = 7
	monthdayMask = 0x3f
	monthMask    = 0xf
	yearMask     = 0x1f
	flagsMask    = 0x3

	hourShift     = 6
	weekdayShift  = 11
	monthdayShift = 14
	monthShift    = 20
	yearShift     = 24
	flagsShift    = 29
)

type GameTime struct {
	Minute   int32
	Hour     int32
	Weekday  int32
	Monthday int32
	Month    int32
	Year     int32
	Flags    int32
}

func (gt *GameTime) SetTime(tm time.Time) {
	gt.Minute = int32(tm.Minute())
	gt.Hour = int32(tm.Hour())
	gt.Weekday = int32(tm.Weekday())
	gt.Monthday = int32(tm.Day() - 1)
	gt.Month = int32(tm.Month()) - 1
	gt.Year = int32(tm.Year() - 2000)
}

func (t *GameTime) Time(loc *time.Location) time.Time {
	var (
		minute   int
		hour     int
		monthday int
		month    time.Month
		year     int
	)

	minute = int(t.Minute)
	hour = int(t.Hour)
	monthday = int(t.Weekday) + 1
	month = time.Month(t.Month)
	year = int(t.Year) + 2000

	return time.Date(
		year,
		month,
		monthday,
		hour,
		minute,
		0,
		0,
		loc,
	)
}

func (t *GameTime) Pack() (value int32) {
	value = t.Minute & minMask
	value |= (t.Hour & hourMask) << hourShift
	value |= (t.Weekday & weekdayMask) << weekdayShift
	value |= (t.Monthday & monthdayMask) << monthdayShift
	value |= (t.Month & monthMask) << monthShift
	value |= (t.Year & yearMask) << yearShift
	value |= (t.Flags & flagsMask) << flagsShift
	return
}

func (t *GameTime) Unpack(value int32) {
	t.Minute = value & minMask
	if t.Minute == minMask {
		t.Minute = -1
	}

	t.Hour = value >> hourShift & hourMask
	if t.Hour == hourMask {
		t.Hour = -1
	}

	t.Weekday = value >> weekdayShift & weekdayMask
	if t.Weekday == weekdayMask {
		t.Weekday = -1
	}

	t.Monthday = value >> monthdayShift & monthdayMask
	if t.Monthday == monthdayMask {
		t.Monthday = -1
	}

	t.Month = value >> monthShift & monthMask
	if t.Month == monthMask {
		t.Month = -1
	}

	t.Year = value >> yearShift & yearMask
	if t.Year == yearMask {
		t.Year = -1
	}

	t.Flags = value >> flagsShift & flagsMask
	if t.Flags == flagsMask {
		t.Flags = -1
	}
}

const (
	// Default = 1 minute / 60 seconds
	DefaultTimeSpeed float32 = 1.0 / 60.0
)

type SetTimeSpeed struct {
	Time  time.Time
	Speed float32
}

func (set *SetTimeSpeed) Encode(build version.Build, out *message.Packet) error {
	out.Type = message.SMSG_LOGIN_SETTIMESPEED

	var gt GameTime
	gt.SetTime(set.Time)

	out.WriteInt32(gt.Pack())
	out.WriteFloat32(set.Speed)

	return nil
}

func (set *SetTimeSpeed) Decode(build version.Build, in *message.Packet) error {
	var gt GameTime
	gt.Unpack(in.ReadInt32())

	set.Time = gt.Time(time.Local)
	set.Speed = in.ReadFloat32()
	return nil
}

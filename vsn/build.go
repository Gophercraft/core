// Package vsn provides utilities for handling protocol versions, as well as Gophercraft software versions.
package vsn

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Gophercraft/log"
)

// A Build refers to the snapshot of a protocol that Gophercraft is trying to interact with.
type Build uint32

var (
	names = map[Build]string{
		3368:  "Alpha",
		5875:  "Vanilla",
		8606:  "TBC",
		12340: "WoTLK",
		33369: "BfA",
	}

	Max Build = 0xFFFFFFFF
)

const (
	Alpha   Build = 3368
	V1_12_1 Build = 5875
	V1_12_2 Build = 6005
	V2_0_1  Build = 6180
	V2_4_3  Build = 8606
	V3_0_2  Build = 9056
	V3_3_5a Build = 12340
	V4_0_1  Build = 13164
	V4_3_4  Build = 15595
	V5_4_8  Build = 18414
	V6_2_4  Build = 21742
	V8_3_0  Build = 33369
)

const NewAuthSystem = V6_2_4
const NewCryptSystem = V8_3_0

func (b Build) String() string {
	info := details[b]
	if info == nil {
		return fmt.Sprintf("unknown version (%d)", b)
	}

	str := b.SemVer()
	if name := names[b]; name != "" {
		str += " " + name
	}

	return fmt.Sprintf("%s (%d)", str, b)
}

func (b Build) SemVer() string {
	info := details[b]
	if info == nil {
		return "?.?.?"
	}
	str := fmt.Sprintf("%d.%d.%d", info.MajorVersion, info.MinorVersion, info.BugfixVersion)
	if info.HotfixVersion != "" {
		str += info.HotfixVersion
	}
	return str
}

func (b Build) DBD() string {
	return fmt.Sprintf("%s.%d", b.SemVer(), b)
}

func (hasFeature Build) AddedIn(update Build) bool {
	if hasFeature >= update {
		return true
	}

	return false
}

func (hasFeature Build) RemovedIn(update Build) bool {
	if hasFeature < update {
		return true
	}

	return false
}

// Accepts DBD formatted builds e.g. 0.5.3.3368
func ParseDBD(str string) (Build, error) {
	parts := strings.Split(str, ".")
	last := parts[len(parts)-1]
	u, err := strconv.ParseUint(last, 0, 32)
	if err != nil {
		return 0, err
	}
	return Build(u), nil
}

// var Exp = map[BuildRange]uint8{
// 	{0, 6005}:      0,
// 	{6180, 8606}:   1,
// 	{9056, 12340}:  2,
// 	{13164, 15595}: 3,
// 	{,18414}
// }

func (b Build) Exp() int {
	info := b.BuildInfo()
	if info == nil {
		log.Warn("No info was found for build. Unable to determine expansion value", b)
		return 0
	}

	exp := int(info.MajorVersion) - 1
	if exp < 0 {
		exp = 0
	}
	return exp
}

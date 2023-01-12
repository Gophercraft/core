package vsn

import (
	"fmt"
	"strconv"
	"strings"
)

var GophercraftVersion CoreVersion

func init() {
	var err error
	GophercraftVersion, err = ParseCoreVersion("0.5.1")
	if err != nil {
		panic(err)
	}
}

type Rev uint16

// Refers to the current version of Gophercraft core.
type CoreVersion struct {
	Major Rev
	Minor Rev
	Patch Rev
}

func (cv CoreVersion) String() string {
	return fmt.Sprintf("%d.%d.%d", cv.Major, cv.Minor, cv.Patch)
}

func (i CoreVersion) LessThan(j CoreVersion) bool {
	return i.Major < j.Major || i.Minor < j.Minor || i.Patch < j.Patch
}

func (i CoreVersion) Cmp(j CoreVersion) int {
	r := 0
	switch {
	case i.LessThan(j):
		r = -1
	case i == j:
		r = 0
	case j.LessThan(i):
		r = 1
	}
	return r
}

func rev(s string) (Rev, error) {
	u, err := strconv.ParseUint(s, 10, 16)
	if err != nil {
		return 0, err
	}
	return Rev(u), nil
}

func ParseCoreVersion(str string) (CoreVersion, error) {
	strs := strings.SplitN(str, ".", 3)
	cv := CoreVersion{}
	var err error
	cv.Major, err = rev(strs[0])
	if err != nil {
		return cv, err
	}
	cv.Minor, err = rev(strs[1])
	if err != nil {
		return cv, err
	}
	cv.Patch, err = rev(strs[2])
	return cv, err
}

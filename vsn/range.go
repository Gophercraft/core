package vsn

import "fmt"

type BuildRange [2]Build

func (br BuildRange) String() string {
	return fmt.Sprintf("%s-%s", br[0].DBD(), br[1].DBD())
}

func Range(min, max Build) BuildRange {
	if min > max {
		panic(fmt.Errorf("vsn: invalid range: %s-%s", min, max))
	}
	return BuildRange{min, max}
}

func (br BuildRange) Contains(v Build) bool {
	return v >= br[0] && v <= br[1]
}

func (b BuildRange) Unary() bool {
	return b[0] == b[1]
}

package version

import (
	"math"
	"strconv"
	"strings"
)

type Selector string

func (sl Selector) Match(b Build) bool {
	s := strings.Split(string(sl), "-")

	min := float64(0)
	max := math.Inf(1)

	if s[0] != "" {
		var err error
		min, err = strconv.ParseFloat(s[0], 64)
		if err != nil {
			panic(err)
		}
	}

	if s[1] != "" {
		var err error
		max, err = strconv.ParseFloat(s[1], 64)
		if err != nil {
			panic(err)
		}
	}

	fb := float64(b)

	return fb >= min && fb <= max
}

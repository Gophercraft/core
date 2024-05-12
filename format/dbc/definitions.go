package dbc

import (
	"fmt"

	"github.com/Gophercraft/core/version"
)

func LocStringSize(v version.Build) (int, error) {
	switch {
	case v < 6692:
		// const char *xx_lang[8];
		// int         xx_flag;
		return 9, nil
	case version.Range(6692, 12340).Contains(v):
		// const char *xx_lang[16];
		// int         xx_flag;
		return 17, nil
		// no bitmask
	case version.Range(13164, version.Max).Contains(v):
		return 1, nil
	default:
		return 0, fmt.Errorf("dbc: LocStringSize unknown build %s", v)
	}
}

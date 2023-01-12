package dbc

import (
	"fmt"

	"github.com/Gophercraft/core/vsn"
)

func LocStringSize(v vsn.Build) (int, error) {
	switch {
	case v < 6692:
		// const char *xx_lang[8];
		// int         xx_flag;
		return 9, nil
	case vsn.Range(6692, 12340).Contains(v):
		// const char *xx_lang[16];
		// int         xx_flag;
		return 17, nil
		// no bitmask
	case vsn.Range(13164, vsn.Max).Contains(v):
		return 1, nil
	default:
		return 0, fmt.Errorf("dbc: LocStringSize unknown build %s", v)
	}
}

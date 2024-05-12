package models

type ItemQuality uint8

const (
	Poor ItemQuality = iota
	Normal
	Uncommon
	Rare
	Epic
	Legendary
	Artifact
	Heirloom
	WowToken
)

var (
	colorQuality = map[ItemQuality]string{
		Poor:      "ff9d9d9d",
		Normal:    "ffffffff",
		Uncommon:  "ff1eff00",
		Rare:      "ff0070dd",
		Epic:      "ffa335ee",
		Legendary: "ffff8000",
		Artifact:  "ffe6cc80",
		Heirloom:  "ffe6cc80",
		WowToken:  "ff00ccff",
	}
)

func (i ItemQuality) Color() string {
	c, ok := colorQuality[i]
	if !ok {
		return colorQuality[Normal]
	}

	return c
}

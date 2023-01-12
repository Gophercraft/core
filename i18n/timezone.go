package i18n

import (
	"fmt"
	"strconv"
)

type Timezone uint32

var (
	TimezoneNames = map[Timezone]string{
		1:  "Development",
		2:  "United States",
		3:  "Oceanic",
		4:  "Latin America",
		5:  "Tournament",
		6:  "Korea",
		7:  "Tournament 1",
		8:  "English",
		9:  "German",
		10: "French",
		11: "Spanish",
		12: "Russian",
		13: "Tournament 2",
		14: "Taiwan",
		15: "Tournament 3",
		16: "China",
		17: "CN1",
		18: "CN2",
		19: "CN3",
		20: "CN4",
		21: "CN5",
		22: "CN6",
		23: "CN7",
		24: "CN8",
		25: "Tournament 4",
		26: "Test Server",
		27: "Tournament 5",
		28: "QA Server",
		29: "CN9",
		30: "Test Server 2",
		31: "CN10",
		32: "CTC",
		33: "CNC",
		34: "CN1/4",
		35: "CN/2/6/9",
		36: "CN3/7",
		37: "Russian Tournament",
		38: "CN5/8",
		39: "CN11",
		40: "CN12",
		41: "CN13",
		42: "CN14",
		43: "CN15",
		44: "CN16",
		45: "CN17",
		46: "CN18",
		47: "CN19",
		48: "CN20",
		49: "Brazil",
		50: "Italian",
		51: "Hyrule",
		52: "QA2 Test",
		53: "",
		54: "",
		55: "Recommended Realm",
		56: "Test",
		57: "Recommended Realm 2",
		58: "",
		59: "Future Test",
	}
)

func (t *Timezone) DecodeWord(str string) error {
	for tz, name := range TimezoneNames {
		if name == str {
			*t = tz
			return nil
		}
	}

	u, err := strconv.ParseUint(str, 10, 32)
	if err != nil {
		return err
	}

	*t = Timezone(u)

	return nil
}

func (t *Timezone) EncodeWord() (string, error) {
	tzn, ok := TimezoneNames[*t]
	if ok {
		return tzn, nil
	}

	return fmt.Sprintf("%d", *t), nil
}

func (t Timezone) String() string {
	ts, err := t.EncodeWord()
	if err != nil {
		return fmt.Sprintf("Timezone(%d)", t)
	}
	return ts
}
